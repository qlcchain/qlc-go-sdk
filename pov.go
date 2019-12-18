package qlcchain

import (
	"encoding/json"
	"fmt"
	"time"

	rpc "github.com/qlcchain/jsonrpc2"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type PovApi struct {
	url    string
	client *rpc.Client
}

type PovApiStatus struct {
	PovEnabled   bool   `json:"povEnabled"`
	SyncState    int    `json:"syncState"`
	SyncStateStr string `json:"syncStateStr"`
}

type PovApiHeader struct {
	*types.PovHeader
	AlgoName       string  `json:"algoName"`
	AlgoEfficiency uint    `json:"algoEfficiency"`
	NormBits       uint32  `json:"normBits"`
	NormDifficulty float64 `json:"normDifficulty"`
	AlgoDifficulty float64 `json:"algoDifficulty"`
}

type PovApiBatchHeader struct {
	Count   int             `json:"count"`
	Headers []*PovApiHeader `json:"headers"`
}

type PovApiBlock struct {
	*types.PovBlock
	AlgoName       string  `json:"algoName"`
	AlgoEfficiency uint    `json:"algoEfficiency"`
	NormBits       uint32  `json:"normBits"`
	NormDifficulty float64 `json:"normDifficulty"`
	AlgoDifficulty float64 `json:"algoDifficulty"`
}

type PovApiTxLookup struct {
	TxHash   types.Hash         `json:"txHash"`
	TxLookup *types.PovTxLookup `json:"txLookup"`

	CoinbaseTx *types.PovCoinBaseTx `json:"coinbaseTx"`
	AccountTx  *types.StateBlock    `json:"accountTx"`
}

type PovApiState struct {
	AccountState *types.PovAccountState `json:"accountState"`
	RepState     *types.PovRepState     `json:"repState"`
}

type PovApiHashInfo struct {
	ChainHashPS   uint64 `json:"chainHashPS"`
	Sha256dHashPS uint64 `json:"sha256dHashPS"`
	ScryptHashPS  uint64 `json:"scryptHashPS"`
	X11HashPS     uint64 `json:"x11HashPS"`
}

type PovApiGetMiningInfo struct {
	SyncState          int               `json:"syncState"`
	CurrentBlockHeight uint64            `json:"currentBlockHeight"`
	CurrentBlockHash   types.Hash        `json:"currentBlockHash"`
	CurrentBlockSize   uint32            `json:"currentBlockSize"`
	CurrentBlockTx     uint32            `json:"currentBlockTx"`
	CurrentBlockAlgo   types.PovAlgoType `json:"currentBlockAlgo"`
	PooledTx           uint32            `json:"pooledTx"`
	Difficulty         float64           `json:"difficulty"`
	HashInfo           *PovApiHashInfo   `json:"hashInfo"`
}

type PovMinerStatItem struct {
	MainBlockNum       uint32        `json:"mainBlockNum"`
	MainRewardAmount   types.Balance `json:"mainRewardAmount"`
	StableBlockNum     uint32        `json:"stableBlockNum"`
	StableRewardAmount types.Balance `json:"stableRewardAmount"`
	FirstBlockTime     time.Time     `json:"firstBlockTime"`
	LastBlockTime      time.Time     `json:"lastBlockTime"`
	FirstBlockHeight   uint64        `json:"firstBlockHeight"`
	LastBlockHeight    uint64        `json:"lastBlockHeight"`
	IsHourOnline       bool          `json:"isHourOnline"`
	IsDayOnline        bool          `json:"isDayOnline"`
}

type PovMinerStats struct {
	MinerCount      int `json:"minerCount"`
	HourOnlineCount int `json:"hourOnlineCount"`
	DayOnlineCount  int `json:"dayOnlineCount"`

	MinerStats map[types.Address]*PovMinerStatItem `json:"minerStats"`

	TotalBlockNum     uint32 `json:"totalBlockNum"`
	LatestBlockHeight uint64 `json:"latestBlockHeight"`
}

type PovRepStats struct {
	MainBlockNum       uint32        `json:"mainBlockNum"`
	MainRewardAmount   types.Balance `json:"mainRewardAmount"`
	StableBlockNum     uint32        `json:"stableBlockNum"`
	StableRewardAmount types.Balance `json:"stableRewardAmount"`
}

type PovApiGetLastNHourItem struct {
	Hour uint32

	AllBlockNum    uint32
	AllTxNum       uint32
	AllMinerReward types.Balance
	AllRepReward   types.Balance

	Sha256dBlockNum uint32
	X11BlockNum     uint32
	ScryptBlockNum  uint32
	AuxBlockNum     uint32

	MaxTxPerBlock uint32
	MinTxPerBlock uint32
	AvgTxPerBlock uint32
}

type PovApiGetLastNHourInfo struct {
	MaxTxPerBlock uint32
	MinTxPerBlock uint32
	AvgTxPerBlock uint32

	MaxTxPerHour uint32
	MinTxPerHour uint32
	AvgTxPerHour uint32

	MaxBlockPerHour uint32
	MinBlockPerHour uint32
	AvgBlockPerHour uint32

	AllBlockNum uint32
	AllTxNum    uint32

	Sha256dBlockNum uint32
	X11BlockNum     uint32
	ScryptBlockNum  uint32
	AuxBlockNum     uint32

	HourItemList []*PovApiGetLastNHourItem
}

type PovApiGetWork struct {
	WorkHash      types.Hash     `json:"workHash"`
	Version       uint32         `json:"version"`
	Previous      types.Hash     `json:"previous"`
	Bits          uint32         `json:"bits"`
	Height        uint64         `json:"height"`
	MinTime       uint32         `json:"minTime"`
	MerkleBranch  []*types.Hash  `json:"merkleBranch"`
	CoinBaseData1 types.HexBytes `json:"coinbaseData1"`
	CoinBaseData2 types.HexBytes `json:"coinbaseData2"`
}

type PovApiSubmitWork struct {
	WorkHash  types.Hash `json:"workHash"`
	BlockHash types.Hash `json:"blockHash"`

	MerkleRoot    types.Hash     `json:"merkleRoot"`
	Timestamp     uint32         `json:"timestamp"`
	Nonce         uint32         `json:"nonce"`
	CoinbaseExtra types.HexBytes `json:"coinbaseExtra"`
	CoinbaseHash  types.Hash     `json:"coinbaseHash"`

	AuxPow *types.PovAuxHeader `json:"auxPow"`
}

// NewPovAPI creates pov module for client
func NewPovAPI(url string, c *rpc.Client) *PovApi {
	return &PovApi{url: url, client: c}
}

// GetFittestHeader returns fittest pov header info
// If node is in pov syncing, will return error
func (p *PovApi) GetPovStatus() (*PovApiStatus, error) {
	var rspData PovApiStatus
	err := p.client.Call(&rspData, "pov_getPovStatus")
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetFittestHeader returns fittest pov header info
// If node is in pov syncing, will return error
func (p *PovApi) GetFittestHeader(gap uint64) (*PovApiHeader, error) {
	var rspData PovApiHeader
	err := p.client.Call(&rspData, "pov_getFittestHeader", gap)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetLatestHeader returns latest pov header info
func (p *PovApi) GetLatestHeader() (*PovApiHeader, error) {
	var rspData PovApiHeader
	err := p.client.Call(&rspData, "pov_getLatestHeader")
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetHeaderByHeight returns pov header info by height
func (p *PovApi) GetHeaderByHeight(height uint64) (*PovApiHeader, error) {
	var rspData PovApiHeader
	err := p.client.Call(&rspData, "pov_getHeaderByHeight", height)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetHeaderByHash returns pov header info by hash
func (p *PovApi) GetHeaderByHash(blockHash types.Hash) (*PovApiHeader, error) {
	var rspData PovApiHeader
	err := p.client.Call(&rspData, "pov_getHeaderByHash", blockHash)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// BatchGetHeadersByHeight returns a lots of pov headers info by range
func (p *PovApi) BatchGetHeadersByHeight(height uint64, count uint64, asc bool) (*PovApiBatchHeader, error) {
	var rspData PovApiBatchHeader
	err := p.client.Call(&rspData, "pov_batchGetHeadersByHeight", height, count, asc)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetLatestBlock returns latest pov block info
func (p *PovApi) GetLatestBlock(txOffset uint32, txLimit uint32) (*PovApiBlock, error) {
	var rspData PovApiBlock
	err := p.client.Call(&rspData, "pov_getLatestBlock", txOffset, txLimit)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetBlockByHash returns pov block info by hash
func (p *PovApi) GetBlockByHash(blockHash types.Hash, txOffset uint32, txLimit uint32) (*PovApiBlock, error) {
	var rspData PovApiBlock
	err := p.client.Call(&rspData, "pov_getBlockByHash", blockHash, txOffset, txLimit)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetBlockByHeight returns pov block info by height
func (p *PovApi) GetBlockByHeight(height uint64, txOffset uint32, txLimit uint32) (*PovApiBlock, error) {
	var rspData PovApiBlock
	err := p.client.Call(&rspData, "pov_getBlockByHeight", height, txOffset, txLimit)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetTransaction returns pov tx lookup info by tx hash
func (p *PovApi) GetTransaction(txHash types.Hash) (*PovApiTxLookup, error) {
	var rspData PovApiTxLookup
	err := p.client.Call(&rspData, "pov_getTransaction", txHash)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetTransactionByBlockHashAndIndex returns pov tx lookup info by block hash and tx index
func (p *PovApi) GetTransactionByBlockHashAndIndex(blockHash types.Hash, index uint32) (*PovApiTxLookup, error) {
	var rspData PovApiTxLookup
	err := p.client.Call(&rspData, "pov_getTransactionByBlockHashAndIndex", blockHash, index)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetTransactionByBlockHeightAndIndex returns pov tx lookup info by block height and tx index
func (p *PovApi) GetTransactionByBlockHeightAndIndex(height uint64, index uint32) (*PovApiTxLookup, error) {
	var rspData PovApiTxLookup
	err := p.client.Call(&rspData, "pov_getTransactionByBlockHeightAndIndex", height, index)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetLatestAccountState returns pov account state in latest block
func (p *PovApi) GetLatestAccountState(address types.Address) (*PovApiState, error) {
	var rspData PovApiState
	err := p.client.Call(&rspData, "pov_getLatestAccountState", address)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetAccountStateByBlockHash returns pov account state by block hash
func (p *PovApi) GetAccountStateByBlockHash(address types.Address, blockHash types.Hash) (*PovApiState, error) {
	var rspData PovApiState
	err := p.client.Call(&rspData, "pov_getAccountStateByBlockHash", address, blockHash)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetAccountStateByBlockHeight returns pov account state by block height
func (p *PovApi) GetAccountStateByBlockHeight(address types.Address, height uint64) (*PovApiState, error) {
	var rspData PovApiState
	err := p.client.Call(&rspData, "pov_getAccountStateByBlockHeight", address, height)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetHashInfo returns pov network hash info
func (p *PovApi) GetHashInfo(height uint64, lookup uint64) (*PovApiHashInfo, error) {
	var rspData PovApiHashInfo
	err := p.client.Call(&rspData, "pov_getHashInfo", height, lookup)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetMiningInfo returns pov mining info
func (p *PovApi) GetMiningInfo() (*PovApiGetMiningInfo, error) {
	var rspData PovApiGetMiningInfo
	err := p.client.Call(&rspData, "pov_getMiningInfo")
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetMinerStats returns pov miner statistic
func (p *PovApi) GetMinerStats(addrs []types.Address) (*PovMinerStats, error) {
	var rspData PovMinerStats
	err := p.client.Call(&rspData, "pov_getMinerStats", addrs)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetMinerDayStat returns pov miner day statistic
func (p *PovApi) GetMinerDayStat(dayIndex int) (*types.PovMinerDayStat, error) {
	var rspData types.PovMinerDayStat
	err := p.client.Call(&rspData, "pov_getMinerDayStat", dayIndex)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetMinerDayStatByHeight returns pov miner day statistic
func (p *PovApi) GetMinerDayStatByHeight(height uint64) (*types.PovMinerDayStat, error) {
	var rspData types.PovMinerDayStat
	err := p.client.Call(&rspData, "pov_getMinerDayStatByHeight", height)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetDiffDayStat returns pov difficulty day statistic
func (p *PovApi) GetDiffDayStat(dayIndex int) (*types.PovDiffDayStat, error) {
	var rspData types.PovDiffDayStat
	err := p.client.Call(&rspData, "pov_getDiffDayStat", dayIndex)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetDiffDayStatByHeight returns pov difficulty day statistic
func (p *PovApi) GetDiffDayStatByHeight(height uint64) (*types.PovDiffDayStat, error) {
	var rspData types.PovDiffDayStat
	err := p.client.Call(&rspData, "pov_getDiffDayStatByHeight", height)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetRepStats returns pov rep statistic
func (p *PovApi) GetRepStats(addrs []types.Address) (map[types.Address]*PovRepStats, error) {
	var rspData map[types.Address]*PovRepStats
	err := p.client.Call(&rspData, "pov_getRepStats", addrs)
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

// GetLastNHourInfo returns pov last n hour statistic
func (p *PovApi) GetLastNHourInfo(endHeight uint64, timeSpan uint32) (*PovApiGetLastNHourInfo, error) {
	var rspData PovApiGetLastNHourInfo
	err := p.client.Call(&rspData, "pov_getLastNHourInfo", endHeight, timeSpan)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// GetWork returns pov next block work info
// If node is in pov syncing, will return error
func (p *PovApi) GetWork(minerAddr types.Address, algoName string) (*PovApiGetWork, error) {
	var rspData PovApiGetWork
	err := p.client.Call(&rspData, "pov_getWork", minerAddr, algoName)
	if err != nil {
		return nil, err
	}
	return &rspData, nil
}

// SubmitWork sumbits new block work to node
// If node is in pov syncing, will return error
func (p *PovApi) SubmitWork(work *PovApiSubmitWork) error {
	err := p.client.Call(nil, "pov_submitWork", work)
	if err != nil {
		return err
	}
	return nil
}

// NewBlock support publish/subscription, ch is PovApiHeader channel,
// once there is new block stored to the chain, set the block to channel
func (p *PovApi) SubscribeNewBlock(ch chan *PovApiHeader) (*Subscribe, error) {
	subscribe := NewSubscribe(p.url)
	request := `{"id":1,"method":"pov_subscribe","params":["newBlock"]}`
	if err := subscribe.subscribe(request); err != nil {
		return nil, fmt.Errorf("subscribe fail: %s", err)
	}

	go func() {
		for {
			if result, stopped := subscribe.publish(); !stopped {
				rBytes, err := json.Marshal(result)
				if err != nil {
					fmt.Println(err)
					continue
				}
				blk := new(PovApiHeader)
				err = json.Unmarshal(rBytes, &blk)
				if err != nil {
					fmt.Println(err)
					continue
				}
				ch <- blk
			} else {
				break
			}
		}
	}()
	return subscribe, nil
}

// Unsubscribe close a pub-sub connection
func (p *PovApi) Unsubscribe(subscribe *Subscribe) error {
	request := fmt.Sprintf(`{"id":1,"method":"pov_unsubscribe","params":["%s"]}`, subscribe.subscribeID)
	return subscribe.Unsubscribe(request)
}
