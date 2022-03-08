package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockQueryTeamService struct {
	mock.Mock
}

func (m *MockQueryTeamService) GetTeamInfo(teamId *big.Int, ctx context.Context) (owner string, battlePoint int, minePoint int, reinforceCount int, err error) {
	args := m.Called(teamId, ctx)
	if args.Get(4) != nil {
		return "", -1, -1, -1, args.Error(4)
	}
	return args.Get(0).(string), args.Get(1).(int), args.Get(2).(int), args.Get(3).(int), args.Error(4)
}
