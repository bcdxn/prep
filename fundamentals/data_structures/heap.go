package ds

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	h []T
}

// Add an item to the heap and maintain heap properties.
func (h *MinHeap[T]) Push(item T) {
	if len(h.h) < 2 {
		var zero T
		h.h = []T{zero, item}
		return
	}

	// add item to the end to maintain complete property
	h.h = append(h.h, item)
	// percolate up
	i := len(h.h) - 1
	for {
		if h.h[i] < h.h[parent(i)] {
			h.h[i], h.h[parent(i)] = h.h[parent(i)], h.h[i]
			i = parent(i)
		} else {
			break
		}
	}
}

// Remove the item from the top of the heap and maintain heap properties.
func (h *MinHeap[T]) Pop() (T, error) {
	if len(h.h) < 2 {
		var zero T
		return zero, errors.New("cannot perform pop on empty heap")
	}
	// take current top of heap
	item := h.h[1]
	// copy last item in heap to top and truncate last item
	h.h[1] = h.h[len(h.h)-1]
	h.h = h.h[:len(h.h)-1]
	// percolate down
	i := 1
	for left(i) < len(h.h) {
		// has a right child, right child < left child, item > right child
		if right(i) < len(h.h) && h.h[right(i)] < h.h[left(i)] && h.h[i] > h.h[right(i)] {
			h.h[i], h.h[right(i)] = h.h[right(i)], h.h[i]
			i = right(i)
		} else if h.h[i] > h.h[left(i)] {
			h.h[i], h.h[left(i)] = h.h[left(i)], h.h[i]
			i = left(i)
		} else {
			break
		}
	}
	return item, nil
}

// Return a string representation of the heap.
func (h *MinHeap[T]) String() string {
	return fmt.Sprintf("%v", h.h)
}

// get the index into the backing slice of the parent of a given node index.
func parent(i int) int {
	return i / 2
}

// get index into the backing slice of the left child of a given node index.
func left(i int) int {
	return i * 2
}

// get the index into the backing slice of the right child of a given node index.
func right(i int) int {
	return i*2 + 1
}
