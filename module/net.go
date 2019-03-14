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

func (q *NetApi) OnlineRepresentatives() []types.Address {
	return nil
}
