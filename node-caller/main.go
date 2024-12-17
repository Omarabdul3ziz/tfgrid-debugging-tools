package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Omarabdulaziz/tfgrid_debugging_tools/utils"
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var dest uint
	var cmd, payload string
	flag.UintVar(&dest, "dest", 0, "destination")
	flag.StringVar(&cmd, "cmd", "", "command")
	flag.StringVar(&payload, "payload", "", "payload")
	flag.Parse()

	if dest == 0 || cmd == "" {
		return errors.New("missing flag/envvar")
	}

	net, mne := utils.LoadFromEnv()
	if !utils.IsValidMnemonic(mne) {
		return errors.New("invalid mnemonic")
	}

	chain, relay, err := utils.GetUrlsForNet(net)
	if err != nil {
		return err
	}

	man := substrate.NewManager(chain)
	rmb, err := peer.NewRpcClient(context.Background(), mne, man, peer.WithRelay(relay), peer.WithSession("debugging-tools")) // Todo: add twin id to session
	if err != nil {
		return err
	}

	var res interface{}
	if err := rmb.Call(context.Background(), uint32(dest), cmd, payload, &res); err != nil {
		return err
	}

	output, err := utils.Jsonify(res)
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}
