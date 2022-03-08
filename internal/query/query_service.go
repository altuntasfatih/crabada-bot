package query

import (
	"context"
	"crabada-bot/internal/storage"
	"crabada-bot/pkg/crabada"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

type TeamQueryService interface {
	GetTeamInfo(teamId *big.Int, ctx context.Context) (owner string, battlePoint int, minePoint int, reinforceCount int, err error)
}

type teamQueryService struct {
	cacheEnabled   bool
	contract       crabada.CallerContract
	storageService storage.StorageService
}

func NewTeamQueryService(cacheEnabled bool, callerContract crabada.CallerContract, storageService storage.StorageService) TeamQueryService {
	return &teamQueryService{
		cacheEnabled:   cacheEnabled,
		contract:       callerContract,
		storageService: storageService,
	}
}

func (t *teamQueryService) GetTeamInfo(teamId *big.Int, ctx context.Context) (owner string, battlePoint int, minePoint int, reinforceCount int, err error) {

	if t.cacheEnabled {
		return t.getFromCache(teamId.String())
	}
	return t.getFromContract(teamId)
}

func (t *teamQueryService) getFromCache(teamId string) (string, int, int, int, error) {
	err, item := t.storageService.Get(teamId)
	if err != nil {
		return "", -1, -1, -1, fmt.Errorf("error while doing storageService.Get for key [%s]\n\t%v", teamId, err)
	}

	return "", item.BattlePoint, item.MinePoint, item.ReinforceCount, nil
}

func (t *teamQueryService) getFromContract(teamId *big.Int) (string, int, int, int, error) {
	rivalTeam, err := t.contract.GetTeamInfo(&bind.CallOpts{}, teamId)
	if err != nil {
		return "", -1, -1, -1, err
	}
	return rivalTeam.Owner.String(), int(rivalTeam.BattlePoint), int(rivalTeam.TimePoint), 0, nil
}