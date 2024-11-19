package utils

import (
	"encoding/json"
	"errors"
	"os"
)

func LoadFromEnv() (net string, mne string) {
	return os.Getenv("NETWORK"), os.Getenv("MNEMONIC")
}

func IsValidMnemonic(mne string) bool {
	return mne != ""
}

func GetUrlsForEnv(net string) (chain string, relay string, err error) {
	switch net {
	case "dev":
		return "wss://tfchain.dev.grid.tf/ws", "wss://relay.dev.grid.tf", nil
	case "qa":
		return "wss://tfchain.qa.grid.tf/ws", "wss://relay.qa.grid.tf", nil
	case "test":
		return "wss://tfchain.test.grid.tf/ws", "wss://relay.test.grid.tf", nil
	case "main":
		return "wss://tfchain.grid.tf/ws", "wss://relay.grid.tf", nil
	default:
		return "", "", errors.New("invalid net")
	}
}

func Jsonify(data any) (string, error) {
	pres, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pres), nil
}
