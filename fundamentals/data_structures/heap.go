package ds

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	s []T
}

func (h *MinHeap[T]) Push(item T) {
	// Make backing slice 0-indexed to make the pointer arithmetic simpler
	if len(h.s) < 2 {
		var zero T
		h.s = []T{zero}
	}
	// Maintain 'complete tree' property of heap
	h.s = append(h.s, item)

	// Maintain 'order' property of heap by percolating up
	i := len(h.s) - 1
	for i > 0 {
		p := parent(i)
		if h.s[i] < h.s[p] {
			h.s[i], h.s[p] = h.s[p], h.s[i]
			i = p
		} else {
			break
		}
	}
}

func (h *MinHeap[T]) Pop() (T, error) {
	if len(h.s) < 2 {
		var zero T
		return zero, errors.New("cannot perform 'Pop' on empty heap")
	}
	item := h.s[1]

	if len(h.s) < 3 {
		var zero T
		h.s = []T{zero}
		return item, nil
	}

	// move 'bottom' item in heap to the root
	h.s[1] = h.s[len(h.s)-1]
	h.s = h.s[:len(h.s)-1]

	// Maintain 'complete tree' property of heap by perocolating new root down
	i := 1
	for left(i) < len(h.s) {
		l := left(i)
		r := right(i)
		if r < len(h.s) && h.s[r] < h.s[l] && h.s[i] > h.s[r] {
			// perclate down to right child
			h.s[i], h.s[r] = h.s[r], h.s[i]
			i = r
		} else if h.s[l] < h.s[i] {
			// percolate down to left child
			h.s[i], h.s[l] = h.s[l], h.s[i]
			i = l
		} else {
			// percolated to correct position
			break
		}
	}

	return item, nil
}

func (h *MinHeap[T]) String() string {
	return fmt.Sprintf("%v", h.s[1:])
}

func Heapify[T constraints.Ordered](items []T) MinHeap[T] {
	// shift the first element to the end of the backing array to get 1-based index
	var zero T
	s := append(items, items[0])
	s[0] = zero

	h := MinHeap[T]{s}
	// half of the tree will not have children so we can skip those as they will, by definition,
	// follow the heap order property
	i := (len(s) - 1) / 2
	for i > 0 {
		j := i
		// for each node in the heap, percolate down until it satisifies the heap order property
		for left(j) < len(h.s) {
			l := left(j)
			r := right(j)
			if r < len(h.s) && h.s[r] < h.s[l] && h.s[j] > h.s[r] {
				h.s[j], h.s[r] = h.s[r], h.s[j]
				j = r
			} else if h.s[j] > h.s[l] {
				h.s[l], h.s[j] = h.s[j], h.s[l]
				j = l
			} else {
				break
			}
		}
		i--
	}

	return h
}

func left(i int) int {
	return i * 2
}

func right(i int) int {
	return i*2 + 1
}

func parent(i int) int {
	return i / 2
}
