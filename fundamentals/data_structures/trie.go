package ds

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

// Insert a word into the trie.
func (t *Trie) Insert(word string) {
	node := t.root

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = newTrieNode()
		}
		node = node.children[c]
	}

	node.word = true
}

// Searches for a word in the Trie and returns true if it exists.
func (t *Trie) Search(word string) bool {
	node := t.root

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return node.word
}

// Search the tree to see if there are words in the tree with the given prefix.
func (t *Trie) StartsWith(word string) bool {
	node := t.root

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return true
}

type TrieNode struct {
	children map[rune]*TrieNode
	word     bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: map[rune]*TrieNode{},
		word:     false,
	}
}
