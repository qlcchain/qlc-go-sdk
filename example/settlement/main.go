package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"strings"
	"time"

	qlcchain "github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/pkg/random"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	"github.com/qlcchain/qlc-go-sdk/pkg/util"
)

var (
	flagNodeUrl         string
	gasAccount          *types.Account
	pccwAccount         *types.Account
	cslAccount          *types.Account
	createContractParam = &qlcchain.CreateContractParam{
		PartyA: qlcchain.Contractor{
			Name: "PCCWG",
		},
		PartyB: qlcchain.Contractor{
			Name: "HKTCSL",
		},
		Services: []qlcchain.ContractService{{
			ServiceId:   hash().String(),
			Mcc:         1,
			Mnc:         2,
			TotalAmount: 10,
			UnitPrice:   0.0426,
			Currency:    "USD",
		}, {
			ServiceId:   hash().String(),
			Mcc:         22,
			Mnc:         1,
			TotalAmount: 30,
			UnitPrice:   0.023,
			Currency:    "USD",
		}},
		StartDate: time.Now().Add(time.Second * 10).Unix(),
		EndDate:   time.Now().AddDate(1, 0, 0).Unix(),
	}
)

func init() {
	_, priv, err := types.KeypairFromSeed("", 0)
	if err != nil {
		fmt.Println(err)
	}
	gasAccount = types.NewAccount(priv)
	b1, _ := hex.DecodeString("b912bf9a3d2efdb1cb5a20052a8b1e47de3f484992677bcb126660c3c83ee206ba1f06f8eccee2277f9b54f34ef485b08dadde36ee2814cf8e361bf7ec2e9bfb")
	pccwAccount = types.NewAccount(b1)
	b2, _ := hex.DecodeString("e2994c94c8e2c60f96c9c7b7c0035a9e43894526a461630de78a9f253f65ce3ab3a99e4455903ee39d47a4f0954fae59f9a67a3206ed2351b6d821010b51aeed")
	cslAccount = types.NewAccount(b2)
}

