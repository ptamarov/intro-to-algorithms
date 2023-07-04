package bsts

import (
	"fmt"
)

type node struct {
	value  int
	size   int
	left   *node
	right  *node
	parent *node
}

func (t *node) isEqualTo(s *node) bool {
	if t == nil {
		return s == nil
	}

	if s == nil {
		return false
	}

	if t.value != s.value {
		return false
	}

	return t.left.isEqualTo(s.left) && t.right.isEqualTo(s.right)
}

func (t *node) String() string {
	if t == nil {
		return ""
	}
	out := ""

	r := t
	if r.left != nil {
		out += r.left.String() + " "
	}
	out += fmt.Sprint(r.value)
	if r.right != nil {
		out += " " + r.right.String()
	}
	return "(" + out + ")"
}

func (t *node) mininum() *node {
	if t == nil {
		return nil
	}

	currentNode := t
	for currentNode.right != nil {
		currentNode = currentNode.right
	}
	return currentNode
}

func (t *node) maximum() *node {
	if t == nil {
		return nil
	}

	currentNode := t
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func (t *node) walk() []int {
	out := []int{}
	toVisit := []*node{}
	currentNode := t

	for {
		if currentNode != nil {
			// Go to leftmost node,
			toVisit = append(toVisit, currentNode)
			currentNode = currentNode.left
		} else if currentNode == nil && len(toVisit) != 0 {
			// Node has no left, so read value by popping from stack.
			currentNode = toVisit[len(toVisit)-1]
			toVisit = toVisit[:len(toVisit)-1]
			out = append(out, currentNode.value)
			// Go to right node (if also nil, will pop from stack again.)
			currentNode = currentNode.right
		} else {
			return out
		}
	}
}

func (t *node) search(value int) *node {
	for t != nil {
		v := t.value
		switch {
		case v < value:
			t = t.left
		case v > value:
			t = t.right
		default:
			return t
		}
	}
	return nil
}

func (t *node) insertNode(s *node) error {
	// Nothing to do if node to be inserted is nil.
	if s == nil {
		return nil
	}
	// Cannot insert into a nil tree.
	if t == nil {
		return OperationErr{"cannot insert node into nil tree"}
	}

	currentNode := t
	value := s.value

outer:
	for {
		if currentNode.value < value {
			if currentNode.left == nil {
				currentNode.left = s
				return nil
			} else {
				currentNode = currentNode.left
				continue outer
			}
		}
		if currentNode.value > value {
			if currentNode.right == nil {
				currentNode.right = s
				return nil
			} else {
				currentNode = currentNode.right
				continue outer
			}
		} else {
			return OperationErr{fmt.Sprintf("value %d already exists in tree", value)}
		}
	}
}

func (t *node) successor() *node {
	if t == nil {
		return nil
	}

	if t.left != nil {
		// Get the smallest value of all values larger than t.value.
		return t.left.mininum()
	}

	// If t has no left child, then go up and left as much as possible

	c := t

	for c.parent != nil && c.parent.left == c {
		c = c.parent
	}

	// when this  loop stops, either c.parent is nil, or
	// c.parent is not nil, but c.parent.right = c
	// --> c.parent is the successor.

	if c.parent == nil {
		return t // if we reached the root it means we were at the largest element of the tree
	} else {
		return c.parent
	}
}

func (t *node) predecessor() *node {
	if t == nil {
		return nil
	}

	if t.right != nil {
		// Get the smallest value of all values larger than t.Value.
		return t.right.maximum()
	}

	c := t

	// Go up right, this will produce values that are smaller.
	// If at some point p is nil, node must have been the maximum of the tree (leftmost node)
	// Else, can go up left, which means that tree grows so this is the successor.

	for c.parent != nil && c == c.parent.right {
		c = c.parent
	}

	if c.parent == nil {
		return t
	} else {
		return c.parent
	}
}
