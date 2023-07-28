package main

import "fmt"

// Explain and implment a cycle detection algorithm using DFS.
// Do a DFS and return true if at some point you see a vertex that was already visited.

func ContainsCycle(g Graph) bool {

	var newNode Node
	var toVisit Nodes
	visited := make(map[Node]bool)

	for _, node := range g.nodes {
		if _, ok := visited[node]; !ok {
			toVisit = Nodes{node}
			for len(toVisit) > 0 {
				toVisit, newNode = Pop(toVisit, false)
				if _, ok := visited[newNode]; !ok {
					visited[newNode] = true
					newNodes := g.adj[newNode]
					toVisit = append(toVisit, newNodes...)
				} else {
					fmt.Println("Found cycle at ", newNode)
					return false
				}
			}
		} else {
			fmt.Println("Found cycle at ", node)
			return false
		}
	}
	return true
}
