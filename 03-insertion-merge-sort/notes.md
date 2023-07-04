## Why sorting? 

1. Sorting is simply necessary in real life: e.g. phonebooks.

2. Sorting makes other problems easier: 
    - finding the median of an array can be done in constant time in a sorted array.
    - item look-up can be done in `log(N)` time if the array is sorted (binary search).
    - data compression: sorting makes it easy to find duplicates, then replace repetitions by counts.
    - computer graphics: sorting by opaqueness and transparency (independently) is important. 


## Insertion sort

Basic idea: assume that `xs` is an array and that `xs[:i]` is sorted. Then, insert `xs[i]` in the correct place, and
obtain a sorted subarray `xs[:i+1]`. Continue until `xs` is exhausted. 

```
func InsertionSort(xs []int) {
	for i := 1; i < len(xs); i++ {
		key := i
		for j := i - 1; j > -1; j-- {
			if xs[key] >= xs[j] {
				continue
			} else {
				xs[j], xs[key] = xs[key], xs[j]
				key = j
			}
		}
	}
}
```

__Observation:__ since `xs[:i]` is sorted, it is much more convenient to instead do a binary search to insert `xs[i]`. But this
is not enough because the insertion of the element into the subarray still needs a linear amount of work. This is still an 
improvement if compares are much more expensive that insertions.

## Merge sort 

This is a good example of a _divide and conquer_ approach. We take an array `xs`, split it into two subarrays and apply the algorithm 
recursively to these two arrays to obtain two sorted arrays. Then, marge these two arrays into a single sorted array.

Notice that insertion sort involves _in place_ sorting, but a priori the recursive merge sort is creating copies of the original
array, which means that extra  `O(n)` auxiliary memory space is necessary. Insertion sort only needs `O(1)` auxiliary memory space.
One trick to address this is to only store half of `xs` (think about it). In place merge sort is another more complicated version (out
of scope), but it does have an overhead cost. 

_Note_: in Python, merge sort takes about `2.2 n log(n) µs`, whiile insertion sort takes about `0.2 n^2 µs`. Insertion sort in `C` 
is about 20 times faster, with constant factor `0.01`. In particular, once `n >= 4000`, merge sort in Python will beat insert sort in
`C`.