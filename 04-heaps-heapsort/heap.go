package main

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

// HeapSort sorts the heap in increasing order.
func HeapSort(xs []int) {
	for i := len(xs); i > 0; i-- {
		BuildMaxHeap(xs[:i])
		xs[0], xs[i-1] = xs[i-1], xs[0]
	}
}

// HeapMaxInsert adds an element to a max-heap, preserving the heap invariant.
func HeapMaxInsert(xs []int, value int) []int {
	ys := append(xs, value)
	BuildMaxHeap(ys)
	return ys
}

// HeapminInsert adds an element to a min-heap, preserving the heap invariant.
func HeapMinInsert(xs *[]int, value int) {
	*xs = append(*xs, value)
	BuildMinHeap(*xs)
}

// HeapMax returns the maximum value of a heap without modifying it.
func HeapMax(xs []int) int {
	v := xs[0]
	return v
}

// HeapPop pops the root element of a heap, preserving the heap invariant.
func HeapPop(xs []int) int {
	if len(xs) == 0 {
		panic("cannot pop from empty heap")
	}

	xs[0], xs[len(xs)-1] = xs[len(xs)-1], xs[0]

	var val int
	xs, val = xs[:len(xs)-1], xs[len(xs)-1]

	BuildMaxHeap(xs)
	return val
}

// BuildMaxHeap makes xs into a max-heap, in place, in linear time.
func BuildMaxHeap(xs []int) {
	n := len(xs)
	for i := (n - 2) / 2; i > -1; i-- {
		heapify(xs, i, max)
	}
}

// BuildMinHeap makes xs into a min-heap, in place, in linear time.
func BuildMinHeap(xs []int) {
	n := len(xs)
	for i := (n - 2) / 2; i > -1; i-- {
		heapify(xs, i, min)
	}
}

// heapify takes a heap and an index i, and checks if the heap invariant
// with respect to f holds at i. It assumes that the left and right children
// of i are f-heaps, and ensures the heap invariant at i is satisfied.
// Usage: f is min or max on integers.
func heapify(xs []int, i int, f func(int, int) int) {
	n := len(xs)

	if 2*i+1 >= n { // Do nothing if node has no left child (it then also has no right child)
		return
	}

	rootValue := xs[i]
	leftValue := xs[2*i+1]

	if 2*i+2 < n { // Case when there is a right node is different.
		rightValue := xs[2*i+2]

		if rootValue == f(rootValue, f(leftValue, rightValue)) {
			return
		}

		if leftValue == f(leftValue, rightValue) {
			xs[i], xs[2*i+1] = leftValue, rootValue
			heapify(xs, 2*i+1, f)
			return
		}

		xs[i], xs[2*i+2] = rightValue, rootValue
		heapify(xs, 2*i+2, f)
		return
	}

	if rootValue == f(rootValue, leftValue) { // No right node, so check what happens at left node.
		return
	} else {
		xs[i], xs[2*i+1] = leftValue, rootValue
		heapify(xs, 2*i+1, f)
		return
	}

}
