# Last time

We saw that BSTs support the following operations:

1. `Insert`, `Delete`
3. `Minimum`, `Maximum`
5. `Sucessor`, `Predecessor`

All of this can be done in `O(h)` time where `h` is the height of the tree. 

__Problem__: `h` could be `N` where `N` is the size of the tree if we do not _balance_ it, i.e. distribute nodes
evenly among subtrees. Thus, it is convenient to keep `h` as close as possible to `log(N)` at every time (possibly
up to a constant factor).


If `n` is a node in a tree, the _depth_ of `n` is its distance to the root. The _height_ of `n` is the longest
path from `n` (down) to a leaf. In the image below, `26` has height `0`, `29` has height `1`, `20` has height `2`, etc.


```
                          ________ 41(3) ________
                        /                         \
                __ 20(2) ____                   65(1)
               /              \                 /
                11(0)     ___29(1)          50(0)
                        /        
                        26(0)        
```

In general, we have that `ht(n)- 1` is equal to `Max(ht(n.Left), ht(n.Right))`. 

By augmenting our data structre, we can augment our BSTs to keep track of the height of each node at a constant
overhead cost. The rule is local at each node, so this works.

_Convention_: the height of an empty node (i.e. empty pointer) is `-1`. Will be useful later so as not to worry about subcases.

## AVL trees

_Goal_: keep track of heights, and fix them if they get too big. We say a binary search tree is an `AVL tree` if
the heights of left and right children of _every_ node differ by at most `1`.  

For example, the tree above is AVL (notice the `-1` convention kicks in). One could also balance by
subtree size, this gives a different notion of a balanced search tree.

_Obs_: zero difference is sometimes impossible to achieve. We will show that keeping a difference of at most `1` is not only achievable,
but can be maintained at a cost of `log(N)`. 

_Construction_: define `T(0)` to be the tree with a single node, and let `T(1)` be the tree with a root that has a single right child.
Both of these trees are AVL. Define recursively the tree `T(n+1) = [ T(n-1) T(n) ]`. Then `T(n+1)` is an AVL tree where the right child
at every node has height one more than its left child. Hence, `T(n)` has size `Fib(n) - 1` and height `n`. Since `F_n` grows like `1.6 ** n`,
it is enough to show this is the worst case scenario, which turns out to be true. 

___Claim 1___: AVL trees are balanced. 

Let `N(h)` be the minimum number of nodes an AVL tree of height `h` can have (worst case scenario). Then 

```
    N(h) = N(h-1) + N(h-2) + 1 
```

If we put `M(h) = N(h) + 1` then `M(0) = 2; M(1) = 3` and  `M(h) = M(h-1) + M(h-2)`, so we see that `N(h)` is up to a constant a
Fibonacci number. Hence, the worst case scenario gives exponentially many nodes, so AVL trees are balanced. 

Notice that one does not need to solve the recurrence explicitly, we can just notice that `N(h) >= 2 * N(h-2)`. If we assume
that `N(x) >= 2 ** (x/2)` for all `x < h+2`, then clearly `N(h+2) >= 2 **((h+2)/2)`. So we already know that there are at least
exponentially many nodes. 

## AVL insert

1. Do the usual BST insertion. 
2. Fix the AVL property. 

_Note_: if we know we have an AVL tree, then it is enough to remember which side is the bigger one, or if they are even. Thus, you
can just store `-1, 0, 1`. 

```
                          ________ 41(3) ________
                        /                         \
                __ 20(2) ____                   65(1)
               /              \                 /
                11(0)     ___29(1)          50(0)
                        /        
                        26(0)        
```

If we insert `23`, we will end up with the following non-AVL tree. The algorithm will now walk up the tree and try to find a 
vertex that has turned bad (in this case, 29 is too heavy to the left). 

```
                          ________ 41(4) ________
                        /                         \
                __ 20(3) ____                   65(1)
               /              \                 /
                11(0)     ___29(2)          50(0)
                        /        
                      26(1)
                     /        
                   23(0)
```

This can be fixed by "rotating" 29 to the right:


