package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	words := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(words))
}

func groupAnagrams(words []string) [][]string {
	anagramGroups := make([][]string, 0)
	anagrams := make(map[string]int)

	// loop given words
	for _, word := range words {
		currAnagram := make(map[rune]int)
		// build anagram set
		for _, ch := range word {
			if count, ok := currAnagram[ch]; ok {
				currAnagram[ch] = count + 1
			} else {
				currAnagram[ch] = 1
			}
		}
		// sort characters
		chars := make([]rune, 0, len(currAnagram))
		for ch := range currAnagram {
			chars = append(chars, ch)
		}
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		// build a key from the anagram a2b3c1
		key := ""
		for _, ch := range chars {
			key += string(ch) + strconv.Itoa(currAnagram[ch])
		}
		if index, ok := anagrams[key]; ok {
			anagramGroups[index] = append(anagramGroups[index], word)
		} else {
			anagramGroups = append(anagramGroups, []string{word})
			anagrams[key] = len(anagramGroups) - 1
		}
	}

	return anagramGroups
}
