package qlcchain

import (
	"math/big"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type QGasSwapApi struct {
	client *QLCClient
}

// QGasSwapApi creates QGasSwap module for client
func NewQGasSwapAPI(c *QLCClient) *QGasSwapApi {
	return &QGasSwapApi{client: c}
}

type QGasPledgeParam struct {
	FromAddress types.Address
	Amount      types.Balance
	ToAddress   types.Address
}

func (q *QGasSwapApi) GetPledgeSendBlock(param *QGasPledgeParam) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := q.client.getClient().Call(&sb, "qgasswap_getPledgeSendBlock", param)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (q *QGasSwapApi) GetPledgeRewardBlock(sendHash types.Hash) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := q.client.getClient().Call(&sb, "qgasswap_getPledgeRewardBlock", sendHash)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

type QGasWithdrawParam struct {
	ToAddress   types.Address
	Amount      types.Balance
	FromAddress types.Address
	LinkHash    types.Hash
}

func (q *QGasSwapApi) ParseWithdrawParam(data []byte) (*QGasWithdrawParam, error) {
	var r QGasWithdrawParam
	err := q.client.getClient().Call(&r, "qgasswap_parseWithdrawParam", data)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (q *QGasSwapApi) GetWithdrawSendBlock(param *QGasWithdrawParam) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := q.client.getClient().Call(&sb, "qgasswap_getWithdrawSendBlock", param)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (q *QGasSwapApi) GetWithdrawRewardBlock(sendHash types.Hash) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := q.client.getClient().Call(&sb, "qgasswap_getWithdrawRewardBlock", sendHash)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

type QGasSwapInfo struct {
	SwapType    string
	FromAddress types.Address
	Amount      types.Balance
	ToAddress   types.Address
	SendHash    types.Hash
	RewardHash  types.Hash
	LinkHash    types.Hash
	Time        string
}

func (q *QGasSwapApi) GetAllSwapInfos(count int, offset int, isPledge *bool) ([]*QGasSwapInfo, error) {
	var r []*QGasSwapInfo
	err := q.client.getClient().Call(&r, "qgasswap_getAllSwapInfos", count, offset, isPledge)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (q *QGasSwapApi) GetSwapInfosByAddress(addr types.Address, count int, offset int, isPledge *bool) ([]*QGasSwapInfo, error) {
	var r []*QGasSwapInfo
	err := q.client.getClient().Call(&r, "qgasswap_getSwapInfosByAddress", addr, count, offset, isPledge)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (q *QGasSwapApi) GetSwapAmountByAddress(addr types.Address) (map[string]*big.Int, error) {
	r := make(map[string]*big.Int)
	err := q.client.getClient().Call(&r, "qgasswap_getSwapAmountByAddress", addr)
	if err != nil {
		return nil, err
	}
	return r, nil
}
