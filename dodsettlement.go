package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	_ "github.com/qlcchain/qlc-go-sdk/pkg/util"
)

type DoDSettlementAPI struct {
	client *rpc.Client
}

func NewDoDSettlementAPI(c *rpc.Client) *DoDSettlementAPI {
	return &DoDSettlementAPI{
		client: c,
	}
}

type DoDSettleUser struct {
	Address types.Address `json:"address" msg:"a,extension"`
	Name    string        `json:"name" msg:"n"`
}

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
request
confirmed
rejected
)
*/
type DoDSettleContractState int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
success
complete
fail
)
*/
type DoDSettleOrderState int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
invoice
stableCoin
)
*/
type DoDSettlePaymentType int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
PAYG
DOD
)
*/
type DoDSettleBillingType int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
year
month
week
day
hour
minute
second
)
*/
type DoDSettleBillingUnit int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
gold
silver
bronze
)
*/
type DoDSettleServiceClass int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
confirm
reject
)
*/
type DoDSettleResponseAction int

//go:generate go-enum -f=$GOFILE --marshal --names
/*
ENUM(
null
create
change
terminate
)
*/
type DoDSettleOrderType int

type DoDSettleCreateOrderParam struct {
	Buyer       *DoDSettleUser              `json:"buyer" msg:"b"`
	Seller      *DoDSettleUser              `json:"seller" msg:"s"`
	Connections []*DoDSettleConnectionParam `json:"connections,omitempty" msg:"c"`
}
type DoDSettleConnectionParam struct {
	DoDSettleConnectionStaticParam
	DoDSettleConnectionDynamicParam
}

type DoDSettleConnectionStaticParam struct {
	ItemId         string `json:"itemId,omitempty" msg:"ii"`
	BuyerProductId string `json:"buyerProductId,omitempty" msg:"bp"`
	ProductId      string `json:"productId,omitempty" msg:"pi"`
	SrcCompanyName string `json:"srcCompanyName,omitempty" msg:"scn"`
	SrcRegion      string `json:"srcRegion,omitempty" msg:"sr"`
	SrcCity        string `json:"srcCity,omitempty" msg:"sc"`
	SrcDataCenter  string `json:"srcDataCenter,omitempty" msg:"sdc"`
	SrcPort        string `json:"srcPort,omitempty" msg:"sp"`
	DstCompanyName string `json:"dstCompanyName,omitempty" msg:"dcn"`
	DstRegion      string `json:"dstRegion,omitempty" msg:"dr"`
	DstCity        string `json:"dstCity,omitempty" msg:"dc"`
	DstDataCenter  string `json:"dstDataCenter,omitempty" msg:"ddc"`
	DstPort        string `json:"dstPort,omitempty" msg:"dp"`
}

type DoDSettleConnectionDynamicParam struct {
	OrderId        string                `json:"orderId,omitempty" msg:"oi"`
	QuoteId        string                `json:"quoteId,omitempty" msg:"q"`
	QuoteItemId    string                `json:"quoteItemId,omitempty" msg:"qi"`
	ConnectionName string                `json:"connectionName,omitempty" msg:"cn"`
	PaymentType    DoDSettlePaymentType  `json:"paymentType,omitempty" msg:"pt"`
	BillingType    DoDSettleBillingType  `json:"billingType,omitempty" msg:"bt"`
	Currency       string                `json:"currency,omitempty" msg:"cr"`
	ServiceClass   DoDSettleServiceClass `json:"serviceClass,omitempty" msg:"scs"`
	Bandwidth      string                `json:"bandwidth,omitempty" msg:"bw"`
	BillingUnit    DoDSettleBillingUnit  `json:"billingUnit,omitempty" msg:"bu"`
	Price          float64               `json:"price,omitempty" msg:"p"`
	Addition       float64               `json:"addition" msg:"ad"`
	StartTime      int64                 `json:"startTime" msg:"st"`
	StartTimeStr   string                `json:"startTimeStr,omitempty" msg:"-"`
	EndTime        int64                 `json:"endTime" msg:"et"`
	EndTimeStr     string                `json:"endTimeStr,omitempty" msg:"-"`
}
type DoDSettleResponseParam struct {
	RequestHash types.Hash              `json:"requestHash" msg:"-"`
	Action      DoDSettleResponseAction `json:"action" msg:"c"`
}
type DoDSettleChangeConnectionParam struct {
	ProductId string `json:"productId" msg:"p"`
	DoDSettleConnectionDynamicParam
}
type DoDSettleTerminateOrderParam struct {
	Buyer       *DoDSettleUser                    `json:"buyer" msg:"b"`
	Seller      *DoDSettleUser                    `json:"seller" msg:"s"`
	Connections []*DoDSettleChangeConnectionParam `json:"connections" msg:"c"`
}
type DoDSettleChangeOrderParam struct {
	Buyer       *DoDSettleUser                    `json:"buyer" msg:"b"`
	Seller      *DoDSettleUser                    `json:"seller" msg:"s"`
	Connections []*DoDSettleChangeConnectionParam `json:"connections" msg:"c"`
}
type DoDSettleResourceReadyParam struct {
	Address    types.Address `json:"address" msg:"-"`
	InternalId types.Hash    `json:"internalId" msg:"i,extension"`
	ProductId  []string      `json:"productId" msg:"p"`
}

