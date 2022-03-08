package mocks

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockPlayer struct {
	mock.Mock
}

func (m *MockPlayer) Attack(gameId *big.Int, ctx context.Context) (*types.Transaction, error) {
	args := m.Called(gameId, ctx)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.Transaction), args.Error(1)
}

func (m *MockPlayer) IsAttackAbe(battlePoint int, minePoint int, reinforceCount int) bool {
	args := m.Called(battlePoint, minePoint, reinforceCount)
	return args.Get(0).(bool)
}

func (m *MockPlayer) GetTeamId() *big.Int {
	args := m.Called()
	return args.Get(0).(*big.Int)
}

func (m *MockPlayer) GenerateTransactionOpt(ctx context.Context) *bind.TransactOpts {
	args := m.Called(ctx)
	return args.Get(0).(*bind.TransactOpts)
}
