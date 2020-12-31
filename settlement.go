package qlcchain

import (
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	"github.com/qlcchain/qlc-go-sdk/pkg/util"
)

type SettlementAPI struct {
	client *QLCClient
}

func NewSettlementAPI(c *QLCClient) *SettlementAPI {
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
	Services  []ContractService `json:"services"`
	StartDate int64             `json:"startDate"`
	EndDate   int64             `json:"endDate"`
}

func (s *SettlementAPI) ToAddress(param *CreateContractParam) (types.Address, error) {
	var r types.Address
	err := s.client.getClient().Call(&r, "settlement_toAddress", param)
	if err != nil {
		return types.ZeroAddress, err
	}
	return r, nil
}

func (s *SettlementAPI) GetCreateContractBlock(param *CreateContractParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "settlement_getCreateContractBlock", param)
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

func (s *SettlementAPI) GetSettlementRewardsBlock(send *types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "settlement_getSettlementRewardsBlock", send)
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
	Address         types.Address `json:"address"`
}

type StopParam struct {
	ContractAddress types.Address `json:"contractAddress"`
	StopName        string        `json:"stopName" validate:"nonzero"`
	Address         types.Address `json:"address"`
}

type UpdateStopParam struct {
	ContractAddress types.Address `json:"contractAddress"`
	StopName        string        `json:"stopName" validate:"nonzero"`
	New             string        `json:"newName" validate:"nonzero"`
	Address         types.Address `json:"address"`
}

func (s *SettlementAPI) GetSignContractBlock(param *SignContractParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "settlement_getSignContractBlock", param)
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

func (s *SettlementAPI) GetProcessCDRBlock(addr *types.Address, params []*CDRParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "settlement_getProcessCDRBlock", addr, params)
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
	err := s.client.getClient().Call(&blk, "settlement_getAddPreStopBlock", param)
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
	err := s.client.getClient().Call(&blk, "settlement_getRemovePreStopBlock", param)
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
	err := s.client.getClient().Call(&blk, "settlement_getUpdatePreStopBlock", param)
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
	err := s.client.getClient().Call(&blk, "settlement_getAddNextStopBlock", param)
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
	err := s.client.getClient().Call(&blk, "settlement_getRemoveNextStopBlock", param)
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
	err := s.client.getClient().Call(&blk, "settlement_getUpdateNextStopBlock", param)
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

type TerminateParam struct {
	ContractAddress types.Address `json:"contractAddress"`
	Address         types.Address `json:"address"`
	Request         bool          `json:"request"`
}

func (s *SettlementAPI) GetTerminateContractBlock(param *TerminateParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "settlement_getTerminateContractBlock", param)
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

type ContractStatus int

type SettlementContract struct {
	CreateContractParam
	PreStops    []string       `json:"preStops"`
	NextStops   []string       `json:"nextStops"`
	ConfirmDate int64          `json:"confirmDate"`
	Status      ContractStatus `json:"status"`
	Address     types.Address  `json:"address"`
	Terminator  *types.Address `json:"-"`
}

