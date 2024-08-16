package analysis

import (
	"css-var-lsp/lsp"
	"testing"
)

type testingTable struct {
	input    string
	expected string
	position lsp.Position
}

func TestSelectedWord(t *testing.T) {
	expected := []testingTable{
		{input: "This is the test Text\nHans ist ein Name\n this should be selected",
			expected: "this",
			position: lsp.Position{
				Line:      3,
				Character: 3,
			},
		},
		{input: "test\nTest\ntEst",
			expected: "tEst",
			position: lsp.Position{
				Line:      3,
				Character: 3,
			},
		},
	}
	for _, test := range expected {
		selected, err := selectedWord(test.input, test.position)
		if err != nil {
			t.Fatal(err)
		}
		if selected != test.expected {
			t.Fatalf("word selected was: \"%s\", expected \"%s\"", selected, test.expected)
		}
	}
}

func TestWordIsolation(t *testing.T) {
	expected := []testingTable{
		{input: "Das ist ein Test ob mein File lesen geht",
			expected: "Das",
			position: lsp.Position{
				Line:      0,
				Character: 1,
			},
		},
		{input: "Das ist ein Test ob mein File lesen geht",
			expected: "Test",
			position: lsp.Position{
				Line:      0,
				Character: 15,
			},
		},
		{input: "Das ist ein Test ob mein File lesen geht",
			expected: "geht",
			position: lsp.Position{
				Line:      0,
				Character: 40,
			},
		},
	}
	for _, test := range expected {
		selected := isolateWord(test.input, test.position.Character)
		if selected != test.expected {
			t.Fatalf("word selected was: \"%s\", expected \"%s\"", selected, test.expected)
		}
	}

}

func TestGetLine(t *testing.T) {
	expected := []testingTable{
		{input: "This is the test Text\nHans ist ein Name\n this should be selected",
			expected: " this should be selected",
			position: lsp.Position{
				Line:      3,
				Character: 3,
			},
		},
	}
	for _, test := range expected {
		line := getLine(test.input, test.position.Line)
		if line != test.expected {
			t.Fatalf("line expected: \"%s\"\nline got: \"%s\"", test.expected, line)
		}
	}
}
