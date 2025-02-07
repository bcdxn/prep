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

	fmt.Println(h.String())
	h.Push(17)
	fmt.Println(h.String())
}
