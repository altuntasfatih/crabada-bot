package chain

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type Client struct {
	*ethclient.Client
}

type Service interface {
	IsTxSuccess(hash common.Hash, ctx context.Context) error
	GetNonce(address common.Hash, ctx context.Context) (*big.Int, error)
}

func NewChainService(chainAddress string) (*Client, error) {
	client, err := ethclient.Dial(chainAddress)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) IsTxSuccess(hash common.Hash, ctx context.Context) error {

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return errors.New("context is cancelled")
		case <-ticker.C:
			txResult, err := c.TransactionReceipt(ctx, hash)
			if err != nil {
				if err == ethereum.NotFound {
					continue
				} else {
					return errors.New(fmt.Sprintf("query of transaciton is failed, txHash: %s", hash.String()))
				}
			}

			if txResult.Status == 1 {
				return nil
			}
			return errors.New(fmt.Sprintf("transaction failed: txHash: %s", hash.String()))
		}
	}
}

func (c *Client) GetNonce(address common.Hash, ctx context.Context) (*big.Int, error) {
	txCount, err := c.TransactionCount(ctx, address)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(txCount + 1)), nil
}