// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crabada

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CrabadaMetaData contains all meta data concerning the Crabada contract.
var CrabadaMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"battlePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timePoint\",\"type\":\"uint256\"}],\"name\":\"AddCrabada\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"CloseGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"battlePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timePoint\",\"type\":\"uint256\"}],\"name\":\"CreateTeam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTeamAmountLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"ExpandTeam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"turn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defenseTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"soldierId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attackTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"attackPoint\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"defensePoint\",\"type\":\"uint16\"}],\"name\":\"Fight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"battlePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timePoint\",\"type\":\"uint256\"}],\"name\":\"RemoveCrabada\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SetLendingPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerCraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerTusReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loserCraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loserTusReward\",\"type\":\"uint256\"}],\"name\":\"SettleGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"craReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tusReward\",\"type\":\"uint256\"}],\"name\":\"StartGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"battlePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"statVersion\",\"type\":\"uint256\"}],\"name\":\"UpdateStatTeam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"STEP_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amuletContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attackCooldownDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attackDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseClosenessFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCraReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseLuckyWinRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseModifier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTusReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusCraClass\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bonusRewardClasses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusTimePointDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusTimePointThreshold\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusTusClass\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"classBonusPercent\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craIncentiveEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crabaInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"lockTo\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"status\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"expandTeamCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expandTeamPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingFeeHolerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingFeePercent\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingLockDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lendingPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lootingCraBaseReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lootingPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lootingTusBaseReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerRevengeEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modifierSetting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinforceLockDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rngContract\",\"outputs\":[{\"internalType\":\"contractIRNG\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"teamCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrabada\",\"name\":\"_crabada\",\"type\":\"address\"},{\"internalType\":\"contractIToken\",\"name\":\"_cra\",\"type\":\"address\"},{\"internalType\":\"contractIToken\",\"name\":\"_tus\",\"type\":\"address\"},{\"internalType\":\"contractIGameStat\",\"name\":\"_gameStat\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"crabadaIds\",\"type\":\"uint256[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"crabadaIds\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setLendingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"}],\"name\":\"createTeam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"}],\"name\":\"addCrabadaToTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"removeCrabadaFromTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"}],\"name\":\"startGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"}],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowPrice\",\"type\":\"uint256\"}],\"name\":\"reinforceAttack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowPrice\",\"type\":\"uint256\"}],\"name\":\"reinforceDefense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"settleGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"closeGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"}],\"name\":\"getStats\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"battlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timePoint\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"}],\"name\":\"getTeamInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"battlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"currentGameId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"lockTo\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getGameBasicInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"craReward\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tusReward\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"status\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getGameBattleInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"attackTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastAttackTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastDefTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"attackId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attackId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defId2\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getLootingStatsInfo\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"attackPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defensePoint\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getBattleTimePoint\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"attackPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defensePoint\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"craRewardAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tusRewardAmount\",\"type\":\"uint128\"}],\"name\":\"setRewardMining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"craRewardAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tusRewardAmount\",\"type\":\"uint128\"}],\"name\":\"setRewardLooting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setLootingPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setStepDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setMiningDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setReinforceLockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setAttackCooldownDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setLendingFeeHolerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"cooldown\",\"type\":\"uint128\"}],\"name\":\"setLendingCooldownDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"setLendingFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"setCRAIncentiveStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"classId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBonused\",\"type\":\"bool\"}],\"name\":\"setBonusClass\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"percent\",\"type\":\"uint128\"}],\"name\":\"setBonusClassPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newThreshold\",\"type\":\"uint128\"}],\"name\":\"setBonusTimePointThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setBonusTimePointDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setAttackDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIToken\",\"name\":\"_tusToken\",\"type\":\"address\"},{\"internalType\":\"contractIToken\",\"name\":\"_craToken\",\"type\":\"address\"},{\"internalType\":\"contractICrabada\",\"name\":\"_crabada\",\"type\":\"address\"},{\"internalType\":\"contractIGameStat\",\"name\":\"_gameStatContract\",\"type\":\"address\"}],\"name\":\"changeContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"craAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crabadaAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"statAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"looterBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minerBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minerMP\",\"type\":\"uint256\"}],\"name\":\"calculateLuckyWinRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseLuckyWinRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseModifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_modifierSetting\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseClosenessFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRNG\",\"name\":\"_rngContract\",\"type\":\"address\"}],\"name\":\"setMinerRevengeSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"enableMinerRevenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"expandTeamLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amuletAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setupTeamExpandSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrabadaABI is the input ABI used to generate the binding from.
// Deprecated: Use CrabadaMetaData.ABI instead.
var CrabadaABI = CrabadaMetaData.ABI

// Crabada is an auto generated Go binding around an Ethereum contract.
type Crabada struct {
	CrabadaCaller     // Read-only binding to the contract
	CrabadaTransactor // Write-only binding to the contract
	CrabadaFilterer   // Log filterer for contract events
}

// CrabadaCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrabadaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabadaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrabadaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabadaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrabadaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabadaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrabadaSession struct {
	Contract     *Crabada          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrabadaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrabadaCallerSession struct {
	Contract *CrabadaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CrabadaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrabadaTransactorSession struct {
	Contract     *CrabadaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CrabadaRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrabadaRaw struct {
	Contract *Crabada // Generic contract binding to access the raw methods on
}

// CrabadaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrabadaCallerRaw struct {
	Contract *CrabadaCaller // Generic read-only contract binding to access the raw methods on
}

// CrabadaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrabadaTransactorRaw struct {
	Contract *CrabadaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrabada creates a new instance of Crabada, bound to a specific deployed contract.
func NewCrabada(address common.Address, backend bind.ContractBackend) (*Crabada, error) {
	contract, err := bindCrabada(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crabada{CrabadaCaller: CrabadaCaller{contract: contract}, CrabadaTransactor: CrabadaTransactor{contract: contract}, CrabadaFilterer: CrabadaFilterer{contract: contract}}, nil
}

// NewCrabadaCaller creates a new read-only instance of Crabada, bound to a specific deployed contract.
func NewCrabadaCaller(address common.Address, caller bind.ContractCaller) (*CrabadaCaller, error) {
	contract, err := bindCrabada(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrabadaCaller{contract: contract}, nil
}

// NewCrabadaTransactor creates a new write-only instance of Crabada, bound to a specific deployed contract.
func NewCrabadaTransactor(address common.Address, transactor bind.ContractTransactor) (*CrabadaTransactor, error) {
	contract, err := bindCrabada(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrabadaTransactor{contract: contract}, nil
}

// NewCrabadaFilterer creates a new log filterer instance of Crabada, bound to a specific deployed contract.
func NewCrabadaFilterer(address common.Address, filterer bind.ContractFilterer) (*CrabadaFilterer, error) {
	contract, err := bindCrabada(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrabadaFilterer{contract: contract}, nil
}

// bindCrabada binds a generic wrapper to an already deployed contract.
func bindCrabada(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrabadaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crabada *CrabadaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crabada.Contract.CrabadaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crabada *CrabadaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crabada.Contract.CrabadaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crabada *CrabadaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crabada.Contract.CrabadaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crabada *CrabadaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crabada.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crabada *CrabadaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crabada.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crabada *CrabadaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crabada.Contract.contract.Transact(opts, method, params...)
}

// STEPDURATION is a free data retrieval call binding the contract method 0xf67718cb.
//
// Solidity: function STEP_DURATION() view returns(uint128)
func (_Crabada *CrabadaCaller) STEPDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "STEP_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STEPDURATION is a free data retrieval call binding the contract method 0xf67718cb.
//
// Solidity: function STEP_DURATION() view returns(uint128)
func (_Crabada *CrabadaSession) STEPDURATION() (*big.Int, error) {
	return _Crabada.Contract.STEPDURATION(&_Crabada.CallOpts)
}

// STEPDURATION is a free data retrieval call binding the contract method 0xf67718cb.
//
// Solidity: function STEP_DURATION() view returns(uint128)
func (_Crabada *CrabadaCallerSession) STEPDURATION() (*big.Int, error) {
	return _Crabada.Contract.STEPDURATION(&_Crabada.CallOpts)
}

// AmuletContract is a free data retrieval call binding the contract method 0xd34fa7e2.
//
// Solidity: function amuletContract() view returns(address)
func (_Crabada *CrabadaCaller) AmuletContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "amuletContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmuletContract is a free data retrieval call binding the contract method 0xd34fa7e2.
//
// Solidity: function amuletContract() view returns(address)
func (_Crabada *CrabadaSession) AmuletContract() (common.Address, error) {
	return _Crabada.Contract.AmuletContract(&_Crabada.CallOpts)
}

// AmuletContract is a free data retrieval call binding the contract method 0xd34fa7e2.
//
// Solidity: function amuletContract() view returns(address)
func (_Crabada *CrabadaCallerSession) AmuletContract() (common.Address, error) {
	return _Crabada.Contract.AmuletContract(&_Crabada.CallOpts)
}

// AttackCooldownDuration is a free data retrieval call binding the contract method 0x6c1fcc04.
//
// Solidity: function attackCooldownDuration() view returns(uint128)
func (_Crabada *CrabadaCaller) AttackCooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "attackCooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttackCooldownDuration is a free data retrieval call binding the contract method 0x6c1fcc04.
//
// Solidity: function attackCooldownDuration() view returns(uint128)
func (_Crabada *CrabadaSession) AttackCooldownDuration() (*big.Int, error) {
	return _Crabada.Contract.AttackCooldownDuration(&_Crabada.CallOpts)
}

// AttackCooldownDuration is a free data retrieval call binding the contract method 0x6c1fcc04.
//
// Solidity: function attackCooldownDuration() view returns(uint128)
func (_Crabada *CrabadaCallerSession) AttackCooldownDuration() (*big.Int, error) {
	return _Crabada.Contract.AttackCooldownDuration(&_Crabada.CallOpts)
}

// AttackDuration is a free data retrieval call binding the contract method 0xfd90eb2e.
//
// Solidity: function attackDuration() view returns(uint128)
func (_Crabada *CrabadaCaller) AttackDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "attackDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttackDuration is a free data retrieval call binding the contract method 0xfd90eb2e.
//
// Solidity: function attackDuration() view returns(uint128)
func (_Crabada *CrabadaSession) AttackDuration() (*big.Int, error) {
	return _Crabada.Contract.AttackDuration(&_Crabada.CallOpts)
}

// AttackDuration is a free data retrieval call binding the contract method 0xfd90eb2e.
//
// Solidity: function attackDuration() view returns(uint128)
func (_Crabada *CrabadaCallerSession) AttackDuration() (*big.Int, error) {
	return _Crabada.Contract.AttackDuration(&_Crabada.CallOpts)
}

// BaseClosenessFactor is a free data retrieval call binding the contract method 0x1de24004.
//
// Solidity: function baseClosenessFactor() view returns(uint256)
func (_Crabada *CrabadaCaller) BaseClosenessFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "baseClosenessFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseClosenessFactor is a free data retrieval call binding the contract method 0x1de24004.
//
// Solidity: function baseClosenessFactor() view returns(uint256)
func (_Crabada *CrabadaSession) BaseClosenessFactor() (*big.Int, error) {
	return _Crabada.Contract.BaseClosenessFactor(&_Crabada.CallOpts)
}

// BaseClosenessFactor is a free data retrieval call binding the contract method 0x1de24004.
//
// Solidity: function baseClosenessFactor() view returns(uint256)
func (_Crabada *CrabadaCallerSession) BaseClosenessFactor() (*big.Int, error) {
	return _Crabada.Contract.BaseClosenessFactor(&_Crabada.CallOpts)
}

// BaseCraReward is a free data retrieval call binding the contract method 0x7d2694e4.
//
// Solidity: function baseCraReward() view returns(uint128)
func (_Crabada *CrabadaCaller) BaseCraReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "baseCraReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCraReward is a free data retrieval call binding the contract method 0x7d2694e4.
//
// Solidity: function baseCraReward() view returns(uint128)
func (_Crabada *CrabadaSession) BaseCraReward() (*big.Int, error) {
	return _Crabada.Contract.BaseCraReward(&_Crabada.CallOpts)
}

// BaseCraReward is a free data retrieval call binding the contract method 0x7d2694e4.
//
// Solidity: function baseCraReward() view returns(uint128)
func (_Crabada *CrabadaCallerSession) BaseCraReward() (*big.Int, error) {
	return _Crabada.Contract.BaseCraReward(&_Crabada.CallOpts)
}

// BaseLuckyWinRate is a free data retrieval call binding the contract method 0xf6bebc7b.
//
// Solidity: function baseLuckyWinRate() view returns(uint256)
func (_Crabada *CrabadaCaller) BaseLuckyWinRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "baseLuckyWinRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLuckyWinRate is a free data retrieval call binding the contract method 0xf6bebc7b.
//
// Solidity: function baseLuckyWinRate() view returns(uint256)
func (_Crabada *CrabadaSession) BaseLuckyWinRate() (*big.Int, error) {
	return _Crabada.Contract.BaseLuckyWinRate(&_Crabada.CallOpts)
}

// BaseLuckyWinRate is a free data retrieval call binding the contract method 0xf6bebc7b.
//
// Solidity: function baseLuckyWinRate() view returns(uint256)
func (_Crabada *CrabadaCallerSession) BaseLuckyWinRate() (*big.Int, error) {
	return _Crabada.Contract.BaseLuckyWinRate(&_Crabada.CallOpts)
}

// BaseModifier is a free data retrieval call binding the contract method 0x2eaa9f24.
//
// Solidity: function baseModifier() view returns(uint256)
func (_Crabada *CrabadaCaller) BaseModifier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "baseModifier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseModifier is a free data retrieval call binding the contract method 0x2eaa9f24.
//
// Solidity: function baseModifier() view returns(uint256)
func (_Crabada *CrabadaSession) BaseModifier() (*big.Int, error) {
	return _Crabada.Contract.BaseModifier(&_Crabada.CallOpts)
}

// BaseModifier is a free data retrieval call binding the contract method 0x2eaa9f24.
//
// Solidity: function baseModifier() view returns(uint256)
func (_Crabada *CrabadaCallerSession) BaseModifier() (*big.Int, error) {
	return _Crabada.Contract.BaseModifier(&_Crabada.CallOpts)
}

// BaseTusReward is a free data retrieval call binding the contract method 0x580c484e.
//
// Solidity: function baseTusReward() view returns(uint128)
func (_Crabada *CrabadaCaller) BaseTusReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "baseTusReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTusReward is a free data retrieval call binding the contract method 0x580c484e.
//
// Solidity: function baseTusReward() view returns(uint128)
func (_Crabada *CrabadaSession) BaseTusReward() (*big.Int, error) {
	return _Crabada.Contract.BaseTusReward(&_Crabada.CallOpts)
}

// BaseTusReward is a free data retrieval call binding the contract method 0x580c484e.
//
// Solidity: function baseTusReward() view returns(uint128)
func (_Crabada *CrabadaCallerSession) BaseTusReward() (*big.Int, error) {
	return _Crabada.Contract.BaseTusReward(&_Crabada.CallOpts)
}

// BonusCraClass is a free data retrieval call binding the contract method 0xd0f0cd1f.
//
// Solidity: function bonusCraClass() view returns(uint8)
func (_Crabada *CrabadaCaller) BonusCraClass(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "bonusCraClass")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BonusCraClass is a free data retrieval call binding the contract method 0xd0f0cd1f.
//
// Solidity: function bonusCraClass() view returns(uint8)
func (_Crabada *CrabadaSession) BonusCraClass() (uint8, error) {
	return _Crabada.Contract.BonusCraClass(&_Crabada.CallOpts)
}

// BonusCraClass is a free data retrieval call binding the contract method 0xd0f0cd1f.
//
// Solidity: function bonusCraClass() view returns(uint8)
func (_Crabada *CrabadaCallerSession) BonusCraClass() (uint8, error) {
	return _Crabada.Contract.BonusCraClass(&_Crabada.CallOpts)
}

// BonusRewardClasses is a free data retrieval call binding the contract method 0xb0cb529a.
//
// Solidity: function bonusRewardClasses(uint256 ) view returns(bool)
func (_Crabada *CrabadaCaller) BonusRewardClasses(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "bonusRewardClasses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BonusRewardClasses is a free data retrieval call binding the contract method 0xb0cb529a.
//
// Solidity: function bonusRewardClasses(uint256 ) view returns(bool)
func (_Crabada *CrabadaSession) BonusRewardClasses(arg0 *big.Int) (bool, error) {
	return _Crabada.Contract.BonusRewardClasses(&_Crabada.CallOpts, arg0)
}

// BonusRewardClasses is a free data retrieval call binding the contract method 0xb0cb529a.
//
// Solidity: function bonusRewardClasses(uint256 ) view returns(bool)
func (_Crabada *CrabadaCallerSession) BonusRewardClasses(arg0 *big.Int) (bool, error) {
	return _Crabada.Contract.BonusRewardClasses(&_Crabada.CallOpts, arg0)
}

// BonusTimePointDuration is a free data retrieval call binding the contract method 0x5f9e2083.
//
// Solidity: function bonusTimePointDuration() view returns(uint128)
func (_Crabada *CrabadaCaller) BonusTimePointDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "bonusTimePointDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusTimePointDuration is a free data retrieval call binding the contract method 0x5f9e2083.
//
// Solidity: function bonusTimePointDuration() view returns(uint128)
func (_Crabada *CrabadaSession) BonusTimePointDuration() (*big.Int, error) {
	return _Crabada.Contract.BonusTimePointDuration(&_Crabada.CallOpts)
}

// BonusTimePointDuration is a free data retrieval call binding the contract method 0x5f9e2083.
//
// Solidity: function bonusTimePointDuration() view returns(uint128)
func (_Crabada *CrabadaCallerSession) BonusTimePointDuration() (*big.Int, error) {
	return _Crabada.Contract.BonusTimePointDuration(&_Crabada.CallOpts)
}

// BonusTimePointThreshold is a free data retrieval call binding the contract method 0x6186a78b.
//
// Solidity: function bonusTimePointThreshold() view returns(uint128)
func (_Crabada *CrabadaCaller) BonusTimePointThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "bonusTimePointThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusTimePointThreshold is a free data retrieval call binding the contract method 0x6186a78b.
//
// Solidity: function bonusTimePointThreshold() view returns(uint128)
func (_Crabada *CrabadaSession) BonusTimePointThreshold() (*big.Int, error) {
	return _Crabada.Contract.BonusTimePointThreshold(&_Crabada.CallOpts)
}

// BonusTimePointThreshold is a free data retrieval call binding the contract method 0x6186a78b.
//
// Solidity: function bonusTimePointThreshold() view returns(uint128)
func (_Crabada *CrabadaCallerSession) BonusTimePointThreshold() (*big.Int, error) {
	return _Crabada.Contract.BonusTimePointThreshold(&_Crabada.CallOpts)
}

// BonusTusClass is a free data retrieval call binding the contract method 0x4baa6a6b.
//
// Solidity: function bonusTusClass() view returns(uint8)
func (_Crabada *CrabadaCaller) BonusTusClass(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "bonusTusClass")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BonusTusClass is a free data retrieval call binding the contract method 0x4baa6a6b.
//
// Solidity: function bonusTusClass() view returns(uint8)
func (_Crabada *CrabadaSession) BonusTusClass() (uint8, error) {
	return _Crabada.Contract.BonusTusClass(&_Crabada.CallOpts)
}

// BonusTusClass is a free data retrieval call binding the contract method 0x4baa6a6b.
//
// Solidity: function bonusTusClass() view returns(uint8)
func (_Crabada *CrabadaCallerSession) BonusTusClass() (uint8, error) {
	return _Crabada.Contract.BonusTusClass(&_Crabada.CallOpts)
}

// CalculateLuckyWinRate is a free data retrieval call binding the contract method 0x0fe2b4d5.
//
// Solidity: function calculateLuckyWinRate(uint256 looterBP, uint256 minerBP, uint256 minerMP) view returns(uint256 rate)
func (_Crabada *CrabadaCaller) CalculateLuckyWinRate(opts *bind.CallOpts, looterBP *big.Int, minerBP *big.Int, minerMP *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "calculateLuckyWinRate", looterBP, minerBP, minerMP)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateLuckyWinRate is a free data retrieval call binding the contract method 0x0fe2b4d5.
//
// Solidity: function calculateLuckyWinRate(uint256 looterBP, uint256 minerBP, uint256 minerMP) view returns(uint256 rate)
func (_Crabada *CrabadaSession) CalculateLuckyWinRate(looterBP *big.Int, minerBP *big.Int, minerMP *big.Int) (*big.Int, error) {
	return _Crabada.Contract.CalculateLuckyWinRate(&_Crabada.CallOpts, looterBP, minerBP, minerMP)
}

// CalculateLuckyWinRate is a free data retrieval call binding the contract method 0x0fe2b4d5.
//
// Solidity: function calculateLuckyWinRate(uint256 looterBP, uint256 minerBP, uint256 minerMP) view returns(uint256 rate)
func (_Crabada *CrabadaCallerSession) CalculateLuckyWinRate(looterBP *big.Int, minerBP *big.Int, minerMP *big.Int) (*big.Int, error) {
	return _Crabada.Contract.CalculateLuckyWinRate(&_Crabada.CallOpts, looterBP, minerBP, minerMP)
}

// ClassBonusPercent is a free data retrieval call binding the contract method 0x9de03ede.
//
// Solidity: function classBonusPercent() view returns(uint128)
func (_Crabada *CrabadaCaller) ClassBonusPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "classBonusPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClassBonusPercent is a free data retrieval call binding the contract method 0x9de03ede.
//
// Solidity: function classBonusPercent() view returns(uint128)
func (_Crabada *CrabadaSession) ClassBonusPercent() (*big.Int, error) {
	return _Crabada.Contract.ClassBonusPercent(&_Crabada.CallOpts)
}

// ClassBonusPercent is a free data retrieval call binding the contract method 0x9de03ede.
//
// Solidity: function classBonusPercent() view returns(uint128)
func (_Crabada *CrabadaCallerSession) ClassBonusPercent() (*big.Int, error) {
	return _Crabada.Contract.ClassBonusPercent(&_Crabada.CallOpts)
}

// CraIncentiveEnable is a free data retrieval call binding the contract method 0x1ef17d36.
//
// Solidity: function craIncentiveEnable() view returns(bool)
func (_Crabada *CrabadaCaller) CraIncentiveEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "craIncentiveEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CraIncentiveEnable is a free data retrieval call binding the contract method 0x1ef17d36.
//
// Solidity: function craIncentiveEnable() view returns(bool)
func (_Crabada *CrabadaSession) CraIncentiveEnable() (bool, error) {
	return _Crabada.Contract.CraIncentiveEnable(&_Crabada.CallOpts)
}

// CraIncentiveEnable is a free data retrieval call binding the contract method 0x1ef17d36.
//
// Solidity: function craIncentiveEnable() view returns(bool)
func (_Crabada *CrabadaCallerSession) CraIncentiveEnable() (bool, error) {
	return _Crabada.Contract.CraIncentiveEnable(&_Crabada.CallOpts)
}

// CrabaInfos is a free data retrieval call binding the contract method 0x62428e4e.
//
// Solidity: function crabaInfos(uint256 ) view returns(address owner, uint128 lockTo, uint128 status)
func (_Crabada *CrabadaCaller) CrabaInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner  common.Address
	LockTo *big.Int
	Status *big.Int
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "crabaInfos", arg0)

	outstruct := new(struct {
		Owner  common.Address
		LockTo *big.Int
		Status *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LockTo = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CrabaInfos is a free data retrieval call binding the contract method 0x62428e4e.
//
// Solidity: function crabaInfos(uint256 ) view returns(address owner, uint128 lockTo, uint128 status)
func (_Crabada *CrabadaSession) CrabaInfos(arg0 *big.Int) (struct {
	Owner  common.Address
	LockTo *big.Int
	Status *big.Int
}, error) {
	return _Crabada.Contract.CrabaInfos(&_Crabada.CallOpts, arg0)
}

// CrabaInfos is a free data retrieval call binding the contract method 0x62428e4e.
//
// Solidity: function crabaInfos(uint256 ) view returns(address owner, uint128 lockTo, uint128 status)
func (_Crabada *CrabadaCallerSession) CrabaInfos(arg0 *big.Int) (struct {
	Owner  common.Address
	LockTo *big.Int
	Status *big.Int
}, error) {
	return _Crabada.Contract.CrabaInfos(&_Crabada.CallOpts, arg0)
}

// ExpandTeamCounts is a free data retrieval call binding the contract method 0x24aa3f61.
//
// Solidity: function expandTeamCounts(address ) view returns(uint256)
func (_Crabada *CrabadaCaller) ExpandTeamCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "expandTeamCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpandTeamCounts is a free data retrieval call binding the contract method 0x24aa3f61.
//
// Solidity: function expandTeamCounts(address ) view returns(uint256)
func (_Crabada *CrabadaSession) ExpandTeamCounts(arg0 common.Address) (*big.Int, error) {
	return _Crabada.Contract.ExpandTeamCounts(&_Crabada.CallOpts, arg0)
}

// ExpandTeamCounts is a free data retrieval call binding the contract method 0x24aa3f61.
//
// Solidity: function expandTeamCounts(address ) view returns(uint256)
func (_Crabada *CrabadaCallerSession) ExpandTeamCounts(arg0 common.Address) (*big.Int, error) {
	return _Crabada.Contract.ExpandTeamCounts(&_Crabada.CallOpts, arg0)
}

// ExpandTeamPrice is a free data retrieval call binding the contract method 0x94a76005.
//
// Solidity: function expandTeamPrice() view returns(uint256)
func (_Crabada *CrabadaCaller) ExpandTeamPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "expandTeamPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpandTeamPrice is a free data retrieval call binding the contract method 0x94a76005.
//
// Solidity: function expandTeamPrice() view returns(uint256)
func (_Crabada *CrabadaSession) ExpandTeamPrice() (*big.Int, error) {
	return _Crabada.Contract.ExpandTeamPrice(&_Crabada.CallOpts)
}

// ExpandTeamPrice is a free data retrieval call binding the contract method 0x94a76005.
//
// Solidity: function expandTeamPrice() view returns(uint256)
func (_Crabada *CrabadaCallerSession) ExpandTeamPrice() (*big.Int, error) {
	return _Crabada.Contract.ExpandTeamPrice(&_Crabada.CallOpts)
}

// GetBattleTimePoint is a free data retrieval call binding the contract method 0x0f769a0c.
//
// Solidity: function getBattleTimePoint(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaCaller) GetBattleTimePoint(opts *bind.CallOpts, gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "getBattleTimePoint", gameId)

	outstruct := new(struct {
		AttackPoint  uint16
		DefensePoint uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttackPoint = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.DefensePoint = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetBattleTimePoint is a free data retrieval call binding the contract method 0x0f769a0c.
//
// Solidity: function getBattleTimePoint(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaSession) GetBattleTimePoint(gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	return _Crabada.Contract.GetBattleTimePoint(&_Crabada.CallOpts, gameId)
}

// GetBattleTimePoint is a free data retrieval call binding the contract method 0x0f769a0c.
//
// Solidity: function getBattleTimePoint(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaCallerSession) GetBattleTimePoint(gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	return _Crabada.Contract.GetBattleTimePoint(&_Crabada.CallOpts, gameId)
}

// GetContractAddresses is a free data retrieval call binding the contract method 0x953874d2.
//
// Solidity: function getContractAddresses() view returns(address tusAddr, address craAddr, address crabadaAddr, address statAddress)
func (_Crabada *CrabadaCaller) GetContractAddresses(opts *bind.CallOpts) (struct {
	TusAddr     common.Address
	CraAddr     common.Address
	CrabadaAddr common.Address
	StatAddress common.Address
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "getContractAddresses")

	outstruct := new(struct {
		TusAddr     common.Address
		CraAddr     common.Address
		CrabadaAddr common.Address
		StatAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TusAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CraAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CrabadaAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.StatAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetContractAddresses is a free data retrieval call binding the contract method 0x953874d2.
//
// Solidity: function getContractAddresses() view returns(address tusAddr, address craAddr, address crabadaAddr, address statAddress)
func (_Crabada *CrabadaSession) GetContractAddresses() (struct {
	TusAddr     common.Address
	CraAddr     common.Address
	CrabadaAddr common.Address
	StatAddress common.Address
}, error) {
	return _Crabada.Contract.GetContractAddresses(&_Crabada.CallOpts)
}

// GetContractAddresses is a free data retrieval call binding the contract method 0x953874d2.
//
// Solidity: function getContractAddresses() view returns(address tusAddr, address craAddr, address crabadaAddr, address statAddress)
func (_Crabada *CrabadaCallerSession) GetContractAddresses() (struct {
	TusAddr     common.Address
	CraAddr     common.Address
	CrabadaAddr common.Address
	StatAddress common.Address
}, error) {
	return _Crabada.Contract.GetContractAddresses(&_Crabada.CallOpts)
}

// GetGameBasicInfo is a free data retrieval call binding the contract method 0xf0344e36.
//
// Solidity: function getGameBasicInfo(uint256 gameId) view returns(uint256 teamId, uint128 craReward, uint128 tusReward, uint32 startTime, uint32 duration, uint32 status)
func (_Crabada *CrabadaCaller) GetGameBasicInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	TeamId    *big.Int
	CraReward *big.Int
	TusReward *big.Int
	StartTime uint32
	Duration  uint32
	Status    uint32
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "getGameBasicInfo", gameId)

	outstruct := new(struct {
		TeamId    *big.Int
		CraReward *big.Int
		TusReward *big.Int
		StartTime uint32
		Duration  uint32
		Status    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeamId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CraReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TusReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Duration = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetGameBasicInfo is a free data retrieval call binding the contract method 0xf0344e36.
//
// Solidity: function getGameBasicInfo(uint256 gameId) view returns(uint256 teamId, uint128 craReward, uint128 tusReward, uint32 startTime, uint32 duration, uint32 status)
func (_Crabada *CrabadaSession) GetGameBasicInfo(gameId *big.Int) (struct {
	TeamId    *big.Int
	CraReward *big.Int
	TusReward *big.Int
	StartTime uint32
	Duration  uint32
	Status    uint32
}, error) {
	return _Crabada.Contract.GetGameBasicInfo(&_Crabada.CallOpts, gameId)
}

// GetGameBasicInfo is a free data retrieval call binding the contract method 0xf0344e36.
//
// Solidity: function getGameBasicInfo(uint256 gameId) view returns(uint256 teamId, uint128 craReward, uint128 tusReward, uint32 startTime, uint32 duration, uint32 status)
func (_Crabada *CrabadaCallerSession) GetGameBasicInfo(gameId *big.Int) (struct {
	TeamId    *big.Int
	CraReward *big.Int
	TusReward *big.Int
	StartTime uint32
	Duration  uint32
	Status    uint32
}, error) {
	return _Crabada.Contract.GetGameBasicInfo(&_Crabada.CallOpts, gameId)
}

// GetGameBattleInfo is a free data retrieval call binding the contract method 0x183ce75d.
//
// Solidity: function getGameBattleInfo(uint256 gameId) view returns(uint256 attackTeamId, uint32 attackTime, uint32 lastAttackTime, uint32 lastDefTime, uint256 attackId1, uint256 attackId2, uint256 defId1, uint256 defId2)
func (_Crabada *CrabadaCaller) GetGameBattleInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	AttackTeamId   *big.Int
	AttackTime     uint32
	LastAttackTime uint32
	LastDefTime    uint32
	AttackId1      *big.Int
	AttackId2      *big.Int
	DefId1         *big.Int
	DefId2         *big.Int
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "getGameBattleInfo", gameId)

	outstruct := new(struct {
		AttackTeamId   *big.Int
		AttackTime     uint32
		LastAttackTime uint32
		LastDefTime    uint32
		AttackId1      *big.Int
		AttackId2      *big.Int
		DefId1         *big.Int
		DefId2         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttackTeamId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AttackTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LastAttackTime = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.LastDefTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AttackId1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AttackId2 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DefId1 = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DefId2 = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGameBattleInfo is a free data retrieval call binding the contract method 0x183ce75d.
//
// Solidity: function getGameBattleInfo(uint256 gameId) view returns(uint256 attackTeamId, uint32 attackTime, uint32 lastAttackTime, uint32 lastDefTime, uint256 attackId1, uint256 attackId2, uint256 defId1, uint256 defId2)
func (_Crabada *CrabadaSession) GetGameBattleInfo(gameId *big.Int) (struct {
	AttackTeamId   *big.Int
	AttackTime     uint32
	LastAttackTime uint32
	LastDefTime    uint32
	AttackId1      *big.Int
	AttackId2      *big.Int
	DefId1         *big.Int
	DefId2         *big.Int
}, error) {
	return _Crabada.Contract.GetGameBattleInfo(&_Crabada.CallOpts, gameId)
}

// GetGameBattleInfo is a free data retrieval call binding the contract method 0x183ce75d.
//
// Solidity: function getGameBattleInfo(uint256 gameId) view returns(uint256 attackTeamId, uint32 attackTime, uint32 lastAttackTime, uint32 lastDefTime, uint256 attackId1, uint256 attackId2, uint256 defId1, uint256 defId2)
func (_Crabada *CrabadaCallerSession) GetGameBattleInfo(gameId *big.Int) (struct {
	AttackTeamId   *big.Int
	AttackTime     uint32
	LastAttackTime uint32
	LastDefTime    uint32
	AttackId1      *big.Int
	AttackId2      *big.Int
	DefId1         *big.Int
	DefId2         *big.Int
}, error) {
	return _Crabada.Contract.GetGameBattleInfo(&_Crabada.CallOpts, gameId)
}

// GetLootingStatsInfo is a free data retrieval call binding the contract method 0xab0c8f8d.
//
// Solidity: function getLootingStatsInfo(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaCaller) GetLootingStatsInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "getLootingStatsInfo", gameId)

	outstruct := new(struct {
		AttackPoint  uint16
		DefensePoint uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttackPoint = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.DefensePoint = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetLootingStatsInfo is a free data retrieval call binding the contract method 0xab0c8f8d.
//
// Solidity: function getLootingStatsInfo(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaSession) GetLootingStatsInfo(gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	return _Crabada.Contract.GetLootingStatsInfo(&_Crabada.CallOpts, gameId)
}

// GetLootingStatsInfo is a free data retrieval call binding the contract method 0xab0c8f8d.
//
// Solidity: function getLootingStatsInfo(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaCallerSession) GetLootingStatsInfo(gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	return _Crabada.Contract.GetLootingStatsInfo(&_Crabada.CallOpts, gameId)
}

// GetStats is a free data retrieval call binding the contract method 0x7b303965.
//
// Solidity: function getStats(uint256 crabadaId) view returns(uint16 battlePoint, uint16 timePoint)
func (_Crabada *CrabadaCaller) GetStats(opts *bind.CallOpts, crabadaId *big.Int) (struct {
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "getStats", crabadaId)

	outstruct := new(struct {
		BattlePoint uint16
		TimePoint   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BattlePoint = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.TimePoint = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetStats is a free data retrieval call binding the contract method 0x7b303965.
//
// Solidity: function getStats(uint256 crabadaId) view returns(uint16 battlePoint, uint16 timePoint)
func (_Crabada *CrabadaSession) GetStats(crabadaId *big.Int) (struct {
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	return _Crabada.Contract.GetStats(&_Crabada.CallOpts, crabadaId)
}

// GetStats is a free data retrieval call binding the contract method 0x7b303965.
//
// Solidity: function getStats(uint256 crabadaId) view returns(uint16 battlePoint, uint16 timePoint)
func (_Crabada *CrabadaCallerSession) GetStats(crabadaId *big.Int) (struct {
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	return _Crabada.Contract.GetStats(&_Crabada.CallOpts, crabadaId)
}

// GetTeamInfo is a free data retrieval call binding the contract method 0x969215ba.
//
// Solidity: function getTeamInfo(uint256 teamId) view returns(address owner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabada *CrabadaCaller) GetTeamInfo(opts *bind.CallOpts, teamId *big.Int) (struct {
	Owner         common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "getTeamInfo", teamId)

	outstruct := new(struct {
		Owner         common.Address
		CrabadaId1    *big.Int
		CrabadaId2    *big.Int
		CrabadaId3    *big.Int
		BattlePoint   uint16
		TimePoint     uint16
		CurrentGameId *big.Int
		LockTo        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CrabadaId1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId2 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId3 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BattlePoint = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.TimePoint = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.CurrentGameId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LockTo = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTeamInfo is a free data retrieval call binding the contract method 0x969215ba.
//
// Solidity: function getTeamInfo(uint256 teamId) view returns(address owner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabada *CrabadaSession) GetTeamInfo(teamId *big.Int) (struct {
	Owner         common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Crabada.Contract.GetTeamInfo(&_Crabada.CallOpts, teamId)
}

// GetTeamInfo is a free data retrieval call binding the contract method 0x969215ba.
//
// Solidity: function getTeamInfo(uint256 teamId) view returns(address owner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabada *CrabadaCallerSession) GetTeamInfo(teamId *big.Int) (struct {
	Owner         common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Crabada.Contract.GetTeamInfo(&_Crabada.CallOpts, teamId)
}

// LendingFeeHolerAddress is a free data retrieval call binding the contract method 0x6dad6b4b.
//
// Solidity: function lendingFeeHolerAddress() view returns(address)
func (_Crabada *CrabadaCaller) LendingFeeHolerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "lendingFeeHolerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendingFeeHolerAddress is a free data retrieval call binding the contract method 0x6dad6b4b.
//
// Solidity: function lendingFeeHolerAddress() view returns(address)
func (_Crabada *CrabadaSession) LendingFeeHolerAddress() (common.Address, error) {
	return _Crabada.Contract.LendingFeeHolerAddress(&_Crabada.CallOpts)
}

// LendingFeeHolerAddress is a free data retrieval call binding the contract method 0x6dad6b4b.
//
// Solidity: function lendingFeeHolerAddress() view returns(address)
func (_Crabada *CrabadaCallerSession) LendingFeeHolerAddress() (common.Address, error) {
	return _Crabada.Contract.LendingFeeHolerAddress(&_Crabada.CallOpts)
}

// LendingFeePercent is a free data retrieval call binding the contract method 0x4699f846.
//
// Solidity: function lendingFeePercent() view returns(uint128)
func (_Crabada *CrabadaCaller) LendingFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "lendingFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendingFeePercent is a free data retrieval call binding the contract method 0x4699f846.
//
// Solidity: function lendingFeePercent() view returns(uint128)
func (_Crabada *CrabadaSession) LendingFeePercent() (*big.Int, error) {
	return _Crabada.Contract.LendingFeePercent(&_Crabada.CallOpts)
}

// LendingFeePercent is a free data retrieval call binding the contract method 0x4699f846.
//
// Solidity: function lendingFeePercent() view returns(uint128)
func (_Crabada *CrabadaCallerSession) LendingFeePercent() (*big.Int, error) {
	return _Crabada.Contract.LendingFeePercent(&_Crabada.CallOpts)
}

// LendingLockDuration is a free data retrieval call binding the contract method 0x24a9c445.
//
// Solidity: function lendingLockDuration() view returns(uint128)
func (_Crabada *CrabadaCaller) LendingLockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "lendingLockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendingLockDuration is a free data retrieval call binding the contract method 0x24a9c445.
//
// Solidity: function lendingLockDuration() view returns(uint128)
func (_Crabada *CrabadaSession) LendingLockDuration() (*big.Int, error) {
	return _Crabada.Contract.LendingLockDuration(&_Crabada.CallOpts)
}

// LendingLockDuration is a free data retrieval call binding the contract method 0x24a9c445.
//
// Solidity: function lendingLockDuration() view returns(uint128)
func (_Crabada *CrabadaCallerSession) LendingLockDuration() (*big.Int, error) {
	return _Crabada.Contract.LendingLockDuration(&_Crabada.CallOpts)
}

// LendingPool is a free data retrieval call binding the contract method 0x2ccb74c8.
//
// Solidity: function lendingPool(uint256 ) view returns(uint256)
func (_Crabada *CrabadaCaller) LendingPool(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "lendingPool", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendingPool is a free data retrieval call binding the contract method 0x2ccb74c8.
//
// Solidity: function lendingPool(uint256 ) view returns(uint256)
func (_Crabada *CrabadaSession) LendingPool(arg0 *big.Int) (*big.Int, error) {
	return _Crabada.Contract.LendingPool(&_Crabada.CallOpts, arg0)
}

// LendingPool is a free data retrieval call binding the contract method 0x2ccb74c8.
//
// Solidity: function lendingPool(uint256 ) view returns(uint256)
func (_Crabada *CrabadaCallerSession) LendingPool(arg0 *big.Int) (*big.Int, error) {
	return _Crabada.Contract.LendingPool(&_Crabada.CallOpts, arg0)
}

// LootingCraBaseReward is a free data retrieval call binding the contract method 0xf2839523.
//
// Solidity: function lootingCraBaseReward() view returns(uint128)
func (_Crabada *CrabadaCaller) LootingCraBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "lootingCraBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LootingCraBaseReward is a free data retrieval call binding the contract method 0xf2839523.
//
// Solidity: function lootingCraBaseReward() view returns(uint128)
func (_Crabada *CrabadaSession) LootingCraBaseReward() (*big.Int, error) {
	return _Crabada.Contract.LootingCraBaseReward(&_Crabada.CallOpts)
}

// LootingCraBaseReward is a free data retrieval call binding the contract method 0xf2839523.
//
// Solidity: function lootingCraBaseReward() view returns(uint128)
func (_Crabada *CrabadaCallerSession) LootingCraBaseReward() (*big.Int, error) {
	return _Crabada.Contract.LootingCraBaseReward(&_Crabada.CallOpts)
}

// LootingPercent is a free data retrieval call binding the contract method 0xfec8d7d1.
//
// Solidity: function lootingPercent() view returns(uint256)
func (_Crabada *CrabadaCaller) LootingPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "lootingPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LootingPercent is a free data retrieval call binding the contract method 0xfec8d7d1.
//
// Solidity: function lootingPercent() view returns(uint256)
func (_Crabada *CrabadaSession) LootingPercent() (*big.Int, error) {
	return _Crabada.Contract.LootingPercent(&_Crabada.CallOpts)
}

// LootingPercent is a free data retrieval call binding the contract method 0xfec8d7d1.
//
// Solidity: function lootingPercent() view returns(uint256)
func (_Crabada *CrabadaCallerSession) LootingPercent() (*big.Int, error) {
	return _Crabada.Contract.LootingPercent(&_Crabada.CallOpts)
}

// LootingTusBaseReward is a free data retrieval call binding the contract method 0x4f3b9225.
//
// Solidity: function lootingTusBaseReward() view returns(uint128)
func (_Crabada *CrabadaCaller) LootingTusBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "lootingTusBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LootingTusBaseReward is a free data retrieval call binding the contract method 0x4f3b9225.
//
// Solidity: function lootingTusBaseReward() view returns(uint128)
func (_Crabada *CrabadaSession) LootingTusBaseReward() (*big.Int, error) {
	return _Crabada.Contract.LootingTusBaseReward(&_Crabada.CallOpts)
}

// LootingTusBaseReward is a free data retrieval call binding the contract method 0x4f3b9225.
//
// Solidity: function lootingTusBaseReward() view returns(uint128)
func (_Crabada *CrabadaCallerSession) LootingTusBaseReward() (*big.Int, error) {
	return _Crabada.Contract.LootingTusBaseReward(&_Crabada.CallOpts)
}

// MinerRevengeEnable is a free data retrieval call binding the contract method 0xb78803a3.
//
// Solidity: function minerRevengeEnable() view returns(bool)
func (_Crabada *CrabadaCaller) MinerRevengeEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "minerRevengeEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinerRevengeEnable is a free data retrieval call binding the contract method 0xb78803a3.
//
// Solidity: function minerRevengeEnable() view returns(bool)
func (_Crabada *CrabadaSession) MinerRevengeEnable() (bool, error) {
	return _Crabada.Contract.MinerRevengeEnable(&_Crabada.CallOpts)
}

// MinerRevengeEnable is a free data retrieval call binding the contract method 0xb78803a3.
//
// Solidity: function minerRevengeEnable() view returns(bool)
func (_Crabada *CrabadaCallerSession) MinerRevengeEnable() (bool, error) {
	return _Crabada.Contract.MinerRevengeEnable(&_Crabada.CallOpts)
}

// MiningDuration is a free data retrieval call binding the contract method 0x7dd8a48f.
//
// Solidity: function miningDuration() view returns(uint128)
func (_Crabada *CrabadaCaller) MiningDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "miningDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningDuration is a free data retrieval call binding the contract method 0x7dd8a48f.
//
// Solidity: function miningDuration() view returns(uint128)
func (_Crabada *CrabadaSession) MiningDuration() (*big.Int, error) {
	return _Crabada.Contract.MiningDuration(&_Crabada.CallOpts)
}

// MiningDuration is a free data retrieval call binding the contract method 0x7dd8a48f.
//
// Solidity: function miningDuration() view returns(uint128)
func (_Crabada *CrabadaCallerSession) MiningDuration() (*big.Int, error) {
	return _Crabada.Contract.MiningDuration(&_Crabada.CallOpts)
}

// ModifierSetting is a free data retrieval call binding the contract method 0x091992b9.
//
// Solidity: function modifierSetting() view returns(uint256)
func (_Crabada *CrabadaCaller) ModifierSetting(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "modifierSetting")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModifierSetting is a free data retrieval call binding the contract method 0x091992b9.
//
// Solidity: function modifierSetting() view returns(uint256)
func (_Crabada *CrabadaSession) ModifierSetting() (*big.Int, error) {
	return _Crabada.Contract.ModifierSetting(&_Crabada.CallOpts)
}

// ModifierSetting is a free data retrieval call binding the contract method 0x091992b9.
//
// Solidity: function modifierSetting() view returns(uint256)
func (_Crabada *CrabadaCallerSession) ModifierSetting() (*big.Int, error) {
	return _Crabada.Contract.ModifierSetting(&_Crabada.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crabada *CrabadaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crabada *CrabadaSession) Owner() (common.Address, error) {
	return _Crabada.Contract.Owner(&_Crabada.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crabada *CrabadaCallerSession) Owner() (common.Address, error) {
	return _Crabada.Contract.Owner(&_Crabada.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 crabadaId) view returns(address)
func (_Crabada *CrabadaCaller) OwnerOf(opts *bind.CallOpts, crabadaId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "ownerOf", crabadaId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 crabadaId) view returns(address)
func (_Crabada *CrabadaSession) OwnerOf(crabadaId *big.Int) (common.Address, error) {
	return _Crabada.Contract.OwnerOf(&_Crabada.CallOpts, crabadaId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 crabadaId) view returns(address)
func (_Crabada *CrabadaCallerSession) OwnerOf(crabadaId *big.Int) (common.Address, error) {
	return _Crabada.Contract.OwnerOf(&_Crabada.CallOpts, crabadaId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Crabada *CrabadaCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Crabada *CrabadaSession) Paused() (bool, error) {
	return _Crabada.Contract.Paused(&_Crabada.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Crabada *CrabadaCallerSession) Paused() (bool, error) {
	return _Crabada.Contract.Paused(&_Crabada.CallOpts)
}

// ReinforceLockDuration is a free data retrieval call binding the contract method 0xe55cebde.
//
// Solidity: function reinforceLockDuration() view returns(uint128)
func (_Crabada *CrabadaCaller) ReinforceLockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "reinforceLockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReinforceLockDuration is a free data retrieval call binding the contract method 0xe55cebde.
//
// Solidity: function reinforceLockDuration() view returns(uint128)
func (_Crabada *CrabadaSession) ReinforceLockDuration() (*big.Int, error) {
	return _Crabada.Contract.ReinforceLockDuration(&_Crabada.CallOpts)
}

// ReinforceLockDuration is a free data retrieval call binding the contract method 0xe55cebde.
//
// Solidity: function reinforceLockDuration() view returns(uint128)
func (_Crabada *CrabadaCallerSession) ReinforceLockDuration() (*big.Int, error) {
	return _Crabada.Contract.ReinforceLockDuration(&_Crabada.CallOpts)
}

// RngContract is a free data retrieval call binding the contract method 0xf0d86d55.
//
// Solidity: function rngContract() view returns(address)
func (_Crabada *CrabadaCaller) RngContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "rngContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RngContract is a free data retrieval call binding the contract method 0xf0d86d55.
//
// Solidity: function rngContract() view returns(address)
func (_Crabada *CrabadaSession) RngContract() (common.Address, error) {
	return _Crabada.Contract.RngContract(&_Crabada.CallOpts)
}

// RngContract is a free data retrieval call binding the contract method 0xf0d86d55.
//
// Solidity: function rngContract() view returns(address)
func (_Crabada *CrabadaCallerSession) RngContract() (common.Address, error) {
	return _Crabada.Contract.RngContract(&_Crabada.CallOpts)
}

// TeamCounts is a free data retrieval call binding the contract method 0x4f35f123.
//
// Solidity: function teamCounts(address ) view returns(uint256)
func (_Crabada *CrabadaCaller) TeamCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crabada.contract.Call(opts, &out, "teamCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamCounts is a free data retrieval call binding the contract method 0x4f35f123.
//
// Solidity: function teamCounts(address ) view returns(uint256)
func (_Crabada *CrabadaSession) TeamCounts(arg0 common.Address) (*big.Int, error) {
	return _Crabada.Contract.TeamCounts(&_Crabada.CallOpts, arg0)
}

// TeamCounts is a free data retrieval call binding the contract method 0x4f35f123.
//
// Solidity: function teamCounts(address ) view returns(uint256)
func (_Crabada *CrabadaCallerSession) TeamCounts(arg0 common.Address) (*big.Int, error) {
	return _Crabada.Contract.TeamCounts(&_Crabada.CallOpts, arg0)
}

// AddCrabadaToTeam is a paid mutator transaction binding the contract method 0xc0d8080c.
//
// Solidity: function addCrabadaToTeam(uint256 teamId, uint256 position, uint256 crabadaId) returns()
func (_Crabada *CrabadaTransactor) AddCrabadaToTeam(opts *bind.TransactOpts, teamId *big.Int, position *big.Int, crabadaId *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "addCrabadaToTeam", teamId, position, crabadaId)
}

// AddCrabadaToTeam is a paid mutator transaction binding the contract method 0xc0d8080c.
//
// Solidity: function addCrabadaToTeam(uint256 teamId, uint256 position, uint256 crabadaId) returns()
func (_Crabada *CrabadaSession) AddCrabadaToTeam(teamId *big.Int, position *big.Int, crabadaId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.AddCrabadaToTeam(&_Crabada.TransactOpts, teamId, position, crabadaId)
}

// AddCrabadaToTeam is a paid mutator transaction binding the contract method 0xc0d8080c.
//
// Solidity: function addCrabadaToTeam(uint256 teamId, uint256 position, uint256 crabadaId) returns()
func (_Crabada *CrabadaTransactorSession) AddCrabadaToTeam(teamId *big.Int, position *big.Int, crabadaId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.AddCrabadaToTeam(&_Crabada.TransactOpts, teamId, position, crabadaId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 gameId, uint256 attackTeamId) returns()
func (_Crabada *CrabadaTransactor) Attack(opts *bind.TransactOpts, gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "attack", gameId, attackTeamId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 gameId, uint256 attackTeamId) returns()
func (_Crabada *CrabadaSession) Attack(gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.Attack(&_Crabada.TransactOpts, gameId, attackTeamId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 gameId, uint256 attackTeamId) returns()
func (_Crabada *CrabadaTransactorSession) Attack(gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.Attack(&_Crabada.TransactOpts, gameId, attackTeamId)
}

// ChangeContractAddresses is a paid mutator transaction binding the contract method 0xcc9d080b.
//
// Solidity: function changeContractAddresses(address _tusToken, address _craToken, address _crabada, address _gameStatContract) returns()
func (_Crabada *CrabadaTransactor) ChangeContractAddresses(opts *bind.TransactOpts, _tusToken common.Address, _craToken common.Address, _crabada common.Address, _gameStatContract common.Address) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "changeContractAddresses", _tusToken, _craToken, _crabada, _gameStatContract)
}

// ChangeContractAddresses is a paid mutator transaction binding the contract method 0xcc9d080b.
//
// Solidity: function changeContractAddresses(address _tusToken, address _craToken, address _crabada, address _gameStatContract) returns()
func (_Crabada *CrabadaSession) ChangeContractAddresses(_tusToken common.Address, _craToken common.Address, _crabada common.Address, _gameStatContract common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.ChangeContractAddresses(&_Crabada.TransactOpts, _tusToken, _craToken, _crabada, _gameStatContract)
}

// ChangeContractAddresses is a paid mutator transaction binding the contract method 0xcc9d080b.
//
// Solidity: function changeContractAddresses(address _tusToken, address _craToken, address _crabada, address _gameStatContract) returns()
func (_Crabada *CrabadaTransactorSession) ChangeContractAddresses(_tusToken common.Address, _craToken common.Address, _crabada common.Address, _gameStatContract common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.ChangeContractAddresses(&_Crabada.TransactOpts, _tusToken, _craToken, _crabada, _gameStatContract)
}

// CloseGame is a paid mutator transaction binding the contract method 0x2d6ef310.
//
// Solidity: function closeGame(uint256 gameId) returns()
func (_Crabada *CrabadaTransactor) CloseGame(opts *bind.TransactOpts, gameId *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "closeGame", gameId)
}

// CloseGame is a paid mutator transaction binding the contract method 0x2d6ef310.
//
// Solidity: function closeGame(uint256 gameId) returns()
func (_Crabada *CrabadaSession) CloseGame(gameId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.CloseGame(&_Crabada.TransactOpts, gameId)
}

// CloseGame is a paid mutator transaction binding the contract method 0x2d6ef310.
//
// Solidity: function closeGame(uint256 gameId) returns()
func (_Crabada *CrabadaTransactorSession) CloseGame(gameId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.CloseGame(&_Crabada.TransactOpts, gameId)
}

// CreateTeam is a paid mutator transaction binding the contract method 0xcf034493.
//
// Solidity: function createTeam(uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3) returns(uint256 teamId)
func (_Crabada *CrabadaTransactor) CreateTeam(opts *bind.TransactOpts, crabadaId1 *big.Int, crabadaId2 *big.Int, crabadaId3 *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "createTeam", crabadaId1, crabadaId2, crabadaId3)
}

// CreateTeam is a paid mutator transaction binding the contract method 0xcf034493.
//
// Solidity: function createTeam(uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3) returns(uint256 teamId)
func (_Crabada *CrabadaSession) CreateTeam(crabadaId1 *big.Int, crabadaId2 *big.Int, crabadaId3 *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.CreateTeam(&_Crabada.TransactOpts, crabadaId1, crabadaId2, crabadaId3)
}

// CreateTeam is a paid mutator transaction binding the contract method 0xcf034493.
//
// Solidity: function createTeam(uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3) returns(uint256 teamId)
func (_Crabada *CrabadaTransactorSession) CreateTeam(crabadaId1 *big.Int, crabadaId2 *big.Int, crabadaId3 *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.CreateTeam(&_Crabada.TransactOpts, crabadaId1, crabadaId2, crabadaId3)
}

// Deposit is a paid mutator transaction binding the contract method 0x598b8e71.
//
// Solidity: function deposit(uint256[] crabadaIds) returns()
func (_Crabada *CrabadaTransactor) Deposit(opts *bind.TransactOpts, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "deposit", crabadaIds)
}

// Deposit is a paid mutator transaction binding the contract method 0x598b8e71.
//
// Solidity: function deposit(uint256[] crabadaIds) returns()
func (_Crabada *CrabadaSession) Deposit(crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.Deposit(&_Crabada.TransactOpts, crabadaIds)
}

// Deposit is a paid mutator transaction binding the contract method 0x598b8e71.
//
// Solidity: function deposit(uint256[] crabadaIds) returns()
func (_Crabada *CrabadaTransactorSession) Deposit(crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.Deposit(&_Crabada.TransactOpts, crabadaIds)
}

// EnableMinerRevenge is a paid mutator transaction binding the contract method 0xd604681b.
//
// Solidity: function enableMinerRevenge(bool isEnable) returns()
func (_Crabada *CrabadaTransactor) EnableMinerRevenge(opts *bind.TransactOpts, isEnable bool) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "enableMinerRevenge", isEnable)
}

// EnableMinerRevenge is a paid mutator transaction binding the contract method 0xd604681b.
//
// Solidity: function enableMinerRevenge(bool isEnable) returns()
func (_Crabada *CrabadaSession) EnableMinerRevenge(isEnable bool) (*types.Transaction, error) {
	return _Crabada.Contract.EnableMinerRevenge(&_Crabada.TransactOpts, isEnable)
}

// EnableMinerRevenge is a paid mutator transaction binding the contract method 0xd604681b.
//
// Solidity: function enableMinerRevenge(bool isEnable) returns()
func (_Crabada *CrabadaTransactorSession) EnableMinerRevenge(isEnable bool) (*types.Transaction, error) {
	return _Crabada.Contract.EnableMinerRevenge(&_Crabada.TransactOpts, isEnable)
}

// ExpandTeamLimit is a paid mutator transaction binding the contract method 0x973abe53.
//
// Solidity: function expandTeamLimit(uint256 amount) returns()
func (_Crabada *CrabadaTransactor) ExpandTeamLimit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "expandTeamLimit", amount)
}

// ExpandTeamLimit is a paid mutator transaction binding the contract method 0x973abe53.
//
// Solidity: function expandTeamLimit(uint256 amount) returns()
func (_Crabada *CrabadaSession) ExpandTeamLimit(amount *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.ExpandTeamLimit(&_Crabada.TransactOpts, amount)
}

// ExpandTeamLimit is a paid mutator transaction binding the contract method 0x973abe53.
//
// Solidity: function expandTeamLimit(uint256 amount) returns()
func (_Crabada *CrabadaTransactorSession) ExpandTeamLimit(amount *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.ExpandTeamLimit(&_Crabada.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _crabada, address _cra, address _tus, address _gameStat, address _owner) returns()
func (_Crabada *CrabadaTransactor) Initialize(opts *bind.TransactOpts, _crabada common.Address, _cra common.Address, _tus common.Address, _gameStat common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "initialize", _crabada, _cra, _tus, _gameStat, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _crabada, address _cra, address _tus, address _gameStat, address _owner) returns()
func (_Crabada *CrabadaSession) Initialize(_crabada common.Address, _cra common.Address, _tus common.Address, _gameStat common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.Initialize(&_Crabada.TransactOpts, _crabada, _cra, _tus, _gameStat, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _crabada, address _cra, address _tus, address _gameStat, address _owner) returns()
func (_Crabada *CrabadaTransactorSession) Initialize(_crabada common.Address, _cra common.Address, _tus common.Address, _gameStat common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.Initialize(&_Crabada.TransactOpts, _crabada, _cra, _tus, _gameStat, _owner)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Crabada *CrabadaTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Crabada *CrabadaSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Crabada.Contract.OnERC721Received(&_Crabada.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Crabada *CrabadaTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Crabada.Contract.OnERC721Received(&_Crabada.TransactOpts, operator, from, tokenId, data)
}

// ReinforceAttack is a paid mutator transaction binding the contract method 0x3dc8d5ce.
//
// Solidity: function reinforceAttack(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Crabada *CrabadaTransactor) ReinforceAttack(opts *bind.TransactOpts, gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "reinforceAttack", gameId, crabadaId, borrowPrice)
}

// ReinforceAttack is a paid mutator transaction binding the contract method 0x3dc8d5ce.
//
// Solidity: function reinforceAttack(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Crabada *CrabadaSession) ReinforceAttack(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.ReinforceAttack(&_Crabada.TransactOpts, gameId, crabadaId, borrowPrice)
}

// ReinforceAttack is a paid mutator transaction binding the contract method 0x3dc8d5ce.
//
// Solidity: function reinforceAttack(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Crabada *CrabadaTransactorSession) ReinforceAttack(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.ReinforceAttack(&_Crabada.TransactOpts, gameId, crabadaId, borrowPrice)
}

// ReinforceDefense is a paid mutator transaction binding the contract method 0x08873bfb.
//
// Solidity: function reinforceDefense(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Crabada *CrabadaTransactor) ReinforceDefense(opts *bind.TransactOpts, gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "reinforceDefense", gameId, crabadaId, borrowPrice)
}

// ReinforceDefense is a paid mutator transaction binding the contract method 0x08873bfb.
//
// Solidity: function reinforceDefense(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Crabada *CrabadaSession) ReinforceDefense(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.ReinforceDefense(&_Crabada.TransactOpts, gameId, crabadaId, borrowPrice)
}

// ReinforceDefense is a paid mutator transaction binding the contract method 0x08873bfb.
//
// Solidity: function reinforceDefense(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Crabada *CrabadaTransactorSession) ReinforceDefense(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.ReinforceDefense(&_Crabada.TransactOpts, gameId, crabadaId, borrowPrice)
}

// RemoveCrabadaFromTeam is a paid mutator transaction binding the contract method 0x31c1bf82.
//
// Solidity: function removeCrabadaFromTeam(uint256 teamId, uint256 position) returns()
func (_Crabada *CrabadaTransactor) RemoveCrabadaFromTeam(opts *bind.TransactOpts, teamId *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "removeCrabadaFromTeam", teamId, position)
}

// RemoveCrabadaFromTeam is a paid mutator transaction binding the contract method 0x31c1bf82.
//
// Solidity: function removeCrabadaFromTeam(uint256 teamId, uint256 position) returns()
func (_Crabada *CrabadaSession) RemoveCrabadaFromTeam(teamId *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.RemoveCrabadaFromTeam(&_Crabada.TransactOpts, teamId, position)
}

// RemoveCrabadaFromTeam is a paid mutator transaction binding the contract method 0x31c1bf82.
//
// Solidity: function removeCrabadaFromTeam(uint256 teamId, uint256 position) returns()
func (_Crabada *CrabadaTransactorSession) RemoveCrabadaFromTeam(teamId *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.RemoveCrabadaFromTeam(&_Crabada.TransactOpts, teamId, position)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crabada *CrabadaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crabada *CrabadaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crabada.Contract.RenounceOwnership(&_Crabada.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crabada *CrabadaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crabada.Contract.RenounceOwnership(&_Crabada.TransactOpts)
}

// SetAttackCooldownDuration is a paid mutator transaction binding the contract method 0xbb11c5cb.
//
// Solidity: function setAttackCooldownDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactor) SetAttackCooldownDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setAttackCooldownDuration", duration)
}

// SetAttackCooldownDuration is a paid mutator transaction binding the contract method 0xbb11c5cb.
//
// Solidity: function setAttackCooldownDuration(uint128 duration) returns()
func (_Crabada *CrabadaSession) SetAttackCooldownDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetAttackCooldownDuration(&_Crabada.TransactOpts, duration)
}

// SetAttackCooldownDuration is a paid mutator transaction binding the contract method 0xbb11c5cb.
//
// Solidity: function setAttackCooldownDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactorSession) SetAttackCooldownDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetAttackCooldownDuration(&_Crabada.TransactOpts, duration)
}

// SetAttackDuration is a paid mutator transaction binding the contract method 0x0b17abc2.
//
// Solidity: function setAttackDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactor) SetAttackDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setAttackDuration", duration)
}

// SetAttackDuration is a paid mutator transaction binding the contract method 0x0b17abc2.
//
// Solidity: function setAttackDuration(uint128 duration) returns()
func (_Crabada *CrabadaSession) SetAttackDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetAttackDuration(&_Crabada.TransactOpts, duration)
}

// SetAttackDuration is a paid mutator transaction binding the contract method 0x0b17abc2.
//
// Solidity: function setAttackDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactorSession) SetAttackDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetAttackDuration(&_Crabada.TransactOpts, duration)
}

// SetBonusClass is a paid mutator transaction binding the contract method 0x9bd44391.
//
// Solidity: function setBonusClass(uint256 classId, bool isBonused) returns()
func (_Crabada *CrabadaTransactor) SetBonusClass(opts *bind.TransactOpts, classId *big.Int, isBonused bool) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setBonusClass", classId, isBonused)
}

// SetBonusClass is a paid mutator transaction binding the contract method 0x9bd44391.
//
// Solidity: function setBonusClass(uint256 classId, bool isBonused) returns()
func (_Crabada *CrabadaSession) SetBonusClass(classId *big.Int, isBonused bool) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusClass(&_Crabada.TransactOpts, classId, isBonused)
}

// SetBonusClass is a paid mutator transaction binding the contract method 0x9bd44391.
//
// Solidity: function setBonusClass(uint256 classId, bool isBonused) returns()
func (_Crabada *CrabadaTransactorSession) SetBonusClass(classId *big.Int, isBonused bool) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusClass(&_Crabada.TransactOpts, classId, isBonused)
}

// SetBonusClassPercent is a paid mutator transaction binding the contract method 0x456502d6.
//
// Solidity: function setBonusClassPercent(uint128 percent) returns()
func (_Crabada *CrabadaTransactor) SetBonusClassPercent(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setBonusClassPercent", percent)
}

// SetBonusClassPercent is a paid mutator transaction binding the contract method 0x456502d6.
//
// Solidity: function setBonusClassPercent(uint128 percent) returns()
func (_Crabada *CrabadaSession) SetBonusClassPercent(percent *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusClassPercent(&_Crabada.TransactOpts, percent)
}

// SetBonusClassPercent is a paid mutator transaction binding the contract method 0x456502d6.
//
// Solidity: function setBonusClassPercent(uint128 percent) returns()
func (_Crabada *CrabadaTransactorSession) SetBonusClassPercent(percent *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusClassPercent(&_Crabada.TransactOpts, percent)
}

// SetBonusTimePointDuration is a paid mutator transaction binding the contract method 0xfd897937.
//
// Solidity: function setBonusTimePointDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactor) SetBonusTimePointDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setBonusTimePointDuration", duration)
}

// SetBonusTimePointDuration is a paid mutator transaction binding the contract method 0xfd897937.
//
// Solidity: function setBonusTimePointDuration(uint128 duration) returns()
func (_Crabada *CrabadaSession) SetBonusTimePointDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusTimePointDuration(&_Crabada.TransactOpts, duration)
}

// SetBonusTimePointDuration is a paid mutator transaction binding the contract method 0xfd897937.
//
// Solidity: function setBonusTimePointDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactorSession) SetBonusTimePointDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusTimePointDuration(&_Crabada.TransactOpts, duration)
}

// SetBonusTimePointThreshold is a paid mutator transaction binding the contract method 0xb0351064.
//
// Solidity: function setBonusTimePointThreshold(uint128 newThreshold) returns()
func (_Crabada *CrabadaTransactor) SetBonusTimePointThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setBonusTimePointThreshold", newThreshold)
}

