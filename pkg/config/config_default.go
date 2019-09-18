package config

import (
	"fmt"
	"path"

	"github.com/qlcchain/qlc-go-sdk/pkg/util"
)

type Config ConfigV5

func DefaultConfig(dir string) (*Config, error) {
	v5, err := DefaultConfigV5(dir)
	if err != nil {
		return &Config{}, err
	}
	cfg := Config(*v5)

	return &cfg, nil
}

func DefaultConfigV5(dir string) (*ConfigV5, error) {
	var cfg ConfigV5
	cfg4, _ := DefaultConfigV4(dir)
	cfg.ConfigV4 = *cfg4
	cfg.Version = configVersion
	cfg.Metrics = defaultMetrics()
	return &cfg, nil
}

func defaultMetrics() *MetricsConfig {
	return &MetricsConfig{
		Enable:         false,
		SampleInterval: 1,
		Influx: &Influx{
			Enable:   false,
			URL:      "http://localhost:10086",
			Database: "qlcchain",
			User:     "qlcchain",
			Password: "",
			Interval: 10,
		},
	}
}

func DefaultConfigV4(dir string) (*ConfigV4, error) {
	var cfg ConfigV4
	cfg3, _ := DefaultConfigV3(dir)
	cfg.ConfigV3 = *cfg3
	cfg.Version = configVersion

	cfg.RPC.PublicModules = append(cfg.RPC.PublicModules, "pov", "miner")

	cfg.PoV = defaultPoV()

	return &cfg, nil
}

func defaultPoV() *PoVConfig {
	return &PoVConfig{
		PovEnabled:   false,
		MinerEnabled: false,
		Coinbase:     "",
	}
}

func DefaultConfigV3(dir string) (*ConfigV3, error) {
	var cfg ConfigV3
	cfg2, _ := DefaultConfigV2(dir)
	cfg.ConfigV2 = *cfg2
	cfg.Version = 3
	cfg.RPC.HttpVirtualHosts = []string{"*"}
	cfg.RPC.PublicModules = append(cfg.RPC.PublicModules, "pledge")

	cfg.DB = defaultDb(dir)

	return &cfg, nil
}

var (
	relationDir = "relation"
	pwLen       = 16
)

func defaultDb(dir string) *DBConfig {
	d := path.Join(dir, "ledger", relationDir, "index.db")
	pw := util.RandomFixedString(pwLen)

	//"postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	//postgres
	return &DBConfig{
		ConnectionString: fmt.Sprintf("file:%s?_auth&_auth_user=qlcchain&_auth_pass=%s", d, pw),
		Driver:           "sqlite3",
	}
}

func DefaultConfigV2(dir string) (*ConfigV2, error) {
	pk, id, err := identityConfig()
	if err != nil {
		return nil, err
	}
	var cfg ConfigV2
	modules := []string{"qlcclassic", "ledger", "account", "net", "util", "wallet", "mintage", "contract", "sms"}
	cfg = ConfigV2{
		Version:             2,
		DataDir:             dir,
		StorageMax:          "10GB",
		AutoGenerateReceive: false,
		LogLevel:            "error",
		PerformanceEnabled:  false,
		RPC: &RPCConfigV2{
			Enable:           true,
			HTTPEnabled:      true,
			HTTPEndpoint:     "tcp4://0.0.0.0:10735",
			HTTPCors:         []string{"*"},
			HttpVirtualHosts: []string{},
			WSEnabled:        true,
			WSEndpoint:       "tcp4://0.0.0.0:10736",
			IPCEnabled:       true,
			IPCEndpoint:      defaultIPCEndpoint(),
			PublicModules:    modules,
		},
		P2P: &P2PConfigV2{
			BootNodes:    []string{},
			Listen:       "/ip4/0.0.0.0/tcp/10734",
			SyncInterval: 30,
			Discovery: &DiscoveryConfigV2{
				DiscoveryInterval: 10,
				Limit:             20,
				MDNSEnabled:       false,
				MDNSInterval:      30,
			},
			ID: &IdentityConfigV2{id, pk},
		},
	}

	return &cfg, nil
}
