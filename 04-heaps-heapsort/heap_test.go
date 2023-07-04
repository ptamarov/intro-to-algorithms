package main

import "testing"

func Test_BuildMaxHeap(t *testing.T) {
	xs := []int{5, 12, 2, 8, 0, 10, 11, 4, 9, 6, 1, 7, 3}
	//                  5
	//         12             2
	//      8     0        10   11
	//    4   9  6  1    7   3

	want := []int{12, 9, 11, 8, 6, 10, 2, 4, 5, 0, 1, 7, 3}
	//                  12
	//         9                  11
	//      8      6          10     2
	//    4   5  0  1       7   3

	BuildMaxHeap(xs)

	for i := range want {
		if want[i] != xs[i] {
			t.Errorf("Wanted %d at index %d but got %d", want[i], i, xs[i])
		}
	}
}

func Test_BuildMinHea(t *testing.T) {
	xs := []int{5, 12, 2, 8, 0, 10, 11, 4, 9, 6, 1, 7, 3}
	//                  5
	//         12               2
	//      8     0        10       11
	//    4   9  6  1    7   3

	want := []int{0, 1, 2, 4, 5, 3, 11, 8, 9, 6, 12, 7, 10}
	//                     0
	//         1                   2
	//      4     5         3          11
	//    8   9  6  12    7   10

	BuildMinHeap(xs)

	for i := range want {
		if want[i] != xs[i] {
			t.Errorf("Wanted %d at index %d but got %d", want[i], i, xs[i])
		}
	}
}
