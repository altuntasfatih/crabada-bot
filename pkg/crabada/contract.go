package crabada

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
)

type Contract interface {
	CallerContract
	TransactorContract
	FiltererContract
}

type CallerContract interface {
	GetTeamInfo(opts *bind.CallOpts, teamId *big.Int) (struct {
		Owner         common.Address
		CrabadaId1    *big.Int
		CrabadaId2    *big.Int
		CrabadaId3    *big.Int
		BattlePoint   uint16
		TimePoint     uint16
		CurrentGameId *big.Int
		LockTo        *big.Int
	}, error)

	GetGameBasicInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
		TeamId    *big.Int
		CraReward *big.Int
		TusReward *big.Int
		StartTime uint32
		Duration  uint32
		Status    uint32
	}, error)
}
type TransactorContract interface {
	Attack(opts *bind.TransactOpts, gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error)
}
type FiltererContract interface {
	WatchStartGame(*bind.WatchOpts, chan<- *CrabadaStartGame) (event.Subscription, error)
}
