package main

import (
	"bufio"
	"css-var-lsp/analysis"
	"css-var-lsp/lsp"
	"css-var-lsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/kyu/src/css-var-lsp/log.txt")
	logger.Println("LSP started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("error: %s", err)
			continue
		}
		handleMessage(logger, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldn't parse the init Request: %s", err)
		}
		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Print("Sent the init response")
	case "textDocument/didOpen":
		var request lsp.DidOpenNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("didn't open: %s", err)
			return
		}
		logger.Printf("Opened: %s", request.Params.TextDocument.Uri)
		state.OpenDocument(request.Params.TextDocument.Uri, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.DidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("didChangeNotification: %s", err)
			return
		}
		logger.Printf("Ch: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.OpenDocument(request.Params.TextDocument.URI, change.Text)
		}
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[css-var-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
