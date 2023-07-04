package avltrees

import (
	"fmt"
	"testing"
)

func Test_rotateLeft(t *testing.T) {
	//                  x1
	//          y3              C0
	//      A4       B2

	A := avlnode{value: 4}
	B := avlnode{value: 2}
	C := avlnode{value: 0}
	y := avlnode{value: 3, left: &A, right: &B}
	x := avlnode{value: 1, left: &y, right: &C}
	A.parent = &y
	B.parent = &y
	C.parent = &x
	y.parent = &x
	tree := AVLTree{root: &x}
	t.Log("Rotating the following tree at root:\n" + fmt.Sprint(tree))

	tree.rotateRight(&x)
	//                  a3
	//          X4               b1
	//                      Y2       Z0

	X := avlnode{value: 4}
	Y := avlnode{value: 2}
	Z := avlnode{value: 0}
	b := avlnode{value: 1, left: &Y, right: &Z, height: 1}
	a := avlnode{value: 3, left: &X, right: &b, height: 2}
	X.parent = &a
	Y.parent = &b
	Z.parent = &b
	b.parent = &a

	t.Log("Results in:\n" + fmt.Sprint(tree))

	want := AVLTree{root: &a}

	if !want.IsEqualTo(tree) {
		t.Error("Wanted", fmt.Sprint(want), "but got", fmt.Sprint(tree))

	}

}

func Test_lectureExample1(t *testing.T) {
	// Not a real test...
	var n26 = avlnode{value: 26}
	var n11 = avlnode{value: 11}
	var n50 = avlnode{value: 50}
	var n29 = avlnode{value: 29, right: &n26, height: 1}
	var n65 = avlnode{value: 65, right: &n50, height: 2}
	var n20 = avlnode{value: 20, right: &n11, left: &n29, height: 2}
	var n41 = avlnode{value: 41, right: &n20, left: &n65, height: 3}
	var tree = AVLTree{root: &n41}

	n65.parent = &n41
	n20.parent = &n41
	n29.parent = &n20
	n11.parent = &n20
	n50.parent = &n65
	n26.parent = &n29

	tree.InsertValue(23)

	t.Log("\n" + fmt.Sprint(tree))

}

func Test_lectureExample2(t *testing.T) {
	// Not a real test...
	var n26 = avlnode{value: 26}
	var n11 = avlnode{value: 11}
	var n50 = avlnode{value: 50}
	var n29 = avlnode{value: 29, right: &n26, height: 1}
	var n65 = avlnode{value: 65, right: &n50, height: 2}
	var n20 = avlnode{value: 20, right: &n11, left: &n29, height: 2}
	var n41 = avlnode{value: 41, right: &n20, left: &n65, height: 3}
	var tree = AVLTree{root: &n41}

	n65.parent = &n41
	n20.parent = &n41
	n29.parent = &n20
	n11.parent = &n20
	n50.parent = &n65
	n26.parent = &n29

	tree.InsertValue(23)
	tree.InsertValue(55)

	t.Log("\n" + fmt.Sprint(tree))

}
