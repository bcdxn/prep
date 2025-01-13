package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	q := ds.NewQueue[int]()

	q.Enqueue(1)
	fmt.Println(q.String())
	q.Enqueue(2)
	fmt.Println(q.String())
	q.Enqueue(3)
	fmt.Println(q.String())
	q.Enqueue(4)
	fmt.Println(q.String())
	q.Enqueue(5)
	fmt.Println(q.String())
	q.Dequeue()
	fmt.Println(q.String())
	q.Dequeue()
	fmt.Println(q.String())
	q.Dequeue()
	fmt.Println(q.String())
	q.Dequeue()
	fmt.Println(q.String())
	q.Dequeue()
	fmt.Println(q.String())
}
