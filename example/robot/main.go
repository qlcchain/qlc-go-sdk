package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/crypto/random"
	"os"
	"os/signal"
	"time"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	interval      = flag.Int("interval", 10, "send message interval")
	accounts      arrayFlags
	toAccounts    arrayFlags
	minInterval   = 5
	txAccount     []*types.Account
	txAccountSize int
	maxAmount     = 10
)

func main() {
	//--interval 10 --account a1 --account a2 --to t1 --to t2
	flag.Var(&accounts, "account", "account private key")
	flag.Var(&toAccounts, "to", "account private key")
	flag.Parse()

	if *interval < minInterval {
		fmt.Printf("invalid interval %d[%d,∞]\n", *interval, minInterval)
		return
	}

	if len(accounts) == 0 {
		fmt.Println("can not find any send account")
		return
	}
	if len(toAccounts) == 0 {
		fmt.Println("can not find any receive account")
		return
	}

	for i, a := range accounts {
		bytes, e := hex.DecodeString(a)
		if e != nil {
			fmt.Printf("can not decode (%s) at %d to Account", a, i)
			continue
		}
		account := types.NewAccount(bytes)
		fmt.Println("Tx: ", account.Address().String())
		txAccount = append(txAccount, account)
	}
	txAccountSize = len(txAccount)

	if txAccountSize < 2 {
		fmt.Printf("not enought account(%d) to send Tx\n", txAccountSize)
		return
	}

	fmt.Printf("%d Account will send Tx  every %d second(s)\n", txAccountSize, *interval)

	txDuration, _ := time.ParseDuration(fmt.Sprintf("%ds", *interval))
	txTicker := time.NewTicker(txDuration)
	defer txTicker.Stop()

	rxDuration, _ := time.ParseDuration("2m")
	rxTicker := time.NewTicker(rxDuration)
	defer rxTicker.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		select {
		case <-c:
			fmt.Println("stop ....")
			return
		case t := <-txTicker.C:
			fmt.Println("Current time: ", t)
		case r := <-rxTicker.C:
			go func() {
				//generate receive block
				for _, a := range txAccount {
					fmt.Println(a.Address().String())
				}
			}()
			fmt.Println("generate receive", r)
		}
	}
}

func getAmount() int {
	i, _ := random.Intn(maxAmount)
	return i
}

func randomAccount(a *types.Account, account []*types.Account) *types.Account {
	i, _ := random.Intn(len(account))
	tmp := account[i]
	if a != nil && tmp.Address() == a.Address() {
		return randomAccount(a, account)
	}
	return tmp
}
