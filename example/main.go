package main

import (
	"fmt"
	qlcchain "github.com/qlcchain/qlc-go-sdk"
	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	//client, err := qlcchain.NewQLCClient("ws://127.0.0.1:19736")
	client, err := qlcchain.NewQLCClient("http://127.0.0.1:19735", jsonrpc2.LogMessages(printLog{}))
	if err != nil || client == nil {
		fmt.Println(err)
	}

	addr, err := client.Ledger.Accounts(20, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for idx, val := range addr {
		fmt.Println(idx, "==>", val.String())
	}

	fmt.Println(client.Version())
}

type printLog struct {
}

func (printLog) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v)
}
