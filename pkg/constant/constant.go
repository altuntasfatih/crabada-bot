package constant

import (
	"github.com/pkg/errors"
	"math/big"
	"math/rand"
	"time"
)

const (
	TimeLogField       = "timestamp"
	TimeLogFieldFormat = "2006-01-02T15:04:05.999"
)
const (
	ContractRevertErrorMessage = "execution reverted"

	CrabadaApiUrl   = "https://idle-api.crabada.com"
	AvaxHttpAddress = "https://api.avax.network/ext/bc/C/rpc"
	AvaxWsAdress    = "wss://api.avax.network/ext/bc/C/ws"

	CustomNodeHttpAddress = "http://localhost:9650/ext/bc/C/rpc"
	CustomNodeWsAddress   = "ws://localhost:9650/ext/bc/C/ws"

	ChainWsAddressOne      = "wss://speedy-nodes-nyc.moralis.io/68d097e5da42e67f3c55a7d1/avalanche/mainnet/ws"
	ChainWsAddressTwo      = "wss://speedy-nodes-nyc.moralis.io/5d568b289c49a99967ab9e57/avalanche/mainnet/ws"
	ChainWsAddressThree    = "wss://speedy-nodes-nyc.moralis.io/7e5a0d5ddbd3c03fe188e6b7/avalanche/mainnet/ws"
	ChainWsAddressFour     = "wss://speedy-nodes-nyc.moralis.io/273a047a2c700491e7b2e946/avalanche/mainnet/ws"
	ChainWsAddressFive     = "wss://speedy-nodes-nyc.moralis.io/3cc3eab1e18c3a64dbb30217/avalanche/mainnet/ws"
	ChainWsAddressSix      = "wss://speedy-nodes-nyc.moralis.io/c461c6347568457a09e86b4a/avalanche/mainnet/ws"
	ChainWsAddressSeven    = "wss://speedy-nodes-nyc.moralis.io/5654fe7ac948acfec749207d/avalanche/mainnet/ws"
	ChainWsAddressEight    = "wss://speedy-nodes-nyc.moralis.io/3175b9855b43b60c401215d3/avalanche/mainnet/ws"
	CrabadaContractAddress = "0x82a85407BD612f52577909F4A58bfC6873f14DA8"
)
const ChainId = 43114
const GasLimit = uint64(300000)
const LootTimeout = 3 * time.Minute
const TeamQueryTimeout = 1000 * time.Millisecond
const DefaultIteratorCount = 3

var TeamIsNotSuitableErr = errors.New("team is not suitable")

const (
	CacheReloadEndTeamId      int = 20000
	CacheReloadBattlePointMin int = 500
)

var ZeroValue = big.NewInt(0)

func GetWsChainAddress() string {
	//return CustomNodeWsAddress
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(9)
	switch number {
	case 1:
		return ChainWsAddressOne
	case 2:
		return ChainWsAddressTwo
	case 3:
		return ChainWsAddressThree
	case 4:
		return ChainWsAddressFour
	case 5:
		return ChainWsAddressFive
	case 6:
		return ChainWsAddressSix
	case 7:
		return ChainWsAddressSeven
	case 8:
		return ChainWsAddressEight
	default:
		return ChainWsAddressEight

	}

}