```
                          ________ 41(3) ________
                        /                         \
                __ 20(2) ____                   65(1)
               /              \                 /
                11(0)     ___26(1)___        50(0)
                        /            \
                      23(0)          29(0)
                
```


_Rotation_ is the following operation on trees: `LeftRotate(x)` does the following

```         
           |
      ____ x _____
    /              \ 
    A         ____ y ____
            /             \
            B              C
```

turns into the following with a constant number of pointer changes


```
                |
          _____ y _____
        /              \
  _____ x ____          C
/             \
A             B              
```

Notice that the in order traversal of both trees is `AxByC`, so the order is preserved. The inverse
operation is `RightRotate(y)`. In the previous example, we did `RightRotate(29)`, with `x = 26`, `A = 23`
and `B` and `C` empty subtrees. 

Let us now do `Insert(55)`:


```
                          ________ 41(3) ________
                        /                         \
                __ 20(2) ____                   65(2)
               /              \                 /
                11(0)     ___26(1)___        50(1)
                        /            \         \
                      23(0)          29(0)      55(0)
                
```

The AVL property fails at `65`, since `50` has height `1` and the right child has height `-1`.


```
                          ________ 41(3) ________
                        /                         \
                __ 20(2) ____                   65(2)
               /              \                 /
                11(0)     ___26(1)___        50(1)
                        /            \         \
                      23(0)          29(0)      55(0)
                
```

We rotate at `50` to the left to get



```
                          ________ 41(3) ________
                        /                         \
                __ 20(2) ____                   65(2)
               /              \                 /
                11(0)     ___26(1)___         55(1)
                        /            \         / 
                      23(0)          29(0)   50(0)
                
```

and then we rotate `65` to the right:


```
                          ____ 41(3)_____
                        /                 \
                __ 20(2) _               55(1)
               /          \             /    \
             11(0)       26(1)       50(0)   65(0)
                      /      \         
                    23(0)   29(0)   
                
```

This is called a double rotation.

### General case for AVL insert (i.e. fix AVL property)

1. Insert node using the usual BST insertion. 
2. The node is inserted with height `0`.
3. Go to parent, and re-compute height by `Max(ht(n.Left), ht(n.Right)) + 1`. 
4. Check that AVL property holds for subnodes. 
5. Repeat 3 and 4 until the property fails for some node `x`. 

We now fix the AVL property at node `x`, assuming that it is right over-heavy.

_Case 1_. The right child is right heavy. Then, we do `LeftRotate(x.Right)`. 
This turns

```         
           |
      ____ x(h) _____
    /                \ 
A(h-3)        ____ y(h-1) ____
            /                 \
            B(h-3)           C(h-2)
```

into to:


```
                |
          _____ y(h-1) _____
        /                   \
  _____ x(h-2) ____       C(h-2)
/                  \
A(h-3)            B(h-3)              
``` 

and we now move on to the parent of `y`. 

_Case 2_. If `y` was balanced, then things work out as well:


```
                |
          _____ y(h)_______
        /                   \
  _____ x(h-1) ____       C(h-2)
/                  \
A(h-3)            B(h-2)              
``` 

_Case 3_:  the right c hild `y` is left heavy. Then we need to rotate twice. 

```         
           |
      ____ x(h) _____
    /                \ 
A(h-3)        ____ y(h-1) ____
            /                 \
      ___z(h-2)___         D(h-3)
     /            \
    B(h-3)        C(h-3)
```

first, rotate `z` to the right


```         
           |
      ____ x(h-1) _____
    /                 \ 
A(h-3)        ____ z(h-1) ____
            /                 \
           B(h-3/4)     ____y(h-2)____
                     /               \
                     C(h-3/4)         D(h-3)
```

which still gives an unbalanced tree. Then rotate `x` to the left

```         
               _____ z(h-1) ______
              /                   \
        ____x(h-2)__        _____y(h-2)_____
       /            \      /                \
    A(h-3)       B(h-3)  C(h-3)           D(h-3)
```

which balances out the tree. Then move on to the parent of `z`.

__AVL sort__ implements a sorting algorithm in `O(N log N)` time. Insert the `N` items into an AVL (each successive `i`th insert costs `log(i)`)
and then go through the tree in order.


