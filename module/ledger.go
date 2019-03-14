package module

import (
	"github.com/qlcchain/go-qlc/common/types"
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
	return nil, nil
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
	return nil, nil
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
	return nil, nil
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

func (l *LedgerApi) GenerateSendBlock(para *api.APISendBlockPara, prkStr string) (*types.StateBlock, error) {
	return nil, nil
}

func (l *LedgerApi) GenerateReceiveBlock(sendBlock *types.StateBlock, prkStr string) (*types.StateBlock, error) {
	return nil, nil
}

func (l *LedgerApi) GenerateChangeBlock(account types.Address, representative types.Address, prkStr string) (*types.StateBlock, error) {
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
