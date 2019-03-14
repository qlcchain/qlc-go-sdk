package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
)

type WalletApi struct {
	client *rpc.Client
}

func NewWalletApi(c *rpc.Client) *WalletApi {
	return &WalletApi{client: c}
}

func (w *WalletApi) GetRawKey(address types.Address, passphrase string) (map[string]string, error) {
	return nil, nil
}

func (w *WalletApi) NewSeed() (string, error) {
	return "", nil
}

func (w *WalletApi) NewWallet(passphrase string, seed *string) (types.Address, error) {
	return types.ZeroAddress, nil
}

func (w *WalletApi) List() ([]types.Address, error) {
	return nil, nil
}

func (w *WalletApi) Remove(addr types.Address) error {
	return nil
}

func (w *WalletApi) GetBalances(address types.Address, passphrase string) (map[string]types.Balance, error) {
	return nil, nil
}

func (w *WalletApi) ChangePassword(addr types.Address, pwd string, newPwd string) error {
	return nil
}
