package module

import (
	"math/big"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
	"github.com/qlcchain/go-qlc/rpc/api"
)

type RewardsApi struct {
	client *rpc.Client
}

// NewRewardApi creates reward module for client
func NewRewardApi(c *rpc.Client) *RewardsApi {
	return &RewardsApi{client: c}
}

func (r *RewardsApi) GetUnsignedRewardData(param *api.RewardsParam) (types.Hash, error) {
	var hash types.Hash
	err := r.client.Call(&hash, "rewards_getUnsignedRewardData", param)
	if err != nil {
		return types.ZeroHash, err
	}
	return hash, nil
}

func (r *RewardsApi) GetUnsignedConfidantData(param *api.RewardsParam) (types.Hash, error) {
	var hash types.Hash
	err := r.client.Call(&hash, "rewards_getUnsignedConfidantData", param)
	if err != nil {
		return types.ZeroHash, err
	}
	return hash, nil
}

func (r *RewardsApi) GetSendRewardBlock(param *api.RewardsParam, sign *types.Signature) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := r.client.Call(&sb, "rewards_getSendRewardBlock", param, sign)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (r *RewardsApi) GetSendConfidantBlock(param *api.RewardsParam, sign *types.Signature) (*types.StateBlock, error) {
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
