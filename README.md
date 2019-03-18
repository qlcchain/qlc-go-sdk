# QLC Chain Golang SDK
[![Build Status](https://travis-ci.org/qlcchain/qlc-go-sdk.svg?branch=master)](https://travis-ci.org/qlcchain/qlc-go-sdk)
[![codecov](https://codecov.io/gh/qlcchain/qlc-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/qlcchain/qlc-go-sdk)
[![GoDoc](https://godoc.org/github.com/qlcchain/qlc-go-sdk?status.svg)](https://godoc.org/github.com/qlcchain/qlc-go-sdk)

QLC Chain Golang API

## Example

```go
func main() {
	client, err := qlcchain.NewQLCClient("ws://127.0.0.1:9736")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client.Version())
}
```

## [License](https://github.com/qlcchain/qlc-go-sdk/blob/master/LICENSE)

MIT Copyright (c) 2019 QLC Chain
