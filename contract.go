package qlcchain

import (
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type ContractSendBlockPara struct {
	Address   types.Address `json:"address"`
	TokenName string        `json:"tokenName"`
	To        types.Address `json:"to"`
	Amount    types.Balance `json:"amount"`
	Data      []byte        `json:"data"`

	PrivateFrom    string   `json:"privateFrom,omitempty"`
	PrivateFor     []string `json:"privateFor,omitempty"`
	PrivateGroupID string   `json:"privateGroupID,omitempty"`
	EnclaveKey     []byte   `json:"enclaveKey,omitempty"`
}

type ContractRewardBlockPara struct {
	SendHash types.Hash `json:"sendHash"`
}

type ContractApi struct {
	client *QLCClient
}

// NewContractAPI creates contract module for client
func NewContractAPI(c *QLCClient) *ContractApi {
	return &ContractApi{client: c}
}

// GetAbiByContractAddress return contract abi by contract address
func (c *ContractApi) GetAbiByContractAddress(address types.Address) (string, error) {
	var r string
	err := c.client.getClient().Call(&r, "contract_getAbiByContractAddress", address)
	if err != nil {
		return "", err
	}
	return r, nil
}

// ContractAddressList return all contract addresses
func (c *ContractApi) ContractAddressList() []types.Address {
	return types.ChainContractAddressList
}

// PackContractData parse a ABI interface and pack the given method name to conform the ABI.
func (c *ContractApi) PackContractData(abiStr string, methodName string, params []string) ([]byte, error) {
	var r []byte
	err := c.client.getClient().Call(&r, "contract_packContractData", abiStr, methodName, params)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// PackChainContractData pack the given method name to conform the ABI for chain contract.
func (c *ContractApi) PackChainContractData(contractAddress types.Address, methodName string, params []string) ([]byte, error) {
	var r []byte
	err := c.client.getClient().Call(&r, "contract_packChainContractData", contractAddress, methodName, params)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GenerateSendBlock return new generated ContractSend block
func (c *ContractApi) GenerateSendBlock(para *ContractSendBlockPara) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := c.client.getClient().Call(&blk, "contract_generateSendBlock", para)
	if err != nil {
		return nil, err
	}
	return &blk, nil
}

// GenerateRewardBlock return new generated ContractReward block
func (c *ContractApi) GenerateRewardBlock(para *ContractRewardBlockPara) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := c.client.getClient().Call(&blk, "contract_generateRewardBlock", para)
	if err != nil {
		return nil, err
	}
	return &blk, nil
}
