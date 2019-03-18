package main

import (
	"github.com/qlcchain/qlc-go-sdk/example/robot/message"
	"testing"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/test/mock"
)

func Test_randomAccount(t *testing.T) {
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

func Test_getAmount(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		amount := randomAmount()
		if amount <= 0 {
			t.Fatal("err")
		}
		//t.Log(amount)
	}
}

func Test_hash(t *testing.T) {
	i, p := message.RandomPoem()
	if p == nil {
		t.Fatal("can not get poem")
	}
	t.Log(i, p)
	m := p.Message()
	h := hash(m)
	if h.IsZero() {
		t.Fatal("failed to hash message")
	}
	t.Log(m, h.String())
}

func Test_randomPhone(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_ = randomPhone()
		//t.Log(phone)
	}
}
