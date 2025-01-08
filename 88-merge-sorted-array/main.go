package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(
		nums1,
		3,
		nums2,
		len(nums2),
	)

	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	s := 0 // index to insert next sorted/merged number
	i := 0 // tracks where we are (count) in nums1
	j := 0 // tracks wgere we are in nums2

	for j < n {
		// exhausted all of nums1; simply merge in nums2 without comparison
		if i >= m {
			nums1[s] = nums2[j]
			j++
			s++
			continue
		}

		// exhausted all of nums2; and remaining nums1 is already sorted
		if j >= n {
			return
		}

		// elements must be compared; nums1 item is smaller then our merge step is complete
		if nums1[s] <= nums2[j] {
			s++
			i++
			continue
		}

		// nums2 item is smaller; shift nums1 starting at s; insert nums2 item
		shift(nums1, s)
		nums1[s] = nums2[j]
		s++
		j++
	}
}

func shift(nums1 []int, start int) {
	for i := len(nums1) - 2; i >= start; i-- {
		nums1[i+1] = nums1[i]
	}
}
