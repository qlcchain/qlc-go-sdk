package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type SettlementAPI struct {
	client *rpc.Client
}

func NewSettlementAPI(c *rpc.Client) *SettlementAPI {
	return &SettlementAPI{
		client: c,
	}
}

type Contractor struct {
	Address types.Address `json:"address"`
	Name    string        `json:"name"`
}

type ContractService struct {
	ServiceId   string  `json:"serviceId" validate:"nonzero"`
	Mcc         uint64  `json:"mcc"`
	Mnc         uint64  `json:"mnc"`
	TotalAmount uint64  `json:"totalAmount" validate:"min=1"`
	UnitPrice   float64 `son:"unitPrice" validate:"nonzero"`
	Currency    string  `json:"currency" validate:"nonzero"`
}

type CreateContractParam struct {
	PartyA    Contractor        `json:"partyA"`
	PartyB    Contractor        `json:"partyB"`
	Previous  types.Hash        `json:"previous"`
	Services  []ContractService `json:"services"`
	SignDate  int64             `json:"signDate"`
	StartDate int64             `json:"startDate"`
	EndData   int64             `json:"endData"`
}

func (s *SettlementAPI) ToAddress(param *CreateContractParam) (types.Address, error) {
	var r types.Address
	err := s.client.Call(&r, "settlement_toAddress", param)
	if err != nil {
		return types.ZeroAddress, err
	}
	return r, nil
}

func (s *SettlementAPI) GetCreateContractBlock(param *CreateContractParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getCreateContractBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetContractRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getContractRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

type SignContractParam struct {
	ContractAddress types.Address `json:"contractAddress"`
	ConfirmDate     int64         `json:"confirmDate"`
	Address         types.Address // PartyB address
}

type StopParam struct {
	StopName string `json:"stopName" validate:"nonzero"`
}

type UpdateStopParam struct {
	StopName string `json:"stopName" validate:"nonzero"`
	New      string `json:"newName" validate:"nonzero"`
}

func (s *SettlementAPI) GetSignContractBlock(param *SignContractParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getContractRewardsBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetSignRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getContractRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetAddPreStopBlock(param *StopParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getAddPreStopBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetAddPreStopRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getAddPreStopRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil

}
func (s *SettlementAPI) GetRemovePreStopBlock(param *StopParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getRemovePreStopBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil

}
func (s *SettlementAPI) GetRemovePreStopRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getRemovePreStopRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil

}

func (s *SettlementAPI) GetUpdatePreStopBlock(param *UpdateStopParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getUpdatePreStopBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetUpdatePreStopRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getUpdatePreStopRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetAddNextStopBlock(param *StopParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getAddNextStopBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}
func (s *SettlementAPI) GetAddNextStopRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getAddNextStopRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetRemoveNextStopBlock(param *StopParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getRemoveNextStopBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetRemoveNextStopRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getRemoveNextStopRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetUpdateNextStopBlock(param *UpdateStopParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getUpdateNextStopBlock", param)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

func (s *SettlementAPI) GetUpdateNextStopRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "settlement_getUpdateNextStopRewardsBlock", send)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	return &blk, nil
}

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
ActiveStage1
Actived
DestroyStage1
Destroyed
)
*/
type ContractStatus int

type SettlementContract struct {
	CreateContractParam
	PreStops    []string       `msg:"pre" json:"preStops"`
	NextStops   []string       `msg:"nex" json:"nextStops"`
	ConfirmDate int64          `msg:"t2" json:"confirmDate"`
	Status      ContractStatus `msg:"s" json:"status"`
	Address     types.Address  // settlement smart contract address
}

func (s *SettlementAPI) GetContractsByAddress(addr *types.Address, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.Call(&r, "settlement_getContractsByAddress", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetContractsAsPartyA(addr *types.Address, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.Call(&r, "settlement_getContractsAsPartyA", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetContractsAsPartyB(addr *types.Address, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.Call(&r, "settlement_getContractsAsPartyB", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
Send
Error
Empty
)
*/
type SendingStatus int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
Delivered
Unknown
Undelivered
Empty
)
*/
type DLRStatus int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
unknown
stage1
success
failure
missing
duplicate
)
*/
type SettlementStatus int

type CDRParam struct {
	Index         uint64        `json:"index" validate:"min=1"`
	SmsDt         int64         `json:"smsDt" validate:"min=1"`
	Sender        string        `json:"sender" validate:"nonzero"`
	Destination   string        `json:"destination" validate:"nonzero"`
	SendingStatus SendingStatus `json:"sendingStatus" `
	DlrStatus     DLRStatus     `json:"dlrStatus"`
	PreStop       string        `json:"preStop" `
	NextStop      string        `json:"nextStop" `
}

type SettlementCDR struct {
	CDRParam
	From types.Address `json:"from"`
}

type CDRStatus struct {
	Params []SettlementCDR  `json:"params"`
	Status SettlementStatus `json:"status"`
}

func (s *SettlementAPI) GetCDRStatus(addr *types.Address, hash types.Hash) (*CDRStatus, error) {
	var r CDRStatus
	err := s.client.Call(&r, "settlement_getCDRStatus", addr, hash)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *SettlementAPI) GetAllCDRStatus(addr *types.Address, count int, offset *int) ([]*CDRStatus, error) {
	var r []*CDRStatus
	err := s.client.Call(&r, "settlement_getCDRStatus", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}
