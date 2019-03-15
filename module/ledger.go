package module

import (
	"errors"
	"fmt"
	"time"

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

func (l *LedgerApi) AccountBlocksCount(addr types.Address) (int64, error) {
	return 0, nil
}

func (l *LedgerApi) AccountHistoryTopn(address types.Address, count int, offset *int) ([]*api.APIBlock, error) {
	return nil, nil
}

func (l *LedgerApi) AccountInfo(address types.Address) (*api.APIAccount, error) {
	var aa *api.APIAccount
	err := l.client.Call(&aa, "ledger_accountInfo", address)
	if err != nil {
		return nil, err
	}
	return aa, nil
}

func (l *LedgerApi) AccountRepresentative(addr types.Address) (types.Address, error) {
	return types.ZeroAddress, nil
}

func (l *LedgerApi) AccountVotingWeight(addr types.Address) (types.Balance, error) {
	return types.ZeroBalance, nil
}

func (l *LedgerApi) AccountsBalance(addresses []types.Address) (map[types.Address]map[string]map[string]types.Balance, error) {
	return nil, nil
}

func (l *LedgerApi) AccountsFrontiers(addresses []types.Address) (map[types.Address]map[string]types.Hash, error) {
	return nil, nil
}

func (l *LedgerApi) AccountsPending(addresses []types.Address, n int) (map[types.Address][]*api.APIPending, error) {
	var ap map[types.Address][]*api.APIPending
	err := l.client.Call(&ap, "ledger_accountsPending", addresses, n)
	if err != nil {
		return nil, err
	}
	return ap, nil
}

func (l *LedgerApi) AccountsCount() (uint64, error) {
	return 0, nil
}

func (l *LedgerApi) Accounts(count int, offset *int) ([]*types.Address, error) {
	return nil, nil
}

func (l *LedgerApi) BlockAccount(hash types.Hash) (types.Address, error) {
	return types.ZeroAddress, nil
}

func (l *LedgerApi) BlockHash(block types.StateBlock) types.Hash {
	return types.ZeroHash
}

func (l *LedgerApi) BlocksCount() (map[string]uint64, error) {
	return nil, nil
}

func (l *LedgerApi) BlocksCountByType() (map[string]uint64, error) {
	return nil, nil
}

func (l *LedgerApi) BlocksInfo(hash []types.Hash) ([]*api.APIBlock, error) {
	var ab []*api.APIBlock
	err := l.client.Call(&ab, "ledger_blocksInfo", hash)
	if err != nil {
		return nil, err
	}
	return ab, nil
}

func (l *LedgerApi) Blocks(count int, offset *int) ([]*api.APIBlock, error) {
	return nil, nil
}

func (l *LedgerApi) Chain(hash types.Hash, n int) ([]types.Hash, error) {
	return nil, nil
}

func (l *LedgerApi) Delegators(hash types.Address) ([]*api.APIAccountBalance, error) {
	return nil, nil
}

func (l *LedgerApi) DelegatorsCount(hash types.Address) (int64, error) {
	return 0, nil
}

type Signature func(Hash types.Hash) (types.Signature, error)

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
	return nil, errors.New("token not found")
}

func (l *LedgerApi) pending(address types.Address, hash types.Hash) (*api.APIPending, error) {
	pendings, err := l.AccountsPending([]types.Address{address}, 100)
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

func (l *LedgerApi) GenerateReceiveBlock(sendBlock *types.StateBlock, sign Signature) (*types.StateBlock, error) {
	sendHash := sendBlock.GetHash()
	if !sendBlock.GetType().Equal(types.Send) {
		return nil, fmt.Errorf("(%s) is not send block", sendHash.String())
	}

	b, err := l.BlocksInfo([]types.Hash{sendBlock.GetHash()})
	if len(b) < 1 || err != nil {
		return nil, fmt.Errorf("send block(%s) does not exist", sendHash.String())
	}

	rcAddress := types.Address(sendBlock.Link)

	pending, err := l.pending(rcAddress, sendHash)
	if err != nil {
		return nil, err
	}
	rcTm, _ := l.token(sendBlock.GetToken(), rcAddress)

	var blk types.StateBlock
	if rcTm != nil {
		blk = types.StateBlock{
			Type:           types.Receive,
			Address:        rcAddress,
			Balance:        rcTm.Balance.Add(pending.Amount),
			Previous:       rcTm.Header,
			Link:           sendHash,
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
			Link:           sendHash,
			Representative: sendBlock.GetRepresentative(), //Representative: genesis.Owner,
			Token:          sendBlock.GetToken(),
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

func (l *LedgerApi) GenerateChangeBlock(account types.Address, representative types.Address, sign Signature) (*types.StateBlock, error) {
	return nil, nil
}

func (l *LedgerApi) Process(block *types.StateBlock) (types.Hash, error) {
	return types.ZeroHash, nil
}

func (l *LedgerApi) Performance() ([]*types.PerformanceTime, error) {
	return nil, nil
}

func (l *LedgerApi) Representatives(sorting *bool) (*api.APIAccountBalances, error) {
	return nil, nil
}

func (l *LedgerApi) Tokens() ([]*types.TokenInfo, error) {
	return nil, nil
}

func (l *LedgerApi) TransactionsCount() (map[string]uint64, error) {
	return nil, nil
}

func (l *LedgerApi) TokenInfoById(tokenId types.Hash) (*api.ApiTokenInfo, error) {
	return nil, nil
}

func (l *LedgerApi) TokenInfoByName(tokenName string) (*api.ApiTokenInfo, error) {
	var at *api.ApiTokenInfo
	err := l.client.Call(&at, "ledger_tokenInfoByName", tokenName)
	if err != nil {
		return nil, err
	}
	return at, nil
}
