package main

import (
	"bufio"
	"css-var-lsp/analysis"
	"css-var-lsp/lsp"
	"css-var-lsp/rpc"
	"css-var-lsp/util"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	logger := util.GetLogger("/home/kyu/src/css-var-lsp/log.txt")
	logger.Println("LSP started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout

	state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("error: %s", err)
			continue
		}
		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
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
		writeResponse(writer, msg)
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
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}
		response, err := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		if err == nil {
			writeResponse(writer, *response)
		}
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
