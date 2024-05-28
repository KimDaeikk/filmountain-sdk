package connectors

import (
	"context"
	"net/http"

	"github.com/KimDaeikk/filmountain-sdk/config"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

// FVM상의 스마트 컨트랙트와 상호작용을 위한 연결
func ConnectEthClient(c *config.AppConfig) (*ethclient.Client, error) {
	node := &c.LotusTestNode
	if !c.OnTestnet {
		node = &c.LotusNode
	}
	return connectEthClient(node.Address, node.Token)
}

func connectEthClient(dialAddr string, token string) (*ethclient.Client, error) {
	if token == "" {
		return ethclient.Dial(dialAddr)
	}

	tokenHeader := ethrpc.WithHeader("Authorization", "Bearer "+token)
	httpClient := ethrpc.WithHTTPClient(&http.Client{
		Timeout: 0,
	})

	client, err := ethrpc.DialOptions(context.Background(), dialAddr, httpClient, tokenHeader)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return ethclient.NewClient(client), nil
}
