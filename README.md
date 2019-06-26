# QLC Chain Golang SDK
[![Build Status](https://travis-ci.com/qlcchain/qlc-go-sdk.svg?branch=master)](https://travis-ci.com/qlcchain/qlc-go-sdk)
[![codecov](https://codecov.io/gh/qlcchain/qlc-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/qlcchain/qlc-go-sdk)
[![GoDoc](https://godoc.org/github.com/qlcchain/qlc-go-sdk?status.svg)](https://godoc.org/github.com/qlcchain/qlc-go-sdk)

QLC Chain Golang API

## Example

```go
func main() {
	//client, err := qlcchain.NewQLCClient("http://127.0.0.1:19736")
	client, err := qlcchain.NewQLCClient("ws://127.0.0.1:19735", jsonrpc2.LogMessages(printLog{}))
	if err != nil || client == nil {
		fmt.Println(err)
		return
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
	fmt.Printf(format, v...)
}

// Output:
// 0 ==> qlc_13117gep55b5xpa7nbaz71s1ykz36bfqc6bieyzycif3ceykg4dtsmo19867
// 1 ==> qlc_14hshh15cduhhgcus9oqcbptw1qa3iwma4nujow3cuow5ni4df16gtirkb35
// 2 ==> qlc_18mbwzt7kjks1ydzk53hw6xropyz3mb4dgwq75tyzy4pcuc4mso1635mfdfz
// 3 ==> qlc_1b1dm6g17a5xrb8wtkas6pn4xz4m6uiq4nahxckhrob876t4oh8scawbhx8s
// 4 ==> qlc_1dzyqpd7h8ag9mamthxq695s38f71i8icjm6yn7y98rinasukhcb9tkbaqx9
// 5 ==> qlc_1gnggt8b6cwro3b4z9gootipykqd6x5gucfd7exsi4xqkryiijciegfhon4u
// 6 ==> qlc_1kk5xst583y8hpn9c48ruizs5cxprdeptw6s5wm6ezz6i1h5srpz3mnjgxao
// 7 ==> qlc_1mnw9gbzdaxz7sz18pyjcffiqaocxnunfdtu1u3fc4wjkib97rp1wcdw6ato
// 8 ==> qlc_1p11fp649uuan5ib9rpprd1dad3ue9qcqban1y8kwagdu56eea44nhq8do8o
// 9 ==> qlc_1pr8ojutnibmj4ej5aptng546hnqc4webd89tb4b31jz4tyqqots5ne6p553
// 10 ==> qlc_1s6rb1wr74r747b8k6wx6m8swgppef46ccnixy1zgtejhfosxaro15x1s8ab
// 11 ==> qlc_1u1d7mgo8hq5nad8jwesw6azfk53a31ge5minwxdfk8t1fqknypqgk8mi3z7
// 12 ==> qlc_1wyei6waj76k4b38prdubc19sr8dync5pz996cwufiy8ksxz6oshudred7q6
// 13 ==> qlc_1zbo3axuh166w4tno77hqye3n1nx5kehzhn8z71xixig7b5ggfxfyfi7f3er
// 14 ==> qlc_34yujh9i5kewnoqjs1s1fyj67u5tkifew4bc1gec4ftwktwgzbhkpjc7t4ge
// 15 ==> qlc_35ejpaokgu514segi1frsiekhpepbznhybe737qq1peczn4yb9hyic9uipbe
// 16 ==> qlc_361j3uiqdkjrzirttrpu9pn7eeussymty4rz4gifs9ijdx1p46xnpu3je7sy
// 17 ==> qlc_38um9p7z54koeu84cqsgmtjhdnf99jrkyzmh1rnky98wmt46ogboedahnsda
// 18 ==> qlc_3947hepb6ipq1m8b1jdbi6h7te3epqiakeb1j59ppwmi7wnj5optopsdgo5g
// 19 ==> qlc_39grwmri6nwwdfyjqpyumexm4px758ssg7njonfk8qqaxyrhn6kqjocn8scy
}
```

## Build tag
```
// for testnet
go build -tags testnet -o build/testnet_example example/main.go

// for mainnet
go build -o build/example example/main.go
```

## [License](https://github.com/qlcchain/qlc-go-sdk/blob/master/LICENSE)

MIT Copyright (c) 2019 QLC Chain
