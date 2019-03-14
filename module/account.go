package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
)

type AccountApi struct {
	client *rpc.Client
}

func NewAccountApi(c *rpc.Client) *AccountApi {
	return &AccountApi{client: c}
}

func (a *AccountApi) Create(seedStr string, i *uint32) (map[string]string, error) {
	var resp map[string]string
	err := a.client.Call(&resp, "account_create")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountApi) ForPublicKey(pubStr string) (types.Address, error) {
	return types.ZeroAddress, nil
}

func (a *AccountApi) PublicKey(addr types.Address) string {
	return ""
}

func (a *AccountApi) Validate(addr string) bool {
	return false
}
