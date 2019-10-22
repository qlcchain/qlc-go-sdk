package qlcchain

import (
	"encoding/hex"

	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type AccountApi struct {
	client *rpc.Client
}

// NewAccountAPI creates account module for client
func NewAccountAPI(c *rpc.Client) *AccountApi {
	return &AccountApi{client: c}
}

// Create gets account by index from seed
func (a *AccountApi) Create(seedStr string, index uint32) (map[string]string, error) {
	var resp map[string]string
	err := a.client.Call(&resp, "account_create", seedStr, index)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ForPublicKey returns address for public key
func (a *AccountApi) ForPublicKey(pubStr string) (types.Address, error) {
	pub, err := hex.DecodeString(pubStr)
	if err != nil {
		return types.ZeroAddress, err
	}
	addr := types.PubToAddress(pub)
	return addr, nil
}

// NewSeed generates new seed
func (a *AccountApi) NewSeed() (string, error) {
	seed, err := types.NewSeed()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(seed[:]), nil
}

// PublicKey returns public key for address
func (a *AccountApi) PublicKey(addr types.Address) string {
	pub := hex.EncodeToString(addr.Bytes())
	return pub
}

// Validate accepts a address string and checks if it's valid.
func (a *AccountApi) Validate(addr string) bool {
	return types.IsValidHexAddress(addr)
}
