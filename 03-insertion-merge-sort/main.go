package main

import "fmt"

func main() {
	xs := []int{5, 12, 2, 8, 0, 10, 11, 4, 9, 6, 1, 7, 3}

	InsertionSort(xs)

	fmt.Println(xs)

	a := []int{29, 30, 31, 33}
	b := []int{1, 9, 11, 12, 28}

	fmt.Println(MergeArrays(a, b))

	fmt.Println(MergeSort(xs))
}
