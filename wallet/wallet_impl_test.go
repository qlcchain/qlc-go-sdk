package wallet

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/config"
	"github.com/qlcchain/go-qlc/wallet"
)

func setupTestCase(t *testing.T) (func(t *testing.T), *Wallet) {
	t.Parallel()
	start := time.Now()

	dir := filepath.Join(config.QlcTestDataDir(), "wallet1", uuid.New().String())
	t.Log("setup store test case", dir)
	_ = os.RemoveAll(dir)

	store, err := NewWallet(dir)
	if err != nil {
		t.Fatal("create store failed")
	}
	t.Logf("NewWalletStore cost %s", time.Since(start))
	return func(t *testing.T) {
		err := store.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = os.RemoveAll(dir)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("teardown wallet test case: %s", time.Since(start))
	}, store
}

func TestNewWalletStore(t *testing.T) {
	dir := filepath.Join(config.QlcTestDataDir(), "wallet_test")
	store1, err := NewWallet(dir)
	if err != nil {
		t.Fatal(err)
	}
	store2, err := NewWallet(dir)
	if err != nil {
		t.Fatal(err)
	}
	if store1 == nil || store2 == nil {
		t.Fatal("error create store")
	}
	t.Logf("store1:%p, store2:%p", store1, store2)
	if !reflect.DeepEqual(store1, store2) {
		t.Fatal("store1!=store2")
	}
	defer func() {
		err := store1.Close()
		if err != nil {
			t.Fatal(err)
		}
		//store2.Close()
		_ = os.RemoveAll(dir)
	}()
}

func TestNewWalletStore2(t *testing.T) {
	dir1 := filepath.Join(config.QlcTestDataDir(), "wallet_test1")
	dir2 := filepath.Join(config.QlcTestDataDir(), "wallet_test2")
	store1, err := NewWallet(dir1)
	if err != nil {
		t.Fatal(err)
	}
	store2, err := NewWallet(dir2)
	if err != nil {
		t.Fatal(err)
	}
	if store1 == nil || store2 == nil {
		t.Fatal("error create store")
	}
	t.Logf("store1:%p, store2:%p", store1, store2)
	if reflect.DeepEqual(store1, store2) {
		t.Fatal("store1==store2")
	}
	defer func() {
		err := store1.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = store2.Close()
		if err != nil {
			t.Fatal(err)
		}
		_ = os.RemoveAll(dir1)
		_ = os.RemoveAll(dir2)
	}()
}

func TestWalletStore_NewWallet(t *testing.T) {
	teardownTestCase, store := setupTestCase(t)
	defer teardownTestCase(t)

	ids, err := store.WalletIds()
	if err != nil {
		t.Fatal(err)
	}

	if len(ids) != 0 {
		bytes, _ := json.Marshal(ids)
		t.Fatal("invalid ids", string(bytes))
	}

	id, err := store.NewWallet()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id.String())
	id2, err := store.CurrentId()
	if err != nil {
		t.Fatal(err)
	}
	if id != id2 {
		t.Fatal("id!=id2")
	}

	ids2, err := store.WalletIds()
	if err != nil {
		t.Fatal(err)
	}
	if len(ids2) != 1 || ids2[0] != id2 {
		t.Fatal("ids2 failed")
	}

	err = store.RemoveWallet(id2)
	if err != nil {
		t.Fatal(err)
	}

	ids3, err := store.WalletIds()
	if err != nil {
		t.Fatal(err)
	}

	for _, id := range ids3 {
		t.Log(id.String())
	}

	if len(ids3) > 0 {
		t.Fatal("invalid ids3 =>", len(ids3))
	}

	_, err = store.CurrentId()
	if err != wallet.ErrEmptyCurrentId {
		t.Fatal(err)
	}
}

func TestWalletStore_NewWalletBySeed(t *testing.T) {
	teardownTestCase, store := setupTestCase(t)
	defer teardownTestCase(t)
	seed1, _ := types.NewSeed()
	seed2, _ := types.NewSeed()
	account1, err := store.NewWalletBySeed(seed1.String(), "1111")
	if err != nil {
		t.Fatal(err)
	}
	account2, err := store.NewWalletBySeed(seed2.String(), "1111")
	if err != nil {
		t.Fatal(err)
	}
	accounts, err := store.WalletIds()
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) != 2 {
		t.Fatal()
	}
	if _, err := indexOf(accounts, account1); err != nil {
		t.Fatal(err)
	}
	if _, err := indexOf(accounts, account2); err != nil {
		t.Fatal(err)
	}
}

func indexOf(ids []types.Address, id types.Address) (int, error) {
	index := -1

	for i, _id := range ids {
		if id == _id {
			index = i
			break
		}
	}

	if index < 0 {
		return -1, fmt.Errorf("can not find id(%s)", id.String())
	}

	return index, nil
}
