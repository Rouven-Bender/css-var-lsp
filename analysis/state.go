package analysis

import (
	"css-var-lsp/lsp"
	"fmt"
	"log"
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

func (s *State) Hover(id int, uri string, position lsp.Position, log *log.Logger) (*lsp.HoverResponse, error) {
	document := s.Documents[uri]
	word, err := selectedWord(document, position, log)
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

func selectedWord(content string, position lsp.Position, log *log.Logger) (string, error) {
	selectedChar := content[position.Character]
	if selectedChar == byte(' ') {
		return "", fmt.Errorf("cursor on whitespace")
	}

	line := getLine(content, position.Line)

	return isolateWord(line, position.Character), nil
}

func getLine(content string, line int) string {
	var cursor int = 0
	for i := 1; i < line; i++ { // test if the 0 is the problem
		idx := strings.Index(content[cursor:], "\n")
		if idx != -1 {
			cursor += idx + 1
		} else {
			break
		}
	}
	//log.Print(content[cursor:])
	return strings.Clone(content[cursor:])
}

func isolateWord(line string, cursor int) string {
	if len(line) < cursor {
		return ""
	}
	var startW int = 0
	var endW int = 0

	for i := 1; cursor-i > 0; i++ {
		if byte(line[cursor-i]) == byte(' ') {
			startW = cursor - i
			break
		}
	}

	for i := 1; cursor+i < len(line); i++ {
		if byte(line[cursor+i]) == byte(' ') {
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
	return strings.Trim(line[startW:endW], " ")
}
