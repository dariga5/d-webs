package loger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Logers struct {
	Loggers []LogerCfg `json:"Logers"`
}

func (lgrs *Logers) IsEmpty() bool {

	for _, item := range lgrs.Loggers {

		if !item.IsEmpty() {
			return false
		}
	}

	return true
}

type LogerCfg struct {
	Mode   string   `json:"Mode"`
	Params []string `json:"Params"`
}

func (cfg *LogerCfg) IsEmpty() bool {

	if cfg.Mode != "" {
		return false
	}

	for _, item := range cfg.Params {

		if item != "" {
			return false
		}
	}

	return true
}

func InitConfig(path string) Logers {

	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var cfgs Logers

	json.Unmarshal([]byte(byteValue), &cfgs)

	return cfgs
}
