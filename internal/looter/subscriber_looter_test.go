package looter

import (
	"context"
	"crabada-bot/mocks"
	"crabada-bot/pkg/crabada"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"math/big"
	"testing"
)

func init() {
	log.Logger = zerolog.New(ioutil.Discard).With().Timestamp().Logger()
}

func TestLooterBot_StartLoot_Close_When_Subscription_Failed_(t *testing.T) {
	//given

	expectedError := errors.New("subscription failed")
	contract := new(mocks.MockFiltererContract)
	contract.On("WatchStartGame", mock.Anything, mock.Anything).Return(nil, expectedError)

	player := new(mocks.MockPlayer)
	looter := &looterBot{
		player:           player,
		filtererContract: contract}

	//when
	_, err := looter.StartLoot(context.Background())

	require.NotNil(t, err)
	require.Equal(t, err, expectedError)
}
func TestLooterBot_QueryTeam(t *testing.T) {
	//given
	teamId := big.NewInt(1)
	expectedGameId := 350

	contract := new(mocks.MockCallerContract)
	contract.On("GetTeamInfo", mock.Anything, teamId).Return(struct {
		Owner         common.Address
		CrabadaId1    *big.Int
		CrabadaId2    *big.Int
		CrabadaId3    *big.Int
		BattlePoint   uint16
		TimePoint     uint16
		CurrentGameId *big.Int
		LockTo        *big.Int
	}{
		CurrentGameId: big.NewInt(int64(expectedGameId)),
	}, nil)
	looter := &looterBot{callerContract: contract}
	lootChannel := make(chan *big.Int, 2)

	//when
	looter.queryTeam(teamId, lootChannel, context.Background())

	//then
	actualGameId := <-lootChannel
	require.Equal(t, actualGameId.Int64(), int64(expectedGameId))
}

func TestLooterBot_HandleEvent_Not_Attack(t *testing.T) {
	//given
	ctx := context.Background()

	teamId := big.NewInt(50)
	event := &crabada.CrabadaStartGame{TeamId: teamId, GameId: big.NewInt(100)}

	teamQueryService := new(mocks.MockQueryTeamService)
	teamQueryService.On("GetTeamInfo", teamId, ctx).Return("", 750, 100, 4, nil)

	player := new(mocks.MockPlayer)
	player.On("IsAttackAbe", 750, 100, 4).Return(false)

	looter := &looterBot{
		queryService: teamQueryService,
		player:       player,
	}

	//when
	looter.handle(event, nil, ctx)

	player.AssertNotCalled(t, "Attack", mock.Anything, mock.Anything)
}

func TestLooterBot_HandleEvent_Attack(t *testing.T) {
	//given
	battlePoint := 650
	minePoint := 170
	reinforceCount := 5
	ctx := context.Background()
	teamId := big.NewInt(50)
	gameId := big.NewInt(100)
	event := &crabada.CrabadaStartGame{TeamId: teamId, GameId: gameId}

	teamQueryService := new(mocks.MockQueryTeamService)
	teamQueryService.On("GetTeamInfo", teamId, ctx).Return("", battlePoint, minePoint, reinforceCount, nil)

	chainService := new(mocks.MockChainService)
	chainService.On("IsTxSuccess", mock.Anything, ctx).Return(nil)

	player := new(mocks.MockPlayer)
	player.On("IsAttackAbe", battlePoint, minePoint, reinforceCount).Return(true)
	player.On("Attack", gameId, ctx).Return(types.NewTx(&types.DynamicFeeTx{}), nil)

	looter := &looterBot{
		queryService: teamQueryService,
		chainService: chainService,
		player:       player,
	}

	lootChannel := make(chan *big.Int, 2)

	//when
	looter.handle(event, lootChannel, ctx)

	actualGameId := <-lootChannel
	require.Equal(t, actualGameId, gameId)
}
