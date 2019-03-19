package module

import (
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
	"github.com/qlcchain/go-qlc/rpc/api"
)

type SMSApi struct {
	client *rpc.Client
}

func NewSMSApi(c *rpc.Client) *SMSApi {
	return &SMSApi{client: c}
}

func (s *SMSApi) PhoneBlocks(phone string) (map[string][]*api.APIBlock, error) {
	var r map[string][]*api.APIBlock
	err := s.client.Call(&r, "sms_phoneBlocks", phone)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SMSApi) MessageBlock(hash types.Hash) (*api.APIBlock, error) {
	var ab api.APIBlock
	err := s.client.Call(&ab, "sms_messageBlock", hash)
	if err != nil {
		return nil, err
	}
	return &ab, nil
}

func (s *SMSApi) MessageHash(message string) (types.Hash, error) {
	var h types.Hash
	err := s.client.Call(&h, "sms_messageHash", message)
	if err != nil {
		return types.ZeroHash, err
	}
	return h, nil
}

func (s *SMSApi) MessageStore(message string) (types.Hash, error) {
	var h types.Hash
	err := s.client.Call(&h, "sms_messageStore", message)
	if err != nil {
		return types.ZeroHash, err
	}
	return h, nil
}

func (s *SMSApi) MessageInfo(mHash types.Hash) (string, error) {
	var str string
	err := s.client.Call(&str, "sms_messageInfo", mHash)
	if err != nil {
		return "", err
	}
	return str, nil
}
