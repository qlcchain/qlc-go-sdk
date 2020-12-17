package qlcchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	"github.com/qlcchain/qlc-go-sdk/pkg/util"
)

type LedgerApi struct {
	url        string
	subscribes map[types.Address]*BlockSubscription
	client     *rpc.Client
}

type APIBlock struct {
	*types.StateBlock
	TokenName string        `json:"tokenName"`
	Amount    types.Balance `json:"amount"`
	Hash      types.Hash    `json:"hash"`

	PovConfirmHeight uint64 `json:"povConfirmHeight"`
	PovConfirmCount  uint64 `json:"povConfirmCount"`
}

type APIAccount struct {
	Address        types.Address   `json:"account"`
	CoinBalance    *types.Balance  `json:"coinBalance,omitempty"`
	CoinVote       *types.Balance  `json:"vote,omitempty"`
	CoinNetwork    *types.Balance  `json:"network,omitempty"`
	CoinStorage    *types.Balance  `json:"storage,omitempty"`
	CoinOracle     *types.Balance  `json:"oracle,omitempty"`
	Representative *types.Address  `json:"representative,omitempty"`
	Tokens         []*APITokenMeta `json:"tokens"`
}

type APIAccountsBalance struct {
	Balance types.Balance  `json:"balance"`
	Vote    *types.Balance `json:"vote,omitempty"`
	Network *types.Balance `json:"network,omitempty"`
	Storage *types.Balance `json:"storage,omitempty"`
	Oracle  *types.Balance `json:"oracle,omitempty"`
	Pending types.Balance  `json:"pending"`
}

type APITokenMeta struct {
	*types.TokenMeta
	TokenName string        `json:"tokenName"`
	Pending   types.Balance `json:"pending"`
}

type APIPending struct {
	*types.PendingKey
	*types.PendingInfo
	TokenName string          `json:"tokenName"`
	Timestamp int64           `json:"timestamp"`
	BlockType types.BlockType `json:"blockType"`
}

type ApiTokenInfo struct {
	types.TokenInfo
}

type APIAccountBalance struct {
	Address types.Address `json:"address"`
	Balance types.Balance `json:"balance"`
}

type APIRepresentative struct {
	Address types.Address `json:"address"`
	Balance types.Balance `json:"balance"`
	Vote    types.Balance `json:"vote"`
	Network types.Balance `json:"network"`
	Storage types.Balance `json:"storage"`
	Oracle  types.Balance `json:"oracle"`
	Total   types.Balance `json:"total"`
}

type APISendBlockPara struct {
	From      types.Address `json:"from"`
	TokenName string        `json:"tokenName"`
	To        types.Address `json:"to"`
	Amount    types.Balance `json:"amount"`
	Sender    string        `json:"sender"`
	Receiver  string        `json:"receiver"`
	Message   types.Hash    `json:"message"`
}

// NewLedgerAPI creates ledger module for client
func NewLedgerAPI(url string, c *rpc.Client) *LedgerApi {
	return &LedgerApi{
		url:    url,
		client: c,

		subscribes: make(map[types.Address]*BlockSubscription),
	}
}

func (l *LedgerApi) Stop() {
	for _, s := range l.subscribes {
		s.subscribe.Close()
	}
}

