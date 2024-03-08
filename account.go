package greenfield

import (
	"context"
	"github.com/bnb-chain/greenfield-go-sdk/client"
	exterrors "github.com/pkg/errors"
)

const (
	//rpcAddr = "https://greenfield-chain.bnbchain.org:443"
	//chainId = "greenfield_1017-1"
	rpcAddr = "https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org:443"
	chainId = "greenfield_5600-1"
)

func GetAccountBalance(address string) (string, error) {
	cli, err := client.New(chainId, rpcAddr, client.Option{})
	if err != nil {
		return "", exterrors.WithStack(err)
	}
	ctx := context.Background()
	balance, err := cli.GetAccountBalance(ctx, address)
	if err != nil {
		return "", exterrors.WithStack(err)
	}
	return balance.Amount.String(), nil
}
