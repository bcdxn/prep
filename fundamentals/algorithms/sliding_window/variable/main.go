package main

import "fmt"

func main() {
	var str = "abcabcbb"
	max := maxUniqueLength(str)
	fmt.Println("max length: ", max)
}

func maxUniqueLength(s string) int {
	set := make(map[byte]struct{})
	l := 0
	r := 0

	max := 0

	for r < len(s) {
		// while the character we're pointing to with our right index is in the set, walk in the left
		// pointer, shrinking the window and remove those characters from the set
		_, ok := set[s[r]]
		for ok {
			delete(set, s[l])
			l++
			_, ok = set[s[r]]
		}
		// add current char to set to prevent duplicates and grow our window
		set[s[r]] = struct{}{}
		r++
		// update max if appropriate
		if len(set) > max {
			max = len(set)
		}
	}

	return max
}
