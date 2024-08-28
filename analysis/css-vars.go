package analysis

import "css-var-lsp/analysis/trie"

var Keywords map[string]string = map[string]string{
	"test":            "Test word",
	"--css-var-works": "CSS variable to test -",
}

var Rarity map[string]uint8 = map[string]uint8{
	"test": trie.COMMON,
}
