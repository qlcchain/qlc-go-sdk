package module

import (
	"github.com/qlcchain/go-qlc/rpc"
)

type ContractApi struct {
	client *rpc.Client
}

func NewContractApi(c *rpc.Client) *ContractApi {
	return &ContractApi{client: c}
}

func (c *ContractApi) PackContractData(abiStr string, methodName string, params []string) ([]byte, error) {
	var r []byte
	err := c.client.Call(&r, "contract_packContractData", abiStr, methodName, params)
	if err != nil {
		return nil, err
	}
	return r, nil
}
