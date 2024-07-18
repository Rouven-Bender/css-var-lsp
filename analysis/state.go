package analysis

import (
	"css-var-lsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	//Map of Filenames to Text content
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) (*lsp.HoverResponse, error) {
	document := s.Documents[uri]
	word, err := selectedWord(document, position)
	if err != nil {
		return nil, err
	}

	return &lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: word,
		},
	}, nil
}

func selectedWord(content string, position lsp.Position) (string, error) {
	var leftover string = content
	var i int
	var idx int

	for i = 1; i < position.Line; i++ {
		idx := strings.Index(leftover, "\n")
		if idx == -1 {
			return "", fmt.Errorf("no line break found")
		}
		idx = idx + 2
		if idx > 0 {
			leftover = leftover[idx:]
		} else {
			leftover = leftover[0:]
		}
	}

	selectedChar := leftover[position.Character]
	fmt.Println(string(selectedChar))
	if selectedChar == byte(' ') {
		return "", fmt.Errorf("cursor on whitespace")
	}

	for i = 1; position.Character-i > 0; i++ {
		if leftover[position.Character-i] == byte(' ') {
			break
		} else {
			continue
		}
	}

	idx = position.Character - i
	if idx > 0 {
		leftover = leftover[position.Character-i:]
	} else {
		leftover = leftover[0:]
	}

	for i = 1; position.Character+i < len(leftover); i++ {
		if leftover[position.Character+i] == byte(' ') {
			break
		} else {
			continue
		}
	}
	return leftover[:position.Character+i], nil
}
