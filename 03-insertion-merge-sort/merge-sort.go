package main

func MergeSort(xs []int) []int {
	n := len(xs)

	if n <= 1 {
		return xs // Do nothing if array is empty or if array has one element.
	}

	return MergeArrays(MergeSort(xs[:n/2]), MergeSort(xs[n/2:]))

}
func MergeArrays(xs, ys []int) []int {
	zs := make([]int, len(xs)+len(ys))

	if len(xs) == 0 {
		return ys
	}

	if len(ys) == 0 {
		return xs
	}

	idx, idy, idz := 0, 0, 0 // Set up pointers.

	for ; idx < len(xs) && idy < len(ys); idz++ {
		if xs[idx] <= ys[idy] {
			zs[idz] = xs[idx]
			idx++
		} else {
			zs[idz] = ys[idy]
			idy++
		}
	}

	// Continue if arrays do not have the same length.

	if idx == len(xs) {
		for ; idy < len(ys); idy, idz = idy+1, idz+1 {
			zs[idz] = ys[idy]
		}
	}

	if idy == len(ys) {
		for ; idx < len(xs); idx, idz = idx+1, idz+1 {
			zs[idz] = xs[idx]
		}
	}

	return zs
}
