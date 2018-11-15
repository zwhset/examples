package logger

import (
	"github.com/op/go-logging"
	"os"
)

var (
	Logger = logging.MustGetLogger("examples")

	LOGPATH = "logs/logs/example.log"
)

func init() {
	// File
	fd := logWriter(LOGPATH)
	fileBackend := logging.NewLogBackend(fd, "", 0)
	// Stdout
	outBackend := logging.NewLogBackend(os.Stdout, "", 0)
	format := logging.MustStringFormatter(
		// 2018-11-15 16:29:56.834 main.tLog INFO Info: main: 6
		`[%{time:2006-01-02 15:04:05.000}] [%{level}] %{shortpkg}.%{longfunc}: %{message}`,
	)
	fileBackendFormat := logging.NewBackendFormatter(fileBackend, format)
	outBackendFormat := logging.NewBackendFormatter(outBackend, format)
	logging.SetBackend(fileBackendFormat, outBackendFormat)
}

func logWriter(filename string) *os.File {
	// return
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}

	return fd
}
