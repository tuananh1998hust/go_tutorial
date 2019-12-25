package main

import (
	"os"

	"github.com/op/go-logging"
)

var (
	log = getLogLevel()
)

// getLogLevel return log level
func getLogLevel() *logging.Logger {
	var log = logging.MustGetLogger("example")
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatter)

	return log
}

func main() {
	// log.Debugf("debug %s", "test")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Info("test")
}
