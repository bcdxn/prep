package main

import "fmt"

func main() {
	haystack := []int{1, 3, 3, 4, 5, 6, 7, 8}
	needle := 6
	i := binarySearch(haystack, needle)
	fmt.Println("index found:", i)
}

func binarySearch(haystack []int, needle int) int {
	l := 0
	r := len(haystack) - 1

	for l <= r {
		m := (l + r) / 2

		if needle < haystack[m] {
			r = m - 1
		} else if needle > haystack[m] {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}
