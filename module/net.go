package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
)

type NetApi struct {
	client *rpc.Client
}

// NewNetApi creates net module for client
func NewNetApi(c *rpc.Client) *NetApi {
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
