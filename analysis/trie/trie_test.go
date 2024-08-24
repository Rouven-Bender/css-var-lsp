package trie_test

import (
	"css-var-lsp/analysis/trie"
	"testing"
)

func TestAdding(t *testing.T) {
	rope := trie.NewTrie()
	err := rope.Add("test")
	if err != nil {
		t.Error(err)
	}
	if !rope.Contains("test") {
		t.Errorf("trie does not contain word that was added")
	}
}

func TestFindwords(t *testing.T) {
	rope := trie.NewTrie()
	expected := struct {
		words []string
	}{
		words: []string{"test", "testing", "testcase"},
	}
	for _, word := range expected.words {
		if err := rope.Add(word); err != nil {
			t.Error(err)
		}
	}
	findings, err := rope.StartsWith("tes")
	if err != nil {
		t.Error(err)
	}
outer:
	for _, word := range expected.words {
		for _, found := range findings {
			if word == found {
				break outer
			}
		}
		t.Errorf("word: %s was not found", word)
	}
}