package qlcchain

import (
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type SMSApi struct {
	client *QLCClient
}

// NewSMSApi creates sms module for client
func NewSMSApi(c *QLCClient) *SMSApi {
	return &SMSApi{client: c}
}

// PhoneBlocks accepts a phone number, and returns send blocks and receiver blocks that relevant to the number
func (s *SMSApi) PhoneBlocks(phone string) (map[string][]*APIBlock, error) {
	var r map[string][]*APIBlock
	err := s.client.Call(&r, "sms_phoneBlocks", phone)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// MessageBlock accepts a message hash, and returns blocks that relevant to the hash
func (s *SMSApi) MessageBlocks(hash types.Hash) ([]*APIBlock, error) {
	var ab []*APIBlock
	err := s.client.Call(&ab, "sms_messageBlocks", hash)
	if err != nil {
		return nil, err
	}
	return ab, nil
}

// MessageHash returns hash of message
func (s *SMSApi) MessageHash(message string) (types.Hash, error) {
	var h types.Hash
	err := s.client.Call(&h, "sms_messageHash", message)
	if err != nil {
		return types.ZeroHash, err
	}
	return h, nil
}

// MessageStore stores message and returns message hash
func (s *SMSApi) MessageStore(message string) (types.Hash, error) {
	var h types.Hash
	err := s.client.Call(&h, "sms_messageStore", message)
	if err != nil {
		return types.ZeroHash, err
	}
	return h, nil
}

// MessageInfo returns message for message hash
func (s *SMSApi) MessageInfo(mHash types.Hash) (string, error) {
	var str string
	err := s.client.Call(&str, "sms_messageInfo", mHash)
	if err != nil {
		return "", err
	}
	return str, nil
}
