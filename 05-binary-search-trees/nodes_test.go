package bsts

import (
	"testing"
)

var n4 = node{value: 4}
var n7 = node{value: 7}
var n13 = node{value: 13}
var n14 = node{value: 14, right: &n13}
var n6 = node{value: 6, left: &n7, right: &n4}
var n1 = node{value: 1}
var n3 = node{value: 3, left: &n6, right: &n1}
var n10 = node{value: 10, left: &n14}
var n8 = node{value: 8, left: &n10, right: &n3}
var tree = Tree{root: &n8}

func setUpParents() {
	n10.parent = &n8
	n3.parent = &n8
	n14.parent = &n10
	n13.parent = &n14
	n1.parent = &n3
	n6.parent = &n3
	n7.parent = &n6
	n4.parent = &n6
}

//          ____ 8 ____
//        /             \
//       10           __ 3 __
//      /            /       \
//    14            6         1
//      \         /   \
//       13      7     4

func Test_minimum(t *testing.T) {
	want := Tree{&n1}
	got, err := tree.Mininum()

	if err != nil {
		t.Errorf("Error while finding minimum: %v\n", err)
	}

	if !got.IsEqualTo(want) {
		t.Errorf("Wanted %v but got %v", want, got)
	}

}

func Test_maximum(t *testing.T) {
	want := Tree{&n14}
	got, err := tree.Maximum()

	if err != nil {
		t.Errorf("Error while finding minimum: %v\n", err)
	}

	if !got.IsEqualTo(want) {
		t.Errorf("Wanted %v but got %v", want, got)
	}

}

type TestCaseNodes struct {
	input *node
	want  *node
}

func Test_successor(t *testing.T) {
	setUpParents()

	cases := []TestCaseNodes{
		{input: &n1, want: &n3},
		{input: &n3, want: &n4},
		{input: &n4, want: &n6},
		{input: &n6, want: &n7},
		{input: &n7, want: &n8},
		{input: &n8, want: &n10},
		{input: &n10, want: &n13},
		{input: &n13, want: &n14},
		{input: &n14, want: &n14},
	}

	for _, c := range cases {
		got := c.input.successor()

		if !got.isEqualTo(c.want) {
			t.Errorf("Wanted %v but got %v", &c.want, got)
		}
	}
}

func Test_predecessor(t *testing.T) {
	setUpParents()

	cases := []TestCaseNodes{
		{input: &n4, want: &n3},
		{input: &n6, want: &n4},
		{input: &n7, want: &n6},
		{input: &n8, want: &n7},
		{input: &n10, want: &n8},
		{input: &n13, want: &n10},
		{input: &n14, want: &n13},
		{input: &n3, want: &n1},
		{input: &n1, want: &n1},
	}

	for _, c := range cases {
		got := c.input.predecessor()

		if !got.isEqualTo(c.want) {
			t.Errorf("Wanted %v but got %v", c.want, got)
		}
	}
}

type DeleteCase struct {
	input  *Tree
	delete int
	want   *Tree
}

func Test_delete(t *testing.T) {
	// Case 1: delete a terminal node.
	i1 := BinaryTreeFromList([]int{8, 10, 3, 14, -1, 6, 1, -1, 13, -1, -1, 7, 4})
	w1 := BinaryTreeFromList([]int{8, 10, 3, 14, -1, 6, 1, -1, -1, -1, -1, 7, 4})

	// Case 2: no left
	i2 := BinaryTreeFromList([]int{8, 10, 3, 14, -1, 6, 1, -1, 13, -1, -1, 7, 4})
	w2 := BinaryTreeFromList([]int{8, 10, 3, 13, -1, 6, 1, -1, -1, -1, -1, 7, 4})

	// Case 2: no right
	i3 := BinaryTreeFromList([]int{8, 10, 3, 14, -1, 6, 1, -1, 13, -1, -1, 7, 4})
	w3 := BinaryTreeFromList([]int{8, 14, 3, -1, 13, 6, 1, -1, -1, -1, -1, 7, 4})

	// Case 3: both
	i4 := BinaryTreeFromList([]int{8, 10, 3, 14, -1, 6, 1, -1, 13, -1, -1, 7, 4})
	w4 := BinaryTreeFromList([]int{8, 10, 4, 14, -1, 6, 1, -1, 13, -1, -1, 7})

	cases := []DeleteCase{
		{input: &i1, delete: 13, want: &w1},
		{input: &i2, delete: 14, want: &w2},
		{input: &i3, delete: 10, want: &w3},
		{input: &i4, delete: 3, want: &w4},
	}

	for _, testCase := range cases {
		err := testCase.input.DeleteValue(testCase.delete)
		if err != nil {
			t.Errorf("error while deleting: %s", err)
		}

		if !testCase.input.IsEqualTo(*testCase.want) {
			t.Errorf("wanted %s but got %s", testCase.want, testCase.input)
		}
	}
}
