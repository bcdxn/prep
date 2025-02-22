package main

import "fmt"

func main() {
	nums1 := []int{4, -1, 2, -7, 3, 4}
	nums2 := []int{4, -1, 2, -7}

	sum, l, r := maxSum(nums1)
	fmt.Printf("sum %d, l %d, r %d\n", sum, l, r)
	sum, l, r = maxSum(nums2)
	fmt.Printf("sum %d, l %d, r %d\n", sum, l, r)
}

// maxSum returns the max sum and the start and end indices of the subarray
func maxSum(nums []int) (int, int, int) {
	if len(nums) < 1 {
		return 0, -1, -1
	}

	maxSum := nums[0]
	maxL := 0
	maxR := 0
	currSum := 0
	currL := 0
	currR := 0

	for currR < len(nums) {
		// we gain nothing from adding a negative number to the array;
		// so we can move the left pointer in, past the negative subgraph
		if currSum < 0 {
			currSum = 0
			currL = currR
		}

		currSum += nums[currR]

		if currSum > maxSum {
			maxSum = currSum
			maxL = currL
			maxR = currR
		}

		currR += 1
	}

	return maxSum, maxL, maxR
}
