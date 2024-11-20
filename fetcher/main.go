package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"slices"
	"time"
)

// get the nodes reported in last 2d - nodes that is up now - nodes that is standby = nodes that has a problem

type Node struct {
	NodeID    int `json:"nodeId"`
	TwinID    int `json:"twinId"`
	UpdatedAt int `json:"updatedAt"`
}

type Twin struct {
	Relay string `json:"relay"`
}

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

func fetchTwinRelay(twinID int) (string, error) {
	url := fmt.Sprintf("https://gridproxy.grid.tf/twins?twin_id=%d", twinID)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var twins []Twin
	if err := json.Unmarshal(body, &twins); err != nil {
		return "", err
	}

	if len(twins) > 0 {
		return twins[0].Relay, nil
	}
	return "", fmt.Errorf("relay not found for twin ID %d", twinID)
}

func main() {
	allnodes, err := fetchNodes("")
	if err != nil {
		fmt.Println("Error fetching nodes:", err)
		return
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
		return
	}

	weirdNodes := []Node{}
	for _, node := range upInTwoDays {
		if !slices.Contains(upStandbyNodes, node) {
			weirdNodes = append(weirdNodes, node)
		}
	}

	fmt.Println(weirdNodes, len(weirdNodes))
}
