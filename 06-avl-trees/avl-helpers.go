package avltrees

func (t *avlnode) isEqualTo(s *avlnode) bool {
	if t == nil {
		return s == nil
	}

	if s == nil {
		return false
	}

	if t.value != s.value {
		return false
	}

	if t.height != s.height {
		return false
	}

	return t.left.isEqualTo(s.left) && t.right.isEqualTo(s.right)
}

// getHeight returns the height of the node stored in pointer or -1 if the pointer is nil
func (n *avlnode) getHeight() int {
	if n == nil {
		return -1
	} else {
		return n.height
	}
}

// isLeftHeavy checks if the left child of a node has height two or more than its right child
func (y *avlnode) isLeftHeavy() bool {
	return y.left.getHeight()-y.right.getHeight() >= 2
}

// isRightHeavy checks if the right child of a node has height two or more than its left child
func (y *avlnode) isRightHeavy() bool {
	return y.right.getHeight()-y.left.getHeight() >= 2
}

// updateHeight recomputes the height of a node according to the formula h = max(hLeft, hRight) + 1
func (n *avlnode) updateHeight() error {
	if n == nil {
		return OperationErr{"cannot update height of nil node"}
	}
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	return nil
}
