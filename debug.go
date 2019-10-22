package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"
)

type DebugApi struct {
	client *rpc.Client
}

func NewDebugAPI(c *rpc.Client) *DebugApi {
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
