package qlcchain

import (
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type NetApi struct {
	client *QLCClient
}

// NewNetApi creates net module for client
func NewNetApi(c *QLCClient) *NetApi {
	return &NetApi{client: c}
}

// OnlineRepresentatives returns representatives that online at this moment
func (q *NetApi) OnlineRepresentatives() ([]types.Address, error) {
	var addrs []types.Address
	if err := q.client.Call(&addrs, "net_onlineRepresentatives"); err != nil {
		return nil, err
	}
	return addrs, nil
}
