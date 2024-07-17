package lsp

type DidOpenNotification struct {
	Notification
	Params DidOpenParams `json:"params"`
}

type DidOpenParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
