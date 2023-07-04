package avltrees

type avlnode struct {
	left   *avlnode
	right  *avlnode
	parent *avlnode
	value  int
	height int
}

func (t *avlnode) mininum() *avlnode {
	if t == nil {
		return nil
	}

	currentNode := t
	for currentNode.right != nil {
		currentNode = currentNode.right
	}
	return currentNode
}

func (t *avlnode) maximum() *avlnode {
	if t == nil {
		return nil
	}

	currentNode := t
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func (t *avlnode) walk() []int {
	out := []int{}
	toVisit := []*avlnode{}
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

func (t *avlnode) search(value int) *avlnode {
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

func (t *avlnode) successor() *avlnode {
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

func (t *avlnode) predecessor() *avlnode {
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
