## Motivating example

Reservations for future landings in a busy airport runway. 

- Want to reserve requests for landing (request specifies a landing time `t`)
- Add `t` to the list of landing times if not landings are scheduled within `k` minutes of `t`. 
- The parameter `k` can be static, or can vary depending on other conditions.
- If condition is not satisfied, request is denied. 
- Keep track of time, then remove `t` when it is due (i.e. plane has landed).

__Goal__: be able to do these operations on logarithmic time (with respect to length of list).


1. Unsorted arrays are not good: essentially any complicated operation is at least linear.
2. Sorted arrays: good for finding insertion point and checking constrain, but bad for instertion (still linear)
3. Sorted list: can change pointers in constant time, but we cannot do binary search.
4. Heaps (min/max): invariant is quite weak, so cannot do binary search with them (takes linear time).

## Binary search trees

We define a new ADS with a stronger invariant: a _binary search tree_ is a binary tree in which the key of 
a node is _larger_ than all of the keys in its _left_ subtree and _smaller_ then all of the keys in its _right_
subtree. That is, it grows to the right and decreases to the left.

```
           ________ 20 ________
         /                     \
   __ 10 ____                     30
 /           \
 9         __15___
         /        \
        14         11

```

The basic structure is then:

```
    type node struct {
        left  *node
        va    lint
        right *node
    }
```

__Take away__: inserting takes `O(h)` where `h` is the height. The same is true for finding the minimum, maximum or finding, for
a given value `v` the largest key smaller than `v` or the smallest key larger than `v`.

_New requirement_:  how many planes are scheduled to land before a given time `t`? Call this `Rank(t)`.

_Solution_: Augment the tree by decorating each node with the size of its subtree. Start with a counter `c := 0`.
Now `Rank(t)` can be computed by finding the smallest `t'` larger than `t`, and while doing so:
    - Adding 1 to a counter when we move right (i.e. if we find something smaller),   
    - Also adding the size of the left subtree we ommitted (i.e. adding the number of nodes passed apart from the subtree root).

