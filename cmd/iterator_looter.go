package cmd

import (
	"context"
	"crabada-bot/config"
	"crabada-bot/internal/looter"
	"crabada-bot/internal/player"
	"crabada-bot/internal/query"
	"crabada-bot/internal/storage"
	"crabada-bot/pkg/constant"
	"crabada-bot/pkg/crabada"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var iteratorLooterCmd = &cobra.Command{
	Use:   "loot-iterate",
	Short: "Iterator looter mode",
	Long:  `Iterator looter mode`,
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
		log.Info().Str("config", config.ToString()).Msg("started as loot-iterate mode")

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
		transactorContract, callerContract, _, err := createContractInstances()

		teamQueryService := query.NewTeamQueryService(config.CacheEnabled, callerContract, storageService)

		player, err := player.NewPlayer(transactorContract, apiClient, config)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		looter := looter.NewLootIterator(player, callerContract, apiClient, teamQueryService)
		err = looter.StartLootIterator(ctx, cancel)
		if err != nil {
			log.Error().Err(err).Send()
		}
	},
}

func init() {
	rootCmd.AddCommand(iteratorLooterCmd)

	iteratorLooterCmd.Flags().StringVarP(&privateKey, "key", "k", "", "Account private key")
	iteratorLooterCmd.Flags().IntVarP(&gasTip, "gasTip", "g", 50, "Transaction gas tip for miners.")
	iteratorLooterCmd.Flags().IntVarP(&gasCapLimit, "gasFeeLimit", "l", 150, "Transaction gas fee cap.")
	iteratorLooterCmd.Flags().IntVarP(&battlePointMargin, "bMargin", "b", 200, "Battle point margin with your team battle point.")
	iteratorLooterCmd.Flags().IntVarP(&minePointLimit, "mLimit", "m", 250, "Mine point upper limit ")
	iteratorLooterCmd.Flags().BoolP("cache", "c", false, "Cache active or passive")
	iteratorLooterCmd.Flags().IntVarP(&reinforceLimit, "reinforceLimit", "r", 10, "Reinforce limit")

	if err := iteratorLooterCmd.MarkFlagRequired("key"); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := iteratorLooterCmd.MarkFlagRequired("gasTip"); err != nil {
		log.Fatal().Err(err).Send()
	}
}