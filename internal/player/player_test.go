package player

import (
	"crabada-bot/config"
	"crabada-bot/mocks"
	"crabada-bot/pkg/crabada"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func init() {
	log.Logger = zerolog.New(ioutil.Discard).With().Timestamp().Logger()
}

func TestNewTeam_When_You_Dont_Have_Suit_Team(t *testing.T) {
	//given
	walletAddress := common.HexToAddress("0xFatih")
	teamResult := &crabada.TeamsResult{
		TeamSize: 1,
		Teams:    []*crabada.Team{{TeamId: 1, CrabadaIdOne: 101, CrabadaIdTwo: 102, CrabadaIdThree: 0, BattlePoint: 160, MinePoint: 175, Status: "AVAILABLE", GameRound: 0, GameType: "", ProcessStatus: ""}},
	}
	mockApi := new(mocks.MockCrabadaApi)
	mockApi.On("GetTeamsByAddress", walletAddress.String()).Return(teamResult, nil)

	player := player{&account{Address: walletAddress}, nil, nil, mockApi, &config.AttackParameter{}}

	//when
	err := player.SelectTeam()
	require.NotNil(t, err)
	require.Equal(t, err, TeamDoesNotExist)
}

func TestNewTeam(t *testing.T) {
	//given
	address := common.HexToAddress("0xFatih")
	teamResult := &crabada.TeamsResult{
		TeamSize: 1,
		Teams:    []*crabada.Team{{TeamId: 1, CrabadaIdOne: 101, CrabadaIdTwo: 102, CrabadaIdThree: 103, BattlePoint: 650, MinePoint: 175, Status: "AVAILABLE", GameRound: 0, GameType: "", ProcessStatus: ""}},
	}

	mockApi := new(mocks.MockCrabadaApi)
	mockApi.On("GetTeamsByAddress", address.String()).Return(teamResult, nil)

	player := player{&account{Address: address}, nil, nil, mockApi, nil}
	err := player.SelectTeam()
	require.Nil(t, err)
	require.Equal(t, player.currentTeam, teamResult.Teams[0])
}

func TestTeam_IsAttackAbe(t *testing.T) {
	//given
	selectedTeam := &crabada.Team{BattlePoint: 650, MinePoint: 175, Status: "AVAILABLE"}
	c := &config.AttackParameter{BattlePointMargin: 5, MinePointLimit: 220, ReinforceLimit: 5}

	team := &player{
		currentTeam: selectedTeam,
		config:      c,
	}

	tables := []struct {
		BattlePoint    int
		MinePoint      int
		ReinforceCount int
		ExpectedResult bool
	}{

		{649, 150, 1, true},
		{646, 150, 1, true},
		{645, 175, 1, true},
		{645, 175, 5, true},
		{645, 220, 2, true},

		{646, 221, 1, false},
		{646, 176, 6, false},
		{646, 221, 6, false},
		{650, 100, 1, false},
		{644, 100, 1, false},
	}

	//when
	for _, table := range tables {
		require.Equal(t, team.IsAttackAbe(table.BattlePoint, table.MinePoint, table.ReinforceCount), table.ExpectedResult, table)
	}
}
