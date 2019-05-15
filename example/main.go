package main

import (
	"fmt"
	qlcchain "github.com/qlcchain/qlc-go-sdk"
)

func main() {
	client, err := qlcchain.NewQLCClient("ws://127.0.0.1:9736")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client.Version())
}
