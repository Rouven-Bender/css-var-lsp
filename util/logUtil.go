package util

import (
	"log"
	"os"
)

func GetLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[css-var-lsp]", log.Ltime|log.Lshortfile)
}