func main() {
	flag.StringVar(&flagNodeUrl, "nodeurl", "http://127.0.0.1:19735", "RPC URL of node")
	flag.Parse()

	client, err := qlcchain.NewQLCClient(flagNodeUrl)
	if err != nil || client == nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client.Version())

	if _, err := initStatus(client); err != nil {
		fmt.Println("check PoV status failed")
		return
	}

	fmt.Println("node is ready, Go Go Go ....")

	// prepare accounts
	//pccwAccount := account()
	fmt.Println(pccwAccount.String())
	//cslAccount := account()
	fmt.Println(cslAccount.String())

	token, _ := client.Ledger.GasToken()
	fmt.Println(token.String())

	if err = send(client, gasAccount, pccwAccount, token); err != nil {
		fmt.Println(err)
		return
	}
	if err := send(client, gasAccount, cslAccount, token); err != nil {
		fmt.Println(err)
		return
	}

	// create settlement contract
	pccwAddr := pccwAccount.Address()
	cslAddr := cslAccount.Address()

	offset := 0

	size := printAllContract(client)
	// create settlement contract if zero
	if size == 0 {
		param := createContractParam
		param.PartyA.Address = pccwAddr
		param.PartyB.Address = cslAddr
		mnc, _ := random.Intn(100)
		param.Services[0].Mnc = uint64(mnc)

		if txBlk, err := client.Settlement.GetCreateContractBlock(createContractParam, func(hash types.Hash) (signature types.Signature, err error) {
			return pccwAccount.Sign(hash), nil
		}); err != nil {
			fmt.Println(err)
			return
		} else {
			//fmt.Println(util.ToIndentString(txBlk))
			if err := processBlockAndWaitConfirmed(client, txBlk); err != nil {
				fmt.Println(err)
				return
			}
			//printAllContract(client)
			txHash := txBlk.GetHash()
			if rxBlk, err := client.Settlement.GetSettlementRewardsBlock(&txHash, func(hash types.Hash) (signature types.Signature, err error) {
				return pccwAccount.Sign(hash), nil
			}); err != nil {
				fmt.Println(err)
				return
			} else {
				if err := processBlockAndWaitConfirmed(client, rxBlk); err != nil {
					fmt.Println(err)
					return
				}
			}

			if contracts, err := client.Settlement.GetContractsAsPartyB(&cslAddr, 10, &offset); err != nil {
				fmt.Println(err)
				return
			} else {
				for _, c := range contracts {
					if c.Status != qlcchain.ContractStatusActivated {
						if txBlk, err := client.Settlement.GetSignContractBlock(&qlcchain.SignContractParam{
							ContractAddress: c.Address,
							Address:         cslAddr,
						}, func(hash types.Hash) (signature types.Signature, err error) {
							return cslAccount.Sign(hash), nil
						}); err != nil {
							fmt.Println(err)
							return
						} else {
							if err := processBlockAndWaitConfirmed(client, txBlk); err != nil {
								fmt.Println(err)
								return
							}
						}
					}

					fmt.Println(util.ToIndentString(c))
					if len(c.NextStops) == 0 {
						// update next stop
						stopParam := &qlcchain.StopParam{
							StopName:        "HKTCSL",
							Address:         pccwAddr,
							ContractAddress: c.Address,
						}
						fmt.Println(util.ToIndentString(stopParam))
						if block, err := client.Settlement.GetAddNextStopBlock(stopParam, func(hash types.Hash) (signature types.Signature, err error) {
							return pccwAccount.Sign(hash), nil
						}); err != nil {
							fmt.Println(err)
							return
						} else {
							if err := processBlockAndWaitConfirmed(client, block); err != nil {
								fmt.Println(err)
								return
							}
						}

						// update previous stop
						if block, err := client.Settlement.GetAddPreStopBlock(&qlcchain.StopParam{
							StopName:        "PCCWG",
							Address:         cslAddr,
							ContractAddress: c.Address,
						}, func(hash types.Hash) (signature types.Signature, err error) {
							return cslAccount.Sign(hash), nil
						}); err != nil {
							fmt.Println(err)
							return
						} else {
							if err := processBlockAndWaitConfirmed(client, block); err != nil {
								fmt.Println(err)
								return
							}
						}

					}

				}
			}

			// check settlement contract
			if contracts, err := client.Settlement.GetContractsByAddress(&pccwAddr, 10, &offset); err != nil {
				fmt.Println(err)
				return
			} else {
				if len(contracts) == 0 {
					fmt.Println("can not find any contracts")
					return
				}
				fmt.Println(util.ToIndentString(contracts))
			}
		}
	}

	cdr1 := &qlcchain.CDRParam{
		Index:         1,
		SmsDt:         time.Now().Unix(),
		Sender:        "WeChat",
		Destination:   "85257***343",
		SendingStatus: qlcchain.SendingStatusSent,
		DlrStatus:     qlcchain.DLRStatusDelivered,
		PreStop:       "",
		NextStop:      "HKTCSL",
	}

	if blk, err := client.Settlement.GetProcessCDRBlock(&pccwAddr, []*qlcchain.CDRParam{cdr1}, func(hash types.Hash) (signature types.Signature, err error) {
		return pccwAccount.Sign(hash), nil
	}); err != nil {
		fmt.Println(err)
		return
	} else {
		if err := processBlockAndWaitConfirmed(client, blk); err != nil {
			fmt.Println(err)
			return
		}

		txHash := blk.GetHash()
		if rxBlk, err := client.Settlement.GetSettlementRewardsBlock(&txHash, func(hash types.Hash) (signature types.Signature, err error) {
			return pccwAccount.Sign(hash), nil
		}); err != nil {
			fmt.Println(err)
			return
		} else {
			if err := processBlockAndWaitConfirmed(client, rxBlk); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	cdr2 := &qlcchain.CDRParam{
		Index:         1,
		SmsDt:         time.Now().Unix(),
		Sender:        "WeChat",
		Destination:   "85257***343",
		SendingStatus: qlcchain.SendingStatusSent,
		DlrStatus:     qlcchain.DLRStatusDelivered,
		PreStop:       "PCCWG",
		NextStop:      "",
	}

	if blk, err := client.Settlement.GetProcessCDRBlock(&cslAddr, []*qlcchain.CDRParam{cdr2}, func(hash types.Hash) (signature types.Signature, err error) {
		return cslAccount.Sign(hash), nil
	}); err != nil {
		fmt.Println(err)
		return
	} else {
		if err := processBlockAndWaitConfirmed(client, blk); err != nil {
			fmt.Println(err)
			return
		}

		txHash := blk.GetHash()
		if rxBlk, err := client.Settlement.GetSettlementRewardsBlock(&txHash, func(hash types.Hash) (signature types.Signature, err error) {
			return cslAccount.Sign(hash), nil
		}); err != nil {
			fmt.Println(err)
			return
		} else {
			if err := processBlockAndWaitConfirmed(client, rxBlk); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func printAllContract(c *qlcchain.QLCClient) int {
	offset := 0
	if contracts, err := c.Settlement.GetAllContracts(100, &offset); err != nil {
		fmt.Println(err)
		return 0
	} else {
		for _, c := range contracts {
			fmt.Println(util.ToIndentString(c))
		}
		fmt.Println(strings.Repeat("*", 64))
		return len(contracts)
	}
}

func send(client *qlcchain.QLCClient, from, to *types.Account, token *types.Hash) error {
	_, err := client.Ledger.TokenMeta(*token, to.Address())
	if err != nil {
		if txBlk, err := client.Ledger.GenerateSendBlock(&qlcchain.APISendBlockPara{
			From:      from.Address(),
			TokenName: "QGAS",
			To:        to.Address(),
			Amount:    types.Balance{Int: big.NewInt(1e14)},
			Sender:    "",
			Receiver:  "",
			Message:   types.Hash{},
		}, func(hash types.Hash) (signature types.Signature, err error) {
			return from.Sign(hash), nil
		}); err != nil {
			return err
		} else {
			if err := processBlockAndWaitConfirmed(client, txBlk); err != nil {
				return err
			}
			if rxBlk, err := client.Ledger.GenerateReceiveBlock(txBlk, func(hash types.Hash) (signature types.Signature, err error) {
				return to.Sign(hash), nil
			}); err != nil {
				return err
			} else {
				if err := processBlockAndWaitConfirmed(client, rxBlk); err != nil {
					return err
				}
				if am, err := client.Ledger.AccountInfo(to.Address()); err != nil {
					return err
				} else {
					if am == nil {
						return errors.New("invalid account info")
					} else {
						fmt.Println(util.ToIndentString(am))
					}
				}
			}
		}
	}

	return nil
}

func processBlockAndWaitConfirmed(c *qlcchain.QLCClient, block *types.StateBlock) error {
	_, err := c.Ledger.Process(block)
	if err != nil {
		return fmt.Errorf("process block error: %s", err)
	}
	return waitBlockConfirmed(c, block.GetHash())
}

func waitBlockConfirmed(c *qlcchain.QLCClient, hash types.Hash) error {
	t := time.NewTimer(time.Second * 180)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			return errors.New("consensus confirmed timeout")
		default:
			confirmed, err := c.Ledger.BlockConfirmedStatus(hash)
			if err != nil {
				return err
			}
			if confirmed {
				return nil
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func initStatus(c *qlcchain.QLCClient) (bool, error) {
	ticker := time.NewTicker(500 * time.Second)
	for {
		s, err := c.Pov.GetPovStatus()
		if err != nil {
			return false, err
		} else if s.SyncState == 2 {
			return true, nil
		}
		if !s.PovEnabled {
			return false, errors.New("pov is not enable")
		}

		select {
		case <-ticker.C:
			return false, errors.New("timeout")
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func account() *types.Account {
	seed, _ := types.NewSeed()
	a, _ := seed.Account(0)
	return a
}

func hash() types.Hash {
	h := types.Hash{}
	_ = random.Bytes(h[:])
	return h
}
