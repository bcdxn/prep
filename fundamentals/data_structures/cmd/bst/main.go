package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	t := ds.NewBST[int]()

	t.Insert(5)
	fmt.Println(t.StringInOrder())
	t.Insert(3)
	fmt.Println(t.StringInOrder())
	t.Insert(7)
	fmt.Println(t.StringInOrder())
	t.Insert(1)
	fmt.Println(t.StringInOrder())
	t.Insert(2)
	fmt.Println(t.StringInOrder())

	_, err := t.Find(7)
	if err != nil {
		fmt.Println("did not find 7 in tree")
	} else {
		fmt.Println("found 7 in tree")
	}

	_, err = t.Find(-1)
	if err != nil {
		fmt.Println("did not find -1 in tree")
	} else {
		fmt.Println("found 7 in tree")
	}

	min, _ := t.FindMin()
	fmt.Println("min:", min)
	max, _ := t.FindMax()
	fmt.Println("max:", max)

	t.Remove(5)
	fmt.Println("in order: ", t.StringInOrder())
	fmt.Println("pre order: ", t.StringPreOrder())
	fmt.Println("post order: ", t.StringPostOrder())
	fmt.Println("in order:")
	t.PrintBFS()
}
