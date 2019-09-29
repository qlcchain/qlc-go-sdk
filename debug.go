package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type DebugApi struct {
	client *rpc.Client
}

func NewDebugApi(c *rpc.Client) *DebugApi {
	return &DebugApi{client: c}
}

func (l *DebugApi) BlockCacheCount() (map[string]uint64, error) {
	var r map[string]uint64
	err := l.client.Call(&r, "debug_blockCacheCount")
	if err != nil {
		return nil, err
	}
	return r, nil
}

type APIPendingInfo struct {
	*types.PendingKey
	*types.PendingInfo
	TokenName string `json:"tokenName"`
	Timestamp int64  `json:"timestamp"`
	Used      bool   `json:"used"`
}

func (l *DebugApi) AccountPending(address types.Address) (*APIPendingInfo, error) {
	var r APIPendingInfo
	err := l.client.Call(&r, "debug_accountPending", address)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
