
## First approach to a dictionary: direct access table.

1. Keys are non-negative integers.
2. Store items in an array.

_Issues_: 
- cannot store items using keys other than integers.
- huge memory hog: uses too much memory. For example, you need one slot in your array per key.

_Solution to first issue: pre-hashing_: map keys (whatever they are, so long as they are hashable) to integers. In theory keys are finite and discrete (string of bits). Ideally, a pre-hash function should be _injective_, but this is not the case, and it cannot be the case since we want to resolve the memory issue.

One can implement (for example, in Python) what `__hash__` should do on custom objects. This should be done carefully to avoid it changing. For example, lists are not hashable since they can change over time. The same is true for sets. 

_More interesting problem_: reducing space through hashing (related to "hatching", cutting into pieces and mix around). Idea is to take all possible keys (integers) and reduce them to a reasonably small size `m` for the table. Notice that although the set of _possible_ keys is very large, the set of keys that are in the hash table is in general smaller (though still large). 

_Idea_: we would like for `m` to be around the number of keys in the dictionary at a given moment. This is optimal, i.e. proportional to the size of information being stored.

__Obvious problem__: there will be collisions! This is _guaranteed to happen_.

## Chaining 

The way to solve the collision problem is to store items with the same hash in a linked list, i.e. the hash table contains pointers to linked lists, and a null pointer is stored if there are no items with that hash value.

Why would we expect this to be good? Worst case, all keys will just go to one huge linked list! This can happen, and hashing is theoretically worst case `\Theta(n)`, but the hash functions are designed in a way that collisions occur with low probability, so that these linked lists will be evenly distributed and have more or less a constant length. This can be proved, and of course requires some assumptions. 

__Assumption__: simple uniform hashing (not quite a correct assumption, but good for purposes): each key is equally likely to be hashed to any slot of the table, uniformly and independently of the other keys. 

## Analysis of hashing with chaining

__Problem__: What is the expected length of a linked list in a hash table of size `m` that contains `n` keys? It is `n/m`, since the probability of one key going to a given slot of `1/m`, and there are `n` independent keys. Easy! In the real world, there is no independence. This ratio is called to _load factor_ of the hash table. This is constant as long as `m` is `\Theta(n)`. Thus, the running time of an operation is `O(1+α)`, where the 1 comes from computing the hash function. 

## Hashing functions 

1. Division method: use `h(k) = k mod m`. This brings problems if `k` and `m` share factors, but can work if `m` is prime and not to close to a power of 10 or 2. 

2. Multiplication method: use `h(k) = (a * k mod 2^w) >> (w-r)`, where `m=2^r`. Assume working with a `w` bit machine. This will take `a` copies of `k` and shift them according to the non-zero bits of `a`, add them up, then throw away the left word (result is twice as long as `k`) and then take the `r` first bits of the last `w` bits. Number `a` should be odd and not too close to a power of 2. 


## Universal hashing

The hashing function is `h(k) = ((ak + b) mod p) mod m`, where `p` is a prime number larger than the universal set of keys, and `a`, `b` are random numbers in `0, ..., p-1`. The probability that this choice of hash function has a collision for worst-case keys is `1/m`. 