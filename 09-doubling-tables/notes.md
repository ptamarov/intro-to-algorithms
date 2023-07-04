We know that a dictionary supports

1. `Insert`
2. `Delete`
3. `Lookup`

and that using a hash table, we can achieve a `O(1)` loopup operation with high probability. 

We saw that the idea is to have a table of a certain size `m`, and a hash function that assigns keys to 
one integer in `{0, ..., m-1}`. We argued that if the table stores `n` distinct keys and if these
are distributed evenly among the `m` buckets, then one can do a look up in `O(1+α)` where `α = n/m`.
Hence, we want `m` to be large enough in relation to the number of keys stores, but not too large
to be taking up too much unnecessary memory.

## How do we choose `m`?

We want `m = Θ(n)` so that `α = Θ(1)`. We will make sure this holds at any given time by growing or
shrinking the table. _New concept_: amortization.

1. Start with some small `m`, say `m = 8`. 
2. Grow or shrink as necessary. 

__What does it mean to grow the table?__ 

We must allocate memory for the new buckets, and then build new hash function and rehash all previous
keys with the new hash function, i.e. re-insert the items. How much time does this take? Suppose we
had `n` keys in `m` buckets and we want to create a new table of size `m'`. Then we pay `m` to visit
every bucket, `n` to read every key and `m'` to initialize the new `m'` buckets.

A naive guess would be that it is enough to allot only a new bucket, but this is _wrong_: we would
have to grow the table with each insertion, making the cost of `n` insertions quadratic in `n`.

A good guess is to _double_ the size of the table: the cost of `n` insertions is linear, but not all
of the `n` insertions are expensive: only about `log n` of the `n` insertions are costly (and cost `2^t` for some
`t`), and the other ones take constant time. So, on average, the cost of one insertion is `O(1)`.

We say that an operation takes `T(n)` amortized time if `k` operations take at most `k T(n)` time. Thus,
insertion is `O(1)` amortized, since we have just argued that `n` insertions take at most linear time.

## String matching

Give two strings `s` and `t`, determine if `s` occurs as a substring of `t`. In general, `s` is small, and `t` is __huge__: 
for example, `s` can be a keyword and `t` could be the inbox of an email address from the past 10 years. The naive approach
is `O(M(N-M))`. where `M = len(s)` and `N = len(t)`. We will show that it can be made `O(M+N)`.

[Complete.]