func (s *SettlementAPI) GetAllContracts(count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.getClient().Call(&r, "settlement_getAllContracts", count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetContractsByAddress(addr *types.Address, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.getClient().Call(&r, "settlement_getContractsByAddress", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetContractsAsPartyA(addr *types.Address, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.getClient().Call(&r, "settlement_getContractsAsPartyA", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetContractsAsPartyB(addr *types.Address, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.getClient().Call(&r, "settlement_getContractsAsPartyB", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetContractAddressByPartyANextStop(addr *types.Address, stopName string) (*types.Address, error) {
	var contractAddress types.Address
	err := s.client.getClient().Call(&contractAddress, "settlement_getContractAddressByPartyANextStop", addr, stopName)
	if err != nil {
		return nil, err
	}
	return &contractAddress, nil
}

func (s *SettlementAPI) GetContractAddressByPartyBPreStop(addr *types.Address, stopName string) (*types.Address, error) {
	var contractAddress types.Address
	err := s.client.getClient().Call(&contractAddress, "settlement_getContractAddressByPartyBPreStop", addr, stopName)
	if err != nil {
		return nil, err
	}
	return &contractAddress, nil
}

func (s *SettlementAPI) GetContractsByStatus(addr *types.Address, status string, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.getClient().Call(&r, "settlement_getContractsByStatus", addr, status, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetExpiredContracts(addr *types.Address, count int, offset *int) ([]*SettlementContract, error) {
	var r []*SettlementContract
	err := s.client.getClient().Call(&r, "settlement_getExpiredContracts", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type SendingStatus int
type DLRStatus int

type SettlementStatus int

type CDRParam struct {
	Index         uint64        `json:"index" validate:"min=1"`
	SmsDt         int64         `json:"smsDt" validate:"min=1"`
	Account       string        `json:"account"`
	Sender        string        `json:"sender" validate:"nonzero"`
	Customer      string        `json:"customer"`
	Destination   string        `json:"destination" validate:"nonzero"`
	SendingStatus SendingStatus `json:"sendingStatus" `
	DlrStatus     DLRStatus     `json:"dlrStatus"`
	PreStop       string        `json:"preStop" `
	NextStop      string        `json:"nextStop" `
}

func (z *CDRParam) ToHash() (types.Hash, error) {
	return types.HashBytes(util.BE_Uint64ToBytes(z.Index), []byte(z.Sender), []byte(z.Destination))
}

func (s *SettlementAPI) GetNextStopNames(addr *types.Address) ([]string, error) {
	var r []string
	err := s.client.getClient().Call(&r, "settlement_getNextStopNames", addr)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetPreStopNames(addr *types.Address) ([]string, error) {
	var r []string
	err := s.client.getClient().Call(&r, "settlement_getPreStopNames", addr)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type SettlementCDR struct {
	CDRParam
	From types.Address `json:"from"`
}

type CDRStatus struct {
	Address *types.Address        `json:"contractAddress"`
	Params  map[string][]CDRParam `json:"params"`
	Status  SettlementStatus      `json:"status"`
}

func (s *SettlementAPI) GetCDRStatus(addr *types.Address, hash types.Hash) (*CDRStatus, error) {
	var r CDRStatus
	err := s.client.getClient().Call(&r, "settlement_getCDRStatus", addr, hash)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *SettlementAPI) GetCDRStatusByCdrData(addr *types.Address, index uint64, sender, destination string) (*CDRStatus, error) {
	var r CDRStatus
	err := s.client.getClient().Call(&r, "settlement_getCDRStatusByCdrData", addr, index, sender, destination)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *SettlementAPI) GetAllCDRStatus(addr *types.Address, count int, offset *int) ([]*CDRStatus, error) {
	var r []*CDRStatus
	err := s.client.getClient().Call(&r, "settlement_getAllCDRStatus", addr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetMultiPartyCDRStatus(firstAddr, secondAddr *types.Address, count int, offset *int) ([]*CDRStatus, error) {
	var r []*CDRStatus
	err := s.client.getClient().Call(&r, "settlement_getMultiPartyCDRStatus", firstAddr, secondAddr, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetCDRStatusByDate(addr *types.Address, start, end int64, count int, offset *int) ([]*CDRStatus, error) {
	var r []*CDRStatus
	err := s.client.getClient().Call(&r, "settlement_getCDRStatusByDate", addr, start, end, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type SummaryRecord struct {
	Total   uint64  `json:"total"`
	Success uint64  `json:"success"`
	Fail    uint64  `json:"fail"`
	Result  float64 `json:"result"`
}

type MatchingRecord struct {
	Orphan   SummaryRecord `json:"orphan"`
	Matching SummaryRecord `json:"matching"`
}

type CompareRecord struct {
	PartyA MatchingRecord `json:"partyA"`
	PartyB MatchingRecord `json:"partyB"`
}

type SummaryResult struct {
	Contract *SettlementContract       `json:"contract"`
	Records  map[string]*CompareRecord `json:"records"`
	Total    CompareRecord             `json:"total"`
}

func (s *SettlementAPI) GetSummaryReport(addr *types.Address, start, end int64) (*SummaryResult, error) {
	var r SummaryResult
	err := s.client.getClient().Call(&r, "settlement_getSummaryReport", addr, start, end)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GetSummaryReportByAccount generate summary report by PCCWG account
// @param addr settlement contract address
// @param account PCCWG account
// @param start report start date (UTC unix time)
// @param end report end data (UTC unix time)
// @return summary report if error not exist
func (s *SettlementAPI) GetSummaryReportByAccount(addr *types.Address, account string, start, end int64) (*SummaryResult, error) {
	var r SummaryResult
	err := s.client.getClient().Call(&r, "settlement_getSummaryReportByAccount", addr, account, start, end)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GetSummaryReportByCustomer generate summary report by PCCWG customer name
// @param addr settlement contract address
// @param customer PCCWG customer name
// @param start report start date (UTC unix time)
// @param end report end data (UTC unix time)
// @return summary report if error not exist
func (s *SettlementAPI) GetSummaryReportByCustomer(addr *types.Address, customer string, start, end int64) (*SummaryResult, error) {
	var r SummaryResult
	err := s.client.getClient().Call(&r, "settlement_getSummaryReportByCustomer", addr, customer, start, end)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

type InvoiceRecord struct {
	Address                  types.Address `json:"contractAddress"`
	StartDate                int64         `json:"startDate"`
	EndDate                  int64         `json:"endDate"`
	Customer                 string        `json:"customer"`
	CustomerSr               string        `json:"customerSr"`
	Country                  string        `json:"country"`
	Operator                 string        `json:"operator"`
	ServiceId                string        `json:"serviceId"`
	MCC                      uint64        `json:"mcc"`
	MNC                      uint64        `json:"mnc"`
	Currency                 string        `json:"currency"`
	UnitPrice                float64       `json:"unitPrice"`
	SumOfBillableSMSCustomer uint64        `json:"sumOfBillableSMSCustomer"`
	SumOfTOTPrice            float64       `json:"sumOfTOTPrice"`
}

func (s *SettlementAPI) GenerateInvoices(addr *types.Address, start, end int64) ([]*InvoiceRecord, error) {
	var r []*InvoiceRecord
	err := s.client.getClient().Call(&r, "settlement_generateInvoices", addr, start, end)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GenerateInvoicesByAccount generate invoice by PCCWG account
// @param addr settlement contract address
// @param account PCCWG account
// @param start report start date (UTC unix time)
// @param end report end data (UTC unix time)
// @return settlement invoice
func (s *SettlementAPI) GenerateInvoicesByAccount(addr *types.Address, account string, start, end int64) ([]*InvoiceRecord, error) {
	var r []*InvoiceRecord
	err := s.client.getClient().Call(&r, "settlement_generateInvoicesByAccount", addr, account, start, end)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GenerateInvoicesByCustomer generate invoice by PCCWG customer name
// @param addr settlement contract address
// @param customer PCCWG customer name
// @param start report start date (UTC unix time)
// @param end report end data (UTC unix time)
// @return settlement invoice
func (s *SettlementAPI) GenerateInvoicesByCustomer(addr *types.Address, customer string, start, end int64) ([]*InvoiceRecord, error) {
	var r []*InvoiceRecord
	err := s.client.getClient().Call(&r, "settlement_generateInvoicesByCustomer", addr, customer, start, end)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GenerateInvoicesByContract(addr *types.Address, start, end int64) ([]*InvoiceRecord, error) {
	var r []*InvoiceRecord
	err := s.client.getClient().Call(&r, "settlement_generateInvoicesByContract", addr, start, end)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GenerateMultiPartyInvoice(firstAddr, secondAddr *types.Address, start, end int64) ([]*InvoiceRecord, error) {
	var r []*InvoiceRecord
	err := s.client.getClient().Call(&r, "settlement_generateMultiPartyInvoice", firstAddr, secondAddr, start, end)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type MultiPartySummaryResult struct {
	Contracts []*SettlementContract     `json:"contracts"`
	Records   map[string]*CompareRecord `json:"records"`
	Total     *CompareRecord            `json:"total"`
}

func (s *SettlementAPI) GenerateMultiPartySummaryReport(firstAddr, secondAddr *types.Address, start, end int64) (*MultiPartySummaryResult, error) {
	var r MultiPartySummaryResult
	err := s.client.getClient().Call(&r, "settlement_generateMultiPartySummaryReport", firstAddr, secondAddr, start, end)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

type Compensation struct {
	Low  float32 `msg:"l" json:"low"`
	High float32 `msg:"h" json:"high"`
	Rate float32 `msg:"r" json:"rate"`
}

type SLAType int

type SLA struct {
	SLAType       SLAType         `json:"type"`
	Priority      uint            `json:"priority"`
	Value         float32         `json:"value"`
	Compensations []*Compensation `json:"compensations,omitempty"`
}

type Asset struct {
	Mcc         uint64 `json:"mcc"`
	Mnc         uint64 `json:"mnc"`
	TotalAmount uint64 `json:"totalAmount"`
	SLAs        []*SLA `json:"sla,omitempty"`
}

type RegisterAssetParam struct {
	Owner     Contractor `json:"owner"`
	Assets    []*Asset   `json:"assets"`
	StartDate int64      `json:"startDate"`
	EndDate   int64      `json:"endDate"`
	Status    string     `json:"status"`
}

func (s *SettlementAPI) GetRegisterAssetBlock(param *RegisterAssetParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "settlement_getRegisterAssetBlock", param)
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

type AssetStatus int

type APIAsset struct {
	Asset
	AssetID types.Hash `json:"assetID"`
}

type AssetParam struct {
	Owner     Contractor    `json:"owner"`
	Assets    []*APIAsset   `json:"assets"`
	SignDate  int64         `json:"signDate"`
	StartDate int64         `json:"startDate"`
	EndDate   int64         `json:"endDate"`
	Status    AssetStatus   `json:"status"`
	Address   types.Address `json:"address"`
}

func (s *SettlementAPI) GetAllAssets(count int, offset *int) ([]*AssetParam, error) {
	var r []*AssetParam
	err := s.client.getClient().Call(&r, "settlement_getAllAssets", count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetAssetsByOwner(owner *types.Address, count int, offset *int) ([]*AssetParam, error) {
	var r []*AssetParam
	err := s.client.getClient().Call(&r, "settlement_getAssetsByOwner", owner, count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *SettlementAPI) GetAsset(address types.Address) (*AssetParam, error) {
	var r AssetParam
	err := s.client.getClient().Call(&r, "settlement_getAsset", address)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
