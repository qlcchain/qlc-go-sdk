package module

import (
	"encoding/hex"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/rpc"
)

type AccountApi struct {
	client *rpc.Client
}

func NewAccountApi(c *rpc.Client) *AccountApi {
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

// PublicKey returns public key for address
func (a *AccountApi) PublicKey(addr types.Address) string {
	pub := hex.EncodeToString(addr.Bytes())
	return pub
}

// Validate accepts a address string and checks if it's valid.
func (a *AccountApi) Validate(addr string) bool {
	return types.IsValidHexAddress(addr)
}
