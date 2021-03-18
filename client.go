package qlcchain

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"sync"
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
	QGasSwap      *QGasSwapApi
	ctx           context.Context
	cancel        context.CancelFunc
	endpoint      string
	isWsConnected bool
	mutex         sync.RWMutex
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
		ctx:           ctx,
		cancel:        cancel,
		endpoint:      endpoint,
		isWsConnected: isWsConnected(endpoint),
		mutex:         sync.RWMutex{},
	}
	c.Account = NewAccountAPI(c)
	c.Ledger = NewLedgerAPI(endpoint, c)
	c.Contract = NewContractAPI(c)
	c.Mintage = NewMintageAPI(c)
	c.Pledge = NewPledgeAPI(c)
	c.Rewards = NewRewardAPI(c)
	c.Network = NewNetAPI(c)
	c.Util = NewUtilAPI(c)
	c.Destroy = NewDestroyAPI(c)
	c.Debug = NewDebugAPI(c)
	c.Pov = NewPovAPI(endpoint, c)
	c.Miner = NewMinerAPI(c)
	c.Rep = NewRepAPI(c)
	c.Settlement = NewSettlementAPI(c)
	c.Privacy = NewPrivacyAPI(c)
	c.DoDBilling = NewDoDBillingApi(c)
	c.DoDSettlement = NewDoDSettlementAPI(c)
	c.QGasSwap = NewQGasSwapAPI(c)

	c.wsConnected()
	return c, nil
}

// Version returns version for sdk
func (c *QLCClient) Version() string {
	return fmt.Sprintf("%s.%s.%s", VERSION, GITREV, BUILDTIME)
}

func (c *QLCClient) wsConnected() {
	if c.isWsConnected {
		go func() {
			cTicker := time.NewTicker(5 * time.Second)
			for {
				select {
				case <-cTicker.C:
					_, err := c.Ledger.ChainToken()
					if err != nil {
						c.mutex.Lock()
						c.client.Close()
						client, err := rpc.Dial(c.endpoint)
						if err == nil {
							c.client = client
						}
						c.mutex.Unlock()
					}
				case <-c.ctx.Done():
					return
				}
			}
		}()
	}
}

func (c *QLCClient) getClient() *rpc.Client {
	if c.isWsConnected {
		c.mutex.RLock()
		defer c.mutex.RUnlock()
		return c.client
	} else {
		return c.client
	}
}

func isWsConnected(endpoint string) bool {
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	if u.Scheme == "ws" || u.Scheme == "wss" {
		return true
	}
	return false
}
