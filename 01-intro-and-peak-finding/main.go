package main

import "fmt"

func main() {
	xs := []int{4, 6, 4, 3, 2, 5, 6, 4, 4, 2, 2, 3, 1, 4}
	// expect 1
	fmt.Println(OneDimPeakFast(xs))

	xxs := [][]int{{10, 8, 10, 10}, {14, 13, 12, 11}, {15, 9, 11, 21}, {16, 17, 19, 20}}
	// expect 2,3
	fmt.Println(TwoDimensionalPeak(xxs))

}
