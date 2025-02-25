package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	fmt.Printf("k: %d -- %v\n", k, nums) // 5 -- 1, 2, 2
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k = removeDuplicates(nums)
	fmt.Printf("k: %d -- %v\n", k, nums) // 5 -- 0, 1, 2, 3, 4, 2, 2, 3, 3, 4

}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	k := 1
	unique := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[unique] {
			k++
			unique++
			nums[unique] = nums[i]
		}
	}

	return k
}
