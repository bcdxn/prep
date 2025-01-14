package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	ll := ds.NewLinkedList[int]()

	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	fmt.Println("ll: ", ll.String())

	found, _ := ll.Find(4)
	fmt.Println("found: ", found)

	ll.Remove(1)
	fmt.Println("ll: ", ll.String())
	ll.Remove(5)
	fmt.Println("ll: ", ll.String())
	ll.Remove(3)
	fmt.Println("ll: ", ll.String())
}
