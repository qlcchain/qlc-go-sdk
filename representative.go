package qlcchain

import (
	"math/big"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type RepApi struct {
	client *QLCClient
}

type RepRewardParam struct {
	Account      types.Address `json:"account"`
	Beneficial   types.Address `json:"beneficial"`
	StartHeight  uint64        `json:"startHeight"`
	EndHeight    uint64        `json:"endHeight"`
	RewardBlocks uint64        `json:"rewardBlocks"`
	RewardAmount *big.Int      `json:"rewardAmount"`
}

type RepAvailRewardInfo struct {
	LastEndHeight     uint64        `json:"lastEndHeight"`
	LatestBlockHeight uint64        `json:"latestBlockHeight"`
	NodeRewardHeight  uint64        `json:"nodeRewardHeight"`
	AvailStartHeight  uint64        `json:"availStartHeight"`
	AvailEndHeight    uint64        `json:"availEndHeight"`
	AvailRewardBlocks uint64        `json:"availRewardBlocks"`
	AvailRewardAmount types.Balance `json:"availRewardAmount"`
	NeedCallReward    bool          `json:"needCallReward"`
}

type RepHistoryRewardInfo struct {
	LastEndHeight  uint64        `json:"lastEndHeight"`
	RewardBlocks   uint64        `json:"rewardBlocks"`
	RewardAmount   types.Balance `json:"rewardAmount"`
	LastRewardTime int64         `json:"lastRewardTime"`
}

// NewRepAPI creates representative module for client
func NewRepAPI(c *QLCClient) *RepApi {
	return &RepApi{client: c}
}

// GetAvailRewardInfo returns representative available reward info
func (r *RepApi) GetAvailRewardInfo(account types.Address) (*RepAvailRewardInfo, error) {
	var rspData RepAvailRewardInfo
	err := r.client.getClient().Call(&rspData, "rep_getAvailRewardInfo", account)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardSendBlock returns representative contract send block
func (r *RepApi) GetRewardSendBlock(param *RepRewardParam) (*types.StateBlock, error) {
	var rspData types.StateBlock
	err := r.client.getClient().Call(&rspData, "rep_getRewardSendBlock", param)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardSendBlock returns representative contract reward block
func (r *RepApi) GetRewardRecvBlock(input *types.StateBlock) (*types.StateBlock, error) {
	var rspData types.StateBlock
	err := r.client.getClient().Call(&rspData, "rep_getRewardRecvBlock", input)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardRecvBlockBySendHash returns representative contract reward block
func (r *RepApi) GetRewardRecvBlockBySendHash(sendHash types.Hash) (*types.StateBlock, error) {
	var rspData types.StateBlock
	err := r.client.getClient().Call(&rspData, "rep_getRewardRecvBlockBySendHash", sendHash)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardHistory returns representative history reward info
func (r *RepApi) GetRewardHistory(account types.Address) (*RepHistoryRewardInfo, error) {
	var rspData RepHistoryRewardInfo
	err := r.client.getClient().Call(&rspData, "rep_getRewardHistory", account)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}