// SetBonusTimePointThreshold is a paid mutator transaction binding the contract method 0xb0351064.
//
// Solidity: function setBonusTimePointThreshold(uint128 newThreshold) returns()
func (_Crabada *CrabadaSession) SetBonusTimePointThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusTimePointThreshold(&_Crabada.TransactOpts, newThreshold)
}

// SetBonusTimePointThreshold is a paid mutator transaction binding the contract method 0xb0351064.
//
// Solidity: function setBonusTimePointThreshold(uint128 newThreshold) returns()
func (_Crabada *CrabadaTransactorSession) SetBonusTimePointThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetBonusTimePointThreshold(&_Crabada.TransactOpts, newThreshold)
}

// SetCRAIncentiveStatus is a paid mutator transaction binding the contract method 0xe2c41cee.
//
// Solidity: function setCRAIncentiveStatus(bool isEnable) returns()
func (_Crabada *CrabadaTransactor) SetCRAIncentiveStatus(opts *bind.TransactOpts, isEnable bool) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setCRAIncentiveStatus", isEnable)
}

// SetCRAIncentiveStatus is a paid mutator transaction binding the contract method 0xe2c41cee.
//
// Solidity: function setCRAIncentiveStatus(bool isEnable) returns()
func (_Crabada *CrabadaSession) SetCRAIncentiveStatus(isEnable bool) (*types.Transaction, error) {
	return _Crabada.Contract.SetCRAIncentiveStatus(&_Crabada.TransactOpts, isEnable)
}

