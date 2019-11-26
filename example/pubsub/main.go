package main

import (
	"flag"
	"log"
	"time"

	qlcchain "github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

func main() {
	var endPoint string
	flag.StringVar(&endPoint, "endpoint", "ws://127.0.0.1:29736", "RPC Server endpoint")
	flag.Parse()

	client, err := qlcchain.NewQLCClient(endPoint)
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Close()

	ch := make(chan *types.StateBlock)
	subscribe, err := client.Ledger.NewBlock(ch)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		time.Sleep(30 * time.Second)
		if err := client.Ledger.Unsubscribe(subscribe); err != nil {
			log.Println(err)
			return
		}
	}()

	chPov := make(chan *qlcchain.PovApiHeader)
	subPov, err := client.Pov.SubscribeNewBlock(chPov)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		time.Sleep(30 * time.Second)
		if err := client.Pov.Unsubscribe(subPov); err != nil {
			log.Println(err)
			return
		}
	}()

	for {
		select {
		case result := <-ch:
			log.Println("result: ", result)
		case result := <-chPov:
			log.Println("pov result: ", result)
		case <-subscribe.Stopped:
			log.Println("subscribe stopped")
			return
		case <-subPov.Stopped:
			log.Println("pov subscribe stopped")
			return
		}
	}
}

//
// func main() {
//	flag.Parse()
//
//	url := "ws://127.0.0.1:29736"
//	ws, err := websocket.Dial(url, "", url)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	s := `{"id":1,"method":"ledger_subscribe","params":["newBlock"]}`
//	if err := websocket.Message.Send(ws, s); err != nil {
//		fmt.Println("send error, ", err)
//	}
//
//	var reply string
//	if err := websocket.Message.Receive(ws, &reply); err != nil {
//		return
//	}
//
//	var r = new(wResult)
//	err = json.Unmarshal([]byte(reply), &r)
//	if err != nil {
//		fmt.Errorf("Can not decode data: %v\n", err)
//	}
//
//	var stop bool
//	go func() {
//		time.Sleep(50 * time.Second)
//		s2 := fmt.Sprintf(`{"id":1,"method":"ledger_unsubscribe","params":["%s"]}`, r.Result)
//		if err := websocket.Message.Send(ws, s2); err != nil {
//			fmt.Println("send error, ", err)
//		}
//		//ws.Close()
//		stop = true
//	}()
//
//	for {
//		err = websocket.Message.Receive(ws, &reply)
//		if err != nil {
//			fmt.Println("read2 err:", err)
//			return
//		} else {
//			fmt.Println("receive ,", reply)
//		}
//		if stop {
//			return
//		}
//	}
// }
//
// type wResult struct {
//	Jsonrpc string `json:"jsonrpc"`
//	Id      int    `json:"id"`
//	Result  string `json:"result"`
// }
