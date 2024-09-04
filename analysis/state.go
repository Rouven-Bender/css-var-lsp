package analysis

import (
	"css-var-lsp/analysis/trie"
	"css-var-lsp/lsp"
	"fmt"
	"log"
	"strings"
)

var rope = trie.NewTrie()

type State struct {
	//Map of Filenames to Text content
	Documents map[string]string
	Logger    *log.Logger
	Trie      *trie.Trie
}

func NewState() State {
	return State{
		Documents: map[string]string{},
		Logger:    nil,
		Trie:      trie.NewTrie(),
	}
}

func (s *State) OpenLogger(log *log.Logger) {
	s.Logger = log
}
func (s *State) FillTrie() error {
	for keyword := range Keywords {
		s.Logger.Printf("current word beeing added is: %s", keyword)
		if err := s.Trie.Add(keyword); err != nil {
			return err
		}
	}
	return nil
}
func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) (*lsp.HoverResponse, error) {
	document := s.Documents[uri]
	s.Logger.Printf("ID: %d; position: line: %d, char: %d", id, position.Line, position.Character)
	word, err := selectedWord(document, position)
	if err != nil {
		return nil, err
	}
	info := Keywords[word]

	return &lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: info,
		},
	}, nil
}

func (s *State) TextDocumentCompletion(id int, uri string, position lsp.Position) *lsp.CompletionResponse {
	text := s.Documents[uri]
	word, err := selectedWord(text, position)
	if err != nil {
		s.Logger.Printf("error with text:%s; position: %d, %d: Error: %s", text, position.Line, position.Character, err)
	}
	completions, err := s.Trie.StartsWith(word)
	if err != nil {
		s.Logger.Printf("Completion request failed: %s", err)
	}
	out := []lsp.CompletionItem{}
	for _, completion := range completions {
		out = append(out, lsp.CompletionItem{
			Label:  completion,
			Detail: Keywords[completion],
		})
	}
	return &lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "",
			ID:  &id,
		},
		Result: out,
	}
}

func selectedWord(content string, position lsp.Position) (string, error) {
	selectedChar := content[position.Character]
	if selectedChar == byte(' ') {
		return "", fmt.Errorf("cursor on whitespace")
	}

	line := getLine(content, position.Line)

	word := isolateWord(line, position.Character)
	return word, nil
}

func getLine(content string, line int) string {
	var cursor int = 0
	for i := 0; i < line; i++ { // test if the 0 is the problem
		idx := strings.Index(content[cursor:], "\n")
		if idx != -1 {
			cursor += idx + 1
		} else {
			break
		}
	}
	end := cursor + strings.Index(content[cursor:], "\n")
	if end == -1 || cursor > end {
		end = len(content)
	}
	fmt.Printf("%d %d", cursor, end)
	return strings.Clone(content[cursor:end])
}

func isolateWord(line string, cursor int) string {
	if len(line) < cursor {
		return ""
	}
	var startW int = 0
	var endW int = 0

	for i := 1; cursor-i > 0; i++ {
		if byte(line[cursor-i]) == byte(' ') || byte(line[cursor-i]) == byte('(') {
			startW = cursor - i
			break
		}
	}

	for i := 1; cursor+i < len(line); i++ {
		if byte(line[cursor+i]) == byte(' ') || byte(line[cursor+i]) == byte(')') {
			endW = cursor + i
			break
		}
	}
	if endW < startW {
		endW = len(line)
	}
	if startW < 0 {
		startW = 0
	}
	if startW == endW && endW == 0 {
		return line
	}
	out := strings.Trim(line[startW:endW], " ")
	out = strings.Trim(out, "(")
	out = strings.Trim(out, ")")
	return out
}