// SetCRAIncentiveStatus is a paid mutator transaction binding the contract method 0xe2c41cee.
//
// Solidity: function setCRAIncentiveStatus(bool isEnable) returns()
func (_Crabada *CrabadaTransactorSession) SetCRAIncentiveStatus(isEnable bool) (*types.Transaction, error) {
	return _Crabada.Contract.SetCRAIncentiveStatus(&_Crabada.TransactOpts, isEnable)
}

// SetLendingCooldownDuration is a paid mutator transaction binding the contract method 0xe6986845.
//
// Solidity: function setLendingCooldownDuration(uint128 cooldown) returns()
func (_Crabada *CrabadaTransactor) SetLendingCooldownDuration(opts *bind.TransactOpts, cooldown *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setLendingCooldownDuration", cooldown)
}

// SetLendingCooldownDuration is a paid mutator transaction binding the contract method 0xe6986845.
//
// Solidity: function setLendingCooldownDuration(uint128 cooldown) returns()
func (_Crabada *CrabadaSession) SetLendingCooldownDuration(cooldown *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingCooldownDuration(&_Crabada.TransactOpts, cooldown)
}

// SetLendingCooldownDuration is a paid mutator transaction binding the contract method 0xe6986845.
//
// Solidity: function setLendingCooldownDuration(uint128 cooldown) returns()
func (_Crabada *CrabadaTransactorSession) SetLendingCooldownDuration(cooldown *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingCooldownDuration(&_Crabada.TransactOpts, cooldown)
}

