package qlcchain

import (
	"fmt"

	"github.com/qlcchain/go-qlc/rpc"
	"github.com/qlcchain/qlc-go-sdk/module"
)

type QLCClient struct {
	client   *rpc.Client
	Account  *module.AccountApi
	Contract *module.ContractApi
	Ledger   *module.LedgerApi
	Mintage  *module.MintageApi
	Network  *module.NetApi
	SMS      *module.SMSApi
	Wallet   *module.WalletApi
	Util     *module.UtilApi
}

func NewQLCClient(url string) (*QLCClient, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return &QLCClient{client: client, Account: module.NewAccountApi(client)}, nil
}

func (c *QLCClient) Version() string {
	return fmt.Sprintf("%s.%s.%s", VERSION, GITREV, BUILDTIME)
}
