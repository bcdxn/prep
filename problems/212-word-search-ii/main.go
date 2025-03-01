package main

import "fmt"

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}

	foundWords := findWords(board, words)
	fmt.Println("found words:", foundWords)
}

func findWords(board [][]byte, words []string) []string {
	// build Trie
	t := newTrie()
	for _, word := range words {
		t.Insert(word)
	}
	// DFS on each node, keeping track of found words
	foundWords := []string{}
	for r := range board {
		for c := range board[r] {
			foundWords = append(foundWords, dfsWrapper(board, r, c, t)...)
		}
	}

	wordsMap := map[string]struct{}{}
	uniqueWords := []string{}

	for _, word := range foundWords {
		if _, ok := wordsMap[word]; !ok {
			uniqueWords = append(uniqueWords, word)
			wordsMap[word] = struct{}{}
		}
	}

	return uniqueWords
}

type Trie struct {
	root *TrieNode
}

/* DFS
-------------------------------------------------------------------- */

func dfsWrapper(m [][]byte, r, c int, t *Trie) []string {
	foundWords := map[string]struct{}{}
	rows := len(m)
	cols := len(m[0])
	visited := map[rc]struct{}{}

	var dfs func(
		r, c int,
		word string,
		node *TrieNode,
	)

	dfs = func(
		r, c int,
		word string,
		node *TrieNode,
	) {
		// can't revisit a node
		if _, ok := visited[rc{r, c}]; ok {
			return
		}
		// can't run out of bounds
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		// no word with the given prefix exists; stop DFS
		if _, ok := node.children[rune(m[r][c])]; !ok {
			return
		}
		// update node as visited
		visited[rc{r, c}] = struct{}{}
		// update which node in the Trie we're on
		node = node.children[rune(m[r][c])]
		// build our word as we recurse
		word = word + string(m[r][c])
		// If we're at an isWord node, add the completed word to the set
		if node.isWord {
			foundWords[word] = struct{}{}
		}
		// Recurse
		dfs(r-1, c, word, node)
		dfs(r, c+1, word, node)
		dfs(r+1, c, word, node)
		dfs(r, c-1, word, node)
		// Remove the node from the visited set on wait back up
		delete(visited, rc{r, c})
	}

	// start DFSing
	dfs(
		r, c,
		"",
		t.root,
	)

	wordList := []string{}

	for key := range foundWords {
		wordList = append(wordList, key)
	}

	return wordList
}

type rc struct {
	r int
	c int
}

/* Trie
------------------------------------------------------------------- */

func newTrie() *Trie {
	return &Trie{
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

func (t *Trie) ContainsPrefix(prefix string) bool {
	node := t.root

	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return true
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: map[rune]*TrieNode{},
	}
}
