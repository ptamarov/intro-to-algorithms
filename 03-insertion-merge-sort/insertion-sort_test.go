package main

import "testing"

func Test_insertion(t *testing.T) {
	got := []int{7, 6, 3, 2, 7, 8, 54, 3, 2, 13, 2}
	InsertionSort(got)
	want := []int{2, 2, 2, 3, 3, 6, 7, 7, 8, 13, 54}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("Wanted %d at index %d but got %d", want[i], i, got[i])
		}
	}
}
