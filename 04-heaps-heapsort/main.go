package main

import "fmt"

func main() {
	xs := []int{5, 12, 2, 8, 0, 10, 11, 4, 9, 6, 1, 7, 3}
	// xs := []int{3, 2, 1}

	BuildMinHeap(xs)

	fmt.Println(xs)

	HeapMinInsert(&xs, -1)

	fmt.Println(xs)

}

//               0
//           1          2
//        4     5     3   11
//      8   9  6 12  7 10
