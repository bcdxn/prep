package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 2, 2}
	dupe := findDuplicate(nums)
	fmt.Println("duplicate: ", dupe) // 2
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for true {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := nums[0]

	for slow != slow2 {
		slow = nums[slow]
		slow2 = nums[slow2]
	}

	return slow
}
