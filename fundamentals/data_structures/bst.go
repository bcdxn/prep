package ds

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type BSTNode[T constraints.Ordered] struct {
	Item  T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

type BST[T constraints.Ordered] struct {
	Root *BSTNode[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

// Insert adds an item to the tree (duplicates are not allowed).
func (t *BST[T]) Insert(item T) error {
	r, err := bstInsert(t.Root, item)
	if err != nil {
		return err
	}

	t.Root = r
	return nil
}

// Remove removes an item from the tree if it exists.
func (t *BST[T]) Remove(item T) error {
	if t.Root == nil {
		return errors.New("cannot perform Remove on an empty tree")
	}
	t.Root = bstRemove(t.Root, item)
	return nil
}

// Find looks up an element in the tree and returns an error if it does not exist.
func (t *BST[T]) Find(item T) (T, error) {
	return bstFind(t.Root, item)
}

// FindMin returns the smallest element in the tree.
func (t *BST[T]) FindMin() (T, error) {
	n := bstFindMin(t.Root)
	if n != nil {
		return n.Item, nil
	}

	var item T
	return item, errors.New("cannot perform FindMin on empty tree")
}

// FindMax  returns the largest element in the tree.
func (t *BST[T]) FindMax() (T, error) {
	n := bstFindMax(t.Root)
	if n != nil {
		return n.Item, nil
	}

	var item T
	return item, errors.New("cannot perform FindMax on empty tree")
}

func (t *BST[T]) StringInOrder() string {
	s := make([]string, 0)
	return "[ " + strings.Join(bstStringInOrder(t.Root, s), " ") + " ]"
}

func (t *BST[T]) StringPreOrder() string {
	s := make([]string, 0)
	return "[ " + strings.Join(bstStringPreOrder(t.Root, s), " ") + " ]"
}

func (t *BST[T]) StringPostOrder() string {
	s := make([]string, 0)
	return "[ " + strings.Join(bstStringPostOrder(t.Root, s), " ") + " ]"
}

func (t *BST[T]) StringPreOrderIterative() string {
	res := []T{}
	stack := []*BSTNode[T]{}

	curr := t.Root

	for curr != nil || len(stack) > 0 {
		if curr == nil {
			curr = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1] // pop
		}
		res = append(res, curr.Item)

		if curr.Right != nil {
			stack = append(stack, curr.Right) // push
		}
		curr = curr.Left
	}

	return fmt.Sprintf("%v", res)
}

func (t *BST[T]) StringInOrderIterative() string {
	res := []T{}
	stack := []*BSTNode[T]{}
	curr := t.Root

	for curr != nil || len(stack) > 0 {
		// traverse left
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
			continue
		}
		// pop item from stack and add it to our in order result slice
		curr = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		res = append(res, curr.Item)
		// traverse right
		curr = curr.Right
	}

	return fmt.Sprintf("%v", res)
}

func (t *BST[T]) PrintBFS() {
	if t.Root == nil {
		fmt.Println("[ ]")
		return
	}

	q := NewQueue[*BSTNode[T]]()
	q.Enqueue(t.Root)

	level := 0
	for q.Size() > 0 {
		str := []string{strconv.Itoa(level) + ": "}
		for range q.Size() {
			n, _ := q.Dequeue()

			if n != nil {
				str = append(str, fmt.Sprintf("%v", n.Item))
				q.Enqueue(n.Left)
				q.Enqueue(n.Right)
			}
		}
		fmt.Println(strings.Join(str, " "))
		level += 1
	}
}

/* private helpers
------------------------------------------------------------------------------------------------- */

func bstInsert[T constraints.Ordered](root *BSTNode[T], item T) (*BSTNode[T], error) {
	if root == nil {
		return &BSTNode[T]{Item: item}, nil
	}

	var node *BSTNode[T]
	var err error
	if item < root.Item {
		node, err = bstInsert(root.Left, item)
		root.Left = node
	} else if item > root.Item {
		node, err = bstInsert(root.Right, item)
		root.Right = node
	} else {
		err = errors.New("duplicates not allowed")
	}

	return root, err
}

func bstRemove[T constraints.Ordered](root *BSTNode[T], item T) *BSTNode[T] {
	if root == nil {
		return root
	}

	if item < root.Item {
		root.Left = bstRemove(root.Left, item)
	} else if item > root.Item {
		root.Right = bstRemove(root.Right, item)
	} else {
		// if the node is a leaf we can simply remove it
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// if the node has only a left child we can return the left child
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		// if the node has only a right child we can return the right child
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		// if the node has two children, we must replace the node with the maximum from the left subtree
		max := bstFindMax(root.Left)
		newRoot := &BSTNode[T]{
			Item:  max.Item,
			Left:  bstRemove(root.Left, max.Item),
			Right: root.Right,
		}
		root = newRoot
	}

	return root
}

func bstFind[T constraints.Ordered](root *BSTNode[T], item T) (T, error) {
	var found T
	if root == nil {
		return found, errors.New("item not found")
	}

	if item < root.Item {
		return bstFind(root.Left, item)
	} else if item > root.Item {
		return bstFind(root.Right, item)
	} else {
		return root.Item, nil
	}
}

func bstFindMin[T constraints.Ordered](root *BSTNode[T]) *BSTNode[T] {
	if root == nil {
		return root
	}

	if root.Left != nil {
		return bstFindMin(root.Left)
	}

	return root
}

func bstFindMax[T constraints.Ordered](root *BSTNode[T]) *BSTNode[T] {
	if root == nil {
		return root
	}

	if root.Right != nil {
		return bstFindMax(root.Right)
	}

	return root
}

func bstStringInOrder[T constraints.Ordered](n *BSTNode[T], s []string) []string {
	if n == nil {
		return s
	}

	s = bstStringInOrder(n.Left, s)
	s = append(s, fmt.Sprintf("%v", n.Item))
	s = bstStringInOrder(n.Right, s)

	return s
}

func bstStringPreOrder[T constraints.Ordered](n *BSTNode[T], s []string) []string {
	if n == nil {
		return s
	}

	s = append(s, fmt.Sprintf("%v", n.Item))
	s = bstStringPreOrder(n.Left, s)
	s = bstStringPreOrder(n.Right, s)

	return s
}

func bstStringPostOrder[T constraints.Ordered](n *BSTNode[T], s []string) []string {
	if n == nil {
		return s
	}

	s = bstStringPostOrder(n.Left, s)
	s = bstStringPostOrder(n.Right, s)
	s = append(s, fmt.Sprintf("%v", n.Item))

	return s
}
