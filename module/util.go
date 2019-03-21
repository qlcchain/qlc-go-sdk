package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/common/util"
	"github.com/qlcchain/go-qlc/rpc"
)

type UtilApi struct {
	client *rpc.Client
}

// NewUtilApi creates unit module for client
func NewUtilApi(c *rpc.Client) *UtilApi {
	return &UtilApi{client: c}
}

// Decrypt decrypts cryptograph to raw by passphrase
func (u *UtilApi) Decrypt(cryptograph string, passphrase string) (string, error) {
	return util.Decrypt(cryptograph, passphrase)
}

// Encrypt encrypts raw to cryptograph by passphrase
func (u *UtilApi) Encrypt(raw string, passphrase string) (string, error) {
	return util.Encrypt(raw, passphrase)
}

// RawToBalance transforms QLC amount from raw to unit
func (u *UtilApi) RawToBalance(balance types.Balance, unit string) (types.Balance, error) {
	var b types.Balance
	if err := u.client.Call(&b, "util_rawToBalance", balance, unit); err != nil {
		return types.ZeroBalance, err
	}
	return b, nil
}

// RawToBalance transforms token (not QLC) amount from raw
func (u *UtilApi) RawToBalanceForToken(balance types.Balance, tokenName string) (types.Balance, error) {
	var b types.Balance
	if err := u.client.Call(&b, "util_rawToBalance", balance, "", tokenName); err != nil {
		return types.ZeroBalance, err
	}
	return b, nil
}

// RawToBalance transforms QLC amount from unit to raw
func (u *UtilApi) BalanceToRaw(balance types.Balance, unit string) (types.Balance, error) {
	var b types.Balance
	if err := u.client.Call(&b, "util_balanceToRaw", balance, unit); err != nil {
		return types.ZeroBalance, err
	}
	return b, nil
}

// RawToBalance transforms token (not QLC) amount to raw
func (u *UtilApi) BalanceToRawForToken(balance types.Balance, tokenName string) (types.Balance, error) {
	var b types.Balance
	if err := u.client.Call(&b, "util_balanceToRaw", balance, "", tokenName); err != nil {
		return types.ZeroBalance, err
	}
	return b, nil
}
