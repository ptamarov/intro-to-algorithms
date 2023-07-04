package main

// Returns the index of a global maximum in the array of integers xs. Complexity Θ(n).

func GlobalMaximumByColumn(xs [][]int, row int) int {
	m := len(xs[0]) // Width of row.

	if m == 0 {
		return -1
	}

	max := xs[row][0]
	jdx := 0

	for i := 1; i < m; i++ {
		val := xs[row][i]
		if val > max {
			max, jdx = val, i
		}
	}

	return jdx
}

func GlobalMaximumByRow(xs [][]int, column int) int {
	n := len(xs) // Height of column.

	if n == 0 {
		return -1
	}

	max := xs[0][column]
	idx := 0

	for i := 1; i < n; i++ {
		val := xs[i][column]
		if val > max {
			max, idx = val, i
		}
	}

	return idx
}

func TwoDimensionalPeak(xs [][]int) (int, int) {
	rows, cols := len(xs), len(xs[0])

	if rows == 0 || cols == 0 {
		return -1, -1
	}

	if rows == 1 {
		return 0, GlobalMaximumByRow(xs, 0)
	}

	i, j := rows/2, GlobalMaximumByRow(xs, rows/2)

	if xs[i-1][j] >= xs[i][j] {
		return TwoDimensionalPeak(xs[:i])
	}

	if i+1 < len(xs) && xs[i][j] <= xs[i+1][j] {
		a, b := TwoDimensionalPeak(xs[i+1:])
		return i + 1 + a, b
	}

	return i, j

}
