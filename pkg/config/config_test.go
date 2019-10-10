package config

import (
	"testing"

	"github.com/qlcchain/qlc-go-sdk/pkg/util"
)

func TestConfig_LogDir(t *testing.T) {
	cfg, err := DefaultConfig(DefaultDataDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(util.ToIndentString(cfg))

	t.Log(cfg.DataDir)
	//t.Log(cfg.LogDir())
	//t.Log(cfg.LedgerDir())
	//t.Log(cfg.WalletDir())
	t.Log(QlcTestDataDir())
}
