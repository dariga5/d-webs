package main

import (
	logger "d-webs/components/logger"
	server "d-webs/components/server"
)

func main() {
	logger.InitLogger("config.json")
	server.InitServer("config.json")
}
