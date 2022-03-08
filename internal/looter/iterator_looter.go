package looter

import (
	"context"
	"crabada-bot/internal/player"
	"crabada-bot/internal/query"
	"crabada-bot/pkg/constant"
	"crabada-bot/pkg/crabada"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"
	"math/big"
	"sync"
	"time"
)

type IteratorBot struct {
	player player.Player

	queryService   query.TeamQueryService
	callerContract crabada.CallerContract

	apiClient crabada.Api
}

func NewLootIterator(player player.Player, callerContract crabada.CallerContract, apiClient crabada.Api, teamService query.TeamQueryService) *IteratorBot {
	return &IteratorBot{
		player:         player,
		callerContract: callerContract,
		apiClient:      apiClient,
		queryService:   teamService,
	}
}

func (b *IteratorBot) StartLootIterator(ctx context.Context, cancelFunc context.CancelFunc) (err error) {
	go b.queryTeam(ctx, cancelFunc, b.player.GetTeamId())

	lastMineInfo, err := b.apiClient.GetLastOpenMine()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for i := 0; i < constant.DefaultIteratorCount; i++ {
		wg.Add(1)
		go b.loot(ctx, &wg, int64(lastMineInfo.GameId+i), constant.DefaultIteratorCount)
	}

	wg.Wait()
	return nil
}

func (b *IteratorBot) loot(ctx context.Context, wg *sync.WaitGroup, startGameId, increaseCount int64) {
	defer func() { wg.Done() }()
	log.Error().Msgf("loot-iterate was started,startGameId %d and increaseCount %d ", startGameId, increaseCount)
	gameId := startGameId
	for {
		select {
		case <-ctx.Done():
			log.Error().Msgf("context is cancelled for %d", gameId)
			return
		default:
			gameInfo, err := b.callerContract.GetGameBasicInfo(&bind.CallOpts{}, big.NewInt(gameId))
			if err != nil {
				if err.Error() == constant.ContractRevertErrorMessage {
				}
				continue
			}
			if _, battlePoint, minePoint, reinforceCount, err := b.queryService.GetTeamInfo(gameInfo.TeamId, ctx); err == nil {
				if b.player.IsAttackAbe(battlePoint, minePoint, reinforceCount) {
					_, err := b.player.Attack(big.NewInt(gameId), ctx)
					if err != nil {
						log.Error().Err(err).Send()
					}
				}
			} else {
				log.Error().Err(err).Send()
			}

		}
		gameId = gameId + increaseCount
		fmt.Printf("Iterate to gameId %d\n", gameId)
	}
}

func (b *IteratorBot) queryTeam(ctx context.Context, cancelFunc context.CancelFunc, teamId *big.Int) {

	ticker := time.NewTicker(300 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				team, err := b.callerContract.GetTeamInfo(&bind.CallOpts{}, teamId)
				if err != nil {
					continue
				}
				if team.CurrentGameId != nil && team.CurrentGameId.Int64() != 0 {
					log.Warn().Msgf("your team is already in the game, close script")
					cancelFunc()
					return
				}
			}
		}
	}()
}