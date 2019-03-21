package module

import (
	"github.com/qlcchain/go-qlc/rpc"
)

type ContractApi struct {
	client *rpc.Client
}

// NewContractApi creates contract module for client
func NewContractApi(c *rpc.Client) *ContractApi {
	return &ContractApi{client: c}
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
