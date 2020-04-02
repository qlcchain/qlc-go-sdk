package qlcchain

import (
	rpc "github.com/qlcchain/jsonrpc2"
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
func (p *PovApi) DistributeRawPayload(param *PrivacyDistributeParam) ([]byte, error) {
	var rspData []byte
	err := p.client.Call(&rspData, "privacy_distributeRawPayload")
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

// GetRawPayload return private raw data by enclave key
func (p *PovApi) GetRawPayload(enclaveKey []byte) ([]byte, error) {
	var rspData []byte
	err := p.client.Call(&rspData, "privacy_getRawPayload")
	if err != nil {
		return nil, err
	}
	return rspData, nil
}
