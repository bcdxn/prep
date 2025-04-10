package main

import "fmt"

func main() {
	haystack := []int{1, 3, 3, 4, 5, 6, 7, 8}

	i := find(3, haystack)
	fmt.Println("found", i)
	i = find(4, haystack)
	fmt.Println("found", i)
	i = find(8, haystack)
	fmt.Println("found", i)
	i = find(1, haystack)
	fmt.Println("found", i)
	i = find(5, haystack)
	fmt.Println("found", i)
	i = find(10, haystack)
	fmt.Println("found", i)
	i = find(2, haystack)
	fmt.Println("found", i)
}

func find(needle int, haystack []int) int {
	l := 0
	r := len(haystack) - 1

	for l <= r {
		m := l + ((r - l) / 2)

		if needle == haystack[m] {
			return m
		} else if needle < haystack[m] {
			r = m - 1
		} else if needle > haystack[m] {
			l = m + 1
		}
	}

	return -1
}
