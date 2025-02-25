package main

import "fmt"

func main() {
	summer := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println("sum: ", summer.SumRange(0, 2)) // 1
	fmt.Println("sum: ", summer.SumRange(2, 5)) // -1
	fmt.Println("sum: ", summer.SumRange(0, 5)) // -3
}

type NumArray struct {
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{}
	total := 0
	for _, num := range nums {
		total += num
		numArray.prefixSums = append(numArray.prefixSums, total)
	}

	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	rightPrefix := this.prefixSums[right]
	leftPrefix := 0

	if left > 0 {
		leftPrefix = this.prefixSums[left-1]
	}

	return rightPrefix - leftPrefix
}

/**
* Your NumArray object will be instantiated and called as such:
* obj := Constructor(nums);
* param_1 := obj.SumRange(left,right);
 */
