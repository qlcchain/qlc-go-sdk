package main

import (
	"fmt"
	"github.com/qlcchain/qlc.go"
)

func main() {
	client, err := qlcchain.NewQLCClient("")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client.Version())
}
