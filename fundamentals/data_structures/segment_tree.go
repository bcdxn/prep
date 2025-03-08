package ds

// SegmentTree is a data structure that can be used to find the sum between various indexes in an
// array in O(log(N)) time. It can be built in O(N) time and updated in O(log(N)) time
type SegmentTree struct {
	root *segmentTreeNode
}

type segmentTreeNode struct {
	leftBoundary  int
	rightBoundary int
	rangeSum      int
	leftChild     *segmentTreeNode
	rightChild    *segmentTreeNode
}

func NewSegmentTree(nums []int) SegmentTree {
	t := SegmentTree{}
	t.root = build(nums, 0, len(nums)-1)

	return t
}

// build is a recursive helper function to construct a new SegmentTree from a given array
func build(nums []int, l, r int) *segmentTreeNode {
	if l == r {
		return &segmentTreeNode{
			leftBoundary:  l,
			rightBoundary: r,
			rangeSum:      nums[l],
		}
	}

	mid := (l + r) / 2

	root := &segmentTreeNode{
		leftBoundary:  l,
		rightBoundary: r,
	}

	root.leftChild = build(nums, l, mid)
	root.rightChild = build(nums, mid+1, r)
	root.rangeSum = root.leftChild.rangeSum + root.rightChild.rangeSum
	return root
}

func (t *SegmentTree) Update(index, val int) {
	update(t.root, index, val)
}

func update(node *segmentTreeNode, index, val int) {
	if node.leftBoundary == node.rightBoundary {
		node.rangeSum = val
		return
	}

	mid := (node.leftBoundary + node.rightBoundary) / 2

	if index > mid {
		update(node.rightChild, index, val)
	} else {
		update(node.leftChild, index, val)
	}

	node.rangeSum = node.leftChild.rangeSum + node.rightChild.rangeSum
}

// rangeSumQuery calculates the sum of the values between the given left and right indexes in the
// array in O(log(N)) time.
func (t *SegmentTree) RangeSumQuery(l, r int) int {
	return rangeSumQuery(t.root, l, r)
}

// rangeSumQuery is a recursive helper function to calculate the sum between given the indexes.
func rangeSumQuery(root *segmentTreeNode, l, r int) int {
	if root.leftBoundary == l && root.rightBoundary == r {
		return root.rangeSum
	}

	mid := (root.leftBoundary + root.rightBoundary) / 2

	if l > mid {
		return rangeSumQuery(root.rightChild, l, r)
	} else if r <= mid {
		return rangeSumQuery(root.leftChild, l, r)
	} else {
		return rangeSumQuery(root.leftChild, l, mid) +
			rangeSumQuery(root.rightChild, mid+1, r)
	}
}
