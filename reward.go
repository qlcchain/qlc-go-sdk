package qlcchain

import (
	"math/big"

	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type RewardsApi struct {
	client *rpc.Client
}

// NewRewardApi creates reward module for client
func NewRewardApi(c *rpc.Client) *RewardsApi {
	return &RewardsApi{client: c}
}

type RewardsParam struct {
	Id     string        `json:"Id"`
	Amount types.Balance `json:"amount"`
	Self   types.Address `json:"self"`
	To     types.Address `json:"to"`
}

func (r *RewardsApi) GetUnsignedRewardData(param *RewardsParam) (types.Hash, error) {
	var hash types.Hash
	err := r.client.Call(&hash, "rewards_getUnsignedRewardData", param)
	if err != nil {
		return types.ZeroHash, err
	}
	return hash, nil
}

func (r *RewardsApi) GetUnsignedConfidantData(param *RewardsParam) (types.Hash, error) {
	var hash types.Hash
	err := r.client.Call(&hash, "rewards_getUnsignedConfidantData", param)
	if err != nil {
		return types.ZeroHash, err
	}
	return hash, nil
}

func (r *RewardsApi) GetSendRewardBlock(param *RewardsParam, sign *types.Signature) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := r.client.Call(&sb, "rewards_getSendRewardBlock", param, sign)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (r *RewardsApi) GetSendConfidantBlock(param *RewardsParam, sign *types.Signature) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := r.client.Call(&sb, "rewards_getSendConfidantBlock", param, sign)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (r *RewardsApi) GetReceiveRewardBlock(send *types.Hash) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := r.client.Call(&sb, "rewards_getReceiveRewardBlock", send)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

//func (r *RewardsApi) GetReceiveConfidantBlock(send *types.Hash) (*types.StateBlock, error) {
//	var sb types.StateBlock
//	err := r.client.Call(&sb, "rewards_getReceiveConfidantBlock", send)
//	if err != nil {
//		return nil, err
//	}
//	return &sb, nil
//}

func (r *RewardsApi) GetTotalRewards(txId string) (*big.Int, error) {
	var result *big.Int
	err := r.client.Call(&result, "rewards_getTotalRewards", txId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RewardsApi) GetConfidantRewords(confidant types.Address) (map[string]*big.Int, error) {
	var result map[string]*big.Int
	err := r.client.Call(&result, "rewards_getConfidantRewords", confidant)
	if err != nil {
		return nil, err
	}
	return result, nil
}
