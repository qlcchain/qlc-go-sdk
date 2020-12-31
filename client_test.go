package qlcchain

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

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
	c, err := NewQLCClient("ws://127.0.0.1:29736")
	//client, err := NewQLCClient("http://47.244.138.61:9735")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	_, sPri, _ := types.KeypairFromSeed("46b31acd0a3bf072e7bea611a86074e7afae5ff95610f5f870208f2fd9357418", 0)
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
	fmt.Println("send address: ", sAccount.Address())

	sendBlock, err := c.Ledger.GenerateSendBlock(&bp, func(hash types.Hash) (signatures types.Signature, e error) {
		return SignatureFunc(sAccount, hash)
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("send block: ", sendBlock.String())
	fmt.Println("send hash: ", sendBlock.GetHash())
	// hash, err := c.Ledger.Process(sendBlock)
	//if err != nil {
	//	t.Fatal(err)
	//}
	h, err := c.Ledger.Process(sendBlock)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	if b, err := c.Ledger.BlockConfirmedStatus(h); err != nil || !b {
		t.Fatal(err)
	}
	// if hash != sendBlock.GetHash() {
	//	t.Fatal()
	//}
	fmt.Println("receiver address: ", rAccount.Address())

	receBlock, err := c.Ledger.GenerateReceiveBlock(sendBlock, func(hash types.Hash) (signatures types.Signature, e error) {
		return SignatureFunc(rAccount, hash)
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("receiver block: ", receBlock.String())
	fmt.Println("hash: ", receBlock.GetHash())
	// rHash, err := c.Ledger.Process(receBlock)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if rHash != receBlock.GetHash() {
	//	t.Fatal()
	//}
	b, err := c.Ledger.Process(receBlock)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	if b, err := c.Ledger.BlockConfirmedStatus(b); err != nil || !b {
		t.Fatal(err)
	}
	//
	//b, err := c.SMS.MessageBlocks(mHash)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if b[0].GetHash() != sendBlock.GetHash() {
	//	t.Fatal()
	//}
}

func TestQLCClient_BlockConfirmedStatus(t *testing.T) {
	t.Skip()
	c, err := NewQLCClient("http://47.244.138.61:9735")
	if err != nil {
		t.Fatal(err)
	}
	defer c.client.Close()

	hByte, err := hex.DecodeString("078199c2baa601e8d4ce49203afa015c5ed861614066b071b7f3fbf431d5462c")
	if err != nil {
		t.Fatal(err)
	}
	hash, err := types.BytesToHash(hByte)
	if err != nil {
		t.Fatal(err)
	}
	r, err := c.Ledger.BlockConfirmedStatus(hash)
	if err != nil {
		t.Fatal(err)
	}
	if !r {
		t.Fatal("block not confirmed")
	}
}

func TestNewQLCClient_Subscribe(t *testing.T) {
	t.Skip()
	c, err := NewQLCClient("ws://127.0.0.1:29736")
	if err != nil {
		t.Fatal(err)
	}
	defer c.client.Close()
	povCh := make(chan *PovApiHeader)
	sb, err := c.Pov.SubscribeNewBlock(povCh)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sb.subscribeID)
	for {
		select {
		case blk := <-povCh:
			log.Println(time.Now().String(), blk.GetHash())
		}
	}

}

func TestNewQLCClient(t *testing.T) {
	t.Skip()
	c, err := NewQLCClient("ws://127.0.0.1:29736")
	if err != nil {
		t.Fatal(err)
	}
	defer c.client.Close()
	for {
		r, err := c.Ledger.BlocksCount()
		fmt.Println("result: ", r, err)
		time.Sleep(2 * time.Second)
	}
}
