package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
)

type Node struct {
	TwinId    uint   `json:"twinId"`
	NodeId    uint   `json:"nodeId"`
	UpdatedAt uint32 `json:"updatedAt"`
}
type Twin struct {
	Relay string `json:"relay"`
}

type Resp struct {
	Zos string `json:"zos"`
}

const (
	targetVersion = "3.12.10"
	mne           = ""
)

// func fetchNodes() ([]Node, error) {
// 	res, err := http.Get("https://gridproxy.grid.tf/nodes?size=999999")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	by, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var nodes []Node
// 	if err := json.Unmarshal(by, &nodes); err != nil {
// 		return nil, err
// 	}
// 	return nodes, nil
// }

func main() {
	nodes, err := fetchNodes("up")
	if err != nil {
		log.Fatal(err)
	}

	chainMan := substrate.NewManager("wss://tfchain.grid.tf/ws")
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	rmbClient, err := peer.NewRpcClient(context.Background(), mne, chainMan,
		peer.WithRelay("wss://relay.ninja.tf"),
		peer.WithSession("rmb-call-bin"))
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	var updated, all int
	var mu sync.Mutex

	for _, node := range nodes {
		wg.Add(1)
		go func(node Node) {
			defer wg.Done()

			var res struct {
				Zos string
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			err := rmbClient.Call(ctx, uint32(node.TwinId), "zos.system.version", nil, &res)
			if err != nil {
				return
			}
			all++
			if res.Zos == targetVersion {
				// fmt.Println("nodeId", node.NodeId, "version", res.Zos)
				mu.Lock()
				updated++
				mu.Unlock()
			}
		}(node)
	}

	wg.Wait()
	fmt.Println("up nodes: ", len(nodes), "responds to rmb: ", all, "updated: ", updated)
}
