package qlcchain

import (
	"math/big"

	rpc "github.com/qlcchain/jsonrpc2"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type MinerApi struct {
	client *rpc.Client
}

type RewardParam struct {
	Coinbase     types.Address `json:"coinbase"`
	Beneficial   types.Address `json:"beneficial"`
	StartHeight  uint64        `json:"startHeight"`
	EndHeight    uint64        `json:"endHeight"`
	RewardBlocks uint64        `json:"rewardBlocks"`
	RewardAmount *big.Int      `json:"rewardAmount"`
}

type MinerAvailRewardInfo struct {
	LastEndHeight     uint64        `json:"lastEndHeight"`
	LatestBlockHeight uint64        `json:"latestBlockHeight"`
	NodeRewardHeight  uint64        `json:"nodeRewardHeight"`
	AvailStartHeight  uint64        `json:"availStartHeight"`
	AvailEndHeight    uint64        `json:"availEndHeight"`
	AvailRewardBlocks uint64        `json:"availRewardBlocks"`
	AvailRewardAmount types.Balance `json:"availRewardAmount"`
	NeedCallReward    bool          `json:"needCallReward"`
}

type MinerHistoryRewardInfo struct {
	LastEndHeight  uint64        `json:"lastEndHeight"`
	RewardBlocks   uint64        `json:"rewardBlocks"`
	RewardAmount   types.Balance `json:"rewardAmount"`
	LastRewardTime int64         `json:"lastRewardTime"`
}

// NewMinerAPI creates miner module for client
func NewMinerAPI(c *rpc.Client) *MinerApi {
	return &MinerApi{client: c}
}

// MinerAvailRewardInfo returns miner available reward info
func (m *MinerApi) GetAvailRewardInfo(coinbase types.Address) (*MinerAvailRewardInfo, error) {
	var rspData MinerAvailRewardInfo
	err := m.client.Call(&rspData, "miner_getAvailRewardInfo", coinbase)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardSendBlock returns miner contract send block
func (m *MinerApi) GetRewardSendBlock(param *RewardParam) (*types.StateBlock, error) {
	var rspData types.StateBlock
	err := m.client.Call(&rspData, "miner_getRewardSendBlock", param)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardSendBlock returns miner contract reward block
func (m *MinerApi) GetRewardRecvBlock(input *types.StateBlock) (*types.StateBlock, error) {
	var rspData types.StateBlock
	err := m.client.Call(&rspData, "miner_getRewardRecvBlock", input)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardRecvBlockBySendHash returns miner contract reward block
func (m *MinerApi) GetRewardRecvBlockBySendHash(sendHash types.Hash) (*types.StateBlock, error) {
	var rspData types.StateBlock
	err := m.client.Call(&rspData, "miner_getRewardRecvBlockBySendHash", sendHash)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRewardHistory returns miner history reward info
func (m *MinerApi) GetRewardHistory(coinbase types.Address) (*MinerHistoryRewardInfo, error) {
	var rspData MinerHistoryRewardInfo
	err := m.client.Call(&rspData, "miner_getRewardHistory", coinbase)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

func (m *MinerApi) GetRewardData(param *RewardParam) ([]byte, error) {
	var rspData []byte
	err := m.client.Call(&rspData, "miner_getRewardData", param)
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

func (m *MinerApi) UnpackRewardData(data []byte) (*RewardParam, error) {
	var rspData RewardParam
	err := m.client.Call(&rspData, "miner_unpackRewardData", data)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}
