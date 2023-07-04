## One dimensional version

If `xs` is an array of objects that are totally ordered, then a _peak_ in `xs` is one of the following:

- An _inner peak_ is a triple `xs[i-1] <= xs[i] >= x[i+1]` where `1 <= i <= len(xs)-2`. 
- A _boundary peak_ happens when either `xs[0] >= xs[1]` or `x[n-2] <= x[n-1]` where `n = len(xs)`. 

_Problem_: write an algorithm that finds a peak.

_Fact_: A peak always exists with these definitions, and the proof gives the finding algorithm.

1. Stand at `n/2`. If its predecessor is larger, then there must be a peak: either the array is decreasing up to `n/2` and then the peak is at the boundary, or at some point the array dips, and there is an inner boundary. If its successor is larger, the same argument applies: if the array continues to increase then there is a peak at the boundary, and if not there is an inner peak. 

2. If inequality is replaced by strict inequality, then this can can fail: `[1, 2, 2, 1]` has no inner peak, and no boundary peaks. 

