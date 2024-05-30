package constants

import (
	"github.com/KimDaeikk/filmountain-sdk/config"
)

var DefaultAppConf = config.AppConfig{
	OnTestnet: true,
	LotusNode: config.LotusNode{
		Address: "https://127.0.0.1",
		Token:   "",
	},
	LotusTestNode: config.LotusNode{
		Address: "https://127.0.0.1",
		Token:   "",
	},
}

var MinerContract = 0xabcd