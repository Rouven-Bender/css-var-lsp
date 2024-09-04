package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... there is more here but I don't care
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync   int               `json:"textDocumentSync"`
	HoverProvider      bool              `json:"hoverProvider"`
	CompletionProvider CompletionOptions `json:"completionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CompletionOptions struct {
	TriggerCharacters []string `json:"triggerCharacters"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1, //Full content updates every time
				HoverProvider:    true,
				CompletionProvider: CompletionOptions{
					TriggerCharacters: []string{"-"},
				},
			},
			ServerInfo: ServerInfo{
				Name:    "css-var-lsp",
				Version: "0.0.0.dev",
			},
		},
	}
}
