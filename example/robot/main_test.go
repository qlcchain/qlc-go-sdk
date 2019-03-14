package main

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/test/mock"
	"testing"
)

func TestRandomAccount(t *testing.T) {
	var accounts []*types.Account

	for i := 0; i < 10; i++ {
		a := mock.Account()
		accounts = append(accounts, a)
	}

	a := randomAccount(nil, accounts)
	if a == nil {
		t.Fatal("invalid a")
	}

	t.Log(a.Address().String())

	b := randomAccount(a, accounts)
	if b.Address() == a.Address() {
		t.Fatal("invalid a and b ", a, b)
	}
	t.Log(a.Address().String(), b.Address().String())

	b2 := randomAccount(b, accounts)

	if b.Address() == b2.Address() {
		t.Fatal("invalid a and b ", b2, b)
	}
	t.Log(b2.Address().String(), b.Address().String())
}
