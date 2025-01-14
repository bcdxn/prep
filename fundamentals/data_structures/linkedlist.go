package ds

import (
	"errors"
	"fmt"
	"strings"
)

type linkedListNode[T comparable] struct {
	Item T
	Next *linkedListNode[T]
}

type LinkedList[T comparable] struct {
	Head *linkedListNode[T]
	Tail *linkedListNode[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Insert(item T) {
	node := &linkedListNode[T]{
		Item: item,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Tail.Next = node
		l.Tail = node
	}

}

func (l *LinkedList[T]) Remove(item T) (T, error) {
	var prev *linkedListNode[T]
	node := l.Head

	for node != nil {
		if node.Item == item {
			item := node.Item

			if prev != nil {
				prev.Next = node.Next
			} else {
				l.Head = node.Next
			}

			if node.Next == nil {
				l.Tail = node
			}

			return item, nil
		}
		prev = node
		node = node.Next
	}

	var notFound T
	return notFound, errors.New("item not found in linked list")
}

func (l *LinkedList[T]) Find(item T) (T, error) {
	node := l.Head
	for node != nil {
		if node.Item == item {
			return node.Item, nil
		}
		node = node.Next
	}

	var notFound T
	return notFound, errors.New("item not found in linked list")
}

func (l *LinkedList[T]) String() string {
	s := make([]string, 0)

	node := l.Head
	for node != nil {
		s = append(s, fmt.Sprintf("%v", node.Item))
		node = node.Next
	}

	return "[" + strings.Join(s, " ") + "]"
}
