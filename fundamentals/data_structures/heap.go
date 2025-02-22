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
	return fmt.Sprintf("%v", h.s)
}

// func Heapify[T constraints.Ordered](s []T) MinHeap[T] {

// }

func left(i int) int {
	return i * 2
}

func right(i int) int {
	return i*2 + 1
}

func parent(i int) int {
	return i / 2
}
