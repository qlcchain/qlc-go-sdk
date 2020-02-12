package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type GenesisInfo struct {
	chain        types.StateBlock
	chainMintage types.StateBlock
	gas          types.StateBlock
	gasMintage   types.StateBlock
}

func GenesisAddress(c *rpc.Client) types.Address {
	gs := getGenesisBlocks(c)
	return gs.chain.Address
}

func ChainToken(c *rpc.Client) types.Hash {
	gs := getGenesisBlocks(c)
	return gs.chain.Token
}

func GenesisBlock(c *rpc.Client) types.StateBlock {
	gs := getGenesisBlocks(c)
	return gs.chain
}

func GenesisMintageBlock(c *rpc.Client) types.StateBlock {
	gs := getGenesisBlocks(c)
	return gs.chainMintage
}

func GenesisBlockHash(c *rpc.Client) types.Hash {
	gs := getGenesisBlocks(c)
	return gs.chain.GetHash()
}

func GenesisMintageHash(c *rpc.Client) types.Hash {
	gs := getGenesisBlocks(c)
	return gs.chainMintage.GetHash()
}

func GasAddress(c *rpc.Client) types.Address {
	gs := getGenesisBlocks(c)
	return gs.gas.Address
}

func GasToken(c *rpc.Client) types.Hash {
	gs := getGenesisBlocks(c)
	return gs.gas.Token
}

func GasBlock(c *rpc.Client) types.StateBlock {
	gs := getGenesisBlocks(c)
	return gs.gas
}

func GasBlockHash(c *rpc.Client) types.Hash {
	gs := getGenesisBlocks(c)
	return gs.gas.GetHash()
}

func GasMintageBlock(c *rpc.Client) types.StateBlock {
	gs := getGenesisBlocks(c)
	return gs.gasMintage
}
func GasMintageHash(c *rpc.Client) types.Hash {
	gs := getGenesisBlocks(c)
	return gs.gasMintage.GetHash()
}

// IsGenesis check block is chain token genesis
func IsGenesisBlock(block *types.StateBlock, c *rpc.Client) bool {
	gs := getGenesisBlocks(c)
	h := block.GetHash()
	return h == gs.chainMintage.GetHash() || h == gs.chain.GetHash() || h == gs.gasMintage.GetHash() || h == gs.gas.GetHash()
}

// IsGenesis check token is chain token genesis
func IsGenesisToken(hash types.Hash, c *rpc.Client) bool {
	gs := getGenesisBlocks(c)
	return hash == gs.chain.Token || hash == gs.gas.Token
}

func AllGenesisBlocks(c *rpc.Client) []types.StateBlock {
	gs := getGenesisBlocks(c)
	return []types.StateBlock{gs.gas, gs.gasMintage, gs.chain, gs.chainMintage}
}

func getGenesisBlocks(c *rpc.Client) *GenesisInfo {
	g := new(GenesisInfo)
	var r map[string]types.StateBlock
	err := c.Call(&r, "ledger_genesisBlocks")
	if err != nil {
		return nil
	}
	g.chain = r["chain"]
	g.chainMintage = r["chain-mintage"]
	g.gas = r["gas"]
	g.gasMintage = r["gas-mintage"]
	return g
}
