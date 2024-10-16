package loger

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var cfg Logers

func log(loglvl string, msg string) {

	if cfg.IsEmpty() {
		cfg = InitConfig("components/loger/exampleConf.json")
	}

	format := (time.Now().Format("02-Jan-2006 15:04:05")) + " " + "[" + loglvl + "]" + " " + msg + "\n"

	for _, loger := range cfg.Loggers {

		switch loger.Mode {

		case "STDOUT":
			{
				fmt.Println(format)
			}

		case "FILE":
			{
				for _, item := range loger.Params {

					var arr = strings.Split(item, "=")

					fileParams := make(map[string]string)

					fileParams[arr[0]] = arr[1]

					err := writeFile(fileParams, format)

					if err != nil {
						fmt.Println(err)
					}
				}
			}
		default:
			{
				fmt.Println(format)
			}
		}
	}
}

func writeFile(params map[string]string, str string) error {

	path := params["path"]

	if path == "" {
		return errors.New("file path for log-file not set")
	}

	if str == "" {
		return errors.New("Message not set")
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0655)

	defer file.Close()

	if err != nil {
		return err
	}

	_, err = file.WriteString(str)

	if err != nil {
		return err
	}

	return nil
}

func Debug(msg string) {
	log("DEBUG", msg)
}

func Info(msg string) {
	log("INFO", msg)
}

func Notice(msg string) {
	log("NOTICE", msg)
}

func Warn(msg string) {
	log("WARN", msg)
}

func Error(msg string) {
	log("ERR", msg)
}

func Crit(msg string) {
	log("CRIT", msg)
}

func Alert(msg string) {
	log("ALERT", msg)
}
