package ds

import (
	"errors"
	"fmt"
	"strings"
)

type Stack[T any] struct {
	s []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		s: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.s = append(s.s, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var item T
	if len(s.s) < 1 {
		return item, errors.New("cannot permform Pop on empty stack")
	}

	last := len(s.s) - 1
	item = s.s[last]
	s.s = s.s[:last]

	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var item T
	if len(s.s) < 1 {
		return item, errors.New("cannot permform Peek on empty stack")
	}

	return s.s[len(s.s)-1], nil
}

func (s *Stack[T]) Height() int {
	return len(s.s)
}

func (s *Stack[T]) String() string {
	str := make([]string, 0, len(s.s))

	for _, val := range s.s {
		str = append(str, fmt.Sprintf("%v", val))
	}

	return "[ " + strings.Join(str, " ") + " ]"
}
