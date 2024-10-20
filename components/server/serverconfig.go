package server

import (
	"d-webs/components/logger"
	"encoding/json"
	"io"
	"os"
)

type serverCfg struct {
	Ports []string `json:"Ports"`
}

func config(path string) *serverCfg {

	jsonFile, err := os.Open(path)

	if err != nil {
		logger.Error(err.Error())
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	cfg := new(serverCfg)

	json.Unmarshal([]byte(byteValue), cfg)

	return cfg
}