// AccountBlocksCount returns number of blocks for a specific account of chain
func (l *LedgerApi) AccountBlocksCount(address types.Address) (int64, error) {
	var count int64
	err := l.client.Call(&count, "ledger_accountBlocksCount", address)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// AccountHistoryTopn returns blocks list for a specific account of chain
// count is number of blocks to return, and offset is index of block where to start
func (l *LedgerApi) AccountHistoryTopn(address types.Address, count int, offset int) ([]*APIBlock, error) {
	var blocks []*APIBlock
	err := l.client.Call(&blocks, "ledger_accountHistoryTopn", address, count, offset)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

// AccountInfo returns account detail info, include each token meta for the account
// If account not found, will return error
func (l *LedgerApi) AccountInfo(address types.Address) (*APIAccount, error) {
	var aa APIAccount
	err := l.client.Call(&aa, "ledger_accountInfo", address)
	if err != nil {
		return nil, err
	}
	return &aa, nil
}

// AccountRepresentative returns the representative address for account
// If account not found, will return error
func (l *LedgerApi) AccountRepresentative(address types.Address) (types.Address, error) {
	var addr types.Address
	err := l.client.Call(&addr, "ledger_accountRepresentative", address)
	if err != nil {
		return types.ZeroAddress, err
	}
	return addr, nil

}

// AccountVotingWeight returns the voting weight for account
// If account not found, will return error
func (l *LedgerApi) AccountVotingWeight(address types.Address) (types.Balance, error) {
	var amount types.Balance
	err := l.client.Call(&amount, "ledger_accountRepresentative", address)
	if err != nil {
		return types.ZeroBalance, err
	}
	return amount, nil
}

// AccountsBalance returns balance and pending(amount that has not yet been received) for each account
func (l *LedgerApi) AccountsBalance(addresses []types.Address) (map[types.Address]map[string]*APIAccountsBalance, error) {
	var r map[types.Address]map[string]*APIAccountsBalance
	err := l.client.Call(&r, "ledger_accountsBalance", addresses)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// AccountsFrontiers returns frontier info for each token of account
func (l *LedgerApi) AccountsFrontiers(addresses []types.Address) (map[types.Address]map[string]types.Hash, error) {
	var r map[types.Address]map[string]types.Hash
	err := l.client.Call(&r, "ledger_accountsFrontiers", addresses)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// AccountsPending returns pending info list for each account
// maximum number of pending for each account return is n, and if n set to -1, will return all pending for each account
func (l *LedgerApi) AccountsPending(addresses []types.Address, n int) (map[types.Address][]*APIPending, error) {
	var r map[types.Address][]*APIPending
	err := l.client.Call(&r, "ledger_accountsPending", addresses, n)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// AccountsCount returns total number of accounts of chain
func (l *LedgerApi) AccountsCount() (uint64, error) {
	var count uint64
	err := l.client.Call(&count, "ledger_accountsCount")
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Accounts returns accounts list of chain
// count is number of accounts to return, and offset is index of account where to start
func (l *LedgerApi) Accounts(count int, offset int) ([]types.Address, error) {
	var r []types.Address
	err := l.client.Call(&r, "ledger_accounts", count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BlockAccount accepts a block hash, and returns account of block owner
func (l *LedgerApi) BlockAccount(hash types.Hash) (types.Address, error) {
	var address types.Address
	err := l.client.Call(&address, "ledger_blockAccount", hash)
	if err != nil {
		return types.ZeroAddress, err
	}
	return address, nil
}

// BlockHash return hash of block
func (l *LedgerApi) BlockHash(block types.StateBlock) types.Hash {
	return block.GetHash()
}

// BlocksCount returns the number of blocks(include smartcontract block) and unchecked blocks of chain
func (l *LedgerApi) BlocksCount() (map[string]uint64, error) {
	var r map[string]uint64
	err := l.client.Call(&r, "ledger_blocksCount")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BlocksCountByType returns number of blocks by type of chain
func (l *LedgerApi) BlocksCountByType() (map[string]uint64, error) {
	var r map[string]uint64
	err := l.client.Call(&r, "ledger_blocksCountByType")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Return block confirmed status, if block confirmed，return `true`，otherwise return `false`
func (l *LedgerApi) BlockConfirmedStatus(hash types.Hash) (bool, error) {
	var r bool
	err := l.client.Call(&r, "ledger_blockConfirmedStatus", hash)
	if err != nil {
		return false, err
	}
	return r, nil
}

// BlockInfo accepts a block hash, and returns block info for the hash
func (l *LedgerApi) BlockInfo(hash types.Hash) (*APIBlock, error) {
	b, err := l.BlocksInfo([]types.Hash{hash})
	if err != nil {
		return nil, err
	}
	if len(b) < 1 {
		return nil, errors.New("block not found")
	}
	return b[0], nil

}

// BlocksInfo accepts blocks hash list, and returns block info for each hash
func (l *LedgerApi) BlocksInfo(hash []types.Hash) ([]*APIBlock, error) {
	var ab []*APIBlock
	err := l.client.Call(&ab, "ledger_blocksInfo", hash)
	if err != nil {
		return nil, err
	}
	return ab, nil
}

// Blocks returns blocks list of chain
// count is number of blocks to return, and offset is index of block where to start
func (l *LedgerApi) Blocks(count int, offset int) ([]*APIBlock, error) {
	var r []*APIBlock
	err := l.client.Call(&r, "ledger_blocks", count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Return confirmed account detail info , include each token in the account
func (l *LedgerApi) ConfirmedAccountInfo(address types.Address) (*APIAccount, error) {
	var aa APIAccount
	err := l.client.Call(&aa, "ledger_confirmedAccountInfo", address)
	if err != nil {
		return nil, err
	}
	return &aa, nil
}

// Chain returns a consecutive block hash list for a specific hash
// maximum number of blocks hash to return is n, and if n set to -1, will return blocks hash to the open block
func (l *LedgerApi) Chain(hash types.Hash, n int) ([]types.Hash, error) {
	var r []types.Hash
	err := l.client.Call(&r, "ledger_chain", hash, n)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Delegators accepts a representative account, and returns its delegator and each delegator's balance
func (l *LedgerApi) Delegators(hash types.Address) ([]*APIAccountBalance, error) {
	var r []*APIAccountBalance
	err := l.client.Call(&r, "ledger_delegators", hash)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// DelegatorsCount gets number of delegators for specific representative account
func (l *LedgerApi) DelegatorsCount(hash types.Address) (int64, error) {
	var count int64
	err := l.client.Call(&count, "ledger_delegatorsCount", hash)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type Signature func(hash types.Hash) (types.Signature, error)

func SignatureFunc(account *types.Account, hash types.Hash) (types.Signature, error) {
	return account.Sign(hash), nil
}

func generateWork(hash types.Hash) types.Work {
	var work types.Work
	worker, _ := types.NewWorker(work, hash)
	return worker.NewWork()
}

func phoneNumberSeri(number string) []byte {
	if number == "" {
		return nil
	}
	b := util.String2Bytes(number)
	return b
}

// GenerateSendBlock returns send block by transaction parameter, sign is a function to sign the block
func (l *LedgerApi) GenerateSendBlock(para *APISendBlockPara, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_generateSendBlock", para)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	blk.Work = generateWork(blk.Root())
	return &blk, nil
}

func (l *LedgerApi) GenerateAndProcessSendBlock(para *APISendBlockPara, sign Signature) (types.Hash, error) {
	blk, err := l.GenerateSendBlock(para, sign)
	if err != nil {
		return types.ZeroHash, err
	}
	return l.Process(blk)
}

// GenerateReceiveBlock returns receive block by send block, sign is a function to sign the block
func (l *LedgerApi) GenerateReceiveBlock(txBlock *types.StateBlock, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_generateReceiveBlock", txBlock)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	blk.Work = generateWork(blk.Root())
	return &blk, nil
}

func (l *LedgerApi) GenerateAndProcessReceiveBlock(txBlock *types.StateBlock, sign Signature) (types.Hash, error) {
	blk, err := l.GenerateReceiveBlock(txBlock, sign)
	if err != nil {
		return types.ZeroHash, err
	}
	return l.Process(blk)
}

// GenerateReceiveBlockByHash returns receive block by send block hash, sign is a function to sign the block
func (l *LedgerApi) GenerateReceiveBlockByHash(txHash types.Hash, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_generateReceiveBlockByHash", txHash)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	blk.Work = generateWork(blk.Root())
	return &blk, nil
}

func (l *LedgerApi) GenerateAndProcessReceiveBlockByHash(txHash types.Hash, sign Signature) (types.Hash, error) {
	blk, err := l.GenerateReceiveBlockByHash(txHash, sign)
	if err != nil {
		return types.ZeroHash, err
	}
	return l.Process(blk)
}

// GenerateChangeBlock returns change block by account and new representative address, sign is a function to sign the block
func (l *LedgerApi) GenerateChangeBlock(account, representative types.Address, sign Signature) (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_generateChangeBlock", account, representative)
	if err != nil {
		return nil, err
	}
	if sign != nil {
		blk.Signature, err = sign(blk.GetHash())
		if err != nil {
			return nil, err
		}
	}
	blk.Work = generateWork(blk.Root())
	return &blk, nil
}

func (l *LedgerApi) GenerateAndProcessChangeBlock(account, representative types.Address, sign Signature) (types.Hash, error) {
	blk, err := l.GenerateChangeBlock(account, representative, sign)
	if err != nil {
		return types.ZeroHash, err
	}
	return l.Process(blk)
}

// Process checks block base info , updates info of chain for the block ,and broadcasts block
func (l *LedgerApi) Process(block *types.StateBlock) (types.Hash, error) {
	var hash types.Hash
	err := l.client.Call(&hash, "ledger_process", block)
	if err != nil {
		return types.ZeroHash, err
	}
	return hash, nil
}

// func (l *LedgerApi) ProcessAndConfirmed(block *types.StateBlock) (bool, error) {
//	var hash types.Hash
//	err := l.client.Call(&hash, "ledger_process", block)
//	if err != nil {
//		return false, err
//	}
//
//	ch := make(chan *types.StateBlock)
//	subscribe, err := l.BlockSubscription(block.GetAddress())
//	if err != nil {
//		return false, err
//	}
//	subscribe.addChan(ch)
//	defer subscribe.removeChan(ch)
//
//	ticker := time.NewTicker(180 * time.Second)
//	for {
//		select {
//		case blk := <-ch:
//			if blk.GetHash() == block.GetHash() {
//				return true, nil
//			}
//		case <-ticker.C:
//			return false, errors.New("consensus timeout")
//		}
//	}
//}

// Pendings returns pending transaction list on chain
func (l *LedgerApi) Pendings() ([]*APIPending, error) {
	var r []*APIPending
	err := l.client.Call(&r, "ledger_pendings")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Performance returns performance time
func (l *LedgerApi) Performance() ([]*types.PerformanceTime, error) {
	var r []*types.PerformanceTime
	err := l.client.Call(&r, "ledger_performance")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Pending return pending info by account and token hash, if pending not found, return error
func (l *LedgerApi) Pending(address types.Address, hash types.Hash) (*APIPending, error) {
	pendings, err := l.AccountsPending([]types.Address{address}, -1)
	if err != nil {
		return nil, err
	}
	rxPendings := pendings[address]
	for _, p := range rxPendings {
		if p.Hash == hash {
			return p, nil
		}
	}
	return nil, errors.New("pending not found")
}

// Representatives returns pairs of representative and its voting weight of chain
// if set sorting false , will return representatives randomly, if set true,
// will sorting representative balance in descending order
func (l *LedgerApi) Representatives(sorting bool) ([]*APIRepresentative, error) {
	var r []*APIRepresentative
	err := l.client.Call(&r, "ledger_representatives", sorting)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// TokenMeta return tokenmeta info by account and token hash
func (l *LedgerApi) TokenMeta(hash types.Hash, address types.Address) (*APITokenMeta, error) {
	am, err := l.AccountInfo(address)
	if err != nil {
		return nil, err
	}
	for _, t := range am.Tokens {
		if t.Type == hash {
			return t, nil
		}
	}
	return nil, fmt.Errorf("account [%s] does not have the  token [%s]", address.String(), hash.String())
}

// Tokens return all token info of chain
func (l *LedgerApi) Tokens() ([]*types.TokenInfo, error) {
	var r []*types.TokenInfo
	err := l.client.Call(&r, "ledger_tokens")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// TransactionsCount returns the number of blocks(not include smartcontract block) and unchecked blocks of chain
func (l *LedgerApi) TransactionsCount() (map[string]uint64, error) {
	var r map[string]uint64
	err := l.client.Call(&r, "ledger_transactionsCount")
	if err != nil {
		return nil, err
	}
	return r, nil
}

// TokenInfoById returns token info by token id
func (l *LedgerApi) TokenInfoById(tokenId types.Hash) (*ApiTokenInfo, error) {
	var at ApiTokenInfo
	err := l.client.Call(&at, "ledger_tokenInfoById", tokenId)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

// TokenInfoById returns token info by token name
func (l *LedgerApi) TokenInfoByName(tokenName string) (*ApiTokenInfo, error) {
	var at ApiTokenInfo
	err := l.client.Call(&at, "ledger_tokenInfoByName", tokenName)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

// NewBlock support publish/subscription, ch is StateBlock channel,
// once there is new block stored to the chain, set the block to channel
func (l *LedgerApi) NewBlock(ch chan *types.StateBlock) (*Subscribe, error) {
	subscribe := NewSubscribe(l.url)
	request := `{"id":1,"method":"ledger_subscribe","params":["newBlock"]}`
	if err := subscribe.subscribe(request); err != nil {
		return nil, fmt.Errorf("subscribe fail: %s", err)
	}

	go func() {
		for {
			if r, stopped := subscribe.publish(); !stopped {
				rBytes, err := json.Marshal(r)
				if err != nil {
					fmt.Println(err)
					continue
				}

				block := new(types.StateBlock)
				err = json.Unmarshal(rBytes, &block)
				if err != nil {
					fmt.Println(err)
					continue
				}

				ch <- block
			} else {
				break
			}
		}
	}()
	return subscribe, nil
}

// NewAccountBlock support publish/subscription, ch is StateBlock channel,
// once there is new account block stored to the chain, set the block to channel
func (l *LedgerApi) NewAccountBlock(ch chan *types.StateBlock, address types.Address) (*Subscribe, error) {
	subscribe := NewSubscribe(l.url)
	request := fmt.Sprintf(`{"id":1,"method":"ledger_subscribe","params":["newAccountBlock","%s"]}`, address)
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
				blk := new(types.StateBlock)
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

// BalanceChange support publish/subscription, ch is AccountMeta channel,
// once the balance of a account change, set the newest account info to channel
func (l *LedgerApi) BalanceChange(ch chan *types.AccountMeta, address types.Address) (*Subscribe, error) {
	subscribe := NewSubscribe(l.url)
	request := fmt.Sprintf(`{"id":1,"method":"ledger_subscribe","params":["balanceChange","%s"]}`, address)
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
				am := new(types.AccountMeta)
				err = json.Unmarshal(rBytes, &am)
				if err != nil {
					fmt.Println(err)
					continue
				}
				ch <- am
			} else {
				break
			}
		}
	}()
	return subscribe, nil
}

// NewPending support publish/subscription, ch is APIPending channel,
// once there is a pending transaction of a account, set the pending info to channel
func (l *LedgerApi) NewPending(ch chan *APIPending, address types.Address) (*Subscribe, error) {
	subscribe := NewSubscribe(l.url)
	request := fmt.Sprintf(`{"id":1,"method":"ledger_subscribe","params":["newPending","%s"]}`, address)
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
				am := new(APIPending)
				err = json.Unmarshal(rBytes, &am)
				if err != nil {
					fmt.Println(err)
					continue
				}
				ch <- am
			} else {
				break
			}
		}
	}()
	return subscribe, nil
}

// Unsubscribe close a pub-sub connection
func (l *LedgerApi) Unsubscribe(subscribe *Subscribe) error {
	request := fmt.Sprintf(`{"id":1,"method":"ledger_unsubscribe","params":["%s"]}`, subscribe.subscribeID)
	return subscribe.Unsubscribe(request)
}

type BlockSubscription struct {
	mu        *sync.Mutex
	subscribe *Subscribe
	chans     []chan *types.StateBlock
	blocks    chan *types.StateBlock
	stoped    chan bool
}

func (l *LedgerApi) BlockSubscription(address types.Address) (*BlockSubscription, error) {
	s, ok := l.subscribes[address]
	if !ok {
		ch := make(chan *types.StateBlock)
		subscribe, err := l.NewAccountBlock(ch, address)
		if err != nil {
			return nil, err
		}

		sub := &BlockSubscription{
			mu:        &sync.Mutex{},
			chans:     make([]chan *types.StateBlock, 0),
			blocks:    make(chan *types.StateBlock, 100),
			stoped:    make(chan bool),
			subscribe: subscribe,
		}
		l.subscribes[address] = sub
		go func() {
			for {
				select {
				case block := <-ch:
					if len(sub.chans) > 0 {
						sub.blocks <- block
					}
				case <-subscribe.Stopped:
					sub.stoped <- true
					return
				}
			}
		}()
		go func() {
			for {
				select {
				case b := <-sub.blocks:
					for _, c := range sub.chans {
						c <- b
					}
				case <-sub.stoped:
					return
				}
			}
		}()
		return sub, nil
	}
	return s, nil
}

func (r *BlockSubscription) addChan(ch chan *types.StateBlock) {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()
	r.chans = append(r.chans, ch)
}

func (r *BlockSubscription) removeChan(ch chan *types.StateBlock) {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()
	for index, c := range r.chans {
		if c == ch {
			r.chans = append(r.chans[:index], r.chans[index+1:]...)
			break
		}
	}
}

func (l *LedgerApi) GenesisAddress() (*types.Address, error) {
	var addr types.Address
	err := l.client.Call(&addr, "ledger_genesisAddress")
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func (l *LedgerApi) GasAddress() (*types.Address, error) {
	var addr types.Address
	err := l.client.Call(&addr, "ledger_gasAddress")
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func (l *LedgerApi) ChainToken() (*types.Hash, error) {
	var h types.Hash
	err := l.client.Call(&h, "ledger_chainToken")
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (l *LedgerApi) GasToken() (*types.Hash, error) {
	var h types.Hash
	err := l.client.Call(&h, "ledger_gasToken")
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (l *LedgerApi) GenesisMintageBlock() (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_genesisMintageBlock")
	if err != nil {
		return nil, err
	}
	return &blk, nil
}

func (l *LedgerApi) GenesisMintageHash() (*types.Hash, error) {
	var h types.Hash
	err := l.client.Call(&h, "ledger_genesisMintageHash")
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (l *LedgerApi) GenesisBlock() (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_genesisBlock")
	if err != nil {
		return nil, err
	}
	return &blk, nil
}

func (l *LedgerApi) GenesisBlockHash() (*types.Hash, error) {
	var h types.Hash
	err := l.client.Call(&h, "ledger_genesisBlockHash")
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (l *LedgerApi) GasBlockHash() (*types.Hash, error) {
	var h types.Hash
	err := l.client.Call(&h, "ledger_gasBlockHash")
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (l *LedgerApi) GasMintageBlock() (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_gasMintageBlock")
	if err != nil {
		return nil, err
	}
	return &blk, nil
}

func (l *LedgerApi) GasBlock() (*types.StateBlock, error) {
	var blk types.StateBlock
	err := l.client.Call(&blk, "ledger_gasBlock")
	if err != nil {
		return nil, err
	}
	return &blk, nil
}

func (l *LedgerApi) IsGenesisBlock() (*bool, error) {
	var b bool
	err := l.client.Call(&b, "ledger_isGenesisBlock")
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (l *LedgerApi) IsGenesisToken() (*bool, error) {
	var b bool
	err := l.client.Call(&b, "ledger_isGenesisToken")
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (l *LedgerApi) AllGenesisBlocks() ([]*types.StateBlock, error) {
	var blks []*types.StateBlock
	err := l.client.Call(&blks, "ledger_allGenesisBlocks")
	if err != nil {
		return nil, err
	}
	return blks, nil
}
