package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	h := ds.MinHeap[int]{}
	h.Push(14)
	h.Push(19)
	h.Push(16)
	h.Push(21)
	h.Push(26)
	h.Push(19)
	h.Push(68)
	h.Push(65)
	h.Push(30)

	fmt.Println(h.String()) // [14 19 16 21 26 19 68 65 30]
	h.Push(17)
	fmt.Println(h.String()) // [14 17 16 21 19 19 68 65 30 26]
	val, _ := h.Pop()
	fmt.Println("popped: ", val)
	fmt.Println(h.String()) // [16 17 19 21 19 26 68 65 30]
	val, _ = h.Pop()
	fmt.Println("popped: ", val)
	fmt.Println(h.String()) // [17 19 19 21 30 26 68 65]

	items := []int{60, 50, 80, 40, 30, 10, 70, 20, 90}
	h1 := ds.Heapify(items)
	fmt.Println("heapified: ", h1.String()) // [10 30 20 50 80 70 40 90 60]
}
