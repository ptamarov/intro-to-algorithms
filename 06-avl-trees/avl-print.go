package avltrees

import (
	"fmt"
	"math"
	"strings"
)

func (node *avlnode) label() string {
	return fmt.Sprintf("%d[%d]", node.value, node.height)
}

func (tree AVLTree) String() string {
	var output strings.Builder
	if tree.root == nil {
		output.WriteString("Empty tree.")
	} else {
		lines, _ := tree.printNode(tree.root)
		for _, line := range lines {
			output.WriteString(line + "\n")
		}
	}
	return output.String()
}

func (tree *AVLTree) printNode(node *avlnode) ([]string, int) {
	var output []string
	rootLabelMidpoint := 0

	if node == nil {
		// Do nothing
	} else if node.left == nil && node.right == nil {
		output = append(output, node.label())
		rootLabelMidpoint = len(node.label()) / 2
	} else {

		label := node.label()

		// Print left subtree
		leftLines, leftLabelMidpoint := tree.printNode(node.left)

		// Print right subtree
		rightLines, rightLabelMidpoint := tree.printNode(node.right)

		// Compute width
		leftWidth := 0

		if len(leftLines) > 0 {
			leftWidth = len(leftLines[0])
		}
		rightWidth := 0
		if len(rightLines) > 0 {
			rightWidth = len(rightLines[0])
		}
		width := int(math.Max(float64(len(label)), float64(leftWidth+rightWidth+1)))

		// Compute label line
		labelLine, labelStartPos := centerBetween(label, width, leftLabelMidpoint, leftWidth+1+rightLabelMidpoint, '_', '_')

		// Compute label midpoint
		rootLabelMidpoint = labelStartPos + len(label)/2

		// Print label
		output = append(output, labelLine)

		// Print edges
		var edgesLine string
		if leftWidth == 0 {
			edgesLine = paddingRight(padder(rightLabelMidpoint+1)+"\\", width)
		} else if rightWidth == 0 {
			edgesLine = paddingRight(padder(leftLabelMidpoint)+"/", width)
		} else {

			left := paddingRight(padder(leftLabelMidpoint)+"/", leftWidth)
			right := padder(rightLabelMidpoint) + "\\"
			edgesLine = paddingRight(left+" "+right, width)
		}
		output = append(output, edgesLine)
		// Append subtrees
		for i := 0; i < int(math.Max(float64(len(leftLines)), float64(len(rightLines)))); i++ {
			left := padder(leftWidth)
			if i < len(leftLines) {
				left = leftLines[i]
			}

			right := padder(rightWidth)
			if i < len(rightLines) {
				right = rightLines[i]
			}

			output = append(output, left+" "+right)
		}
	}

	return output, rootLabelMidpoint
}

func centerBetween(s string, width, left, right int, leftInternalPadding, rightInternalPadding byte) (string, int) {
	startPos := 0
	bestVal := math.MaxInt32

	for curStartPos := 0; curStartPos <= width-len(s); curStartPos++ {
		leftChars := left - curStartPos
		rightChars := curStartPos + len(s) - 1 - right
		val := int(math.Abs(float64(rightChars - leftChars)))
		if val < bestVal {
			startPos = curStartPos
			bestVal = val
		}
	}

	output := strings.Repeat(" ", width)
	for i := 0; i < width; i++ {
		if i < startPos {
			if i > left {
				output = replaceByteAtPosition(output, i, leftInternalPadding)
			} else {
				output = replaceByteAtPosition(output, i, ' ')
			}
		} else if startPos <= i && i < startPos+len(s) {
			output = replaceByteAtPosition(output, i, s[i-startPos])
		} else {
			if i < right {
				output = replaceByteAtPosition(output, i, rightInternalPadding)
			} else {
				output = replaceByteAtPosition(output, i, ' ')
			}
		}
	}

	return output, startPos
}

// replaceByteAtPosition replaces a character of a string at give index by replacement
// does nothing if index is out of bounds
func replaceByteAtPosition(s string, index int, replacement byte) string {
	if index < 0 || index >= len(s) {
		return s
	}
	return s[:index] + string(replacement) + s[index+1:]
}

// paddingRight pads a string with empty spaces to the right so that it has length w.
func paddingRight(s string, w int) string {
	padding := w - len(s)
	spaces := strings.Repeat(" ", padding)
	return s + spaces
}

// padder takes an integer w and return a string with w spaces. Returns no space if w is not positive.
func padder(w int) string {
	if w <= 0 {
		return ""
	}
	return strings.Repeat(" ", w)
}
