package ds

func NewUnionFind(nodes []int) *UnionFind {
	uf := UnionFind{
		parents: make(map[int]int),
		ranks:   make(map[int]int),
	}

	for _, n := range nodes {
		uf.parents[n] = n
		uf.ranks[n] = 0
	}

	return &uf
}

// Find will return the root ancestor of the given node. For every search we'll compress the path,
// flattening the tree and making subsequent calls to find more efficient.
func (uf *UnionFind) Find(n int) int {
	node := uf.parents[n]

	for uf.parents[node] != node {
		uf.parents[node] = uf.parents[uf.parents[node]]
		node = uf.parents[node]
	}

	return node
}

// Union the nodes within the same tree (at the root parent of each node).
func (uf *UnionFind) Union(a, b int) bool {
	p1 := uf.Find(a) // find parent of A
	p2 := uf.Find(b) // find parent of B

	// These two nodes are already a part of the same set and so they cannot be unioned
	if p1 == p2 {
		return false
	}

	// union by rank (higher rank should be the parent)
	if uf.ranks[p1] > uf.ranks[p2] {
		uf.parents[p2] = p1
	} else if uf.ranks[p1] < uf.ranks[p2] {
		uf.parents[p1] = p2
	} else {
		uf.parents[p2] = p1
		uf.ranks[p2] += 1
	}

	return true
}

type UnionFind struct {
	parents map[int]int // key == node; value == node's parent (which will be compressed)
	ranks   map[int]int
}
