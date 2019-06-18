package qlcchain

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	ws "github.com/sourcegraph/jsonrpc2/websocket"
)

type QLCClient struct {
	client   *jsonrpc2.Conn
	Account  *AccountApi
	Contract *ContractApi
	Ledger   *LedgerApi
	Mintage  *MintageApi
	Pledger  *PledgeApi
	Rewards  *RewardsApi
	Network  *NetApi
	SMS      *SMSApi
	Util     *UtilApi
}

func (c *QLCClient) Close() error {
	if c != nil && c.client != nil {
		return c.client.Close()
	}
	return nil
}

// NewQLCClient creates a new client
func NewQLCClient(url string, opts ...jsonrpc2.ConnOpt) (*QLCClient, error) {
	conn, err := dial(url, opts...)
	if err != nil {
		return nil, err
	}

	client := &QLCClient{client: conn}

	client.Account = NewAccountApi(client)
	client.Ledger = NewLedgerApi(client)
	client.SMS = NewSMSApi(client)
	client.Contract = NewContractApi(client)
	client.Mintage = NewMintageApi(client)
	client.Pledger = NewPledgeApi(client)
	client.Rewards = NewRewardApi(client)
	client.Network = NewNetApi(client)
	client.Util = NewUtilApi(client)

	return client, nil
}

func dial(rawurl string, opts ...jsonrpc2.ConnOpt) (*jsonrpc2.Conn, error) {
	return dialContext(context.Background(), rawurl, opts...)
}

// DialContext creates a new RPC client, just like Dial.
//
// The context is used to cancel or time out the initial connection establishment. It does
// not affect subsequent interactions with the client.
func dialContext(ctx context.Context, rawurl string, opts ...jsonrpc2.ConnOpt) (*jsonrpc2.Conn, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http", "https":
		conn, err := net.Dial("tcp", u.Host)
		if err != nil {
			return nil, err
		}
		return jsonrpc2.NewConn(ctx, jsonrpc2.NewBufferedStream(conn, jsonrpc2.VarintObjectCodec{}), nil, opts...), nil
	case "ws", "wss":
		c, _, err := websocket.DefaultDialer.Dial(rawurl, nil)
		if err != nil {
			return nil, err
		}
		return jsonrpc2.NewConn(ctx, ws.NewObjectStream(c), nil, opts...), nil
	default:
		return nil, fmt.Errorf("no known transport for URL scheme %q", u.Scheme)
	}
}

func (c *QLCClient) Call(result interface{}, method string, params ...interface{}) error {
	ctx := context.Background()
	return c.client.Call(ctx, method, params, result)
}

// Version returns version for sdk
func (c *QLCClient) Version() string {
	return fmt.Sprintf("%s.%s.%s", VERSION, GITREV, BUILDTIME)
}
