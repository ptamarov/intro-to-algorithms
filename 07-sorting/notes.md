## Comparison model

Only operations given on ADTs are comparisons (`isLarger`, `isSmaller`, or `isEqual`) giving a boolean value. The time cost is defined
as the number of comparison that are needed. We will show that:

1. Searching needs at least `log N` time.
2. Sorting needs at least `N log N` time. 


## Searching 

Assume that one can pre-process the `N` items so that they are sorted. We have a target `t`, and consider the _decision tree_ of our algorithm.
At the root, one comparison operation is made, and two possible outputs are given (a `bool`). Continuing in this way, we draw a
tree containing all possible outcomes of our algorithm, which must contain at least one leaf for each one of the `O(N)` outcomes of our
searching process (it could be possible that there are _more_ than the exact amount of leaves, i.e. some leaves contain the same answer)
but at least we know that there are at least linear-in-`N`-many leaves. Thus, the tree must have at least `log N` nodes, which gives us
the requisite lower bound for the running time. 

## Sorting

We run the same argument, but now notice that there are _at least_ `N!` many leaves, corresponding to the `N!` permutations of a list
of `N` elements. Then, we notice that `log N!` is at least `(N/2 -1) log N`, which gives us the necessary lower bound. 

## RAM model

__Counting sort__. We assume that our `N` inputs are among the `k` integers `{0, ..., k-1}`. We then go through our array in linear
time, and store how many times integer `i` occurs in an array with keys `{0, ..., k-1}`. Then, we traverse this array and return at
each step as many copies of `i` as stored in the array. More generally, we can assume that our input list contains items `A` that have a
`key(A)` in `{0, ..., k-1}`. The pseudo-code is:

```
def countingSort(A: list[int]) -> list[int]:

    N = len(A)
    L = [[] for _ in range(k)]

    for j in range(N):
        L[key[A[j]]].append(A[j])

    out = []

    for i in range(k):
        out.extend(L[i])

    return out 
```

_Analysis_: 

1. Creating the list of empty lists is `O(k)`.
2. The appends take constant time, so the first loop is `O(N)`.
3. Checking `L[i]` takes constant time, while each extend takes `len(L[i])` time, for a total of `O(N+k)`.

Thus, counting sort is `O(N+k)`, which is good if `N >> k`. This is not a good algorithm in general!

## Radix sort

Radix sort uses counting sort as a subroutine. It will be linear in `N` as long as `k` is polynomial in `N`.

Let us consider a base `b` which we will fix later. Consider each of the `k` integers in `{0, ..., k-1}` 
to be sorted in base `b` (but do not compute their base `b` representation). Notice that the number of 
digits needed is at most `log_b(k)+ 1`. Then:

1. Sort the integers by their least significant digit i.e. do `xs[i] mod b`. 
2. Sort them by the second significant digit, i.e. now do `(xs[i] // b) mod b`.
3. ...
4. Sort them by their most significant digit.

Each step is a counting sort that takes `O(N+b)` time, and we do this `log_b(k)+1` times, so the running time
is `O((N+b) log_b(k))`. Now we argue that this running time is minimum when `N≈b`, so we get a running time of
`O(N log_N(k))`. If `k` is polynomial in `N`, then `log_N(k)` is a constant, and we reach the requisite
conclusion that the running time is `O(N)`. 



