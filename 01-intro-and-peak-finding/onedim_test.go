package main

import "testing"

type TestCase struct {
	array []int
	want  int
}

func Test_trivialArrays(t *testing.T) {
	tests := []TestCase{
		{array: []int{}, want: -1},
		{array: []int{10}, want: 0},
		{array: []int{1, 2}, want: 1},
		{array: []int{2, 2}, want: 0},
		{array: []int{2, 1}, want: 0},
	}

	for _, test := range tests {
		got := OneDimPeakFast(test.array)
		want := test.want
		if want != got {
			t.Errorf("Wanted %d but got %d when array is empty.", want, got)
		}
	}
}
