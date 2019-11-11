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
	Pov      *PovApi
	Miner    *MinerApi
	Rep      *RepApi
}

func (c *QLCClient) Close() error {
	if c != nil && c.client != nil {
		c.client.Close()
		c.Ledger.Stop()
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
		Account:  NewAccountAPI(client),
		Ledger:   NewLedgerAPI(url, client),
		SMS:      NewSMSAPI(client),
		Contract: NewContractAPI(client),
		Mintage:  NewMintageAPI(client),
		Pledge:   NewPledgeAPI(client),
		Rewards:  NewRewardAPI(client),
		Network:  NewNetAPI(client),
		Util:     NewUtilAPI(client),
		Destroy:  NewDestroyAPI(client),
		Debug:    NewDebugAPI(client),
		Pov:      NewPovAPI(url, client),
		Miner:    NewMinerAPI(client),
		Rep:      NewRepAPI(client),
	}, nil

}

// Version returns version for sdk
func (c *QLCClient) Version() string {
	return fmt.Sprintf("%s.%s.%s", VERSION, GITREV, BUILDTIME)
}
