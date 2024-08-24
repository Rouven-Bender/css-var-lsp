package analysis

import "css-var-lsp/analysis/trie"

var Keywords map[string]string = map[string]string{
	"test": "Test word",
}

var Rarity map[string]uint8 = map[string]uint8{
	"test": trie.COMMON,
}
