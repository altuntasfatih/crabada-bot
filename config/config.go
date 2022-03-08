package config

import "fmt"

type Config struct {
	CacheEnabled bool
	PrivateKey   string
	GasTip       int
	GasFeeCap    int
	*AttackParameter
}

type AttackParameter struct {
	BattlePointMargin int
	MinePointLimit    int
	ReinforceLimit    int
}

func (c *Config) ToString() string {
	return fmt.Sprintf("config{CacheEnabled: %v,PrivateKey: %s, GasTip: %d, GasFeeCap: %d, BattlePointMargin: %d, MinePointLimit: %d, ReinforceLimit: %d}", c.CacheEnabled, c.PrivateKey, c.GasTip, c.GasFeeCap, c.BattlePointMargin, c.MinePointLimit, c.ReinforceLimit)
}