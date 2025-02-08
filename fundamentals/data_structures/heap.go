package ds

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	h []T
}

func (h *MinHeap[T]) Push(item T) {
	if len(h.h) < 2 {
		var zero T
		h.h = []T{zero, item}
		return
	}

	// add element at next index to keep *complete* tree
	h.h = append(h.h, item)
	// percolate up
	i := len(h.h) - 1
	for h.h[i] < h.h[i/2] {
		h.h[i], h.h[i/2] = h.h[i/2], h.h[i]
		i = i / 2
	}
}

func (h *MinHeap[T]) Pop() (T, error) {
	if len(h.h) < 2 {
		var zero T
		return zero, errors.New("cannot pop from empty heap")
	}

	if len(h.h) == 2 {
		item := h.h[1]
		h.h = []T{}
		return item, nil
	}

	item := h.h[1]

	// replace root with last element in heap
	h.h[1] = h.h[len(h.h)-1]
	h.h = h.h[:len(h.h)-1]
	// percolate down
	i := 1
	for 2*i < len(h.h) {
		if 2*i+1 < len(h.h) && h.h[i] > h.h[2*i+1] && h.h[2*i+1] < h.h[2*i] {
			// has right child AND item > right child AND right child < left
			h.h[i], h.h[2*i+1] = h.h[2*i+1], h.h[i]
			i = 2*i + 1
		} else if 2*i < len(h.h) && h.h[i] > h.h[2*i] {
			// item is greater than left child
			h.h[i], h.h[2*i] = h.h[2*i], h.h[i]
			i = 2 * i
		} else {
			break
		}
	}

	return item, nil
}

func (h *MinHeap[T]) String() string {
	return fmt.Sprintf("%v", h.h)
}
