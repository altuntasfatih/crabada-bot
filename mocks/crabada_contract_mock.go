package mocks

import (
	"crabada-bot/pkg/crabada"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockCallerContract struct {
	mock.Mock
}
type MockFiltererContract struct {
	mock.Mock
}
type MockTransactorContract struct {
	mock.Mock
}
type MockContract struct {
	MockCallerContract
	MockFiltererContract
	MockTransactorContract
}

func (m *MockFiltererContract) WatchStartGame(opts *bind.WatchOpts, sink chan<- *crabada.CrabadaStartGame) (event.Subscription, error) {
	args := m.Called(opts, sink)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(event.Subscription), args.Error(1)
}

func (m *MockCallerContract) GetTeamInfo(opts *bind.CallOpts, teamId *big.Int) (struct {
	Owner         common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	args := m.Called(opts, teamId)
	return args.Get(0).(struct {
		Owner         common.Address
		CrabadaId1    *big.Int
		CrabadaId2    *big.Int
		CrabadaId3    *big.Int
		BattlePoint   uint16
		TimePoint     uint16
		CurrentGameId *big.Int
		LockTo        *big.Int
	}), args.Error(1)
}

func (m *MockCallerContract) GetGameBasicInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	TeamId    *big.Int
	CraReward *big.Int
	TusReward *big.Int
	StartTime uint32
	Duration  uint32
	Status    uint32
}, error) {
	args := m.Called(opts, gameId)
	return args.Get(0).(struct {
		TeamId    *big.Int
		CraReward *big.Int
		TusReward *big.Int
		StartTime uint32
		Duration  uint32
		Status    uint32
	}), args.Error(1)

}

func (m *MockTransactorContract) Attack(opts *bind.TransactOpts, gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error) {
	args := m.Called(opts, gameId, attackTeamId)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.Transaction), args.Error(1)
}
