package module

import (
	"errors"
	"fmt"
	"time"

	"github.com/qlcchain/go-qlc/common"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/common/util"
	"github.com/qlcchain/go-qlc/rpc"
	"github.com/qlcchain/go-qlc/rpc/api"
)

type LedgerApi struct {
	client *rpc.Client
}

func NewLedgerApi(c *rpc.Client) *LedgerApi {
	return &LedgerApi{client: c}
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
// If offset not set ,default is 0
func (l *LedgerApi) AccountHistoryTopn(address types.Address, count int, offset *int) ([]*api.APIBlock, error) {
	var blocks []*api.APIBlock
	err := l.client.Call(&blocks, "ledger_accountHistoryTopn", address, count, offset)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

// AccountInfo returns account detail info, include each token meta for the account
// If account not found, will return error
func (l *LedgerApi) AccountInfo(address types.Address) (*api.APIAccount, error) {
	var aa api.APIAccount
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
func (l *LedgerApi) AccountsBalance(addresses []types.Address) (map[types.Address]map[string]map[string]types.Balance, error) {
	var r map[types.Address]map[string]map[string]types.Balance
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
func (l *LedgerApi) AccountsPending(addresses []types.Address, n int) (map[types.Address][]*api.APIPending, error) {
	var r map[types.Address][]*api.APIPending
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
// If offset not set ,default is 0
func (l *LedgerApi) Accounts(count int, offset *int) ([]*types.Address, error) {
	var r []*types.Address
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

// BlocksInfo accepts blocks hash list, and returns block info for each hash
func (l *LedgerApi) BlocksInfo(hash []types.Hash) ([]*api.APIBlock, error) {
	var ab []*api.APIBlock
	err := l.client.Call(&ab, "ledger_blocksInfo", hash)
	if err != nil {
		return nil, err
	}
	return ab, nil
}

// Blocks returns blocks list of chain
// count is number of blocks to return, and offset is index of block where to start
// If offset not set ,default is 0
func (l *LedgerApi) Blocks(count int, offset *int) ([]*api.APIBlock, error) {
	var r []*api.APIBlock
	err := l.client.Call(&r, "ledger_blocks", count, offset)
	if err != nil {
		return nil, err
	}
	return r, nil
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
func (l *LedgerApi) Delegators(hash types.Address) ([]*api.APIAccountBalance, error) {
	var r []*api.APIAccountBalance
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
func (l *LedgerApi) GenerateSendBlock(para *api.APISendBlockPara, sign Signature) (*types.StateBlock, error) {
	info, err := l.TokenInfoByName(para.TokenName)
	if err != nil {
		return nil, err
	}
	tm, err := l.token(info.TokenId, para.From)
	if err != nil {
		return nil, errors.New("token not found")
	}
	if tm.Balance.Compare(para.Amount) != types.BalanceCompSmaller {
		blk := types.StateBlock{
			Type:           types.Send,
			Address:        para.From,
			Token:          info.TokenId,
			Balance:        tm.Balance.Sub(para.Amount),
			Previous:       tm.Header,
			Link:           para.To.ToHash(),
			Representative: tm.Representative,
			Sender:         phoneNumberSeri(para.Sender),
			Receiver:       phoneNumberSeri(para.Receiver),
			Message:        para.Message,
			Timestamp:      time.Now().Unix(),
		}
		if sign != nil {
			blk.Signature, err = sign(blk.GetHash())
			if err != nil {
				return nil, err
			}
		}
		blk.Work = generateWork(blk.Root())
		return &blk, nil
	} else {
		return nil, fmt.Errorf("not enought balance(%s) of %s", tm.Balance, para.Amount)
	}
}

func (l *LedgerApi) token(hash types.Hash, address types.Address) (*api.APITokenMeta, error) {
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

func (l *LedgerApi) blockInfo(hash types.Hash) (*api.APIBlock, error) {
	b, err := l.BlocksInfo([]types.Hash{hash})
	if err != nil {
		return nil, err
	}
	if len(b) < 1 {
		return nil, errors.New("block not found")
	}
	return b[0], nil

}

func (l *LedgerApi) pending(address types.Address, hash types.Hash) (*api.APIPending, error) {
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

// GenerateReceiveBlock returns receive block by send block, sign is a function to sign the block
func (l *LedgerApi) GenerateReceiveBlock(txBlock *types.StateBlock, sign Signature) (*types.StateBlock, error) {
	if !txBlock.GetType().Equal(types.Send) {
		return nil, fmt.Errorf("(%s) is not send block", txBlock.GetHash().String())
	}
	return l.GenerateReceiveBlockByHash(txBlock.GetHash(), sign)
}

// GenerateReceiveBlockByHash returns receive block by send block hash, sign is a function to sign the block
func (l *LedgerApi) GenerateReceiveBlockByHash(txHash types.Hash, sign Signature) (*types.StateBlock, error) {
	txBlock, err := l.blockInfo(txHash)
	if err != nil {
		return nil, err
	}

	rcAddress := types.Address(txBlock.Link)
	pending, err := l.pending(rcAddress, txHash)
	if err != nil {
		return nil, err
	}
	rcTm, _ := l.token(txBlock.GetToken(), rcAddress)

	var blk types.StateBlock
	if rcTm != nil {
		blk = types.StateBlock{
			Type:           types.Receive,
			Address:        rcAddress,
			Balance:        rcTm.Balance.Add(pending.Amount),
			Previous:       rcTm.Header,
			Link:           txHash,
			Representative: rcTm.Representative,
			Token:          rcTm.Type,
			Extra:          types.ZeroHash,
			Timestamp:      time.Now().Unix(),
		}

	} else {
		blk = types.StateBlock{
			Type:           types.Open,
			Address:        rcAddress,
			Balance:        pending.Amount,
			Previous:       types.ZeroHash,
			Link:           txHash,
			Representative: txBlock.GetRepresentative(), //Representative: genesis.Owner,
			Token:          txBlock.GetToken(),
			Extra:          types.ZeroHash,
			Timestamp:      time.Now().Unix(),
		}
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

// GenerateChangeBlock returns change block by account and new representative address, sign is a function to sign the block
func (l *LedgerApi) GenerateChangeBlock(account types.Address, representative types.Address, sign Signature) (*types.StateBlock, error) {
	if _, err := l.AccountInfo(representative); err != nil {
		return nil, fmt.Errorf("invalid representative[%s]", representative.String())
	}

	rcTm, err := l.token(common.QLCChainToken, account)
	if err != nil {
		return nil, err
	}

	//get latest chain token block
	block, err := l.blockInfo(rcTm.Header)
	if err != nil {
		return nil, err
	}

	blk := types.StateBlock{
		Type:           types.Change,
		Address:        account,
		Balance:        rcTm.Balance,
		Previous:       rcTm.Header,
		Link:           types.ZeroHash,
		Representative: representative,
		Token:          block.Token,
		Extra:          types.ZeroHash,
		Timestamp:      time.Now().Unix(),
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

// Process checks block base info , updates info of chain for the block ,and broadcasts block
func (l *LedgerApi) Process(block *types.StateBlock) (types.Hash, error) {
	var hash types.Hash
	err := l.client.Call(&hash, "ledger_process", block)
	if err != nil {
		return types.ZeroHash, err
	}
	return hash, nil
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

// Representatives returns pairs of representative and its voting weight of chain
// if set sorting false , will return representatives randomly, if set true,
// will sorting representative balance in descending order
func (l *LedgerApi) Representatives(sorting bool) (*api.APIAccountBalances, error) {
	var r api.APIAccountBalances
	err := l.client.Call(&r, "ledger_representatives", sorting)
	if err != nil {
		return nil, err
	}
	return &r, nil
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
func (l *LedgerApi) TokenInfoById(tokenId types.Hash) (*api.ApiTokenInfo, error) {
	var at api.ApiTokenInfo
	err := l.client.Call(&at, "ledger_tokenInfoById", tokenId)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

// TokenInfoById returns token info by token name
func (l *LedgerApi) TokenInfoByName(tokenName string) (*api.ApiTokenInfo, error) {
	var at api.ApiTokenInfo
	err := l.client.Call(&at, "ledger_tokenInfoByName", tokenName)
	if err != nil {
		return nil, err
	}
	return &at, nil
}
