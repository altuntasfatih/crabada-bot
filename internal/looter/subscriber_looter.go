package looter

import (
	"context"
	"crabada-bot/internal/chain"
	"crabada-bot/internal/player"
	"crabada-bot/internal/query"
	"crabada-bot/pkg/crabada"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

var InterruptSignalReceived = errors.New("interrupt signal is received, close app")

type Looter interface {
	StartLoot(ctx context.Context) (gameId *int, err error)
}

type looterBot struct {
	player player.Player

	queryService     query.TeamQueryService
	filtererContract crabada.FiltererContract
	callerContract   crabada.CallerContract

	chainService chain.Service
	subscription event.Subscription
}

func NewLooter(player player.Player, callerContract crabada.CallerContract, filtererContract crabada.FiltererContract, chainService chain.Service, teamService query.TeamQueryService) Looter {
	return &looterBot{
		player:           player,
		callerContract:   callerContract,
		filtererContract: filtererContract,
		chainService:     chainService,
		queryService:     teamService,
	}
}

func (b *looterBot) StartLoot(ctx context.Context) (gameId *int, err error) {
	if channel, err := b.subscribe(); err == nil {
		return b.loot(ctx, channel)
	} else {
		return nil, err
	}
}

func (b *looterBot) loot(ctx context.Context, channel chan *crabada.CrabadaStartGame) (*int, error) {
	lootDone := make(chan *big.Int, 2)
	go b.queryTeam(b.player.GetTeamId(), lootDone, ctx)

	defer func() {
		close(lootDone)
		b.unsubscribe()
	}()

	for {
		select {
		case e := <-channel:
			go b.handle(e, lootDone, ctx)
		case err := <-b.subscription.Err():
			return nil, errors.Wrap(err, "subscription failed")

		case <-ctx.Done():
			return nil, InterruptSignalReceived

		case gameId := <-lootDone:
			log.Info().Msgf("attack to game: %d successful", gameId.Int64())
			i := int(gameId.Int64())
			return &i, nil
		}
	}
}

func (b *looterBot) handle(event *crabada.CrabadaStartGame, lootDoneChannel chan *big.Int, ctx context.Context) {
	teamId := event.TeamId
	_, battlePoint, minePoint, reinforceCount, err := b.queryService.GetTeamInfo(teamId, ctx)
	if err != nil {
		return
	}
	isAttackAble := b.player.IsAttackAbe(battlePoint, minePoint, reinforceCount)
	log.Info().Msgf("IsAttackAble: %t, GameId: %d,TeamId: %d,BattlePoint: %d, MinerPoint: %d, ReinforceCount: %d",
		isAttackAble, event.GameId.Int64(),
		teamId, battlePoint, minePoint, reinforceCount)
	if isAttackAble {
		if gameId, err := b.attack(event.GameId, ctx); err == nil {
			sendGameIdToChannel(gameId, lootDoneChannel, ctx)
		} else {
			log.Error().Err(err).Send()
		}
	}
}
func (b *looterBot) attack(gameId *big.Int, ctx context.Context) (*big.Int, error) {
	tx, err := b.player.Attack(gameId, ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("attack to %d is failed", gameId.Int64()))
	}
	if err := b.chainService.IsTxSuccess(tx.Hash(), ctx); err == nil {
		return gameId, nil
	} else {
		return nil, errors.Wrap(err, fmt.Sprintf("attack to %d is failed", gameId.Int64()))
	}
}
func (b *looterBot) subscribe() (chan *crabada.CrabadaStartGame, error) {
	channel := make(chan *crabada.CrabadaStartGame)
	subscription, err := b.filtererContract.WatchStartGame(&bind.WatchOpts{}, channel)
	if err != nil {
		return nil, err
	}
	b.subscription = subscription
	return channel, nil
}
func (b *looterBot) unsubscribe() {
	if b.subscription != nil {
		b.subscription.Unsubscribe()
	}
}
func (b *looterBot) queryTeam(teamId *big.Int, lootDone chan *big.Int, ctx context.Context) {

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
					log.Warn().Msgf("your team is already in a game, close script")
					sendGameIdToChannel(team.CurrentGameId, lootDone, ctx)
					return
				}
			}
		}
	}()
}

func sendGameIdToChannel(gameId *big.Int, lootDone chan *big.Int, ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Info().Msg("context already cancelled")
	default:
		lootDone <- gameId
		return
	}
}
