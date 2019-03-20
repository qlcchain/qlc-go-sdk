package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
	"github.com/qlcchain/go-qlc/rpc/api"
)

type MintageApi struct {
	client *rpc.Client
}

func NewMintageApi(c *rpc.Client) *MintageApi {
	return &MintageApi{client: c}
}

func (m *MintageApi) GetMintageData(param *api.MintageParams) ([]byte, error) {
	var r []byte
	err := m.client.Call(&r, "mintage_getMintageData", param)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m *MintageApi) GetMintageBlock(param *api.MintageParams) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := m.client.Call(&sb, "mintage_getMintageBlock", param)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (m *MintageApi) GetRewardBlock(input *types.StateBlock) (*types.StateBlock, error) {
	var sb types.StateBlock
	err := m.client.Call(&sb, "mintage_getRewardBlock", input)
	if err != nil {
		return nil, err
	}
	return &sb, nil
}

func (m *MintageApi) GetWithdrawMintageData(tokenId types.Hash) ([]byte, error) {
	var r []byte
	err := m.client.Call(&r, "mintage_getWithdrawMintageData", tokenId)
	if err != nil {
		return nil, err
	}
	return r, nil
}
