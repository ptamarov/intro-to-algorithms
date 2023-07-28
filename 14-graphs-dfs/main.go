package main

import "fmt"

type Node int
type Nodes []Node

func recursiveDFS(s Node, adj []Nodes, parent map[Node]Node, seen Nodes) Nodes {
	_, ok := parent[s]
	if !ok {
		parent[s] = -1
	}
	seen = append(seen, s)
	for _, v := range adj[s] {
		if _, ok := parent[v]; !ok {
			parent[v] = s
			seen = recursiveDFS(v, adj, parent, seen)
		}
	}
	return seen
}

func DFS(s Nodes, adj []Nodes) Nodes {
	parent := make(map[Node]Node)
	seen := Nodes{}

	for _, node := range s {
		if _, ok := parent[node]; !ok {
			seen = recursiveDFS(node, adj, parent, seen)
		}
	}
	fmt.Println(parent) // print a spanning forest for the graph. Marked with * below.

	return seen
}

func main() {
	// 0 --- 1    2 -- 3
	// |     |  / |  / |
	// |     | /  | /  |
	// 4     5 -- 6 -- 7

	adj := []Nodes{{4, 1}, {5, 0}, {3, 5, 6}, {2, 6, 7}, {0}, {1, 2, 6}, {2, 3, 5, 7}, {3, 6}}
	out := DFS(Nodes{0, 1, 2, 3, 4, 5, 6, 7}, adj)

	fmt.Println(out)
	// [1 0 4 5 2 3 6 7]
	// map[0:-1 1:0 2:5 3:2 4:0 5:1 6:3 7:6] i.e. the following spanning tree
	// 0 *** 1     2 *** 3
	// *     *   * |   * |
	// *     *  *  | *   |
	// 4     5 --- 6 *** 7

	// Job scheduling example
	adj1 := []Nodes{{2}, {2, 4}, {}, {4, 7}, {5}, {6}, {}, {7}}
	out1 := ScheduleJob(Nodes{0, 1, 2, 3, 4, 5, 6, 7}, adj1)

	fmt.Println(out1)

}
