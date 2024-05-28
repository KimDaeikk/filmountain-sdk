package connectors

import (
	"context"
	"net/http"

	"github.com/KimDaeikk/filmountain-sdk/config"
	"github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
)

// Lotus 노드와 상호작용을 위한 연결
func ConnectLotusClient(c *config.AppConfig) (*lotusapi.FullNodeStruct, jsonrpc.ClientCloser, error) {
	node := &c.LotusTestNode
	if !c.OnTestnet {
		node = &c.LotusNode
	}
	return connectLotusClient(node.Address, node.Token)
}

func connectLotusClient(lotusDialAddr string, lotusToken string) (*lotusapi.FullNodeStruct, jsonrpc.ClientCloser, error) {
	head := http.Header{}

	if lotusToken != "" {
		head.Add("Authorization", "Bearer "+lotusToken)
	}

	lapi := &lotusapi.FullNodeStruct{}

	closer, err := jsonrpc.NewMergeClient(
		context.Background(),
		lotusDialAddr,
		"Filecoin",
		lotusapi.GetInternalStructs(lapi),
		head,
	)

	if err != nil {
		return nil, nil, err
	}

	return lapi, closer, nil
}
