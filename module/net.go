package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
)

type NetApi struct {
	client *rpc.Client
}

func NewNetApi(c *rpc.Client) *NetApi {
	return &NetApi{client: c}
}

func (q *NetApi) OnlineRepresentatives() ([]types.Address, error) {
	var addrs []types.Address
	err := q.client.Call(&addrs, "net_onlineRepresentatives")
	if err != nil {
		return nil, err
	}
	return addrs, nil
}
