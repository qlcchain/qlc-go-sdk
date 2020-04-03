package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
)

type PrivacyDistributeParam struct {
	RawPayload     []byte   `json:"rawPayload"`
	PrivateFrom    string   `json:"privateFrom"`
	PrivateFor     []string `json:"privateFor"`
	PrivateGroupID string   `json:"privateGroupID"`
}

type PrivacyApi struct {
	client *rpc.Client
}

// NewPrivacyAPI creates privacy module for client
func NewPrivacyAPI(c *rpc.Client) *PrivacyApi {
	return &PrivacyApi{client: c}
}

// DistributeRawPayload push private raw data to parties
func (p *PrivacyApi) DistributeRawPayload(param *PrivacyDistributeParam) ([]byte, error) {
	var rspData []byte
	err := p.client.Call(&rspData, "privacy_distributeRawPayload")
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

// GetRawPayload return private raw data by enclave key
func (p *PrivacyApi) GetRawPayload(enclaveKey []byte) ([]byte, error) {
	var rspData []byte
	err := p.client.Call(&rspData, "privacy_getRawPayload")
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

// GetBlockPrivatePayload return private raw data by block hash
func (p *PrivacyApi) GetBlockPrivatePayload(blockHash types.Hash) ([]byte, error) {
	var rspData []byte
	err := p.client.Call(&rspData, "privacy_getBlockPrivatePayload")
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

// GetDemoKV returns KV in PrivacyKV contract (just for demo in testnet)
func (p *PrivacyApi) GetDemoKV(key []byte) ([]byte, error) {
	var rspData []byte
	err := p.client.Call(&rspData, "privacy_getDemoKV", key)
	if err != nil {
		return nil, err
	}
	return rspData, nil
}
