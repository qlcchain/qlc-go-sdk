package qlcchain

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

type Subscribe struct {
	ws          *websocket.Conn
	subscribeID string
	Stopped     chan bool
}

func NewSubscribe(url string) *Subscribe {
	ws, err := websocket.Dial(url, "", url)
	if err != nil {
		fmt.Println("websocket dial: ", err)
	}
	return &Subscribe{
		ws:      ws,
		Stopped: make(chan bool),
	}
}

type subscribeInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type resultInfo struct {
	Subscription string      `json:"subscription"`
	Result       interface{} `json:"result"`
}

type publishInfo struct {
	Jsonrpc string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  resultInfo `json:"params"`
}

// type unsubscribeInfo struct {
//	Jsonrpc string `json:"jsonrpc"`
//	Id      int    `json:"id"`
//	Result  bool   `json:"result"`
//}

func (s *Subscribe) subscribe(request string) error {
	if err := websocket.Message.Send(s.ws, request); err != nil {
		return fmt.Errorf("send error: %s", err)
	}
	var response string
	if err := websocket.Message.Receive(s.ws, &response); err != nil {
		return fmt.Errorf("receive message: %v ", err)
	}
	reply := new(subscribeInfo)
	err := json.Unmarshal([]byte(response), &reply)
	if err != nil {
		return fmt.Errorf("subscribe, Can not decode data: %s ", err)
	}
	s.subscribeID = reply.Result
	return nil
}

func (s *Subscribe) publish() (interface{}, bool) {
	var response string
	err := websocket.Message.Receive(s.ws, &response)
	if err != nil { // if call Close() or sever stopped, connect closed, can not receive message
		fmt.Println("receive publish message:  ", err)
		s.closeConnection()
		return nil, true
	}
	reply := new(publishInfo)
	err = json.Unmarshal([]byte(response), reply)
	if err != nil {
		fmt.Println("Can not decode publish data:  ", err)
		s.closeConnection()
		return nil, true
	}
	if reply.Params.Result == nil { // if call Unsubscribe(), connect closed, receive message is nil
		s.closeConnection()
		return nil, true
	}
	return reply.Params.Result, false
}

func (s *Subscribe) Unsubscribe(request string) error {
	if err := websocket.Message.Send(s.ws, request); err != nil {
		return fmt.Errorf("unsubscribe error: %s", err)
	}
	return nil
}

func (s *Subscribe) Close() error {
	s.closeConnection()
	return s.ws.Close()
}

func (s *Subscribe) closeConnection() {
	s.Stopped <- true
}
