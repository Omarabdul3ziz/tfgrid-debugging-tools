package main

import (
	"fmt"
	"log"

	"github.com/Omarabdul3ziz/tfgrid-debugging-tools/utils"
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
)

func main() {
	chain, _, _ := utils.GetUrlsForEnv("dev")
	manager := substrate.NewManager(chain)

	con, err := manager.Substrate()
	if err != nil {
		log.Fatal(err)
	}

	v, err := con.GetZosVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	node, err := con.GetNode(11)
	if err != nil {
		log.Fatal(err)
	}

	res, err := utils.Jsonify(node)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
