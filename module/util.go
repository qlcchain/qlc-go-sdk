package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
)

type UtilApi struct {
	client *rpc.Client
}

func NewUtilApi(c *rpc.Client) *UtilApi {
	return &UtilApi{client: c}
}

func (u *UtilApi) Decrypt(cryptograph string, passphrase string) (string, error) {
	return "", nil
}

func (u *UtilApi) Encrypt(raw string, passphrase string) (string, error) {
	return "", nil
}

func (u *UtilApi) RawToBalance(balance types.Balance, unit string, tokenName *string) (types.Balance, error) {
	return types.ZeroBalance, nil
}

func (u *UtilApi) BalanceToRaw(balance types.Balance, unit string, tokenName *string) (types.Balance, error) {
	return types.ZeroBalance, nil
}
