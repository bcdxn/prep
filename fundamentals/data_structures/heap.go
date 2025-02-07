package ds

import (
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

func (h *MinHeap[T]) String() string {
	return fmt.Sprintf("%v", h.h)
}
