package cmd

import (
	"context"
	"crabada-bot/config"
	"crabada-bot/internal/chain"
	"crabada-bot/internal/looter"
	"crabada-bot/internal/player"
	"crabada-bot/internal/query"
	"crabada-bot/internal/storage"
	"crabada-bot/pkg/constant"
	"crabada-bot/pkg/crabada"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	privateKey        string
	gasTip            int
	gasCapLimit       int
	battlePointMargin int
	minePointLimit    int
	reinforceLimit    int
)

var looterCmd = &cobra.Command{
	Use:   "loot",
	Short: "Subscriber looter mode",
	Long:  `Subscriber looter mode`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		interruptAndTimeoutToCancel(ctx, cancel)

		cacheEnabled, _ := cmd.Flags().GetBool("cache")

		config := &config.Config{
			CacheEnabled: cacheEnabled,
			PrivateKey:   privateKey,
			GasTip:       gasTip,
			GasFeeCap:    gasCapLimit,
			AttackParameter: &config.AttackParameter{
				BattlePointMargin: battlePointMargin,
				MinePointLimit:    minePointLimit,
				ReinforceLimit:    reinforceLimit,
			},
		}
		log.Info().Str("config", config.ToString()).Msg("started as subscriber-looter mode")

		apiClient, err := crabada.NewApiClient(constant.CrabadaApiUrl)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		storageService := storage.NewDefaultStorageService(cache.New(storage.ExpireInOneHour, storage.CleanupInOneHour))
		if config.CacheEnabled {
			err = storageService.Load()
			if err != nil {
				log.Fatal().Err(err).Send()
			}
		}
		transactorContract, callerContract, filtererContract, err := createContractInstances()

		teamQueryService := query.NewTeamQueryService(config.CacheEnabled, callerContract, storageService)

		chainService, err := chain.NewChainService(constant.CustomNodeHttpAddress)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		player, err := player.NewPlayer(transactorContract, apiClient, config)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		looter := looter.NewLooter(player, callerContract, filtererContract, chainService, teamQueryService)
		_, err = looter.StartLoot(ctx)
		if err != nil {
			log.Error().Err(err).Send()
		}
	},
}

func createContractInstances() (*crabada.CrabadaTransactor, *crabada.CrabadaCaller, *crabada.CrabadaFilterer, error) {

	address := constant.GetWsChainAddress()
	wsChainClient, err := ethclient.Dial(address)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	/*
		customNodeWsAddress, err := ethclient.Dial(constant.CustomNodeWsAddress)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

	*/

	avaxHttpChainClient, err := ethclient.Dial(constant.AvaxHttpAddress)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	filtererContract, err := crabada.NewCrabadaFilterer(common.HexToAddress(constant.CrabadaContractAddress), wsChainClient)
	if err != nil {
		return nil, nil, nil, err
	}

	transactorContract, err := crabada.NewCrabadaTransactor(common.HexToAddress(constant.CrabadaContractAddress), avaxHttpChainClient)
	if err != nil {
		return nil, nil, nil, err
	}

	callerContract, err := crabada.NewCrabadaCaller(common.HexToAddress(constant.CrabadaContractAddress), wsChainClient)
	if err != nil {
		return nil, nil, nil, err
	}
	return transactorContract, callerContract, filtererContract, nil
}

func interruptAndTimeoutToCancel(ctx context.Context, cancel context.CancelFunc) {
	go func() {
		timeOut := time.NewTicker(constant.LootTimeout)
		signalStream := make(chan os.Signal)
		signal.Notify(signalStream, os.Interrupt, syscall.SIGTERM)
		defer func() {
			timeOut.Stop()
			signal.Stop(signalStream)
			close(signalStream)
		}()

		select {
		case <-signalStream:
			cancel()
			return
		case <-timeOut.C:
			log.Info().Msg("timeout occurred close script ")
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}()
}

func init() {
	rootCmd.AddCommand(looterCmd)

	looterCmd.Flags().StringVarP(&privateKey, "key", "k", "", "Account private key")
	looterCmd.Flags().IntVarP(&gasTip, "gasTip", "g", 50, "Transaction gas tip for miners.")
	looterCmd.Flags().IntVarP(&gasCapLimit, "gasFeeLimit", "l", 150, "Transaction gas fee cap.")
	looterCmd.Flags().IntVarP(&battlePointMargin, "bMargin", "b", 200, "Battle point margin with your team battle point.")
	looterCmd.Flags().IntVarP(&minePointLimit, "mLimit", "m", 250, "Mine point upper limit ")
	looterCmd.Flags().BoolP("cache", "c", false, "Cache active or passive")
	looterCmd.Flags().IntVarP(&reinforceLimit, "reinforceLimit", "r", 10, "Reinforce limit")

	if err := looterCmd.MarkFlagRequired("key"); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := looterCmd.MarkFlagRequired("gasTip"); err != nil {
		log.Fatal().Err(err).Send()
	}
}
