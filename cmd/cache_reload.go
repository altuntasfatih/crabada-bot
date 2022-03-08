package cmd

import (
	"context"
	"crabada-bot/internal/chain"
	"crabada-bot/internal/storage"
	"crabada-bot/pkg/constant"
	"crabada-bot/pkg/crabada"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"math/big"
	"strconv"
	"sync"
)

var goRoutineCount int

type JobParameter struct {
	id               int
	reinforceEnabled bool
	startIndex       int
	finishIndex      int
	wg               *sync.WaitGroup
}

func init() {
	rootCmd.AddCommand(cacheReloadCmd)
	cacheReloadCmd.Flags().BoolP("reinforce", "r", false, "Reinforce look is enabled")
	cacheReloadCmd.Flags().IntVarP(&goRoutineCount, "count", "c", 1, "Goroutine count")
}

var cacheReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload cache",
	Long:  "It reloads cache with new data",
	Run: func(cmd *cobra.Command, args []string) {
		reinforceLookEnabled, _ := cmd.Flags().GetBool("reinforce")
		log.Info().Msgf("Cache reload is started with goroutine count: %d and reinforce look: %v", goRoutineCount, reinforceLookEnabled)

		chainService, err := chain.NewChainService(constant.GetWsChainAddress())
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		contract, err := crabada.NewCrabadaCaller(common.HexToAddress(constant.CrabadaContractAddress), chainService.Client)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		apiClient, err := crabada.NewApiClient(constant.CrabadaApiUrl)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		storageService := storage.NewDefaultStorageService(cache.New(storage.ExpireInOneHour, storage.CleanupInOneHour))
		var wg sync.WaitGroup

		offset := constant.CacheReloadEndTeamId / goRoutineCount
		for i := 0; i < goRoutineCount; i++ {
			wg.Add(1)
			job := &JobParameter{i, reinforceLookEnabled, i * offset, (i * offset) + offset, &wg}
			go startCacheReloadJob(job, contract, apiClient, storageService)
		}
		wg.Wait()

		err = storageService.Persist()
		if err != nil {
			log.Error().Err(fmt.Errorf("error while persisting team info  \n\t%v", err)).Send()
		} else {
			log.Info().Msg("saveTeamInfo task is finished.")
		}
	},
}

func startCacheReloadJob(job *JobParameter, contract *crabada.CrabadaCaller, client crabada.Api, storageService storage.StorageService) {
	defer func() { job.wg.Done() }()
	log.Info().Int("jobId", job.id).Msg("Job started")

	for i := job.startIndex; i <= job.finishIndex; i++ {
		item, err := queryAndSaveTeamToCache(i, job.reinforceEnabled, contract, client)
		if err != nil {
			if errors.Is(err, constant.TeamIsNotSuitableErr) {
				continue
			}
			if errors.Is(err, context.DeadlineExceeded) {
				log.Error().Int("jobId", job.id).Msgf("Context Deadline Occurred, team İd [%d]", i)
				i--
				continue
			}
			if err.Error() == constant.ContractRevertErrorMessage {
				log.Error().Int("jobId", job.id).Msgf("Close job because of execution revered, team İd [%d]", i)
				return
			}
			log.Error().Int("jobId", job.id).Err(err).Msgf("Stop job because of unknown error, team İd [%d]", i)
			return
		}
		storageService.Set(strconv.Itoa(i), item)
	}
}

func queryAndSaveTeamToCache(teamId int, reinforceEnabled bool, contract *crabada.CrabadaCaller, client crabada.Api) (*storage.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.TeamQueryTimeout)
	defer cancel()

	team, err := contract.GetTeamInfo(&bind.CallOpts{Context: ctx}, big.NewInt(int64(teamId)))
	if err != nil {
		return nil, err
	}
	battlePoint := int(team.BattlePoint)
	minePoint := int(team.TimePoint)
	if int(team.BattlePoint) < constant.CacheReloadBattlePointMin {
		return nil, constant.TeamIsNotSuitableErr
	}

	reinforceCount := -1
	if reinforceEnabled {
		reinforceCount = 0
		mineHistory, err := client.GetMineHistory(team.Owner.String())
		if err != nil {
			log.Error().Err(err).Send()
		} else {
			for _, mine := range mineHistory {
				if mine.Round != 0 {
					reinforceCount++
				}
			}
		}
	}

	return &storage.Item{
		BattlePoint:    battlePoint,
		MinePoint:      minePoint,
		ReinforceCount: reinforceCount,
	}, nil
}
