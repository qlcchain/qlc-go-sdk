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
	return nil, nil
}

func (m *MintageApi) GetMintageBlock(param *api.MintageParams) (*types.StateBlock, error) {
	return nil, nil
}

func (m *MintageApi) GetRewardBlock(input *types.StateBlock) (*types.StateBlock, error) {
	return nil, nil
}

func (m *MintageApi) GetWithdrawMintageData(tokenId types.Hash) ([]byte, error) {
	return nil, nil
}
