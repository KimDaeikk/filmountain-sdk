package config

var (
	CfgDir  string
	AppConf = AppConfig{}
)

// viper가 mapstrucutre 태그만 인식
type AppConfig struct {
	OnTestnet     bool      `mapstructure:"on_testnet"`
	LotusNode     LotusNode `mapstructure:"lotus_node"`
	LotusTestNode LotusNode `mapstructure:"lotus_test_node"`
}

type LotusNode struct {
	Address string `mapstructure:"address"`
	Token   string `mapstructure:"token"`
}
