package mocks

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockChainService struct {
	mock.Mock
}

func (m *MockChainService) IsTxSuccess(hash common.Hash, ctx context.Context) error {
	args := m.Called(hash, ctx)
	return args.Error(0)
}

func (m *MockChainService) GetNonce(address common.Hash, ctx context.Context) (*big.Int, error) {
	args := m.Called(address, ctx)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*big.Int), args.Error(1)
}