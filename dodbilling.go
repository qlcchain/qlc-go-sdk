package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	_ "github.com/qlcchain/qlc-go-sdk/pkg/util"
)

type DoDBillingAPI struct {
	client *rpc.Client
}

func NewDoDBillingApi(c *rpc.Client) *DoDBillingAPI {
	return &DoDBillingAPI{
		client: c,
	}
}

//go:generate msgp
type DoDAccount struct {
	Account     types.Address `msg:"-" json:"account"`
	AccountName string        `msg:"-" json:"accountName"`
	AccountInfo string        `msg:"ai" json:"accountInfo"`
	AccountType string        `msg:"at" json:"accountType"`
	UUID        string        `msg:"uu" json:"uuid"`
	Connections []string      `msg:"cs" json:"connections"`
}

func (s *DoDBillingAPI) GetDodBillingSetAccountBlock(param *DoDAccount, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "dodbilling_getSetAccountBlock", param)
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

// need to split the params according to different services
type DoDConnection struct {
	Account         types.Address `json:"account"`
	AccountName     string        `json:"accountName"`
	ConnectionID    string        `json:"connectionID"`
	ServiceType     string        `json:"serviceType"`
	ChargeType      string        `json:"chargeType"`
	PaidRule        string        `json:"paidRule"`
	BuyMode         string        `json:"buyMode"`
	Location        string        `json:"location"`
	StartTime       uint64        `json:"startTime"`
	EndTime         uint64        `json:"endTime"`
	Price           float64       `json:"price"`
	Unit            string        `json:"unit"`
	Currency        string        `json:"currency"`
	Balance         float64       `json:"balance"`
	TempStartTime   uint64        `json:"tempStartTime"`
	TempEndTime     uint64        `json:"tempEndTime"`
	TempPrice       float64       `json:"tempPrice"`
	TempBandwidth   string        `json:"tempBandwidth"`
	Bandwidth       string        `json:"bandwidth"`
	Quota           float64       `json:"quota"`
	UsageLimitation float64       `json:"usageLimitation"`
	MinBandwidth    string        `json:"minBandwidth"`
	ExpireTime      uint64        `json:"expireTime"`
}

func (s *DoDBillingAPI) GetDodBillingSetServiceBlock(param *DoDAccount, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "dodbilling_getSetServiceBlock", param)
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
