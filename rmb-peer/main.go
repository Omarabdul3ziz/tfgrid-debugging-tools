package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Omarabdul3ziz/tfgrid-debugging-tools/utils"
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
)

func echo(ctx context.Context, payload []byte) (interface{}, error) {
	fmt.Println("received", string(payload))

	// decode payload
	var msg string
	if err := json.Unmarshal(payload, &msg); err != nil {
		log.Fatal("failed to decode payload", err)
	}

	// respond
	return msg, nil
}

func perfTest(ctx context.Context, payload []byte) (interface{}, error) {
	fmt.Println("received", string(payload))

	// decode payload
	var msg struct {
		Name string
	}
	if err := json.Unmarshal(payload, &msg); err != nil {
		log.Fatal("failed to decode payload", err)
	}

	// respond
	res := struct {
		Name        string
		Description string
		Result      map[string][]string
	}{
		Name:        msg.Name,
		Description: "description",
		Result: map[string][]string{
			"network": {},
			"cache":   {},
		},
	}
	return res, nil
}

func main() {
	router := peer.NewRouter()
	app := router.SubRoute("local")
	app.WithHandler("echo", echo)
	app.WithHandler("getperf", perfTest)

	net, mne := utils.LoadFromEnv()
	if !utils.IsValidMnemonic(mne) {
		log.Fatal("mnemonic is required")
	}

	chain, relay, _ := utils.GetUrlsForEnv(net)
	mgr := substrate.NewManager(chain)
	_, err := peer.NewPeer(context.Background(), mne, mgr,
		router.Serve,
		peer.WithRelay(relay),
		peer.WithSession("debugging-tools"),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("peer started.")
	// block
	select {}
}
