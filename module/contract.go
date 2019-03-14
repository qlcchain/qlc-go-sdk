package module

import "github.com/qlcchain/go-qlc/rpc"

type ContractApi struct {
	client *rpc.Client
}

func NewContractApi(c *rpc.Client) *ContractApi {
	return &ContractApi{client: c}
}

func (c *ContractApi) PackContractData(abiStr string, methodName string, params []string) ([]byte, error) {
	return nil, nil
}
