package qlcchain

import (
	"fmt"

	rpc "github.com/qlcchain/jsonrpc2"
)

type QLCClient struct {
	client   *rpc.Client
	Account  *AccountApi
	Contract *ContractApi
	Ledger   *LedgerApi
	Mintage  *MintageApi
	Pledge   *PledgeApi
	Rewards  *RewardsApi
	Network  *NetApi
	SMS      *SMSApi
	Util     *UtilApi
	Destroy  *DestroyApi
	Debug    *DebugApi
}

func (c *QLCClient) Close() error {
	if c != nil && c.client != nil {
		c.client.Close()
	}
	return nil
}

// NewQLCClient creates a new client
func NewQLCClient(url string) (*QLCClient, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	return &QLCClient{
		client:   client,
		Account:  NewAccountApi(client),
		Ledger:   NewLedgerApi(client),
		SMS:      NewSMSApi(client),
		Contract: NewContractApi(client),
		Mintage:  NewMintageApi(client),
		Pledge:   NewPledgeApi(client),
		Rewards:  NewRewardApi(client),
		Network:  NewNetApi(client),
		Util:     NewUtilApi(client),
		Destroy:  NewDestroyApi(client),
		Debug:    NewDebugApi(client),
	}, nil

}

// Version returns version for sdk
func (c *QLCClient) Version() string {
	return fmt.Sprintf("%s.%s.%s", VERSION, GITREV, BUILDTIME)
}
