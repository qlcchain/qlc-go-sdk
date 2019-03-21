package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
	"github.com/qlcchain/go-qlc/rpc/api"
)

type MintageApi struct {
	client *rpc.Client
}

// NewMintageApi creates mintage module for client
func NewMintageApi(c *rpc.Client) *MintageApi {
	return &MintageApi{client: c}
}

// GetMintageData returns mintage data by mintage parameters
func (m *MintageApi) GetMintageData(param *api.MintageParams) ([]byte, error) {
	var r []byte
	err := m.client.Call(&r, "mintage_getMintageData", param)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetMintageBlock returns mintage block by mintage parameters
func (m *MintageApi) GetMintageBlock(param *api.MintageParams) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := m.client.Call(&sb, "mintage_getMintageBlock", param)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

// GetRewardBlock returns reward block by mintage block
func (m *MintageApi) GetRewardBlock(input *types.StateBlock) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := m.client.Call(&sb, "mintage_getRewardBlock", input)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

// GetWithdrawMintageData returns withdraw mintage data by token id
func (m *MintageApi) GetWithdrawMintageData(tokenId types.Hash) ([]byte, error) {
	var r []byte
	err := m.client.Call(&r, "mintage_getWithdrawMintageData", tokenId)
	if err != nil {
		return nil, err
	}
	return r, nil
}
