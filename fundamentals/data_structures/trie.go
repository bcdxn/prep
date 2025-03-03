package ds

func NewTrie() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = newTrieNode()
		}
		node = node.children[c]
	}
	node.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root

	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return true
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode, 26),
		isWord:   false,
	}
}

type Trie struct {
	root *trieNode
}

type trieNode struct {
	children map[rune]*trieNode
	isWord   bool
}
