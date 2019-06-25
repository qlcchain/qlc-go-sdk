package main

import (
	"fmt"

	qlcchain "github.com/qlcchain/qlc-go-sdk"
)

func main() {
	client, err := qlcchain.NewQLCClient("ws://127.0.0.1:19736")
	if err != nil {
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
