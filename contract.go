package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type ContractApi struct {
	client *rpc.Client
}

// NewContractApi creates contract module for client
func NewContractApi(c *rpc.Client) *ContractApi {
	return &ContractApi{client: c}
}

// GetAbiByContractAddress return contract abi by contract address
func (c *ContractApi) GetAbiByContractAddress(address types.Address) (string, error) {
	var r string
	err := c.client.Call(&r, "contract_getAbiByContractAddress", address)
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
	err := c.client.Call(&r, "contract_packContractData", abiStr, methodName, params)
	if err != nil {
		return nil, err
	}
	return r, nil
}
