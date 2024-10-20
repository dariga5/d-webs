package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DEBUG string = "DEBUG"
	WARN         = "WARN"
	INFO         = "INFO"
	ALERT        = "ALERT"
	ERROR        = "ERROR"
)

var (
	cfg *loggers
	mut sync.RWMutex
)

func InitLogger(path string) {
	cfg = config(path)
}

func log(logmsg *logInfo) {

	//TODO ADD TIME CREATE LOG
	format := time.Now().Format(time.RFC822) + " " + "[" + logmsg.Loglvl + "]" + " " + logmsg.Message

	for _, item := range cfg.Loggers {
		switch item.Mode {
		case "FILE":
			params := make(map[string]string)

			for _, paramtr := range item.Params {
				arr := strings.Split(paramtr, "=")

				key := arr[0]
				value := arr[1]

				params[key] = value
			}

			path, ok := params["path"]

			if ok {
				mut.Lock()
				file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
				if err == nil {
					_, err = file.WriteString(format + "\n")
					if err != nil {
						fmt.Println(err.Error())
					}
				}
				mut.Unlock()
			}
		default:
			fmt.Println(format)
		}
	}
}

func Debug(msg string) {
	log(&logInfo{
		Loglvl:  DEBUG,
		Message: msg,
	})
}

func Warn(msg string) {
	log(&logInfo{
		Loglvl:  WARN,
		Message: msg,
	})
}

func Info(msg string) {
	log(&logInfo{
		Loglvl:  INFO,
		Message: msg,
	})
}

func Error(msg string) {
	log(&logInfo{
		Loglvl:  ERROR,
		Message: msg,
	})
}

func Alert(msg string) {
	log(&logInfo{
		Loglvl:  ALERT,
		Message: msg,
	})
}
