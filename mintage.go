package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type MintageApi struct {
	client *rpc.Client
}

type MintageParams struct {
	SelfAddr    types.Address `json:"selfAddr"`
	PrevHash    types.Hash    `json:"prevHash"`
	TokenName   string        `json:"tokenName"`
	TokenSymbol string        `json:"tokenSymbol"`
	TotalSupply string        `json:"totalSupply"`
	Decimals    uint8         `json:"decimals"`
	Beneficial  types.Address `json:"beneficial"`
	NEP5TxId    string        `json:"nep5TxId"`
}

type WithdrawParams struct {
	SelfAddr types.Address `json:"selfAddr"`
	TokenId  types.Hash    `json:"tokenId"`
}

// NewMintageAPI creates mintage module for client
func NewMintageAPI(c *rpc.Client) *MintageApi {
	return &MintageApi{client: c}
}

// GetMintageData returns mintage data by mintage parameters
func (m *MintageApi) GetMintageData(param *MintageParams) ([]byte, error) {
	var r []byte
	err := m.client.Call(&r, "mintage_getMintageData", param)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetMintageBlock returns mintage block by mintage parameters
func (m *MintageApi) GetMintageBlock(param *MintageParams) (*types.StateBlock, error) {
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

// GetWithdrawMintageBlock returns withdraw mintage block by withdraw parameters
func (m *MintageApi) GetWithdrawMintageBlock(param *WithdrawParams) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := m.client.Call(&sb, "mintage_getWithdrawMintageBlock", param)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

// GetWithdrawRewardBlock returns withdraw mintage block by mintage block
func (m *MintageApi) GetWithdrawRewardBlock(input *types.StateBlock) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := m.client.Call(&sb, "mintage_getWithdrawRewardBlock", input)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (m *MintageApi) ParseTokenInfo(data []byte) (*types.TokenInfo, error) {
	var ti types.TokenInfo
	err := m.client.Call(&ti, "mintage_parseTokenInfo", data)
	if err != nil {
		return nil, err
	}
	return &ti, nil
}
