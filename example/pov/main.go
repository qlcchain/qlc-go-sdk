package main

import (
	"flag"
	"fmt"

	qlcchain "github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	"github.com/qlcchain/qlc-go-sdk/pkg/util"
)

var flagNodeUrl string

func main() {
	flag.StringVar(&flagNodeUrl, "nodeurl", "http://127.0.0.1:19735", "RPC URL of node")

	flag.Parse()

	client, err := qlcchain.NewQLCClient(flagNodeUrl)
	if err != nil || client == nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client.Version())

	fmt.Println("============ pov header api ============")
	rspHeader, err := client.Pov.GetLatestHeader()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetLatestHeader:\n%s\n", util.ToIndentString(rspHeader))

	rspHeader, err = client.Pov.GetHeaderByHeight(rspHeader.GetHeight())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetHeaderByHeight:\n%s\n", util.ToIndentString(rspHeader))

	rspHeader, err = client.Pov.GetHeaderByHash(rspHeader.GetHash())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetHeaderByHash:\n%s\n", util.ToIndentString(rspHeader))

	fmt.Println("============ pov block api ============")
	rspBody, err := client.Pov.GetLatestBlock(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetLatestBlock:\n%s\n", util.ToIndentString(rspBody))

	rspBody, err = client.Pov.GetBlockByHeight(rspBody.GetHeight(), 0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetBlockByHeight:\n%s\n", util.ToIndentString(rspBody))

	rspBody, err = client.Pov.GetBlockByHash(rspBody.GetHash(), 0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetBlockByHash:\n%s\n", util.ToIndentString(rspBody))

	fmt.Println("============ pov statistic api ============")
	rspMining, err := client.Pov.GetMiningInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetMiningInfo:\n%s\n", util.ToIndentString(rspMining))

	rspMiners, err := client.Pov.GetMinerStats(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetMinerStats:\n%s\n", util.ToIndentString(rspMiners))

	rspHours, err := client.Pov.GetLastNHourInfo(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetLastNHourInfo:\n%s\n", util.ToIndentString(rspHours))

	fmt.Println("============ miner contract api ============")
	minerAddr, _ := types.HexToAddress("qlc_176f1aj1361y5i4yu8ccyp8xphjcbxmmu4ryh4jecnsncse1eiud7uncz8bj")
	rspRwdInfo, err := client.Miner.GetAvailRewardInfo(minerAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetAvailRewardInfo:\n%s\n", util.ToIndentString(rspRwdInfo))

	fmt.Println("============ rep contract api ============")
	repAddr, _ := types.HexToAddress("qlc_176f1aj1361y5i4yu8ccyp8xphjcbxmmu4ryh4jecnsncse1eiud7uncz8bj")
	rspRepInfo, err := client.Rep.GetAvailRewardInfo(repAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("GetAvailRewardInfo:\n%s\n", util.ToIndentString(rspRepInfo))
}
