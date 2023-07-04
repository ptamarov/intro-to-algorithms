package main

func InsertionSort(xs []int) {
	for i := 1; i < len(xs); i++ {
		// Insert xs[i] into xs[i] with a O(i)
		NaiveInsertionAtKey(xs, i)
	}
}

func NaiveInsertionAtKey(xs []int, i int) {
	key := i
	for j := i - 1; j > -1; j-- {
		if xs[key] >= xs[j] {
			return
		} else {
			xs[j], xs[key] = xs[key], xs[j]
			key = j
		}
	}
}
