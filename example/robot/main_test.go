package main

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/test/mock"
	"github.com/qlcchain/qlc-go-sdk/example/robot/message"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_randomAccount(t *testing.T) {
	var accounts []*types.Account

	for i := 0; i < 10; i++ {
		a := mock.Account()
		accounts = append(accounts, a)
	}

	pool := newAccountPool(accounts)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			acc1 := pool.Get()
			acc2 := pool.Get()
			t.Logf("acc1: %s, acc2: %s\n", acc1.Address().String(), acc2.Address().String())
			time.Sleep(10 * time.Millisecond)
			pool.Put(acc1)
			pool.Put(acc2)
		}
	}()
	wg.Wait()
}

func Test_getAmount(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		amount := randomAmount()
		if amount.IsZero() {
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

func TestPhonePool(t *testing.T) {
	pool := newResourcePool(func() interface{} {
		return randomPhone()
	})

	go func() {
		for i := 0; i < 30; i++ {
			pool.Put(strconv.Itoa(i))
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//time.Sleep(10 * time.Millisecond)
			s1 := pool.Get().(string)
			s2 := pool.Get().(string)
			t.Log("g1: s1=>", s1)
			t.Log("g1: s2=>", s2)
			pool.Put(s1)
			//pool.Put(s2)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//time.Sleep(10 * time.Millisecond)
			t.Log("g2: ", pool.Get().(string))
		}
	}()

	wg.Wait()

}