// SetLendingFeeHolerAddress is a paid mutator transaction binding the contract method 0x8ea53ee4.
//
// Solidity: function setLendingFeeHolerAddress(address newAddr) returns()
func (_Crabada *CrabadaTransactor) SetLendingFeeHolerAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setLendingFeeHolerAddress", newAddr)
}

// SetLendingFeeHolerAddress is a paid mutator transaction binding the contract method 0x8ea53ee4.
//
// Solidity: function setLendingFeeHolerAddress(address newAddr) returns()
func (_Crabada *CrabadaSession) SetLendingFeeHolerAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingFeeHolerAddress(&_Crabada.TransactOpts, newAddr)
}

// SetLendingFeeHolerAddress is a paid mutator transaction binding the contract method 0x8ea53ee4.
//
// Solidity: function setLendingFeeHolerAddress(address newAddr) returns()
func (_Crabada *CrabadaTransactorSession) SetLendingFeeHolerAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingFeeHolerAddress(&_Crabada.TransactOpts, newAddr)
}

// SetLendingFeePercent is a paid mutator transaction binding the contract method 0x2d73664c.
//
// Solidity: function setLendingFeePercent(uint32 percent) returns()
func (_Crabada *CrabadaTransactor) SetLendingFeePercent(opts *bind.TransactOpts, percent uint32) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setLendingFeePercent", percent)
}

