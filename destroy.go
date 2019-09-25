package qlcchain

import (
	"math/big"

	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type DestroyApi struct {
	client *rpc.Client
}

// NewSMSApi creates sms module for client
func NewDestroyApi(c *rpc.Client) *DestroyApi {
	return &DestroyApi{client: c}
}

type APIDestroyParam struct {
	Owner    types.Address   `json:"owner"`
	Previous types.Hash      `json:"previous"`
	Token    types.Hash      `json:"token"`
	Amount   *big.Int        `json:"amount"`
	Sign     types.Signature `json:"signature"`
}

type APIDestroyInfo struct {
	Owner     types.Address `json:"owner"`
	Previous  types.Hash    `json:"previous"`
	Token     types.Hash    `json:"token"`
	Amount    *big.Int      `json:"amount"`
	TimeStamp int64         `json:"timestamp"`
}

func (b *DestroyApi) GetSendBlackHoleBlock(param *APIDestroyParam) (*types.StateBlock, error) {
	var r types.StateBlock
	err := b.client.Call(&r, "destroy_getSendBlackHoleBlock", param)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (b *DestroyApi) GetReceiveBlackHoleBlock(send *types.Hash) (*types.StateBlock, error) {
	var r types.StateBlock
	err := b.client.Call(&r, "destroy_getReceiveBlackHoleBlock", send)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (b *DestroyApi) GetTotalDestroyInfo(addr *types.Address) (types.Balance, error) {
	var r types.Balance
	err := b.client.Call(&r, "destroy_getTotalDestroyInfo", addr)
	if err != nil {
		return types.ZeroBalance, err
	}
	return r, nil
}

func (b *DestroyApi) GetDestroyInfoDetail(addr *types.Address) ([]*APIDestroyInfo, error) {
	var r []*APIDestroyInfo
	err := b.client.Call(&r, "destroy_getDestroyInfoDetail", addr)
	if err != nil {
		return nil, err
	}
	return r, nil
}
