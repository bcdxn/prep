package main

import (
	"fmt"
)

func main() {
	words := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(words))
}

func groupAnagrams(words []string) [][]string {
	anagramGroups := make([][]string, 0)
	anagramsMap := make(map[[26]int]int)

	// use an array as hashmap for ascii characters
	// aaabc
	// [3, 1, 1]

	// loop given words
	for _, word := range words {
		currAnagramKey := [26]int{}
		a := int('a')
		// build anagram set
		for _, ch := range word {
			i := int(ch) - a
			currAnagramKey[i] += 1
		}

		if index, ok := anagramsMap[currAnagramKey]; ok {
			anagramGroups[index] = append(anagramGroups[index], word)
		} else {
			anagramGroups = append(anagramGroups, []string{word})
			anagramsMap[currAnagramKey] = len(anagramGroups) - 1
		}
	}

	return anagramGroups
}
