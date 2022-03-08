package mocks

import (
	"crabada-bot/pkg/crabada"
	"github.com/stretchr/testify/mock"
)

type MockCrabadaApi struct {
	mock.Mock
}

func (m *MockCrabadaApi) GetTeamsByAddress(address string) (*crabada.TeamsResult, error) {
	args := m.Called(address)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*crabada.TeamsResult), args.Error(1)
}

func (m *MockCrabadaApi) GetMineHistory(address string) ([]*crabada.MineInfo, error) {
	args := m.Called(address)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*crabada.MineInfo), args.Error(1)
}
func (m *MockCrabadaApi) GetLastOpenMine() (*crabada.MineInfo, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*crabada.MineInfo), args.Error(1)
}
