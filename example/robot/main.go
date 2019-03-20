package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/common/util"
	"github.com/qlcchain/go-qlc/crypto/random"
	"github.com/qlcchain/go-qlc/log"
	"github.com/qlcchain/go-qlc/rpc/api"
	"github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/example/robot/message"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	phonePrefix = []string{"130", "131", "132", "133", "134", "135", "136", "137", "138",
		"139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186",
		"187", "188"}
	txInterval     = flag.Int("txInterval", 10, "send message interval")
	rxInterval     = flag.Int("rxInterval", 120, "receive message interval")
	endPoint       = ""
	accounts       arrayFlags
	minInterval    = 10
	txAccounts     []*types.Account
	txAccountSize  int
	maxAmount      = 10
	currentAccount *types.Account
	mutex          sync.RWMutex
	token          = "QLC"
	logger         = log.NewLogger("qlc_robot")
)

func main() {
	flag.StringVar(&endPoint, "endpoint", "ws://127.0.0.1:19736", "RPC Server endpoint")
	flag.Var(&accounts, "account", "account private key")
	flag.Parse()

	if *txInterval < minInterval {
		logger.Errorf("invalid txInterval %d[%d,∞]\n", *txInterval, minInterval)
		return
	}

	if *rxInterval < minInterval {
		logger.Errorf("invalid rxInterval %d[%d,∞]\n", *rxInterval, minInterval)
		return
	}

	if len(accounts) == 0 {
		logger.Error("can not find any account")
		return
	}

	for i, a := range accounts {
		bytes, e := hex.DecodeString(a)
		if e != nil {
			logger.Errorf("can not decode (%s) at %d to Account", a, i)
			continue
		}
		account := types.NewAccount(bytes)
		logger.Infof("Tx[%d]: %s\n", i, account.Address().String())
		txAccounts = append(txAccounts, account)
	}
	txAccountSize = len(txAccounts)

	if txAccountSize < 2 {
		logger.Errorf("not enough account(%d) to send Tx\n", txAccountSize)
		return
	}

	client, err := qlcchain.NewQLCClient(endPoint)
	if err != nil {
		logger.Error(err)
		return
	}

	// make sure all accounts already open
	err = generateReceives(client)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Infof("%d Account will send Tx every %d second(s)", txAccountSize, *txInterval)

	txDuration, _ := time.ParseDuration(fmt.Sprintf("%ds", *txInterval))
	txTicker := time.NewTicker(txDuration)
	defer txTicker.Stop()

	rxDuration, _ := time.ParseDuration(fmt.Sprintf("%ds", *rxInterval))
	rxTicker := time.NewTicker(rxDuration)
	defer rxTicker.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		select {
		case <-c:
			logger.Info("receive close signal, stop ...")
			return
		case txTime := <-txTicker.C:
			logger.Info("produce send @ ", txTime)
			go func() {
				txAccount := randomAccount(nil, txAccounts)
				setAccount(txAccount)
				defer setAccount(nil)
				rxAccount := randomAccount(txAccount, txAccounts)
				amount := randomAmount()
				_, p := message.RandomPoem()
				m := p.Message()
				mh, err := client.SMS.MessageStore(m)
				if err != nil {
					logger.Error(err)
					return
				}
				param := &api.APISendBlockPara{
					From:      txAccount.Address(),
					TokenName: token,
					To:        rxAccount.Address(),
					Amount:    types.Balance{Int: big.NewInt(int64(amount))},
					Sender:    randomPhone(),
					Receiver:  randomPhone(),
					Message:   mh,
				}
				logger.Debug(util.ToString(param))
				txBlock, err := client.Ledger.GenerateSendBlock(param, func(hash types.Hash) (types.Signature, error) {
					return txAccount.Sign(hash), nil
				})

				if err != nil {
					logger.Error(err)
				}
				logger.Info(util.ToString(txBlock))
				hash, err := client.Ledger.Process(txBlock)
				if err != nil {
					logger.Error(err)
				}
				logger.Info(hash.String())
			}()
		case rxTime := <-rxTicker.C:
			logger.Info("generate receive @ ", rxTime)
			go func() {
				err := generateReceives(client)
				if err != nil {
					logger.Error(err)
				}
			}()
		}
	}
}

func generateReceives(client *qlcchain.QLCClient) error {
	//generate receive block
	cache := make(map[types.Address]*types.Account, 0)
	var addresses []types.Address
	account := getAccount()
	for _, a := range txAccounts {
		addr := a.Address()
		if account != nil && addr == account.Address() {
			continue
		}
		addresses = append(addresses, addr)
		cache[addr] = a
	}

	pendings, err := client.Ledger.AccountsPending(addresses, -1)
	if err != nil {
		return err
	}

	for addr, v := range pendings {
		for _, pending := range v {
			rxBlock, err := client.Ledger.GenerateReceiveBlockByHash(pending.Hash, func(hash types.Hash) (types.Signature, error) {
				if a, ok := cache[addr]; ok {
					return a.Sign(hash), nil
				} else {
					return types.Signature{}, fmt.Errorf("can not find addr[%s]private key", addr.String())
				}
			})
			if err != nil {
				logger.Error(err)
				continue
			}
			if h, err := client.Ledger.Process(rxBlock); err == nil {
				logger.Infof("generate receive %s from %s", pending.Hash.String(), h.String())
			}
		}
	}

	return nil
}

func randomAmount() int {
	i, _ := random.Intn(maxAmount)
	if i == 0 {
		return randomAmount()
	}
	return i
}

func randomAccount(a *types.Account, account []*types.Account) *types.Account {
	i, _ := random.Intn(len(account))
	tmp := account[i]
	if a != nil && tmp.Address() == a.Address() {
		return randomAccount(a, account)
	}
	return tmp
}

func setAccount(account *types.Account) {
	mutex.Lock()
	defer mutex.Unlock()

	currentAccount = account
}

func getAccount() *types.Account {
	mutex.RLock()
	defer mutex.RUnlock()

	return currentAccount
}

func hash(msg string) types.Hash {
	m := fmt.Sprintf("%s powered by qlcchain", msg)
	h, _ := types.HashBytes([]byte(m))
	return h
}

func randomPhone() string {
	i, _ := random.Intn(len(phonePrefix))
	var sb strings.Builder
	sb.WriteString("+86 ")
	sb.WriteString(phonePrefix[i])
	for i := 0; i < 8; i++ {
		tmp, _ := random.Intn(10)
		sb.WriteString(strconv.Itoa(tmp))
	}
	return sb.String()
}
