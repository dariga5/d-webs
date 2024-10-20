package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type logInfo struct {
	Loglvl  string
	Message string
}

type loggers struct {
	Loggers []loggerCfg `json:"Loggers"`
}

type loggerCfg struct {
	Mode   string   `json:"Mode"`
	Params []string `json:"Params"`
}

func config(path string) *loggers {

	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	loggers := new(loggers)

	json.Unmarshal([]byte(byteValue), loggers)

	return loggers
}
