package analysis

import (
	"css-var-lsp/lsp"
	"css-var-lsp/util"
	"fmt"
	"log"
	"strings"
)

var logger *log.Logger = util.GetLogger("/dev/shm/css-var-analysis.log")

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
	logger.Printf("ID: %d; position: line: %d, char: %d", id, position.Line, position.Character)
	word, err := selectedWord(document, position)
	if err != nil {
		return nil, err
	}
	//info := Keywords[word]

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
	selectedChar := content[position.Character]
	if selectedChar == byte(' ') {
		return "", fmt.Errorf("cursor on whitespace")
	}

	line := getLine(content, position.Line)
	logger.Printf("line:%s\n", line)

	word := isolateWord(line, position.Character)
	logger.Printf("word:%s", word)
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
	if end == -1 {
		end = len(content)
	}
	return strings.Clone(content[cursor:end])
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
	if startW == endW && endW == 0 {
		return line
	}
	logger.Printf("%d %d", startW, endW)
	return strings.Trim(line[startW:endW], " ")
}
