package queue

import (
	"fmt"
	"strings"
)

type Queue[T any] struct {
	s []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.s = append(q.s, item)
}

func (q *Queue[T]) Peek() *T {
	if len(q.s) < 1 {
		return nil
	}

	return &q.s[0]
}

func (q *Queue[T]) Dequeue() *T {
	if len(q.s) < 1 {
		return nil
	}
	// store first item in queue to return
	item := q.s[0]
	// remove first item from queue
	q.s = q.s[1:]
	// return stored item
	return &item
}

func (q *Queue[T]) String() string {
	var str []string
	for _, val := range q.s {
		str = append(str, fmt.Sprintf("%v", val))
	}

	return "[" + strings.Join(str, " ") + "]"
}

func (q *Queue[T]) Size() int {
	return len(q.s)
}
