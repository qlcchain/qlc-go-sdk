package qlcchain

import (
	"testing"
)

func TestQLCClient_Version(t *testing.T) {
	t.Log("test")
}

//func TestQLCClient_GenerateSendBlock(t *testing.T) {
//	client, err := NewQLCClient("http://47.244.138.61:9375")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	bp := api.APISendBlockPara{
//		From:      common.GenesisAccountAddress,
//		TokenName: "QLC",
//		To:        mock.Address(),
//		Amount:    types.Balance{Int: big.NewInt(int64(100))},
//		Sender:    "100",
//		Receiver:  "200",
//	}
//
//	s, err := client.Ledger.GenerateSendBlock(&bp, nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(s)
//
//}
