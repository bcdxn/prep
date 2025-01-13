package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	s := ds.NewStack[int]()
	s.Push(1)
	fmt.Println(s.String())
	s.Push(2)
	fmt.Println(s.String())
	s.Push(3)
	fmt.Println(s.String())
	s.Push(4)
	fmt.Println(s.String())
	s.Push(5)
	fmt.Println(s.String())
	val, _ := s.Pop()
	fmt.Println("popped: ", val)
	fmt.Println(s.String())
	val, _ = s.Pop()
	fmt.Println("popped: ", val)
	fmt.Println(s.String())
	val, _ = s.Pop()
	fmt.Println("popped: ", val)
	fmt.Println(s.String())
}
