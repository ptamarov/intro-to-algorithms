package bsts

import (
	"fmt"
)

// Tree holds a pointer to its root node.
type Tree struct {
	root *node
}

// nilTree returns a tree whose root is nil.
func nilTree() Tree {
	return Tree{}
}

// RootValue returns the value stores at the root of a Tree and an error if the root is nil.
func (t Tree) RootValue() (int, error) {
	if t.root == nil {
		return -1, OperationErr{"nil tree has no root value"}
	} else {
		return t.root.value, nil
	}
}

// IsEqualTo checks if two trees store the same values according to the in order transversal.
func (t Tree) IsEqualTo(s Tree) bool {
	return t.root.isEqualTo(s.root)
}

// IsNil checks if the root of a tree is nil.
func (t Tree) IsNil() bool {
	return t.root == nil
}

// String prints the tree recursively as (L V R) where L and R are the result of printing the left
// and right children at the root, and V is the value at the root.
func (t Tree) String() string {
	return t.root.String()
}

// Minimum returns the tree stored at the node of the tree storing the minimum value, or an error if the tree is nil.
func (t Tree) Mininum() (Tree, error) {
	out := t.root.mininum()

	if out == nil {
		return nilTree(), OperationErr{"cannot operate on nil tree"}
	} else {
		return Tree{out}, nil
	}
}

// Maximum returns the tree stored at the node of the tree storing the minimum value, or an error if the tree is nil.
func (t Tree) Maximum() (Tree, error) {
	out := t.root.maximum()

	if out == nil {
		return nilTree(), OperationErr{"cannot operate on nil tree"}
	} else {
		return Tree{out}, nil
	}
}

func (t Tree) Walk() []int {
	if t.root == nil {
		return []int{}
	} else {
		return t.root.walk()
	}
}

func (t Tree) Search(value int) (Tree, error) {
	s := t.root.search(value)
	if s != nil {
		return Tree{s}, nil
	} else {
		return nilTree(), OperationErr{fmt.Sprintf("value %d not found.", value)}
	}
}

func (t Tree) InsertTree(s Tree) error {
	return t.root.insertNode(s.root)
}

func (t Tree) InsertValue(value int) error {
	n := Tree{root: &node{value: value}}
	return t.InsertTree(n)
}

func (t Tree) Successor() (Tree, error) {
	s := t.root.successor()
	if s == nil {
		return nilTree(), OperationErr{"nil tree has no successor."}
	} else {
		return Tree{s}, nil
	}
}

func (t Tree) Predecessor() (Tree, error) {
	s := t.root.predecessor()
	if s == nil {
		return nilTree(), OperationErr{"nil tree has no predecessor."}
	} else {
		return Tree{s}, nil
	}
}

func (t *Tree) DeleteValue(value int) error {
	subtree, err := t.Search(value)

	if err != nil {
		return err
	}

	node := subtree.root

	// Case 1: node has no left nor right children.
	if node.left == nil && node.right == nil {
		switch {
		case node.parent == nil:
			return OperationErr{"tree has only one node, cannot delete root!"}
		case node.parent.right == node:
			node.parent.right = nil
		case node.parent.left == node:
			node.parent.left = nil
		}
		return nil
	}

	// Case 2: node has right but no left child.
	if node.left == nil && node.right != nil {
		switch {
		// The node is the root (has no parent)
		case node.parent == nil:
			t.root = node.right
			t.root.parent = nil
		// The node is the left child of its parent
		case node.parent.left == node:
			node.parent.left = node.right
		// The node is the right child of its parent
		case node.parent.right == node:
			node.parent.right = node.right
		}
		return nil
	}

	// Case 3: node has left but no right child (symmetric to Case 2).
	if node.right == nil && node.left != nil {
		switch {
		case node.parent == nil:
			t.root = node.left
			t.root.parent = nil
		case node.parent.left == node:
			node.parent.left = node.left
		case node.parent.right == node:
			node.parent.right = node.left
		}
		return nil
	}

	// Case 4: node has both a left and a right child. In particular, it must have a successor different from itself.

	// 4.1 Find succesor in tree.
	nextNode := node.successor()

	// Successor cannot have a right child
	// --> It would be an elemenet smaller than successor but larger than original node.

	// All that is left to do is to is to connect the parent of nextNode to its left child and update the value of original node.
	// nextNode is now inaccessible from available tree pointers.

	if nextNode.parent.left == nextNode {
		nextNode.parent.left = nextNode.left
	} else {
		nextNode.parent.right = nextNode.left
	}

	node.value = nextNode.value
	return nil
}

func NewNode(value int, left, right *node) *node {
	new := &node{value: value, size: 1}
	if left != nil {
		new.left = left
		new.size += left.size
	}
	if right != nil {
		new.right = right
		new.size += right.size
	}
	return new
}

// Rank computes the number of nodes in a Tree that are smaller
// than or equal to a given target value.
func (t *Tree) Rank(target int) int {
	if t == nil {
		return 0
	}

	counter := 0
	current := t.root

	for current != nil {
		value := current.value
		switch {
		case value <= target:
			// Must always do this in either case (<, =)
			counter++
			if current.left != nil {
				counter += current.left.size
			}

			// If hit, stop
			if value == target {
				return counter
			} else { // else, continue searching
				current = current.right
			}

		case value > target:
			current = current.left
		}
	}
	return counter

}

// TODO: write down insertion algorithm with time window constraint. Can recycle rank.
// Maybe keep track parents in structure too.
