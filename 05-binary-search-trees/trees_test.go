package bsts

import "testing"

type TestCase struct {
	value int
	rank  int
}

func Test_rank(t *testing.T) {
	n43 := NewNode(43, nil, nil)
	n64 := NewNode(64, nil, nil)
	n83 := NewNode(83, nil, nil)
	n79 := NewNode(79, n64, n83)
	n46 := NewNode(46, n43, nil)
	n49 := NewNode(49, n46, n79)
	tree := Tree{root: n49}

	//         ________49(6)_______
	//        /                    \
	//      46 (2)                  79(3)
	//     /                       /     \
	//  43(1)                    64(1)   83(1)

	tests := []TestCase{
		{value: 20, rank: 0},
		{value: 44, rank: 1},
		{value: 47, rank: 2},
		{value: 49, rank: 3},
		{value: 50, rank: 3},
		{value: 65, rank: 4},
		{value: 78, rank: 4},
		{value: 82, rank: 5},
		{value: 100, rank: 6},
	}

	for _, test := range tests {
		got := tree.Rank(test.value)
		want := test.rank

		if got != want {
			t.Errorf("Wanted rank %d but got %d for value %d", want, got, test.value)
		}
	}
}
