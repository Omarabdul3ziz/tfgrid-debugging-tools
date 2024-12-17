package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
)

type Node struct {
	TwinId uint `json:"twinId"`
	NodeId uint `json:"nodeId"`
}

type Resp struct {
	Zos string `json:"zos"`
}

const (
	targetVersion = "3.12.8"
	mne           = ""
)

func fetchNodes() ([]Node, error) {
	res, err := http.Get("https://gridproxy.grid.tf/nodes?size=999999&healthy=true")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	by, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var nodes []Node
	if err := json.Unmarshal(by, &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

func main() {
	nodes, err := fetchNodes()
	if err != nil {
		log.Fatal(err)
	}

	chainMan := substrate.NewManager("wss://tfchain.grid.tf/ws")
	rmbClient, err := peer.NewRpcClient(context.Background(), mne, chainMan,
		peer.WithRelay("wss://relay.grid.tf"),
		peer.WithSession("rmb-call-bin"))
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	var old, new int
	failed := []Node{}
	var mu sync.Mutex

	for _, node := range nodes {
		wg.Add(1)

		go func(node Node) {
			defer wg.Done()

			var res Resp
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err := rmbClient.Call(ctx, uint32(node.TwinId), "zos.system.version", nil, &res)
			if err != nil {
				failed = append(failed, node)
				return
			}
			// fmt.Println("twinId: ", node.TwinId, "nodeId:", node.TwinId, "version: ", res.Zos)

			mu.Lock()
			if res.Zos == targetVersion {
				new++
			} else {
				old++
			}
			mu.Unlock()
		}(node)
	}

	wg.Wait()
	fmt.Println(old, new, old+new)
	fmt.Println(failed)
}