// SetLendingFeePercent is a paid mutator transaction binding the contract method 0x2d73664c.
//
// Solidity: function setLendingFeePercent(uint32 percent) returns()
func (_Crabada *CrabadaSession) SetLendingFeePercent(percent uint32) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingFeePercent(&_Crabada.TransactOpts, percent)
}

// SetLendingFeePercent is a paid mutator transaction binding the contract method 0x2d73664c.
//
// Solidity: function setLendingFeePercent(uint32 percent) returns()
func (_Crabada *CrabadaTransactorSession) SetLendingFeePercent(percent uint32) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingFeePercent(&_Crabada.TransactOpts, percent)
}

// SetLendingPrice is a paid mutator transaction binding the contract method 0xa9686101.
//
// Solidity: function setLendingPrice(uint256 crabadaId, uint256 price) returns()
func (_Crabada *CrabadaTransactor) SetLendingPrice(opts *bind.TransactOpts, crabadaId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setLendingPrice", crabadaId, price)
}

// SetLendingPrice is a paid mutator transaction binding the contract method 0xa9686101.
//
// Solidity: function setLendingPrice(uint256 crabadaId, uint256 price) returns()
func (_Crabada *CrabadaSession) SetLendingPrice(crabadaId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingPrice(&_Crabada.TransactOpts, crabadaId, price)
}

// SetLendingPrice is a paid mutator transaction binding the contract method 0xa9686101.
//
// Solidity: function setLendingPrice(uint256 crabadaId, uint256 price) returns()
func (_Crabada *CrabadaTransactorSession) SetLendingPrice(crabadaId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetLendingPrice(&_Crabada.TransactOpts, crabadaId, price)
}

// SetLootingPercent is a paid mutator transaction binding the contract method 0x7d9aee2a.
//
// Solidity: function setLootingPercent(uint256 percent) returns()
func (_Crabada *CrabadaTransactor) SetLootingPercent(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setLootingPercent", percent)
}

// SetLootingPercent is a paid mutator transaction binding the contract method 0x7d9aee2a.
//
// Solidity: function setLootingPercent(uint256 percent) returns()
func (_Crabada *CrabadaSession) SetLootingPercent(percent *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetLootingPercent(&_Crabada.TransactOpts, percent)
}

// SetLootingPercent is a paid mutator transaction binding the contract method 0x7d9aee2a.
//
// Solidity: function setLootingPercent(uint256 percent) returns()
func (_Crabada *CrabadaTransactorSession) SetLootingPercent(percent *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetLootingPercent(&_Crabada.TransactOpts, percent)
}

// SetMinerRevengeSettings is a paid mutator transaction binding the contract method 0x6fc431ad.
//
// Solidity: function setMinerRevengeSettings(uint256 _baseLuckyWinRate, uint256 _baseModifier, uint256 _modifierSetting, uint256 _baseClosenessFactor, address _rngContract) returns()
func (_Crabada *CrabadaTransactor) SetMinerRevengeSettings(opts *bind.TransactOpts, _baseLuckyWinRate *big.Int, _baseModifier *big.Int, _modifierSetting *big.Int, _baseClosenessFactor *big.Int, _rngContract common.Address) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setMinerRevengeSettings", _baseLuckyWinRate, _baseModifier, _modifierSetting, _baseClosenessFactor, _rngContract)
}

// SetMinerRevengeSettings is a paid mutator transaction binding the contract method 0x6fc431ad.
//
// Solidity: function setMinerRevengeSettings(uint256 _baseLuckyWinRate, uint256 _baseModifier, uint256 _modifierSetting, uint256 _baseClosenessFactor, address _rngContract) returns()
func (_Crabada *CrabadaSession) SetMinerRevengeSettings(_baseLuckyWinRate *big.Int, _baseModifier *big.Int, _modifierSetting *big.Int, _baseClosenessFactor *big.Int, _rngContract common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.SetMinerRevengeSettings(&_Crabada.TransactOpts, _baseLuckyWinRate, _baseModifier, _modifierSetting, _baseClosenessFactor, _rngContract)
}

// SetMinerRevengeSettings is a paid mutator transaction binding the contract method 0x6fc431ad.
//
// Solidity: function setMinerRevengeSettings(uint256 _baseLuckyWinRate, uint256 _baseModifier, uint256 _modifierSetting, uint256 _baseClosenessFactor, address _rngContract) returns()
func (_Crabada *CrabadaTransactorSession) SetMinerRevengeSettings(_baseLuckyWinRate *big.Int, _baseModifier *big.Int, _modifierSetting *big.Int, _baseClosenessFactor *big.Int, _rngContract common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.SetMinerRevengeSettings(&_Crabada.TransactOpts, _baseLuckyWinRate, _baseModifier, _modifierSetting, _baseClosenessFactor, _rngContract)
}

// SetMiningDuration is a paid mutator transaction binding the contract method 0x4b18970a.
//
// Solidity: function setMiningDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactor) SetMiningDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setMiningDuration", duration)
}

// SetMiningDuration is a paid mutator transaction binding the contract method 0x4b18970a.
//
// Solidity: function setMiningDuration(uint128 duration) returns()
func (_Crabada *CrabadaSession) SetMiningDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetMiningDuration(&_Crabada.TransactOpts, duration)
}

// SetMiningDuration is a paid mutator transaction binding the contract method 0x4b18970a.
//
// Solidity: function setMiningDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactorSession) SetMiningDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetMiningDuration(&_Crabada.TransactOpts, duration)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool isPaused) returns()
func (_Crabada *CrabadaTransactor) SetPause(opts *bind.TransactOpts, isPaused bool) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setPause", isPaused)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool isPaused) returns()
func (_Crabada *CrabadaSession) SetPause(isPaused bool) (*types.Transaction, error) {
	return _Crabada.Contract.SetPause(&_Crabada.TransactOpts, isPaused)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool isPaused) returns()
func (_Crabada *CrabadaTransactorSession) SetPause(isPaused bool) (*types.Transaction, error) {
	return _Crabada.Contract.SetPause(&_Crabada.TransactOpts, isPaused)
}

// SetReinforceLockDuration is a paid mutator transaction binding the contract method 0x2938f7c7.
//
// Solidity: function setReinforceLockDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactor) SetReinforceLockDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setReinforceLockDuration", duration)
}

// SetReinforceLockDuration is a paid mutator transaction binding the contract method 0x2938f7c7.
//
// Solidity: function setReinforceLockDuration(uint128 duration) returns()
func (_Crabada *CrabadaSession) SetReinforceLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetReinforceLockDuration(&_Crabada.TransactOpts, duration)
}

// SetReinforceLockDuration is a paid mutator transaction binding the contract method 0x2938f7c7.
//
// Solidity: function setReinforceLockDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactorSession) SetReinforceLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetReinforceLockDuration(&_Crabada.TransactOpts, duration)
}

// SetRewardLooting is a paid mutator transaction binding the contract method 0x5ba4e7c8.
//
// Solidity: function setRewardLooting(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Crabada *CrabadaTransactor) SetRewardLooting(opts *bind.TransactOpts, craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setRewardLooting", craRewardAmount, tusRewardAmount)
}

// SetRewardLooting is a paid mutator transaction binding the contract method 0x5ba4e7c8.
//
// Solidity: function setRewardLooting(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Crabada *CrabadaSession) SetRewardLooting(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetRewardLooting(&_Crabada.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetRewardLooting is a paid mutator transaction binding the contract method 0x5ba4e7c8.
//
// Solidity: function setRewardLooting(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Crabada *CrabadaTransactorSession) SetRewardLooting(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetRewardLooting(&_Crabada.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetRewardMining is a paid mutator transaction binding the contract method 0x68dedbac.
//
// Solidity: function setRewardMining(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Crabada *CrabadaTransactor) SetRewardMining(opts *bind.TransactOpts, craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setRewardMining", craRewardAmount, tusRewardAmount)
}

// SetRewardMining is a paid mutator transaction binding the contract method 0x68dedbac.
//
// Solidity: function setRewardMining(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Crabada *CrabadaSession) SetRewardMining(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetRewardMining(&_Crabada.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetRewardMining is a paid mutator transaction binding the contract method 0x68dedbac.
//
// Solidity: function setRewardMining(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Crabada *CrabadaTransactorSession) SetRewardMining(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetRewardMining(&_Crabada.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetStepDuration is a paid mutator transaction binding the contract method 0x45228060.
//
// Solidity: function setStepDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactor) SetStepDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setStepDuration", duration)
}

// SetStepDuration is a paid mutator transaction binding the contract method 0x45228060.
//
// Solidity: function setStepDuration(uint128 duration) returns()
func (_Crabada *CrabadaSession) SetStepDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetStepDuration(&_Crabada.TransactOpts, duration)
}

// SetStepDuration is a paid mutator transaction binding the contract method 0x45228060.
//
// Solidity: function setStepDuration(uint128 duration) returns()
func (_Crabada *CrabadaTransactorSession) SetStepDuration(duration *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetStepDuration(&_Crabada.TransactOpts, duration)
}

// SettleGame is a paid mutator transaction binding the contract method 0x312d7bbc.
//
// Solidity: function settleGame(uint256 gameId) returns()
func (_Crabada *CrabadaTransactor) SettleGame(opts *bind.TransactOpts, gameId *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "settleGame", gameId)
}

// SettleGame is a paid mutator transaction binding the contract method 0x312d7bbc.
//
// Solidity: function settleGame(uint256 gameId) returns()
func (_Crabada *CrabadaSession) SettleGame(gameId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SettleGame(&_Crabada.TransactOpts, gameId)
}

// SettleGame is a paid mutator transaction binding the contract method 0x312d7bbc.
//
// Solidity: function settleGame(uint256 gameId) returns()
func (_Crabada *CrabadaTransactorSession) SettleGame(gameId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SettleGame(&_Crabada.TransactOpts, gameId)
}

// SetupTeamExpandSetting is a paid mutator transaction binding the contract method 0xf99e7fe9.
//
// Solidity: function setupTeamExpandSetting(address amuletAddress, uint256 price) returns()
func (_Crabada *CrabadaTransactor) SetupTeamExpandSetting(opts *bind.TransactOpts, amuletAddress common.Address, price *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "setupTeamExpandSetting", amuletAddress, price)
}

// SetupTeamExpandSetting is a paid mutator transaction binding the contract method 0xf99e7fe9.
//
// Solidity: function setupTeamExpandSetting(address amuletAddress, uint256 price) returns()
func (_Crabada *CrabadaSession) SetupTeamExpandSetting(amuletAddress common.Address, price *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetupTeamExpandSetting(&_Crabada.TransactOpts, amuletAddress, price)
}

