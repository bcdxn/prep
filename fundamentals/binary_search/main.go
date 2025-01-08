package main

import "fmt"

func main() {
	haystack := []int{1, 3, 3, 4, 5, 6, 7, 8}
	i := binarySearch(haystack, 6)
	fmt.Println("index found:", i)
}

func binarySearch(haystack []int, needle int) int {
	l := 0
	r := len(haystack) - 1

	for l <= r {
		mid := (l + r) / 2

		if needle < haystack[mid] {
			r = mid - 1
		} else if needle > haystack[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
