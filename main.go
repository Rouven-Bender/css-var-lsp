package main

import (
	"bufio"
	"css-var-lsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/kyu/src/css-var-lsp/log.txt")
	logger.Println("LSP started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[css-var-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
