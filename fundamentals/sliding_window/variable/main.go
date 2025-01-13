package main

import "fmt"

func main() {
	t := "abcabcbb"

	n := longestSubStr(t) // should print 3

	fmt.Println("longest:", n)
}

func longestSubStr(s string) int {
	charSet := make(map[byte]struct{})

	max := 0 // max length of substring
	l := 0   // start of window

	for r := range len(s) {
		_, ok := charSet[s[r]]
		for ok {
			delete(charSet, s[l])
			l++
			_, ok = charSet[s[r]]
		}
		charSet[s[r]] = struct{}{}
		// finding length with index arithmetic
		// currWindowSize := r - l + 1
		// if currWindowSize > max {
		// 	max = currWindowSize
		// }
		// find length using built in len function
		if len(charSet) > max {
			max = len(charSet)
		}
	}

	return max
}
