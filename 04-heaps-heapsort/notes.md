## Heaps

An abstract data type modelling a a priority queue of elements of certain type `Type` which are comparable according a value `Value`, with a specified set of operations. _Goal_: study the complexity of these operations.


```
    func (h *Heap) Insert(x Type) 
    func (h *Heap) Max(x Type) 
    func (h *Heap) ExtractMax(x Type) 
    func (h *Heap) ExtractMax(x Type) 
    func (h *Heap) UpdateKey(x Type, k Value) 
```

A _heap_ is an array that is visualized as a _nearly complete_ binary tree. For example:

```
    16 14 10 8 7 9 3 2 4 1 
```

is visualized as follows

```
                 16
              /      \
            14        10
          /   \      /   \  
        8       7   9     3
      /  \     /
     2    4   1  
```

What do we get out of this visualization: 
    - the root of the tree is the first element of the array.
    - the children of element `i` are at `2i+1` and `2i+2`.
    - the parent of element `i` is at `(i-1) / 2`.

A _max heap_ is a heap where a node is larger or equal to its children. A _min heap_ is a heap where the children are larger.

### Operation `MaxHeapify`

If `h` is a heap (i.e. just an array) then `MaxHeapify(h, i)` requires that at node `i`, the left and right subtrees are already max-heaps.
For example, consider the heap 

```
                 16
              /      \
            4        10
          /   \      /   \  
        14     7   9     3
      /  \     /
     2    8   1  
```

which fails the max-heap condition at node `4` at index `1`. The subtrees at `14` and at `7` are already max-heaps.

1. Look at `4` and exchange `4` and `14` i.e. exchange root node with its largest child. 

```
                 16
              /      \
            14        10
          /   \      /   \  
        4     7   9     3
      /  \     /
     2    8   1  
```

2. Notice that now the subtree at `4` is not a max-heap, so max-heapify it by calling the function at `4`.

```
                 16
              /      \
            14        10
          /   \      /   \  
        8      7    9     3
      /  \     /
     2    4   1  
```

3. Now we _must_ have a heap at `14`, since we put the largest subchild at the root. The array is now a max-heap.

_Complexity_: If the heap has size `n` (since we are assuming the binary is tree is nearly complete), the complexity is `log n`. 

(_Spoiler_: work up from leaves using `MaxHeapify` which guarantess the pre-condition is satisfies at any iteration.)

### Operation `Heapify`

How to take an unordered array and build a max-heap out of it? Important observation is that
if the heap has lenght `n` then the first node that is not a leaf will be at position `(n-1)/2`
(the parent of the node at `n-1`), so we do not have to work further than that.

```
        func BuildMaxHeap(xs []int) {
            n := len(xs)
            for i := (n-1)/2 ; i > 0 ; i-- {
                MaxHeapify(xs, i)
            }
        }
```

__What is the complexity?__ We have about `n/2^{i+1}` with level `i`, where level 0 are the leaves, and if a node is at level `l`. then
the procedure takes about `Θ(l)` time (there is nothing to do at level 0, the leaves). Then, working in a log scale `2^k = n/4`, we are summing
something of the form


```
    c * 2^k (1 / 2^0 + 2 / 2^1 + 3 / 2^2 + ... + (k+1) / 2^k)

```

where the summation converges to a constant (in this case it converges to `4`). Thus the complexity is `Θ(n)`. 

### Implementing `HeapSort`

Now that we can heapify an array, we can use this to sort arrays:

1. Extract `xs[0]` which is the maximum.
2. Swap `xs[-1]` and `xs[0]` and truncate at end.
3. Heapify at root and continue. 