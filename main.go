package main

import (
	logger "d-webs/components/logger"
)

func main() {
	logger.InitLogger("config.json")

	logger.Debug("HELLO FROM DEBUG")
	logger.Warn("HELLO FROM WARN")
	logger.Info("HELLO FROM INFO")
	logger.Alert("HELLO FROM ALERT")
	logger.Error("HELLO FROM ERROR")
}
