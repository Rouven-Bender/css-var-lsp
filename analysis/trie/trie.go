package trie

import (
	"fmt"
	"strings"
)

const (
	WORD          uint8 = 128
	RARE          uint8 = 16
	LOW           uint8 = 8
	MEDIUM        uint8 = 4
	COMMON        uint8 = 2
	UBIQUITOUS    uint8 = 0
	numberOfChars int   = 26
)

type Node struct {
	Value      rune
	LetterArms [numberOfChars]*Node
	Flag       uint8
}

func (n *Node) setWord(toggle bool) {
	if toggle {
		n.Flag = n.Flag | WORD
	} else {
		n.Flag = n.Flag & ^WORD
	}
}

func (n *Node) setRarity(rarity uint8) error {
	switch rarity {
	case RARE:
	case LOW:
	case MEDIUM:
	case COMMON:
	case UBIQUITOUS:
		n.Flag = n.Flag | rarity
		break
	default:
		return fmt.Errorf("Not a valid rarity")
	}
	return nil
}

func (n *Node) isWord() bool {
	return (n.Flag & WORD) == WORD
}

type Trie struct {
	Head Node
}

func NewTrie() *Trie {
	return &Trie{
		Head: Node{
			Value: ' ',
		},
	}
}

func (t *Trie) Add(word string) error {
	if len(word) < 1 {
		return fmt.Errorf("string is empty")
	}
	word = strings.ToLower(word)
	chars := []rune(word)
	node := &t.Head
	for i := 0; i < len(chars); i++ {
		letter := byte(chars[i])
		idx := int(letter - byte(rune('a')))
		if 0 <= idx && idx < numberOfChars {
			if node.LetterArms[idx] != nil {
				node = node.LetterArms[idx]
			} else {
				node.LetterArms[idx] = &Node{
					Value: chars[i],
				}
				node = node.LetterArms[idx]
			}
		} else {
			return fmt.Errorf("letter out of range: %c", chars[i])
		}
	}
	node.setWord(true)
	return nil
}

func (t *Trie) Contains(word string) bool {
	if len(word) < 1 {
		return false
	}
	word = strings.ToLower(word)
	node := &t.Head
	chars := []rune(word)
	for i := 0; i < len(chars); i++ {
		letter := byte(chars[i])
		idx := int(letter - byte(rune('a')))
		if 0 <= idx && idx < numberOfChars {
			if node.LetterArms[idx] != nil {
				node = node.LetterArms[idx]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return node.isWord()
}

func (t *Trie) StartsWith(start string) ([]string, error) {
	if len(start) < 1 {
		return nil, fmt.Errorf("input empty")
	}

	out := []string{}
	start = strings.ToLower(start)
	node := &t.Head
	chars := []rune(start)

	for i := 0; i < len(chars); i++ {
		letter := byte(chars[i])
		idx := int(letter - byte(rune('a')))
		if 0 <= idx && idx < numberOfChars {
			if node.LetterArms[idx] != nil {
				node = node.LetterArms[idx]
			} else {
				return out, nil
			}
		} else {
			return nil, fmt.Errorf("letter out of range")
		}
	}

	//find words
	path := chars[:len(chars)-1] // to remove the last letter of chars as in findwords the value of the current node gets added
	out = findWords(node, path)
	return out, nil
}

func findWords(node *Node, path []rune) []string {
	out := []string{}
	if node == nil {
		return nil
	}
	if node.isWord() {
		word := string(append(path, node.Value))
		out = append(out, word)
	}
	for i := 0; i < numberOfChars; i++ {
		findings := findWords(node.LetterArms[i], append(path, node.Value))
		if findings != nil {
			out = append(out, findings...)
		}
	}
	return out
}