// SetupTeamExpandSetting is a paid mutator transaction binding the contract method 0xf99e7fe9.
//
// Solidity: function setupTeamExpandSetting(address amuletAddress, uint256 price) returns()
func (_Crabada *CrabadaTransactorSession) SetupTeamExpandSetting(amuletAddress common.Address, price *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.SetupTeamExpandSetting(&_Crabada.TransactOpts, amuletAddress, price)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 teamId) returns()
func (_Crabada *CrabadaTransactor) StartGame(opts *bind.TransactOpts, teamId *big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "startGame", teamId)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 teamId) returns()
func (_Crabada *CrabadaSession) StartGame(teamId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.StartGame(&_Crabada.TransactOpts, teamId)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 teamId) returns()
func (_Crabada *CrabadaTransactorSession) StartGame(teamId *big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.StartGame(&_Crabada.TransactOpts, teamId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crabada *CrabadaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crabada *CrabadaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.TransferOwnership(&_Crabada.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crabada *CrabadaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crabada.Contract.TransferOwnership(&_Crabada.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8293744b.
//
// Solidity: function withdraw(address to, uint256[] crabadaIds) returns()
func (_Crabada *CrabadaTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Crabada.contract.Transact(opts, "withdraw", to, crabadaIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8293744b.
//
// Solidity: function withdraw(address to, uint256[] crabadaIds) returns()
func (_Crabada *CrabadaSession) Withdraw(to common.Address, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.Withdraw(&_Crabada.TransactOpts, to, crabadaIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8293744b.
//
// Solidity: function withdraw(address to, uint256[] crabadaIds) returns()
func (_Crabada *CrabadaTransactorSession) Withdraw(to common.Address, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Crabada.Contract.Withdraw(&_Crabada.TransactOpts, to, crabadaIds)
}

// CrabadaAddCrabadaIterator is returned from FilterAddCrabada and is used to iterate over the raw logs and unpacked data for AddCrabada events raised by the Crabada contract.
type CrabadaAddCrabadaIterator struct {
	Event *CrabadaAddCrabada // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaAddCrabadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaAddCrabada)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaAddCrabada)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaAddCrabadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaAddCrabadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaAddCrabada represents a AddCrabada event raised by the Crabada contract.
type CrabadaAddCrabada struct {
	TeamId      *big.Int
	Position    *big.Int
	CrabadaId   *big.Int
	BattlePoint *big.Int
	TimePoint   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddCrabada is a free log retrieval operation binding the contract event 0x95f06f23548778c52447db2dca1c565debe55244f3904032ac45c61df764c9a2.
//
// Solidity: event AddCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) FilterAddCrabada(opts *bind.FilterOpts) (*CrabadaAddCrabadaIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "AddCrabada")
	if err != nil {
		return nil, err
	}
	return &CrabadaAddCrabadaIterator{contract: _Crabada.contract, event: "AddCrabada", logs: logs, sub: sub}, nil
}

// WatchAddCrabada is a free log subscription operation binding the contract event 0x95f06f23548778c52447db2dca1c565debe55244f3904032ac45c61df764c9a2.
//
// Solidity: event AddCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) WatchAddCrabada(opts *bind.WatchOpts, sink chan<- *CrabadaAddCrabada) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "AddCrabada")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaAddCrabada)
				if err := _Crabada.contract.UnpackLog(event, "AddCrabada", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddCrabada is a log parse operation binding the contract event 0x95f06f23548778c52447db2dca1c565debe55244f3904032ac45c61df764c9a2.
//
// Solidity: event AddCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) ParseAddCrabada(log types.Log) (*CrabadaAddCrabada, error) {
	event := new(CrabadaAddCrabada)
	if err := _Crabada.contract.UnpackLog(event, "AddCrabada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaCloseGameIterator is returned from FilterCloseGame and is used to iterate over the raw logs and unpacked data for CloseGame events raised by the Crabada contract.
type CrabadaCloseGameIterator struct {
	Event *CrabadaCloseGame // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaCloseGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaCloseGame)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaCloseGame)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaCloseGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaCloseGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaCloseGame represents a CloseGame event raised by the Crabada contract.
type CrabadaCloseGame struct {
	GameId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCloseGame is a free log retrieval operation binding the contract event 0x827e8bff9f46048f8351964d2c871d09e4f4231513cf2fcb786649c68732e24f.
//
// Solidity: event CloseGame(uint256 gameId)
func (_Crabada *CrabadaFilterer) FilterCloseGame(opts *bind.FilterOpts) (*CrabadaCloseGameIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "CloseGame")
	if err != nil {
		return nil, err
	}
	return &CrabadaCloseGameIterator{contract: _Crabada.contract, event: "CloseGame", logs: logs, sub: sub}, nil
}

// WatchCloseGame is a free log subscription operation binding the contract event 0x827e8bff9f46048f8351964d2c871d09e4f4231513cf2fcb786649c68732e24f.
//
// Solidity: event CloseGame(uint256 gameId)
func (_Crabada *CrabadaFilterer) WatchCloseGame(opts *bind.WatchOpts, sink chan<- *CrabadaCloseGame) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "CloseGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaCloseGame)
				if err := _Crabada.contract.UnpackLog(event, "CloseGame", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCloseGame is a log parse operation binding the contract event 0x827e8bff9f46048f8351964d2c871d09e4f4231513cf2fcb786649c68732e24f.
//
// Solidity: event CloseGame(uint256 gameId)
func (_Crabada *CrabadaFilterer) ParseCloseGame(log types.Log) (*CrabadaCloseGame, error) {
	event := new(CrabadaCloseGame)
	if err := _Crabada.contract.UnpackLog(event, "CloseGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaCreateTeamIterator is returned from FilterCreateTeam and is used to iterate over the raw logs and unpacked data for CreateTeam events raised by the Crabada contract.
type CrabadaCreateTeamIterator struct {
	Event *CrabadaCreateTeam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaCreateTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaCreateTeam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaCreateTeam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaCreateTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaCreateTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaCreateTeam represents a CreateTeam event raised by the Crabada contract.
type CrabadaCreateTeam struct {
	Owner       common.Address
	TeamId      *big.Int
	CrabadaId1  *big.Int
	CrabadaId2  *big.Int
	CrabadaId3  *big.Int
	BattlePoint *big.Int
	TimePoint   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateTeam is a free log retrieval operation binding the contract event 0xe2f710515de5fe7937eb9d7ff8093cb7cd011b496de6574d80e51440784e2a71.
//
// Solidity: event CreateTeam(address owner, uint256 teamId, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) FilterCreateTeam(opts *bind.FilterOpts) (*CrabadaCreateTeamIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "CreateTeam")
	if err != nil {
		return nil, err
	}
	return &CrabadaCreateTeamIterator{contract: _Crabada.contract, event: "CreateTeam", logs: logs, sub: sub}, nil
}

// WatchCreateTeam is a free log subscription operation binding the contract event 0xe2f710515de5fe7937eb9d7ff8093cb7cd011b496de6574d80e51440784e2a71.
//
// Solidity: event CreateTeam(address owner, uint256 teamId, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) WatchCreateTeam(opts *bind.WatchOpts, sink chan<- *CrabadaCreateTeam) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "CreateTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaCreateTeam)
				if err := _Crabada.contract.UnpackLog(event, "CreateTeam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateTeam is a log parse operation binding the contract event 0xe2f710515de5fe7937eb9d7ff8093cb7cd011b496de6574d80e51440784e2a71.
//
// Solidity: event CreateTeam(address owner, uint256 teamId, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) ParseCreateTeam(log types.Log) (*CrabadaCreateTeam, error) {
	event := new(CrabadaCreateTeam)
	if err := _Crabada.contract.UnpackLog(event, "CreateTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Crabada contract.
type CrabadaDepositIterator struct {
	Event *CrabadaDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaDeposit represents a Deposit event raised by the Crabada contract.
type CrabadaDeposit struct {
	CrabadaId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4bcc17093cdf51079c755de089be5a85e70fa374ec656c194480fbdcda224a53.
//
// Solidity: event Deposit(uint256 crabadaId, address owner)
func (_Crabada *CrabadaFilterer) FilterDeposit(opts *bind.FilterOpts) (*CrabadaDepositIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &CrabadaDepositIterator{contract: _Crabada.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4bcc17093cdf51079c755de089be5a85e70fa374ec656c194480fbdcda224a53.
//
// Solidity: event Deposit(uint256 crabadaId, address owner)
func (_Crabada *CrabadaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CrabadaDeposit) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaDeposit)
				if err := _Crabada.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x4bcc17093cdf51079c755de089be5a85e70fa374ec656c194480fbdcda224a53.
//
// Solidity: event Deposit(uint256 crabadaId, address owner)
func (_Crabada *CrabadaFilterer) ParseDeposit(log types.Log) (*CrabadaDeposit, error) {
	event := new(CrabadaDeposit)
	if err := _Crabada.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaExpandTeamIterator is returned from FilterExpandTeam and is used to iterate over the raw logs and unpacked data for ExpandTeam events raised by the Crabada contract.
type CrabadaExpandTeamIterator struct {
	Event *CrabadaExpandTeam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaExpandTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaExpandTeam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaExpandTeam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaExpandTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaExpandTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaExpandTeam represents a ExpandTeam event raised by the Crabada contract.
type CrabadaExpandTeam struct {
	Player             common.Address
	Amount             *big.Int
	NewTeamAmountLimit *big.Int
	Cost               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExpandTeam is a free log retrieval operation binding the contract event 0xb80a6b280a8c9c2877b89b4b236bad922aeba236cab8bd2dbe84d1012b9dbc0f.
//
// Solidity: event ExpandTeam(address player, uint256 amount, uint256 newTeamAmountLimit, uint256 cost)
func (_Crabada *CrabadaFilterer) FilterExpandTeam(opts *bind.FilterOpts) (*CrabadaExpandTeamIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "ExpandTeam")
	if err != nil {
		return nil, err
	}
	return &CrabadaExpandTeamIterator{contract: _Crabada.contract, event: "ExpandTeam", logs: logs, sub: sub}, nil
}

// WatchExpandTeam is a free log subscription operation binding the contract event 0xb80a6b280a8c9c2877b89b4b236bad922aeba236cab8bd2dbe84d1012b9dbc0f.
//
// Solidity: event ExpandTeam(address player, uint256 amount, uint256 newTeamAmountLimit, uint256 cost)
func (_Crabada *CrabadaFilterer) WatchExpandTeam(opts *bind.WatchOpts, sink chan<- *CrabadaExpandTeam) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "ExpandTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaExpandTeam)
				if err := _Crabada.contract.UnpackLog(event, "ExpandTeam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExpandTeam is a log parse operation binding the contract event 0xb80a6b280a8c9c2877b89b4b236bad922aeba236cab8bd2dbe84d1012b9dbc0f.
//
// Solidity: event ExpandTeam(address player, uint256 amount, uint256 newTeamAmountLimit, uint256 cost)
func (_Crabada *CrabadaFilterer) ParseExpandTeam(log types.Log) (*CrabadaExpandTeam, error) {
	event := new(CrabadaExpandTeam)
	if err := _Crabada.contract.UnpackLog(event, "ExpandTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaFightIterator is returned from FilterFight and is used to iterate over the raw logs and unpacked data for Fight events raised by the Crabada contract.
type CrabadaFightIterator struct {
	Event *CrabadaFight // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaFightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaFight)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaFight)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaFightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaFightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaFight represents a Fight event raised by the Crabada contract.
type CrabadaFight struct {
	GameId        *big.Int
	Turn          *big.Int
	AttackTeamId  *big.Int
	DefenseTeamId *big.Int
	SoldierId     *big.Int
	AttackTime    *big.Int
	AttackPoint   uint16
	DefensePoint  uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFight is a free log retrieval operation binding the contract event 0xc9116f8ab51c97fa0f2ebd6cc9f75395464ae81bebc69d2d468b508a460a7136.
//
// Solidity: event Fight(uint256 gameId, uint256 turn, uint256 attackTeamId, uint256 defenseTeamId, uint256 soldierId, uint256 attackTime, uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaFilterer) FilterFight(opts *bind.FilterOpts) (*CrabadaFightIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "Fight")
	if err != nil {
		return nil, err
	}
	return &CrabadaFightIterator{contract: _Crabada.contract, event: "Fight", logs: logs, sub: sub}, nil
}

// WatchFight is a free log subscription operation binding the contract event 0xc9116f8ab51c97fa0f2ebd6cc9f75395464ae81bebc69d2d468b508a460a7136.
//
// Solidity: event Fight(uint256 gameId, uint256 turn, uint256 attackTeamId, uint256 defenseTeamId, uint256 soldierId, uint256 attackTime, uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaFilterer) WatchFight(opts *bind.WatchOpts, sink chan<- *CrabadaFight) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "Fight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaFight)
				if err := _Crabada.contract.UnpackLog(event, "Fight", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFight is a log parse operation binding the contract event 0xc9116f8ab51c97fa0f2ebd6cc9f75395464ae81bebc69d2d468b508a460a7136.
//
// Solidity: event Fight(uint256 gameId, uint256 turn, uint256 attackTeamId, uint256 defenseTeamId, uint256 soldierId, uint256 attackTime, uint16 attackPoint, uint16 defensePoint)
func (_Crabada *CrabadaFilterer) ParseFight(log types.Log) (*CrabadaFight, error) {
	event := new(CrabadaFight)
	if err := _Crabada.contract.UnpackLog(event, "Fight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaLendIterator is returned from FilterLend and is used to iterate over the raw logs and unpacked data for Lend events raised by the Crabada contract.
type CrabadaLendIterator struct {
	Event *CrabadaLend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaLend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaLend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaLend represents a Lend event raised by the Crabada contract.
type CrabadaLend struct {
	Borrower  common.Address
	CrabadaId *big.Int
	GameId    *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLend is a free log retrieval operation binding the contract event 0xaf77103a98a12032881b6214985e03ff85a9c1ad6fed28f4b0f3c52e01bcc58b.
//
// Solidity: event Lend(address borrower, uint256 crabadaId, uint256 gameId, uint256 price)
func (_Crabada *CrabadaFilterer) FilterLend(opts *bind.FilterOpts) (*CrabadaLendIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "Lend")
	if err != nil {
		return nil, err
	}
	return &CrabadaLendIterator{contract: _Crabada.contract, event: "Lend", logs: logs, sub: sub}, nil
}

// WatchLend is a free log subscription operation binding the contract event 0xaf77103a98a12032881b6214985e03ff85a9c1ad6fed28f4b0f3c52e01bcc58b.
//
// Solidity: event Lend(address borrower, uint256 crabadaId, uint256 gameId, uint256 price)
func (_Crabada *CrabadaFilterer) WatchLend(opts *bind.WatchOpts, sink chan<- *CrabadaLend) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "Lend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaLend)
				if err := _Crabada.contract.UnpackLog(event, "Lend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLend is a log parse operation binding the contract event 0xaf77103a98a12032881b6214985e03ff85a9c1ad6fed28f4b0f3c52e01bcc58b.
//
// Solidity: event Lend(address borrower, uint256 crabadaId, uint256 gameId, uint256 price)
func (_Crabada *CrabadaFilterer) ParseLend(log types.Log) (*CrabadaLend, error) {
	event := new(CrabadaLend)
	if err := _Crabada.contract.UnpackLog(event, "Lend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crabada contract.
type CrabadaOwnershipTransferredIterator struct {
	Event *CrabadaOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaOwnershipTransferred represents a OwnershipTransferred event raised by the Crabada contract.
type CrabadaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crabada *CrabadaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrabadaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrabadaOwnershipTransferredIterator{contract: _Crabada.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crabada *CrabadaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrabadaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaOwnershipTransferred)
				if err := _Crabada.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crabada *CrabadaFilterer) ParseOwnershipTransferred(log types.Log) (*CrabadaOwnershipTransferred, error) {
	event := new(CrabadaOwnershipTransferred)
	if err := _Crabada.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Crabada contract.
type CrabadaPausedIterator struct {
	Event *CrabadaPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaPaused represents a Paused event raised by the Crabada contract.
type CrabadaPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Crabada *CrabadaFilterer) FilterPaused(opts *bind.FilterOpts) (*CrabadaPausedIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrabadaPausedIterator{contract: _Crabada.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Crabada *CrabadaFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrabadaPaused) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaPaused)
				if err := _Crabada.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Crabada *CrabadaFilterer) ParsePaused(log types.Log) (*CrabadaPaused, error) {
	event := new(CrabadaPaused)
	if err := _Crabada.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaRemoveCrabadaIterator is returned from FilterRemoveCrabada and is used to iterate over the raw logs and unpacked data for RemoveCrabada events raised by the Crabada contract.
type CrabadaRemoveCrabadaIterator struct {
	Event *CrabadaRemoveCrabada // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaRemoveCrabadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaRemoveCrabada)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaRemoveCrabada)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaRemoveCrabadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaRemoveCrabadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaRemoveCrabada represents a RemoveCrabada event raised by the Crabada contract.
type CrabadaRemoveCrabada struct {
	TeamId      *big.Int
	Position    *big.Int
	CrabadaId   *big.Int
	BattlePoint *big.Int
	TimePoint   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveCrabada is a free log retrieval operation binding the contract event 0x164649e040e4f90ed2500b40e25f8e4c1e7991798b4359bf2a16ea564f8ded60.
//
// Solidity: event RemoveCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) FilterRemoveCrabada(opts *bind.FilterOpts) (*CrabadaRemoveCrabadaIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "RemoveCrabada")
	if err != nil {
		return nil, err
	}
	return &CrabadaRemoveCrabadaIterator{contract: _Crabada.contract, event: "RemoveCrabada", logs: logs, sub: sub}, nil
}

// WatchRemoveCrabada is a free log subscription operation binding the contract event 0x164649e040e4f90ed2500b40e25f8e4c1e7991798b4359bf2a16ea564f8ded60.
//
// Solidity: event RemoveCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) WatchRemoveCrabada(opts *bind.WatchOpts, sink chan<- *CrabadaRemoveCrabada) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "RemoveCrabada")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaRemoveCrabada)
				if err := _Crabada.contract.UnpackLog(event, "RemoveCrabada", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveCrabada is a log parse operation binding the contract event 0x164649e040e4f90ed2500b40e25f8e4c1e7991798b4359bf2a16ea564f8ded60.
//
// Solidity: event RemoveCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Crabada *CrabadaFilterer) ParseRemoveCrabada(log types.Log) (*CrabadaRemoveCrabada, error) {
	event := new(CrabadaRemoveCrabada)
	if err := _Crabada.contract.UnpackLog(event, "RemoveCrabada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaSetLendingPriceIterator is returned from FilterSetLendingPrice and is used to iterate over the raw logs and unpacked data for SetLendingPrice events raised by the Crabada contract.
type CrabadaSetLendingPriceIterator struct {
	Event *CrabadaSetLendingPrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaSetLendingPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaSetLendingPrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaSetLendingPrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaSetLendingPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaSetLendingPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaSetLendingPrice represents a SetLendingPrice event raised by the Crabada contract.
type CrabadaSetLendingPrice struct {
	CrabadaId *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetLendingPrice is a free log retrieval operation binding the contract event 0x85937e28b96ea9433b7abff68e1485fb68b1a6a61edc4fe945676f48f12b56b9.
//
// Solidity: event SetLendingPrice(uint256 crabadaId, uint256 price)
func (_Crabada *CrabadaFilterer) FilterSetLendingPrice(opts *bind.FilterOpts) (*CrabadaSetLendingPriceIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "SetLendingPrice")
	if err != nil {
		return nil, err
	}
	return &CrabadaSetLendingPriceIterator{contract: _Crabada.contract, event: "SetLendingPrice", logs: logs, sub: sub}, nil
}

// WatchSetLendingPrice is a free log subscription operation binding the contract event 0x85937e28b96ea9433b7abff68e1485fb68b1a6a61edc4fe945676f48f12b56b9.
//
// Solidity: event SetLendingPrice(uint256 crabadaId, uint256 price)
func (_Crabada *CrabadaFilterer) WatchSetLendingPrice(opts *bind.WatchOpts, sink chan<- *CrabadaSetLendingPrice) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "SetLendingPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaSetLendingPrice)
				if err := _Crabada.contract.UnpackLog(event, "SetLendingPrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetLendingPrice is a log parse operation binding the contract event 0x85937e28b96ea9433b7abff68e1485fb68b1a6a61edc4fe945676f48f12b56b9.
//
// Solidity: event SetLendingPrice(uint256 crabadaId, uint256 price)
func (_Crabada *CrabadaFilterer) ParseSetLendingPrice(log types.Log) (*CrabadaSetLendingPrice, error) {
	event := new(CrabadaSetLendingPrice)
	if err := _Crabada.contract.UnpackLog(event, "SetLendingPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaSettleGameIterator is returned from FilterSettleGame and is used to iterate over the raw logs and unpacked data for SettleGame events raised by the Crabada contract.
type CrabadaSettleGameIterator struct {
	Event *CrabadaSettleGame // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaSettleGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaSettleGame)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaSettleGame)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaSettleGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaSettleGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaSettleGame represents a SettleGame event raised by the Crabada contract.
type CrabadaSettleGame struct {
	GameId          *big.Int
	WinnerTeamId    *big.Int
	WinnerCraReward *big.Int
	WinnerTusReward *big.Int
	LoserCraReward  *big.Int
	LoserTusReward  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSettleGame is a free log retrieval operation binding the contract event 0xf8abd91a64a8ff03c6198f7f7625d9c2b26e964546d6624e997e2d9ae84107b3.
//
// Solidity: event SettleGame(uint256 gameId, uint256 winnerTeamId, uint256 winnerCraReward, uint256 winnerTusReward, uint256 loserCraReward, uint256 loserTusReward)
func (_Crabada *CrabadaFilterer) FilterSettleGame(opts *bind.FilterOpts) (*CrabadaSettleGameIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "SettleGame")
	if err != nil {
		return nil, err
	}
	return &CrabadaSettleGameIterator{contract: _Crabada.contract, event: "SettleGame", logs: logs, sub: sub}, nil
}

// WatchSettleGame is a free log subscription operation binding the contract event 0xf8abd91a64a8ff03c6198f7f7625d9c2b26e964546d6624e997e2d9ae84107b3.
//
// Solidity: event SettleGame(uint256 gameId, uint256 winnerTeamId, uint256 winnerCraReward, uint256 winnerTusReward, uint256 loserCraReward, uint256 loserTusReward)
func (_Crabada *CrabadaFilterer) WatchSettleGame(opts *bind.WatchOpts, sink chan<- *CrabadaSettleGame) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "SettleGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaSettleGame)
				if err := _Crabada.contract.UnpackLog(event, "SettleGame", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettleGame is a log parse operation binding the contract event 0xf8abd91a64a8ff03c6198f7f7625d9c2b26e964546d6624e997e2d9ae84107b3.
//
// Solidity: event SettleGame(uint256 gameId, uint256 winnerTeamId, uint256 winnerCraReward, uint256 winnerTusReward, uint256 loserCraReward, uint256 loserTusReward)
func (_Crabada *CrabadaFilterer) ParseSettleGame(log types.Log) (*CrabadaSettleGame, error) {
	event := new(CrabadaSettleGame)
	if err := _Crabada.contract.UnpackLog(event, "SettleGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaStartGameIterator is returned from FilterStartGame and is used to iterate over the raw logs and unpacked data for StartGame events raised by the Crabada contract.
type CrabadaStartGameIterator struct {
	Event *CrabadaStartGame // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaStartGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaStartGame)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaStartGame)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaStartGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaStartGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaStartGame represents a StartGame event raised by the Crabada contract.
type CrabadaStartGame struct {
	GameId    *big.Int
	TeamId    *big.Int
	Duration  *big.Int
	CraReward *big.Int
	TusReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartGame is a free log retrieval operation binding the contract event 0x0eef6f7452b7d2ee11184579c086fb47626e796a83df2b2e16254df60ab761eb.
//
// Solidity: event StartGame(uint256 gameId, uint256 teamId, uint256 duration, uint256 craReward, uint256 tusReward)
func (_Crabada *CrabadaFilterer) FilterStartGame(opts *bind.FilterOpts) (*CrabadaStartGameIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "StartGame")
	if err != nil {
		return nil, err
	}
	return &CrabadaStartGameIterator{contract: _Crabada.contract, event: "StartGame", logs: logs, sub: sub}, nil
}

// WatchStartGame is a free log subscription operation binding the contract event 0x0eef6f7452b7d2ee11184579c086fb47626e796a83df2b2e16254df60ab761eb.
//
// Solidity: event StartGame(uint256 gameId, uint256 teamId, uint256 duration, uint256 craReward, uint256 tusReward)
func (_Crabada *CrabadaFilterer) WatchStartGame(opts *bind.WatchOpts, sink chan<- *CrabadaStartGame) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "StartGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaStartGame)
				if err := _Crabada.contract.UnpackLog(event, "StartGame", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartGame is a log parse operation binding the contract event 0x0eef6f7452b7d2ee11184579c086fb47626e796a83df2b2e16254df60ab761eb.
//
// Solidity: event StartGame(uint256 gameId, uint256 teamId, uint256 duration, uint256 craReward, uint256 tusReward)
func (_Crabada *CrabadaFilterer) ParseStartGame(log types.Log) (*CrabadaStartGame, error) {
	event := new(CrabadaStartGame)
	if err := _Crabada.contract.UnpackLog(event, "StartGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Crabada contract.
type CrabadaUnpausedIterator struct {
	Event *CrabadaUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaUnpaused represents a Unpaused event raised by the Crabada contract.
type CrabadaUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Crabada *CrabadaFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrabadaUnpausedIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrabadaUnpausedIterator{contract: _Crabada.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Crabada *CrabadaFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrabadaUnpaused) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaUnpaused)
				if err := _Crabada.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Crabada *CrabadaFilterer) ParseUnpaused(log types.Log) (*CrabadaUnpaused, error) {
	event := new(CrabadaUnpaused)
	if err := _Crabada.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaUpdateStatTeamIterator is returned from FilterUpdateStatTeam and is used to iterate over the raw logs and unpacked data for UpdateStatTeam events raised by the Crabada contract.
type CrabadaUpdateStatTeamIterator struct {
	Event *CrabadaUpdateStatTeam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaUpdateStatTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaUpdateStatTeam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaUpdateStatTeam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaUpdateStatTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaUpdateStatTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaUpdateStatTeam represents a UpdateStatTeam event raised by the Crabada contract.
type CrabadaUpdateStatTeam struct {
	TeamId      *big.Int
	BattlePoint *big.Int
	TimePoint   *big.Int
	StatVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateStatTeam is a free log retrieval operation binding the contract event 0x848c06ba8457c5757fa2725cfc8d01241bc5dbc4ebde7881cffee220fb36ce2c.
//
// Solidity: event UpdateStatTeam(uint256 teamId, uint256 battlePoint, uint256 timePoint, uint256 statVersion)
func (_Crabada *CrabadaFilterer) FilterUpdateStatTeam(opts *bind.FilterOpts) (*CrabadaUpdateStatTeamIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "UpdateStatTeam")
	if err != nil {
		return nil, err
	}
	return &CrabadaUpdateStatTeamIterator{contract: _Crabada.contract, event: "UpdateStatTeam", logs: logs, sub: sub}, nil
}

// WatchUpdateStatTeam is a free log subscription operation binding the contract event 0x848c06ba8457c5757fa2725cfc8d01241bc5dbc4ebde7881cffee220fb36ce2c.
//
// Solidity: event UpdateStatTeam(uint256 teamId, uint256 battlePoint, uint256 timePoint, uint256 statVersion)
func (_Crabada *CrabadaFilterer) WatchUpdateStatTeam(opts *bind.WatchOpts, sink chan<- *CrabadaUpdateStatTeam) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "UpdateStatTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaUpdateStatTeam)
				if err := _Crabada.contract.UnpackLog(event, "UpdateStatTeam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateStatTeam is a log parse operation binding the contract event 0x848c06ba8457c5757fa2725cfc8d01241bc5dbc4ebde7881cffee220fb36ce2c.
//
// Solidity: event UpdateStatTeam(uint256 teamId, uint256 battlePoint, uint256 timePoint, uint256 statVersion)
func (_Crabada *CrabadaFilterer) ParseUpdateStatTeam(log types.Log) (*CrabadaUpdateStatTeam, error) {
	event := new(CrabadaUpdateStatTeam)
	if err := _Crabada.contract.UnpackLog(event, "UpdateStatTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabadaWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Crabada contract.
type CrabadaWithdrawIterator struct {
	Event *CrabadaWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabadaWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabadaWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabadaWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabadaWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabadaWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabadaWithdraw represents a Withdraw event raised by the Crabada contract.
type CrabadaWithdraw struct {
	CrabadaId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(uint256 crabadaId, address owner)
func (_Crabada *CrabadaFilterer) FilterWithdraw(opts *bind.FilterOpts) (*CrabadaWithdrawIterator, error) {

	logs, sub, err := _Crabada.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &CrabadaWithdrawIterator{contract: _Crabada.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(uint256 crabadaId, address owner)
func (_Crabada *CrabadaFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CrabadaWithdraw) (event.Subscription, error) {

	logs, sub, err := _Crabada.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabadaWithdraw)
				if err := _Crabada.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(uint256 crabadaId, address owner)
func (_Crabada *CrabadaFilterer) ParseWithdraw(log types.Log) (*CrabadaWithdraw, error) {
	event := new(CrabadaWithdraw)
	if err := _Crabada.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
