package qlcchain

import (
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	_ "github.com/qlcchain/qlc-go-sdk/pkg/util"
)

type DoDSettlementAPI struct {
	client *QLCClient
}

func NewDoDSettlementAPI(c *QLCClient) *DoDSettlementAPI {
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

type DoDSettleConnectionParam struct {
	DoDSettleConnectionStaticParam
	DoDSettleConnectionDynamicParam
}

type DoDSettleConnectionStaticParam struct {
	BuyerProductId    string `json:"buyerProductId,omitempty" msg:"bp"`
	ProductOfferingId string `json:"productOfferingId,omitempty" msg:"po"`
	ProductId         string `json:"productId,omitempty" msg:"pi"`
	SrcCompanyName    string `json:"srcCompanyName,omitempty" msg:"scn"`
	SrcRegion         string `json:"srcRegion,omitempty" msg:"sr"`
	SrcCity           string `json:"srcCity,omitempty" msg:"sc"`
	SrcDataCenter     string `json:"srcDataCenter,omitempty" msg:"sdc"`
	SrcPort           string `json:"srcPort,omitempty" msg:"sp"`
	DstCompanyName    string `json:"dstCompanyName,omitempty" msg:"dcn"`
	DstRegion         string `json:"dstRegion,omitempty" msg:"dr"`
	DstCity           string `json:"dstCity,omitempty" msg:"dc"`
	DstDataCenter     string `json:"dstDataCenter,omitempty" msg:"ddc"`
	DstPort           string `json:"dstPort,omitempty" msg:"dp"`
}

type DoDSettleConnectionDynamicParam struct {
	OrderId        string                `json:"orderId,omitempty" msg:"oi"`
	InternalId     string                `json:"internalId,omitempty" msg:"-"`
	ItemId         string                `json:"itemId,omitempty" msg:"ii"`
	OrderItemId    string                `json:"orderItemId,omitempty" msg:"oii"`
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

type ContractPrivacyParam struct {
	PrivateFrom    string   `json:"privateFrom"`
	PrivateFor     []string `json:"privateFor"`
	PrivateGroupID string   `json:"privateGroupID"`
}

type DoDSettleCreateOrderParam struct {
	ContractPrivacyParam
	Buyer       *DoDSettleUser              `json:"buyer" msg:"b"`
	Seller      *DoDSettleUser              `json:"seller" msg:"s"`
	Connections []*DoDSettleConnectionParam `json:"connections,omitempty" msg:"c"`
}

type DoDSettleResponseParam struct {
	ContractPrivacyParam
	RequestHash types.Hash              `json:"requestHash" msg:"-"`
	Action      DoDSettleResponseAction `json:"action" msg:"c"`
}

type DoDSettleChangeConnectionParam struct {
	ProductId string `json:"productId" msg:"p"`
	DoDSettleConnectionDynamicParam
}

type DoDSettleTerminateOrderParam struct {
	ContractPrivacyParam
	Buyer       *DoDSettleUser                    `json:"buyer" msg:"b"`
	Seller      *DoDSettleUser                    `json:"seller" msg:"s"`
	Connections []*DoDSettleChangeConnectionParam `json:"connections" msg:"c"`
}

type DoDSettleChangeOrderParam struct {
	ContractPrivacyParam
	Buyer       *DoDSettleUser                    `json:"buyer" msg:"b"`
	Seller      *DoDSettleUser                    `json:"seller" msg:"s"`
	Connections []*DoDSettleChangeConnectionParam `json:"connections" msg:"c"`
}

type DoDSettleProductInfo struct {
	OrderItemId string `json:"orderItemId" msg:"oii"`
	ProductId   string `json:"productId" msg:"pi"`
	Active      bool   `json:"active" msg:"a"`
}

type DoDSettleUpdateProductInfoParam struct {
	ContractPrivacyParam
	Address     types.Address           `json:"address" msg:"-"`
	OrderId     string                  `json:"orderId" msg:"oi"`
	ProductInfo []*DoDSettleProductInfo `json:"productInfo" msg:"p"`
}

type DoDSettleOrderItem struct {
	ItemId      string `json:"itemId" msg:"i"`
	OrderItemId string `json:"orderItemId" msg:"o"`
}

type DoDSettleUpdateOrderInfoParam struct {
	ContractPrivacyParam
	Buyer       types.Address         `json:"buyer" msg:"-"`
	InternalId  types.Hash            `json:"internalId,omitempty" msg:"i,extension"`
	OrderId     string                `json:"orderId,omitempty" msg:"oi"`
	OrderItemId []*DoDSettleOrderItem `json:"orderItemId" msg:"oii"`
	Status      DoDSettleOrderState   `json:"status,omitempty" msg:"s"`
	FailReason  string                `json:"failReason,omitempty" msg:"fr"`
}

func (s *DoDSettlementAPI) GetCreateOrderBlock(param *DoDSettleCreateOrderParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "DoDSettlement_getCreateOrderBlock", param)
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
	err := s.client.getClient().Call(&blk, "DoDSettlement_getCreateOrderRewardBlock", param)
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

func (s *DoDSettlementAPI) GetUpdateOrderInfoBlock(param *DoDSettleUpdateOrderInfoParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "DoDSettlement_getUpdateOrderInfoBlock", param)
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
	err := s.client.getClient().Call(&blk, "DoDSettlement_getUpdateOrderInfoRewardBlock", param)
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
	err := s.client.getClient().Call(&blk, "DoDSettlement_getChangeOrderBlock", param)
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
	err := s.client.getClient().Call(&blk, "DoDSettlement_getChangeOrderRewardBlock", param)
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
	err := s.client.getClient().Call(&blk, "DoDSettlement_getTerminateOrderBlock", param)
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
	err := s.client.getClient().Call(&blk, "DoDSettlement_getTerminateOrderRewardBlock", param)
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

func (s *DoDSettlementAPI) GetUpdateProductInfoBlock(param *DoDSettleUpdateProductInfoParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "DoDSettlement_getUpdateProductInfoBlock", param)
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

func (s *DoDSettlementAPI) GetUpdateProductInfoRewardBlock(param *DoDSettleResponseParam, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := s.client.getClient().Call(&blk, "DoDSettlement_getUpdateProductInfoRewardBlock", param)
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
	InternalId    string                      `json:"internalId,omitempty" msg:"i"`
	OrderType     DoDSettleOrderType          `json:"orderType,omitempty" msg:"ot"`
	OrderState    DoDSettleOrderState         `json:"orderState" msg:"os"`
	ContractState DoDSettleContractState      `json:"contractState" msg:"cs"`
	Connections   []*DoDSettleConnectionParam `json:"connections" msg:"c"`
	Track         []*DoDSettleOrderLifeTrack  `json:"track" msg:"t"`
}

type DoDSettleDisconnectInfo struct {
	OrderId      string  `json:"orderId,omitempty" msg:"oi"`
	OrderItemId  string  `json:"orderItemId,omitempty" msg:"oii"`
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
	SendHash   types.Hash              `json:"sendHash"`
	OrderId    string                  `json:"orderId"`
	InternalId types.Hash              `json:"internalId"`
	Products   []*DoDSettleProductInfo `json:"products"`
}

type DoDPlacingOrderInfo struct {
	InternalId types.Hash          `json:"internalId"`
	OrderInfo  *DoDSettleOrderInfo `json:"orderInfo"`
}

type DoDPlacingOrderResp struct {
	TotalOrders int                    `json:"totalOrders"`
	OrderList   []*DoDPlacingOrderInfo `json:"orderList"`
}

type DoDSettleInvoiceConnDynamic struct {
	DoDSettleConnectionDynamicParam
	InvoiceStartTime    int64              `json:"invoiceStartTime,omitempty"`
	InvoiceStartTimeStr string             `json:"invoiceStartTimeStr,omitempty"`
	InvoiceEndTime      int64              `json:"invoiceEndTime,omitempty"`
	InvoiceEndTimeStr   string             `json:"invoiceEndTimeStr,omitempty"`
	InvoiceUnitCount    int                `json:"invoiceUnitCount,omitempty"`
	OrderType           DoDSettleOrderType `json:"orderType,omitempty"`
	Amount              float64            `json:"amount"`
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

type DoDSettlementOrderInfoResp struct {
	OrderInfo   []*DoDSettleOrderInfo `json:"orderInfo"`
	TotalOrders int                   `json:"totalOrders"`
}

type DoDSettlementProductInfoResp struct {
	ProductInfo   []*DoDSettleConnectionInfo `json:"productInfo"`
	TotalProducts int                        `json:"totalProducts"`
}

func (s *DoDSettlementAPI) GetOrderInfoBySellerAndOrderId(seller types.Address, orderId string) (*DoDSettleOrderInfo, error) {
	var r DoDSettleOrderInfo
	err := s.client.getClient().Call(&r, "DoDSettlement_getOrderInfoBySellerAndOrderId", seller, orderId)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetOrderInfoByInternalId(internalId string) (*DoDSettleOrderInfo, error) {
	var r DoDSettleOrderInfo
	err := s.client.getClient().Call(&r, "DoDSettlement_getOrderInfoByInternalId", internalId)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetProductInfoBySellerAndProductId(seller types.Address, productId string) (*DoDSettleConnectionInfo, error) {
	var r DoDSettleConnectionInfo
	err := s.client.getClient().Call(&r, "DoDSettlement_getProductInfoBySellerAndProductId", seller, productId)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetPendingRequest(address types.Address) ([]*DoDPendingRequestRsp, error) {
	var r []*DoDPendingRequestRsp
	err := s.client.getClient().Call(&r, "DoDSettlement_getPendingRequest", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *DoDSettlementAPI) GetPendingResourceCheck(address types.Address) ([]*DoDPendingResourceCheckInfo, error) {
	var r []*DoDPendingResourceCheckInfo
	err := s.client.getClient().Call(&r, "DoDSettlement_getPendingResourceCheck", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *DoDSettlementAPI) GetPlacingOrder(buyer, seller types.Address, count, offset int) (*DoDPlacingOrderResp, error) {
	var r DoDPlacingOrderResp
	err := s.client.getClient().Call(&r, "DoDSettlement_getPlacingOrder", buyer, seller, count, offset)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetProductIdListByAddress(address types.Address) ([]*DoDSettleProduct, error) {
	var r []*DoDSettleProduct
	err := s.client.getClient().Call(&r, "DoDSettlement_getProductIdListByAddress", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *DoDSettlementAPI) GetOrderIdListByAddress(address types.Address) ([]*DoDSettleOrder, error) {
	var r []*DoDSettleOrder
	err := s.client.getClient().Call(&r, "DoDSettlement_getOrderIdListByAddress", address)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *DoDSettlementAPI) GetProductIdListByAddressAndSeller(address, seller types.Address) ([]*DoDSettleOrder, error) {
	var r []*DoDSettleOrder
	err := s.client.getClient().Call(&r, "DoDSettlement_getProductIdListByAddressAndSeller", address, seller)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *DoDSettlementAPI) GetOrderIdListByAddressAndSeller(address, seller types.Address) ([]*DoDSettleOrder, error) {
	var r []*DoDSettleOrder
	err := s.client.getClient().Call(&r, "DoDSettlement_getOrderIdListByAddressAndSeller", address, seller)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *DoDSettlementAPI) GenerateInvoiceByOrderId(seller types.Address, orderId string, start, end int64, flight, split bool) (*DoDSettleOrderInvoice, error) {
	var r DoDSettleOrderInvoice
	err := s.client.getClient().Call(&r, "DoDSettlement_generateInvoiceByOrderId", seller, orderId, start, end, flight, split)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GenerateInvoiceByBuyer(seller, buyer types.Address, start, end int64, flight, split bool) (*DoDSettleBuyerInvoice, error) {
	var r DoDSettleBuyerInvoice
	err := s.client.getClient().Call(&r, "DoDSettlement_generateInvoiceByBuyer", seller, buyer, start, end, flight, split)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GenerateInvoiceByProductId(seller types.Address, productId string, start, end int64, flight, split bool) (*DoDSettleProductInvoice, error) {
	var r DoDSettleProductInvoice
	err := s.client.getClient().Call(&r, "DoDSettlement_generateInvoiceByProductId", seller, productId, start, end, flight, split)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetOrderCountByAddress(address types.Address) int {
	var length int
	err := s.client.getClient().Call(&length, "DoDSettlement_getOrderCountByAddress", address)
	if err != nil {
		return 0
	}
	return length
}

func (s *DoDSettlementAPI) GetOrderInfoByAddress(address types.Address, count, offset int) (*DoDSettlementOrderInfoResp, error) {
	var r DoDSettlementOrderInfoResp
	err := s.client.getClient().Call(&r, "DoDSettlement_getOrderInfoByAddress", address, count, offset)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetOrderCountByAddressAndSeller(address, seller types.Address) int {
	var length int
	err := s.client.getClient().Call(&length, "DoDSettlement_getOrderCountByAddressAndSeller", address, seller)
	if err != nil {
		return 0
	}
	return length
}

func (s *DoDSettlementAPI) GetOrderInfoByAddressAndSeller(address, seller types.Address, count, offset int) (*DoDSettlementOrderInfoResp, error) {
	var r DoDSettlementOrderInfoResp
	err := s.client.getClient().Call(&r, "DoDSettlement_getOrderInfoByAddressAndSeller", address, count, offset)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetProductCountByAddress(address types.Address) int {
	var length int
	err := s.client.getClient().Call(&length, "DoDSettlement_getProductCountByAddress", address)
	if err != nil {
		return 0
	}
	return length
}

func (s *DoDSettlementAPI) GetProductInfoByAddress(address types.Address, count, offset int) (*DoDSettlementProductInfoResp, error) {
	var r DoDSettlementProductInfoResp
	err := s.client.getClient().Call(&r, "DoDSettlement_getProductInfoByAddress", address, count, offset)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetProductCountByAddressAndSeller(address, seller types.Address) int {
	var length int
	err := s.client.getClient().Call(&length, "DoDSettlement_getProductCountByAddressAndSeller", address, seller)
	if err != nil {
		return 0
	}
	return length
}

func (s *DoDSettlementAPI) GetProductInfoByAddressAndSeller(address, seller types.Address, count, offset int) (*DoDSettlementProductInfoResp, error) {
	var r DoDSettlementProductInfoResp
	err := s.client.getClient().Call(&r, "DoDSettlement_getProductInfoByAddress", address, count, offset)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DoDSettlementAPI) GetInternalIdByOrderId(seller types.Address, orderId string) (types.Hash, error) {
	var r types.Hash
	err := s.client.getClient().Call(&r, "DoDSettlement_getInternalIdByOrderId", seller, orderId)
	if err != nil {
		return types.ZeroHash, err
	}
	return r, nil
}
