package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	q := NewQueue[int]()

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

type Queue[T any] struct {
	s []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		s: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.s = append(q.s, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var item T
	if len(q.s) < 1 {
		return item, errors.New("queue is empty")
	}

	item = q.s[0]
	q.s = q.s[1:]

	return item, nil
}

func (q *Queue[T]) Peek() T {
	return q.s[0]
}

func (q *Queue[T]) Size() int {
	return len(q.s)
}

func (q *Queue[T]) String() string {
	s := make([]string, 0, len(q.s))

	for _, val := range q.s {
		s = append(s, fmt.Sprintf("%v", val))
	}

	return "[ " + strings.Join(s, " ") + " ]"
}
