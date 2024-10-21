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

		logger.Info("SERVER OPEN PORT " + port)

		l, err := net.Listen("tcp", port)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		for {
			conn, err := l.Accept()

			if err != nil {
				logger.Error(err.Error())
			}

			go func(conn net.Conn) {
				buff := make([]byte, 64)

				conn.Read(buff)

				logger.Debug("MSG - " + string(buff[:]))

			}(conn)
		}
	}
}
