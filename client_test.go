package qlcchain

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc/api"
	"github.com/qlcchain/go-qlc/test/mock"
	"github.com/qlcchain/qlc-go-sdk/module"
)

func TestQLCClient_Version(t *testing.T) {
	t.Log("test")
}

func TestQLCClient_GenerateBlock(t *testing.T) {
	t.Skip()
	c, err := NewQLCClient("http://127.0.0.1:9735")
	//client, err := NewQLCClient("http://47.244.138.61:9735")
	if err != nil {
		t.Fatal(err)
	}
	defer c.client.Close()

	_, sPri, _ := types.KeypairFromSeed("363CA1A23BA71D078D03A2A52CE390D5CC5AD29CF15453E0A44C6554DA1471C5", 0)
	sAccount := types.NewAccount(sPri)
	_, rPri, _ := types.KeypairFromSeed("123227955e098c68c1fa78953b03cf144b04567826577c1b8cab877b4902f345", 0)
	rAccount := types.NewAccount(rPri)

	sender := "100"
	receiver := "200"
	mHash := mock.Hash()
	fmt.Println("message hash, ", mHash)
	bp := api.APISendBlockPara{
		From:      sAccount.Address(),
		TokenName: "QLC",
		To:        rAccount.Address(),
		Amount:    types.Balance{Int: big.NewInt(int64(100))},
		Sender:    sender,
		Receiver:  receiver,
		Message:   mHash,
	}

	sendBlock, err := c.Ledger.GenerateSendBlock(&bp, func(hash types.Hash) (signatures types.Signature, e error) {
		return module.SignatureFunc(sAccount, hash)
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("send address: ", sAccount.Address())
	fmt.Println("send block: ", sendBlock.String())
	fmt.Println("hash: ", sendBlock.GetHash())
	hash, err := c.Ledger.Process(sendBlock)
	if err != nil {
		t.Fatal(err)
	}
	if hash != sendBlock.GetHash() {
		t.Fatal()
	}

	receBlock, err := c.Ledger.GenerateReceiveBlock(sendBlock, func(hash types.Hash) (signatures types.Signature, e error) {
		return module.SignatureFunc(rAccount, hash)
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("receiver address: ", rAccount.Address())
	fmt.Println("receiver block: ", receBlock.String())
	fmt.Println("hash: ", receBlock.GetHash())
	rHash, err := c.Ledger.Process(receBlock)
	if err != nil {
		t.Fatal(err)
	}
	if rHash != receBlock.GetHash() {
		t.Fatal()
	}

	b, err := c.SMS.MessageBlock(mHash)
	if err != nil {
		t.Fatal(err)
	}
	if b[0].GetHash() != sendBlock.GetHash() {
		t.Fatal()
	}
}
