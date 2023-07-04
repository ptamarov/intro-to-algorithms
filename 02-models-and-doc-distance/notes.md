## What's an algorithm?

On a high level, an algorithm is a way to define a computational procedure for solving a problem. An algorithm by definition 
has an input and an output. It is an analog of a computer program, which can be then converted to actual computer code.

Programs are written in specific programming language, while algorithms are usually written in _pseudocode_, i.e. express the
algorithm in a clear way that people can understand and one can formally reason about. The analog of a computer in the 
mathematical world is a _model of computation_: what operations is your computer allowed to do? How much do they cost in terms
of memory and running time?

### Random access machine (RAM)

This models random access memory. It is a _giant_ array (maybe 4GB, maybe more) of _words_. How do we compute with it? In constant time, an algorithm can:

1. read a constant words from memory,
2. do a constant number of computations on them,
3. write them out (store them).

What is a _word_? It is some finite number `w` bits. Since `w` is supposed to specify an index into the array of memory, it sould be at least as 
big as `log(M)` where `M` is the memory size.

_Note_: this model has no dynamic memory management (or allocation): memory management is a form of resource management applied to computer 
memory. The essential requirement of memory management is to provide ways to dynamically allocate portions of memory to programs at their 
request, and free it for reuse when no longer needed. 

### Pointer machine 

This model is slightly more abstract than the RAM model but also weaker, since one can implementer it using the RAM model, but not conversely.

- _Objects_ are allocated dynamically.  
- Objects have a constant number of _fields_.
- A field is either a word (e.g. store an `int`) or a _pointer_.
- A pointer is something that points to another object, or `nil`, `null`, `None`.

_Example_: (doubly) linked lists. These are of the form

```
type LinkedList struct {
    next *LinkedList
    value Type
    previous *LinkedList
}
```

Nodes can be destroyed and created using dynamic memory allocation. In the RAM model, a pointer corresponds to a memory address, so one can 
do pointer arithmetic, but in the pointer machine model one can only follow pointers at a constant time cost. 

_Note_: both of these models are old, from the 70s and 80s.

## Python model 

It is the model of computation used for these lectures. It can model RAM since it has arrays, and it can model pointers since it has _references_ 
(although not pointers). Operations are now more rich (and complicated): one can `append` to or `sort` an array, one can add two lists, etc. 
It is important to know which Python operations are _slow_ and which are _fast_.

1. Arrays are `lists`. Consulting a value in a list of modifying a value takes constant time. For example, if `xs` is a list of length `8`, 
the following instruction, which edits the 5th entry, takes constant time:

     ``` xs[4] = xs[7] + 3```

2. If an object has a constant number of attributes, accessing an attribute takes constant time. 
3. If `xs` is a list, appending a single element to it takes constant time. However, adding two lists takes linear time. 
4. `x in xs`: this is linear in `len(xs)`. Computing the length of a list is constant, since length is stored.
5. Sorting a list of lenght `N` takes around `O(tN log(N))` where `t` is the time it takes to compare two elements in the list.
6. Accessing the value of a key in a `dict` takes constant time _with high probability_.
7. `long` integers: adding two integers of length `s` and `t` takes `O(s+t)`. Multiplying them takes `O((s+t) ** log 3)`.

## Document distance

_Problem statement_: determine the distance between two _documents_. We think of a document as a sequence of _words_, where a word
is a string of alphanumeric characters `[A-Z][a-z][0-9]`. Then, look at shared words. There are lots of other possibilities. 

_Example_: catalogue webpages and know when to store less when there are repetitions or almost identical content. 

__Idea__: think of a document as a vector, where the coefficient of that vector is the number of times that word has appeared in the document. 
Thus, every document gives a plot where peaks are words occurring often. We can then use methods from linear algebra to produce
distances. For example, a dot product would give a good measure of _commonality_. To make this scale invariant, divide by the length of the
vectors, which is then measuring angles between vectors. But how do we compute this?

1. Split document into words. 
2. Vectorize the document.
3. Compute the inner product.

