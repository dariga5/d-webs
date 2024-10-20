package server

import (
	logger "d-webs/components/logger"
	"net"
)

var (
	serveCfg *serverCfg
)

func InitServer(path string) {

	serveCfg = config(path)

	if serveCfg.Ports == nil {
		logger.Error("Server. Config parse error. Ports is null")
		return
	}

	for _, port := range serveCfg.Ports {
		l, err := net.Listen("tcp4", port)

		if err != nil {
			logger.Error(err.Error())
			return
		} else {
			logger.Info("Server open port " + port)
		}

		defer func() {
			l.Close()
			logger.Info("Server closed port " + port)
		}()

	}
}
