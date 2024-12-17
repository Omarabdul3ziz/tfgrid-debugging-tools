package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

const (
	chainUrl = "wss://tfchain.dev.grid.tf:443"
	relayUrl = "wss://relay.dev.grid.tf"
	destTwin = 5108
	destNode = 131
	srcTwin  = 29
)

var mnemonic = ""

func zdb(size gridtypes.Unit, name string) gridtypes.Workload {
	return gridtypes.Workload{
		Type: zos.ZDBType,
		Name: gridtypes.Name(name),
		Data: gridtypes.MustMarshal(zos.ZDB{
			Size:     size,
			Mode:     zos.ZDBModeUser,
			Password: "password",
		},
		),
	}
}

func deployment(twin uint32, workloads []gridtypes.Workload) gridtypes.Deployment {
	return gridtypes.Deployment{
		Version:     0,
		TwinID:      twin,
		Workloads:   workloads,
		Description: "1",
		Metadata:    "zdbtest",
		SignatureRequirement: gridtypes.SignatureRequirement{
			WeightRequired: 1,
			Requests: []gridtypes.SignatureRequest{
				{
					TwinID: twin,
					Weight: 1,
				},
			},
		},
	}
}

func createContract(mgr substrate.Manager, dl gridtypes.Deployment) (uint64, error) {
	identity, err := substrate.NewIdentityFromSr25519Phrase(mnemonic)
	if err != nil {
		panic(err)
	}

	if err := dl.Valid(); err != nil {
		return 0, fmt.Errorf("not valid dep %w", err)
	}

	if err := dl.Sign(srcTwin, identity); err != nil {
		return 0, err
	}

	hash, err := dl.ChallengeHash()
	if err != nil {
		return 0, err
	}
	hashHex := hex.EncodeToString(hash)

	sub, err := mgr.Substrate()
	if err != nil {
		return 0, err
	}

	contractId, err := sub.CreateNodeContract(identity, destNode, "", hashHex, 0, nil)
	if err != nil {
		return 0, fmt.Errorf("failed creating contract %w", err)
	}

	return contractId, nil
}

// func deploy() uint64 {
// 	dl := deployment(srcTwin, []gridtypes.Workload{
// 		zdb(1*gridtypes.Gigabyte, "zdbtest"),
// 	})
//
// 	// Create contract on the chain
// 	mgr := substrate.NewManager(chainUrl)
// 	contractId, err := createContract(mgr, dl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	dl.ContractID = contractId
//
// 	// Deploy deployment on the grid
// 	cl, err := peer.NewRpcClient(context.Background(), mnemonic, mgr, peer.WithRelay(relayUrl))
// 	if err != nil {
// 		panic(err)
// 	}
// 	var res interface{}
// 	if err := cl.Call(context.Background(), destTwin, "zos.deployment.deploy", dl, &res); err != nil {
// 		panic(err)
// 	}
//
// 	return contractId
// }

func main() {
	mnemonic = os.Getenv("MNEMONIC")
	dl := deployment(srcTwin, []gridtypes.Workload{
		zdb(1*gridtypes.Gigabyte, "zdbtest"),
	})

	mgr := substrate.NewManager(chainUrl)
	cl, err := peer.NewRpcClient(context.Background(), mnemonic, mgr, peer.WithRelay(relayUrl))
	if err != nil {
		panic(err)
	}

	var res interface{}
	contractId := uint64(0)

	if os.Args[0] != "onlyget" {
		fmt.Println("deploying deployment")
		contractId, err = createContract(mgr, dl)
		if err != nil {
			panic(err)
		}
		dl.ContractID = contractId

		if err := cl.Call(context.Background(), destTwin, "zos.deployment.deploy", dl, &res); err != nil {
			panic(err)
		}

		fmt.Println("deployment successed with contract id:", contractId)
	}

	fmt.Println("getting deployment details, contactID: ", contractId)
	if err := cl.Call(context.Background(), destTwin,
		"zos.deployment.get",
		struct{ ContractId uint64 }{ContractId: contractId},
		&res); err != nil {
		panic(err)
	}

	_bytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(_bytes))
}
