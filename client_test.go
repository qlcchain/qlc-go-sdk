package qlcchain

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/qlcchain/qlc-go-sdk/pkg/random"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

func TestQLCClient_Version(t *testing.T) {
	t.Log("test")
}

func Hash() types.Hash {
	h := types.Hash{}
	_ = random.Bytes(h[:])
	return h
}

func TestQLCClient_GenerateBlock(t *testing.T) {
	t.Skip()
	c, err := NewQLCClient("http://127.0.0.1:19735")
	//client, err := NewQLCClient("http://47.244.138.61:9735")
	if err != nil {
		t.Fatal(err)
	}
	defer c.client.Close()

	_, sPri, _ := types.KeypairFromSeed("343227955e098c68c1fa78953b03cf144b04567826577c1b8cab877b4902f345", 0)
	sAccount := types.NewAccount(sPri)
	_, rPri, _ := types.KeypairFromSeed("123227955e098c68c1fa78953b03cf144b04567826577c1b8cab877b4902f345", 0)
	rAccount := types.NewAccount(rPri)

	sender := "100"
	receiver := "200"
	mHash := Hash()
	fmt.Println("message hash, ", mHash)
	bp := APISendBlockPara{
		From:      sAccount.Address(),
		TokenName: "QLC",
		To:        rAccount.Address(),
		Amount:    types.Balance{Int: big.NewInt(int64(100))},
		Sender:    sender,
		Receiver:  receiver,
		Message:   mHash,
	}

	sendBlock, err := c.Ledger.GenerateSendBlock(&bp, func(hash types.Hash) (signatures types.Signature, e error) {
		return SignatureFunc(sAccount, hash)
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
		return SignatureFunc(rAccount, hash)
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

	b, err := c.SMS.MessageBlocks(mHash)
	if err != nil {
		t.Fatal(err)
	}
	if b[0].GetHash() != sendBlock.GetHash() {
		t.Fatal()
	}
}
