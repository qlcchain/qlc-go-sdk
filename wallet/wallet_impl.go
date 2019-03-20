package wallet

import (
	"errors"
	"sync"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/ledger/db"
	qwallet "github.com/qlcchain/go-qlc/wallet"
)

var (
	cache = make(map[string]*Wallet)
	lock  = sync.RWMutex{}
)

type Wallet struct {
	qwallet.WalletStore
	dir string
}

func NewWallet(dir string) (*Wallet, error) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := cache[dir]; !ok {
		store, err := db.NewBadgerStore(dir)
		if err != nil {
			return nil, err
		}
		w := new(Wallet)
		w.WalletStore = qwallet.WalletStore{Store: store}
		w.dir = dir
		cache[dir] = w
	}
	return cache[dir], nil
}

func (w *Wallet) WalletIds() ([]types.Address, error) {
	return w.WalletStore.WalletIds()
}

// NewWalletBySeed create wallet from hex seed string
func (w *Wallet) NewWalletBySeed(seed, password string) (types.Address, error) {
	return w.WalletStore.NewWalletBySeed(seed, password)
}

// IsWalletExist check is the wallet exist by master address
func (w *Wallet) IsWalletExist(address types.Address) (bool, error) {
	return w.WalletStore.IsWalletExist(address)
}

//NewWallet create new wallet and save to db
func (w *Wallet) NewWallet() (types.Address, error) {
	return w.WalletStore.NewWallet()
}

func (w *Wallet) CurrentId() (types.Address, error) {
	return w.WalletStore.CurrentId()
}

func (w *Wallet) RemoveWallet(id types.Address) error {
	return w.WalletStore.RemoveWallet(id)
}

func (w Wallet) ChangePassword(addr types.Address, pwd string, newPwd string) error {
	session := w.WalletStore.NewSession(addr)
	b, err := session.VerifyPassword(pwd)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("password is invalid")
	}
	err = session.ChangePassword(newPwd)
	if err != nil {
		return err
	}
	return nil
}

func (ws *Wallet) Close() error {
	lock.Lock()
	defer lock.Unlock()
	err := ws.Store.Close()
	delete(cache, ws.dir)
	return err
}
