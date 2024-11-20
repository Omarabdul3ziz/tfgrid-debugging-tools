package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"
)

func fetchNodes(status string) ([]Node, error) {
	resp, err := http.Get(fmt.Sprintf("https://gridproxy.grid.tf/nodes?size=999999&status=%s", status))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var nodes []Node
	if err := json.Unmarshal(body, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

func getweridnodes() ([]Node, error) {
	allnodes, err := fetchNodes("")
	if err != nil {
		fmt.Println("Error fetching nodes:", err)
		return nil, err
	}

	twoDaysAgo := time.Now().Add(-48 * time.Hour).Unix()
	upInTwoDays := []Node{}
	for _, node := range allnodes {
		if int64(node.UpdatedAt) >= twoDaysAgo {
			upInTwoDays = append(upInTwoDays, node)
		}
	}

	upStandbyNodes, err := fetchNodes("up,standby")
	if err != nil {
		fmt.Println("Error fetching nodes:", err)
		return nil, err
	}

	weirdNodes := []Node{}
	for _, node := range upInTwoDays {
		if !slices.Contains(upStandbyNodes, node) {
			weirdNodes = append(weirdNodes, node)
		}
	}

	return weirdNodes, nil
}
