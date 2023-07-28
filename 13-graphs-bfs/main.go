package main

import (
	"fmt"
)

type Node int
type Nodes []Node

const PrintPaths = false

func BFS(s Node, adj []Nodes) Nodes {
	level := make(map[Node]int)
	level[s] = 0

	parent := make(map[Node]Node)
	// BFS also finds a shortest path from a given node to s
	// Iterate parent(node) until it is nil. Level keeps track of number of iterations anyways.
	j := 1

	frontier := []Node{s}
	out := []Node{s}
	for len(frontier) > 0 {
		next := []Node{}
		for _, u := range frontier {
			for _, v := range adj[u] {
				if _, ok := level[v]; !ok {
					level[v] = j
					parent[v] = u
					next = append(next, v)
				}
			}
		}
		out = append(out, next...)
		frontier = next
		j++
	}

	paths := []Nodes{}
	// compute shortest paths:
	for node, value := range level {
		iters := value
		path := Nodes{node}
		for iters > 0 {
			par := parent[node]
			node = par
			path = append(path, node)
			iters--
		}
		paths = append(paths, path)
	}

	if PrintPaths {
		fmt.Println(paths)
	}
	return out
}

func main() {
	// 0 --- 1    2 -- 3
	// |     |  / |  / |
	// |     | /  | /  |
	// 4     5 -- 6 -- 7

	adj := []Nodes{{1, 4}, {0, 5}, {3, 5, 6}, {2, 6, 7}, {0}, {1, 2, 6}, {2, 3, 5, 7}, {3, 6}}

	a := BFS(1, adj)
	fmt.Println(a)
}
