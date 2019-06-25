package wallet

import "github.com/qlcchain/qlc-go-sdk/pkg/types"

type WalletStore interface {
	NewWallet() (types.Address, error)
	NewWalletBySeed(seed, password string) (types.Address, error)
	CurrentId() (types.Address, error)
	WalletIds() ([]types.Address, error)
	IsWalletExist(address types.Address) (bool, error)
	RemoveWallet(id types.Address) error
	ChangePassword(addr types.Address, pwd string, newPwd string) error
}