func (s *DoDSettlementAPI) GetCreateOrderBlock(param *DoDSettleCreateOrderParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getCreateOrderBlock", param)
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

func (s *DoDSettlementAPI) GetCreateOrderRewardBlock(param *DoDSettleResponseParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getCreateOrderRewardBlock", param)
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

func (s *DoDSettlementAPI) GetUpdateOrderInfoRewardBlock(param *DoDSettleResponseParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getUpdateOrderInfoRewardBlock", param)
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

func (s *DoDSettlementAPI) GetChangeOrderBlock(param *DoDSettleChangeOrderParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getChangeOrderBlock", param)
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

func (s *DoDSettlementAPI) GetChangeOrderRewardBlock(param *DoDSettleResponseParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getChangeOrderRewardBlock", param)
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

func (s *DoDSettlementAPI) GetTerminateOrderBlock(param *DoDSettleTerminateOrderParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getTerminateOrderBlock", param)
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

func (s *DoDSettlementAPI) GetTerminateOrderRewardBlock(param *DoDSettleResponseParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getTerminateOrderRewardBlock", param)
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

func (s *DoDSettlementAPI) GetResourceReadyBlock(param *DoDSettleResourceReadyParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getResourceReadyBlock", param)
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

func (s *DoDSettlementAPI) GetResourceReadyRewardBlock(param *DoDSettleResponseParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.Call(&blk, "DoDSettlement_getResourceReadyRewardBlock", param)
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

type DoDSettleOrderLifeTrack struct {
	ContractState DoDSettleContractState `json:"contractState" msg:"cs"`
	OrderState    DoDSettleOrderState    `json:"orderState" msg:"os"`
	Reason        string                 `json:"reason,omitempty" msg:"r"`
	Time          int64                  `json:"time" msg:"t"`
	Hash          types.Hash             `json:"hash" msg:"h,extension"`
}
type DoDSettleOrderInfo struct {
	Buyer         *DoDSettleUser              `json:"buyer" msg:"b"`
	Seller        *DoDSettleUser              `json:"seller" msg:"s"`
	OrderId       string                      `json:"orderId,omitempty" msg:"oi"`
	OrderType     DoDSettleOrderType          `json:"orderType,omitempty" msg:"ot"`
	OrderState    DoDSettleOrderState         `json:"orderState" msg:"os"`
	ContractState DoDSettleContractState      `json:"contractState" msg:"cs"`
	Connections   []*DoDSettleConnectionParam `json:"connections" msg:"c"`
	Track         []*DoDSettleOrderLifeTrack  `json:"track" msg:"t"`
}
type DoDSettleDisconnectInfo struct {
	OrderId      string  `json:"orderId,omitempty" msg:"oi"`
	QuoteId      string  `json:"quoteId,omitempty" msg:"q"`
	QuoteItemId  string  `json:"quoteItemId,omitempty" msg:"qi"`
	Price        float64 `json:"price,omitempty" msg:"p"`
	Currency     string  `json:"currency,omitempty" msg:"cr"`
	DisconnectAt int64   `json:"disconnectAt,omitempty" msg:"d"`
}
type DoDSettleConnectionLifeTrack struct {
	OrderType DoDSettleOrderType               `json:"orderType,omitempty" msg:"ot"`
	OrderId   string                           `json:"orderId,omitempty" msg:"oi"`
	Time      int64                            `json:"time,omitempty" msg:"t"`
	Changed   *DoDSettleConnectionDynamicParam `json:"changed,omitempty" msg:"c"`
}
type DoDSettleConnectionInfo struct {
	DoDSettleConnectionStaticParam
	Active     *DoDSettleConnectionDynamicParam   `json:"active" msg:"ac"`
	Done       []*DoDSettleConnectionDynamicParam `json:"done" msg:"do"`
	Disconnect *DoDSettleDisconnectInfo           `json:"disconnect" msg:"dis"`
	Track      []*DoDSettleConnectionLifeTrack    `json:"track" msg:"t"`
}
type DoDPendingRequestRsp struct {
	Hash  types.Hash          `json:"hash"`
	Order *DoDSettleOrderInfo `json:"order"`
}
type DoDSettleProductWithActiveInfo struct {
	ProductId string `json:"productId"`
	Active    bool   `json:"active"`
}
type DoDPendingResourceCheckInfo struct {
	SendHash   types.Hash                        `json:"sendHash"`
	OrderId    string                            `json:"orderId"`
	InternalId types.Hash                        `json:"internalId"`
	Products   []*DoDSettleProductWithActiveInfo `json:"products"`
}
type DoDPlacingOrderInfo struct {
	InternalId types.Hash          `json:"internalId"`
	OrderInfo  *DoDSettleOrderInfo `json:"orderInfo"`
}
type DoDSettleInvoiceConnDynamic struct {
	DoDSettleConnectionDynamicParam
	InvoiceStartTime    int64   `json:"invoiceStartTime,omitempty"`
	InvoiceStartTimeStr string  `json:"invoiceStartTimeStr,omitempty"`
	InvoiceEndTime      int64   `json:"invoiceEndTime,omitempty"`
	InvoiceEndTimeStr   string  `json:"invoiceEndTimeStr,omitempty"`
	InvoiceUnitCount    int     `json:"invoiceUnitCount,omitempty"`
	Amount              float64 `json:"amount"`
}
type DoDSettleInvoiceConnDetail struct {
	ConnectionAmount float64 `json:"connectionAmount"`
	DoDSettleConnectionStaticParam
	Usage []*DoDSettleInvoiceConnDynamic `json:"usage"`
}
type DoDSettleInvoiceOrderDetail struct {
	OrderId         string                        `json:"orderId"`
	ConnectionCount int                           `json:"connectionCount"`
	OrderAmount     float64                       `json:"orderAmount"`
	Connections     []*DoDSettleInvoiceConnDetail `json:"connections"`
}
type DoDSettleOrderInvoice struct {
	TotalConnectionCount int                          `json:"totalConnectionCount"`
	TotalAmount          float64                      `json:"totalAmount"`
	Currency             string                       `json:"currency"`
	StartTime            int64                        `json:"startTime"`
	EndTime              int64                        `json:"endTime"`
	Buyer                *DoDSettleUser               `json:"buyer"`
	Seller               *DoDSettleUser               `json:"seller"`
	Order                *DoDSettleInvoiceOrderDetail `json:"order"`
}
type DoDSettleBuyerInvoice struct {
	OrderCount           int                            `json:"orderCount"`
	TotalConnectionCount int                            `json:"totalConnectionCount"`
	TotalAmount          float64                        `json:"totalAmount"`
	Currency             string                         `json:"currency"`
	StartTime            int64                          `json:"startTime"`
	EndTime              int64                          `json:"endTime"`
	Buyer                *DoDSettleUser                 `json:"buyer"`
	Seller               *DoDSettleUser                 `json:"seller"`
	Orders               []*DoDSettleInvoiceOrderDetail `json:"orders"`
}
type DoDSettleProductInvoice struct {
	TotalAmount float64                     `json:"totalAmount"`
	Currency    string                      `json:"currency"`
	StartTime   int64                       `json:"startTime"`
	EndTime     int64                       `json:"endTime"`
	Buyer       *DoDSettleUser              `json:"buyer"`
	Seller      *DoDSettleUser              `json:"seller"`
	Connection  *DoDSettleInvoiceConnDetail `json:"connection"`
}
type DoDSettleProduct struct {
	Seller    types.Address `json:"seller" msg:"s,extension"`
	ProductId string        `json:"productId,omitempty" msg:"p"`
}
type DoDSettleOrder struct {
	Seller  types.Address `json:"seller" msg:"s,extension"`
	OrderId string        `json:"orderId,omitempty" msg:"o"`
}

func (s *DoDSettlementAPI) GetOrderInfoBySellerAndOrderId(seller types.Address, orderId string) (*DoDSettleOrderInfo, error) {
	var r DoDSettleOrderInfo
	err := s.client.Call(&r, "DoDSettlement_getOrderInfoBySellerAndOrderId", seller, orderId)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
func (s *DoDSettlementAPI) GetOrderInfoByInternalId(internalId string) (*DoDSettleOrderInfo, error) {
	var r DoDSettleOrderInfo
	err := s.client.Call(&r, "DoDSettlement_getOrderInfoByInternalId", internalId)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
func (s *DoDSettlementAPI) GetConnectionInfoBySellerAndProductId(seller types.Address, productId string) (*DoDSettleConnectionInfo, error) {
	var r DoDSettleConnectionInfo
	err := s.client.Call(&r, "DoDSettlement_getConnectionInfoBySellerAndProductId", seller, productId)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
func (s *DoDSettlementAPI) GetPendingRequest(address types.Address) ([]*DoDPendingRequestRsp, error) {
	var r []*DoDPendingRequestRsp
	err := s.client.Call(&r, "DoDSettlement_getPendingRequest", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *DoDSettlementAPI) GetPendingResourceCheck(address types.Address) ([]*DoDPendingResourceCheckInfo, error) {
	var r []*DoDPendingResourceCheckInfo
	err := s.client.Call(&r, "DoDSettlement_getPendingResourceCheck", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *DoDSettlementAPI) GetPlacingOrder(buyer, seller types.Address) ([]*DoDPlacingOrderInfo, error) {
	var r []*DoDPlacingOrderInfo
	err := s.client.Call(&r, "DoDSettlement_getPlacingOrder", buyer, seller)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *DoDSettlementAPI) GetProductIdListByAddress(address types.Address) ([]*DoDSettleProduct, error) {
	var r []*DoDSettleProduct
	err := s.client.Call(&r, "DoDSettlement_getProductIdListByAddress", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *DoDSettlementAPI) GetOrderIdListByAddress(address types.Address) ([]*DoDSettleOrder, error) {
	var r []*DoDSettleOrder
	err := s.client.Call(&r, "DoDSettlement_getOrderIdListByAddress", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *DoDSettlementAPI) GetProductIdListByAddressAndSeller(address, seller types.Address) ([]*DoDSettleOrder, error) {
	var r []*DoDSettleOrder
	err := s.client.Call(&r, "DoDSettlement_getProductIdListByAddressAndSeller", address, seller)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *DoDSettlementAPI) GetOrderIdListByAddressAndSeller(address, seller types.Address) ([]*DoDSettleOrder, error) {
	var r []*DoDSettleOrder
	err := s.client.Call(&r, "DoDSettlement_getOrderIdListByAddressAndSeller", address, seller)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *DoDSettlementAPI) GenerateInvoiceByOrderId(seller types.Address, orderId string, start, end int64, flight, split bool) (*DoDSettleOrderInvoice, error) {
	var r DoDSettleOrderInvoice
	err := s.client.Call(&r, "DoDSettlement_generateInvoiceByOrderId", seller, orderId, start, end, flight, split)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
func (s *DoDSettlementAPI) GenerateInvoiceByBuyer(seller types.Address, orderId string, start, end int64, flight, split bool) (*DoDSettleBuyerInvoice, error) {
	var r DoDSettleBuyerInvoice
	err := s.client.Call(&r, "DoDSettlement_generateInvoiceByBuyer", seller, orderId, start, end, flight, split)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
func (s *DoDSettlementAPI) GenerateInvoiceByProductId(seller types.Address, orderId string, start, end int64, flight, split bool) (*DoDSettleProductInvoice, error) {
	var r DoDSettleProductInvoice
	err := s.client.Call(&r, "DoDSettlement_generateInvoiceByProductId", seller, orderId, start, end, flight, split)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
