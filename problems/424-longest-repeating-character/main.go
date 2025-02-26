package main

import "fmt"

func main() {
	s := "AABABBA"
	max := characterReplacement(s, 1)
	fmt.Println("max: ", max) // 4
}

func characterReplacement(s string, k int) int {
	frequencies := map[byte]int{}

	max := 0
	l := 0

	// look for window with the most characters that are the same. If the number of characters not
	// the same in that window is <= k then it is a valid window; else bring in left pointer until
	// the window is valid.

	for r := range len(s) {

		if count, ok := frequencies[s[r]]; ok {
			frequencies[s[r]] = count + 1
		} else {
			frequencies[s[r]] = 1
		}

		for (r-l+1)-maxFrequency(frequencies) > k {
			frequencies[s[l]] = frequencies[s[l]] - 1
			l++
		}

		window := r - l + 1

		if window > max {
			max = window
		}
	}

	return max
}

func maxFrequency(frequencies map[byte]int) int {
	max := 0

	for _, val := range frequencies {
		if val > max {
			max = val
		}
	}

	return max
}
