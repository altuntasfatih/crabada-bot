package player

import (
	"context"
	"crabada-bot/config"
	"crabada-bot/pkg/crabada"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
)

var TeamDoesNotExist = errors.New("you don't have suitable team")

type Player interface {
	Attack(gameId *big.Int, ctx context.Context) (*types.Transaction, error)
	IsAttackAbe(battlePoint int, minePoint int, reinforceCount int) bool
	GetTeamId() *big.Int
	Account
}

type player struct {
	*account
	currentTeam *crabada.Team

	transactor crabada.TransactorContract

	api crabada.Api

	config *config.AttackParameter
}

func NewPlayer(transactor crabada.TransactorContract, api crabada.Api, config *config.Config) (Player, error) {
	account := NewAccount(config)
	p := player{account: account, transactor: transactor, api: api, config: config.AttackParameter}
	err := p.SelectTeam()
	return &p, err
}
func (p *player) SelectTeam() error {
	team, err := getTeamForAttack(p.api, p.Address.String())
	if err != nil {
		return err
	}
	log.Info().Msgf("selected team for loot -> %s", team.ToString())
	p.currentTeam = team
	return err
}
func (p *player) IsAttackAbe(battlePoint int, minePoint int, reinforceCount int) bool {

	return p.currentTeam.BattlePoint > battlePoint &&
		(p.currentTeam.BattlePoint-battlePoint) <= p.config.BattlePointMargin &&
		(minePoint <= p.config.MinePointLimit) &&
		p.config.ReinforceLimit >= reinforceCount
}

func (p *player) GetTeamId() *big.Int {
	return big.NewInt(int64(p.currentTeam.TeamId))
}

func (p *player) Attack(gameId *big.Int, ctx context.Context) (*types.Transaction, error) {
	log.Info().Msgf("try to attack %d", gameId.Int64())
	tx, err := p.transactor.Attack(p.GenerateTransactionOpt(ctx), gameId, p.GetTeamId())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("attack to %d is failed", gameId.Int64()))
	}
	return tx, nil
}

func getTeamForAttack(api crabada.Api, address string) (*crabada.Team, error) {
	teamsResult, err := api.GetTeamsByAddress(address)
	if err != nil {
		return nil, err
	}
	return findBestTeamToAttack(teamsResult.Teams)
}

func findBestTeamToAttack(teams []*crabada.Team) (*crabada.Team, error) {
	var suitableTeams []*crabada.Team

	for _, team := range teams {
		if team.IsFreeForAttack() {
			suitableTeams = append(suitableTeams, team)
		}
	}
	if len(suitableTeams) == 0 {
		return nil, TeamDoesNotExist
	}
	maxBattlePoint := suitableTeams[0].BattlePoint
	selectedTeam := suitableTeams[0]
	for _, team := range suitableTeams {
		if team.BattlePoint > maxBattlePoint {
			maxBattlePoint = team.BattlePoint
			selectedTeam = team
		}
	}
	return selectedTeam, nil
}