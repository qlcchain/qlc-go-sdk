package qlcchain

import (
	"errors"
	"fmt"
	"math/big"

	rpc "github.com/qlcchain/jsonrpc2"
	common "github.com/qlcchain/qlc-go-sdk/pkg"
	"github.com/qlcchain/qlc-go-sdk/pkg/ed25519"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type DestroyApi struct {
	client *rpc.Client
}

// NewDestroyApi creates destroy module for client
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

func (param *APIDestroyParam) Signature(acc *types.Account) (types.Signature, error) {
	if acc.Address() == param.Owner {
		var data []byte

		data = append(data, param.Owner[:]...)
		data = append(data, param.Previous[:]...)
		data = append(data, param.Token[:]...)
		data = append(data, param.Amount.Bytes()...)
		var sig types.Signature
		copy(sig[:], ed25519.Sign(acc.PrivateKey(), data))
		return sig, nil
	} else {
		return types.ZeroSignature, fmt.Errorf("invalid address, exp: %s, act: %s",
			param.Owner.String(), acc.Address().String())
	}
}

// Verify destroy params
func (param *APIDestroyParam) Verify() (bool, error) {
	if param.Owner.IsZero() {
		return false, errors.New("invalid account")
	}

	if param.Previous.IsZero() {
		return false, errors.New("invalid previous")
	}

	if param.Token != common.GasToken() {
		return false, errors.New("invalid token to be destroyed")
	}

	if param.Amount == nil || param.Amount.Sign() <= 0 {
		return false, errors.New("invalid amount")
	}

	var data []byte

	data = append(data, param.Owner[:]...)
	data = append(data, param.Previous[:]...)
	data = append(data, param.Token[:]...)
	data = append(data, param.Amount.Bytes()...)

	return param.Owner.Verify(data, param.Sign[:]), nil
}

type SignatureParam func() (types.Signature, error)

// GetSendBlock returns destory contract send block by destory parameters
func (b *DestroyApi) GetSendBlock(param *APIDestroyParam, sign SignatureParam) (*types.StateBlock, error) {
	signature, err := sign()
	if err != nil {
		return nil, err
	}

	param.Sign = signature
	if b, err := param.Verify(); err != nil {
		return nil, err
	} else if !b {
		return nil, errors.New("invalid sign")
	}

	param.Sign = signature
	var r types.StateBlock
	err = b.client.Call(&r, "destroy_getSendBlock", param)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GetRewardBlock returns contract reward block by  destory contract send
func (b *DestroyApi) GetRewardsBlock(send *types.Hash) (*types.StateBlock, error) {
	var r types.StateBlock
	err := b.client.Call(&r, "destroy_getRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GetTotalDestroyInfo returns total amount of qgas destroyed
func (b *DestroyApi) GetTotalDestroyInfo(addr *types.Address) (types.Balance, error) {
	var r types.Balance
	err := b.client.Call(&r, "destroy_getTotalDestroyInfo", addr)
	if err != nil {
		return types.ZeroBalance, err
	}
	return r, nil
}

// GetDestroyInfoDetail returns detail info of qgas destroyed
func (b *DestroyApi) GetDestroyInfoDetail(addr *types.Address) ([]*APIDestroyInfo, error) {
	var r []*APIDestroyInfo
	err := b.client.Call(&r, "destroy_getDestroyInfoDetail", addr)
	if err != nil {
		return nil, err
	}
	return r, nil
}
