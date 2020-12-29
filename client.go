package qlcchain

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	rpc "github.com/qlcchain/jsonrpc2"
)

type QLCClient struct {
	client        *rpc.Client
	Account       *AccountApi
	Contract      *ContractApi
	Ledger        *LedgerApi
	Mintage       *MintageApi
	Pledge        *PledgeApi
	Rewards       *RewardsApi
	Network       *NetApi
	Util          *UtilApi
	Destroy       *DestroyApi
	Debug         *DebugApi
	Pov           *PovApi
	Miner         *MinerApi
	Rep           *RepApi
	Settlement    *SettlementAPI
	Privacy       *PrivacyApi
	DoDBilling    *DoDBillingAPI
	DoDSettlement *DoDSettlementAPI
	ctx           context.Context
	cancel        context.CancelFunc
	endpoint      string
}

func (c *QLCClient) Close() error {
	if c != nil && c.client != nil {
		c.cancel()
		c.client.Close()
		c.Ledger.Stop()
	}
	return nil
}

// NewQLCClient creates a new client
func NewQLCClient(endpoint string) (*QLCClient, error) {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, fmt.Errorf("dial: %s", err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	c := &QLCClient{
		client:        client,
		Account:       NewAccountAPI(client),
		Ledger:        NewLedgerAPI(endpoint, client),
		Contract:      NewContractAPI(client),
		Mintage:       NewMintageAPI(client),
		Pledge:        NewPledgeAPI(client),
		Rewards:       NewRewardAPI(client),
		Network:       NewNetAPI(client),
		Util:          NewUtilAPI(client),
		Destroy:       NewDestroyAPI(client),
		Debug:         NewDebugAPI(client),
		Pov:           NewPovAPI(endpoint, client),
		Miner:         NewMinerAPI(client),
		Rep:           NewRepAPI(client),
		Settlement:    NewSettlementAPI(client),
		Privacy:       NewPrivacyAPI(client),
		DoDBilling:    NewDoDBillingApi(client),
		DoDSettlement: NewDoDSettlementAPI(client),
		ctx:           ctx,
		cancel:        cancel,
		endpoint:      endpoint,
	}
	c.wsConnected()
	return c, nil
}

// Version returns version for sdk
func (c *QLCClient) Version() string {
	return fmt.Sprintf("%s.%s.%s", VERSION, GITREV, BUILDTIME)
}

func (c *QLCClient) wsConnected() {
	u, err := url.Parse(c.endpoint)
	if err != nil {
		log.Fatal(err)
	}
	if u.Scheme == "ws" || u.Scheme == "wss" {
		go func() {
			cTicker := time.NewTicker(5 * time.Second)
			for {
				select {
				case <-cTicker.C:
					_, err := c.Ledger.Tokens()
					if err != nil {
						client, err := rpc.Dial(c.endpoint)
						if err == nil {
							c.client = client
						}
					}
				case <-c.ctx.Done():
					return
				}
			}
		}()
	}
}
