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
	Util     *module.UtilApi
}

// NewQLCClient creates a new client
func NewQLCClient(url string) (*QLCClient, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return &QLCClient{client: client, Account: module.NewAccountApi(client),
		Ledger:   module.NewLedgerApi(client),
		SMS:      module.NewSMSApi(client),
		Contract: module.NewContractApi(client),
		Mintage:  module.NewMintageApi(client),
		Network:  module.NewNetApi(client),
		Util:     module.NewUtilApi(client),
	}, nil
}

// Version returns version for sdk
func (c *QLCClient) Version() string {
	return fmt.Sprintf("%s.%s.%s", VERSION, GITREV, BUILDTIME)
}
