package ds

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type BSTNode[T constraints.Ordered] struct {
	Item  T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

type BST[T constraints.Ordered] struct {
	root *BSTNode[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

func (t *BST[T]) StringInOrder() string {
	s := make([]string, 0)
	return "[ " + strings.Join(bstPrintInOrder(t.root, s), " ") + " ]"
}

func (t *BST[T]) Insert(item T) error {
	r, err := bstInsert(t.root, item)
	if err != nil {
		return err
	}

	t.root = r
	return nil
}

func (t *BST[T]) Remove(item T) error {
	if t.root == nil {
		return errors.New("cannot perform Remove on an empty tree")
	}
	t.root = bstRemove(t.root, item)
	return nil
}

func (t *BST[T]) Find(item T) (T, error) {
	return bstFind(t.root, item)
}

func (t *BST[T]) FindMin() (T, error) {
	n := bstFindMin(t.root)
	if n != nil {
		return n.Item, nil
	}

	var item T
	return item, errors.New("cannot perform FindMin on empty tree")
}

func (t *BST[T]) FindMax() (T, error) {
	n := bstFindMax(t.root)
	if n != nil {
		return n.Item, nil
	}

	var item T
	return item, errors.New("cannot perform FindMax on empty tree")
}

/* private helpers
------------------------------------------------------------------------------------------------- */

func bstPrintInOrder[T constraints.Ordered](n *BSTNode[T], s []string) []string {
	if n == nil {
		return s
	}

	if n.Left != nil {
		s = bstPrintInOrder(n.Left, s)
	}

	s = append(s, fmt.Sprintf("%v", n.Item))

	if n.Right != nil {
		s = bstPrintInOrder(n.Right, s)
	}

	return s
}

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
