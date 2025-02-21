package main

import "fmt"

func main() {
	haystack := []int{1, 3, 3, 4, 5, 6, 7, 8}

	index := bs(6, haystack)

	fmt.Printf("found needle %d at index %d\n", 6, index)
}

func bs(needle int, haystack []int) int {
	left := 0
	right := len(haystack) - 1

	for left <= right {
		// mid := (left + right) / 2        // -- simplified expression to calculate midpoint
		mid := left + (right-left)/2

		if needle < haystack[mid] {
			// search left half of haystack
			right = mid - 1
		} else if needle > haystack[mid] {
			// search left half of haystack
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
