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

func (s *SMSApi) PhoneBlocks(sender string) (map[string][]*api.APIBlock, error) {
	return nil, nil
}

func (s *SMSApi) MessageBlock(hash types.Hash) (*api.APIBlock, error) {
	return nil, nil
}

func (s *SMSApi) MessageHash(message string) (types.Hash, error) {
	return types.ZeroHash, nil
}

func (s *SMSApi) MessageStore(message string) (types.Hash, error) {
	return types.ZeroHash, nil
}

func (s *SMSApi) MessageInfo(mHash types.Hash) (string, error) {
	return "", nil
}
