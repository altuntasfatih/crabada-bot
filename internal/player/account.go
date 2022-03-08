package player

import (
	"context"
	"crabada-bot/config"
	"crabada-bot/pkg/constant"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	"math/big"
)

type Account interface {
	GenerateTransactionOpt(ctx context.Context) *bind.TransactOpts
}

type account struct {
	Address    common.Address
	publicKey  *ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey

	gasTip    *big.Int
	gasFeeCap *big.Int

	SignerFn bind.SignerFn
	nonce    *big.Int
}

func NewAccount(config *config.Config) *account {
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal().Msg("error casting public key to ECDSA")
	}
	// 300
	// 200+ 50 < 300
	//(baseFee+MinerTip < Limit)
	signer := types.NewLondonSigner(big.NewInt(constant.ChainId))
	return &account{
		Address:    crypto.PubkeyToAddress(*publicKeyECDSA),
		privateKey: privateKey,
		publicKey:  publicKeyECDSA,
		gasTip:     big.NewInt(int64(config.GasTip * 1000000000)),
		gasFeeCap:  big.NewInt(int64(config.GasFeeCap * 1000000000)),
		SignerFn: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
}

func (a *account) GenerateTransactionOpt(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:      a.Address,
		GasLimit:  constant.GasLimit,
		GasTipCap: a.gasTip,
		GasFeeCap: a.gasFeeCap,
		Value:     constant.ZeroValue,
		Context:   ctx,
		Signer:    a.SignerFn,
	}
}