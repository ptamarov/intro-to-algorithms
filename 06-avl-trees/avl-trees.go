package avltrees

import (
	"fmt"
)

type AVLTree struct {
	root *avlnode
}

// nilTree returns a tree whose root is nil.
func nilTree() AVLTree {
	return AVLTree{}
}

// RootValue returns the value stores at the root of a Tree and an error if the root is nil.
func (t AVLTree) RootValue() (int, error) {
	if t.root == nil {
		return -1, OperationErr{"nil tree has no root value"}
	} else {
		return t.root.value, nil
	}
}

// IsEqualTo checks if two trees store the same values according to the in order transversal.
func (t AVLTree) IsEqualTo(s AVLTree) bool {
	return t.root.isEqualTo(s.root)
}

// IsNil checks if the root of a tree is nil.
func (t AVLTree) IsNil() bool {
	return t.root == nil
}

// Minimum returns the tree stored at the node of the tree storing the minimum value, or an error if the tree is nil.
func (t AVLTree) Mininum() (AVLTree, error) {
	out := t.root.mininum()

	if out == nil {
		return nilTree(), OperationErr{"cannot operate on nil tree"}
	} else {
		return AVLTree{out}, nil
	}
}

// Maximum returns the tree stored at the node of the tree storing the minimum value, or an error if the tree is nil.
func (t AVLTree) Maximum() (AVLTree, error) {
	out := t.root.maximum()

	if out == nil {
		return nilTree(), OperationErr{"cannot operate on nil tree"}
	} else {
		return AVLTree{out}, nil
	}
}

// Walk walks the tree and returns an integer slice of its stored values in order.
func (t AVLTree) Walk() []int {
	if t.root == nil {
		return []int{}
	} else {
		return t.root.walk()
	}
}

// Search searches for a node in the tree with a given value, and returns the subtree rooted at a node with
// that value, or an error if the value does not exist in the tree.
func (t AVLTree) Search(value int) (AVLTree, error) {
	s := t.root.search(value)
	if s != nil {
		return AVLTree{s}, nil
	} else {
		return nilTree(), OperationErr{fmt.Sprintf("value %d not found.", value)}
	}
}

func (t *AVLTree) InsertTree(s AVLTree) error {
	if s.root == nil {
		return nil // nothing to do if s is nil (include warning?)
	}
	if t.root == nil {
		return OperationErr{"cannot insert node into nil tree"} // cannot insert into nil tree

	}

	currentNode := t.root
	value := s.root.value

outer:
	for {
		if currentNode.value < value {
			if currentNode.left == nil {
				currentNode.left = s.root
				s.root.parent = currentNode
				return t.rebalanceNode(s.root)
			} else {
				currentNode = currentNode.left
				continue outer
			}
		}
		if currentNode.value > value {
			if currentNode.right == nil {
				currentNode.right = s.root
				s.root.parent = currentNode
				return t.rebalanceNode(s.root)
			} else {
				currentNode = currentNode.right
				continue outer
			}
		} else {
			return OperationErr{fmt.Sprintf("value %d already exists in tree", value)}
		}
	}
}

func (t *AVLTree) InsertValue(value int) error {
	n := AVLTree{root: &avlnode{value: value}}
	return t.InsertTree(n)
}

func (t AVLTree) Successor() (AVLTree, error) {
	s := t.root.successor()
	if s == nil {
		return nilTree(), OperationErr{"nil tree has no successor."}
	} else {
		return AVLTree{s}, nil
	}
}

func (t AVLTree) Predecessor() (AVLTree, error) {
	s := t.root.predecessor()
	if s == nil {
		return nilTree(), OperationErr{"nil tree has no predecessor."}
	} else {
		return AVLTree{s}, nil
	}
}

func (t *AVLTree) DeleteValue(value int) error {
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
		return t.rebalanceNode(node.parent)
	}

	// Case 2: node has right but no left child.
	if node.left == nil && node.right != nil {
		switch {

		case node.parent == nil:
			t.root = node.right
			t.root.parent = nil

		case node.parent.left == node:
			node.parent.left = node.right

		case node.parent.right == node:
			node.parent.right = node.right
		}
		return t.rebalanceNode(node.parent)
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
		return t.rebalanceNode(node.parent)
	}

	// Case 4: node has both a left and a right child.
	// In particular, it must have a successor different from itself.

	nextNode := node.successor()

	// All that is left to do is to is to connect the parent of nextNode
	// to its left child and update the value of original node.
	// nextNode is now inaccessible from available tree pointers.

	if nextNode.parent.left == nextNode {
		nextNode.parent.left = nextNode.left
	} else {
		nextNode.parent.right = nextNode.left
	}

	node.value = nextNode.value
	return t.rebalanceNode(nextNode.parent) // must rebalance tree at parent of successor
}

func (t *AVLTree) rotateRight(node *avlnode) error {
	if t == nil || t.root == nil {
		return OperationErr{"nil pointer or nil root in tree"}
	}
	if node.left == nil {
		return OperationErr{"no left child, cannot rotate right."}
	}

	leftChild := node.left
	b := leftChild.right // store for later, can be nil

	if node.parent == nil {
		leftChild.parent = nil
		t.root = leftChild
	} else {
		leftChild.parent = node.parent
	}

	if node.parent != nil {
		if node.parent.left == node {
			node.parent.left = leftChild
		} else {
			node.parent.right = leftChild
		}
	}
	node.parent = leftChild
	leftChild.right = node

	node.left = b
	if b != nil {
		b.parent = node
	}

	node.updateHeight()
	leftChild.updateHeight()

	return nil
}

func (t *AVLTree) rotateLeft(node *avlnode) error {
	if t == nil || t.root == nil {
		return OperationErr{"nil pointer or nil root in tree"}
	}
	if node.right == nil {
		return OperationErr{"no right child, cannot rotate left."}
	}

	rightChild := node.right
	b := rightChild.left // store for later

	if node.parent == nil {
		rightChild.parent = nil
		t.root = rightChild
	} else {
		rightChild.parent = node.parent
	}

	if node.parent != nil {
		if node.parent.right == node {
			node.parent.right = rightChild
		} else {
			node.parent.left = rightChild
		}
	}

	node.parent = rightChild
	rightChild.left = node

	node.right = b
	if b != nil {
		b.parent = node
	}

	node.updateHeight()
	rightChild.updateHeight()

	return nil
}

func (t *AVLTree) rebalanceNode(node *avlnode) error {
	current := node
	for current != nil {
		err := current.updateHeight() // current is not nil, so no error needs to be checked
		if err != nil {
			return err
		}

		if current.isLeftHeavy() {
			z := current.left
			if z.left.getHeight() >= z.right.getHeight() {
				if err := t.rotateRight(current); err != nil {
					return err
				}
			} else {
				if err := t.rotateLeft(z); err != nil {
					return err
				}
				if err := t.rotateRight(current); err != nil {
					return err
				}
			}
		} else if current.isRightHeavy() {
			z := current.right
			if z.right.getHeight() >= z.left.getHeight() {
				if err := t.rotateLeft(current); err != nil {
					return err
				}
			} else {
				if err := t.rotateRight(z); err != nil {
					return err
				}
				if err := t.rotateLeft(current); err != nil {
					return err
				}
			}
		}
		current = current.parent
	}
	return nil
}
