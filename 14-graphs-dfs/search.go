package main

type Graph struct {
	nodes Nodes
	adj   []Nodes
}

// SearchGraph searches the graph by traversing its nodes in the input order.
// If fifo is set to false it performs a DFS, if not it performs a BFS.
func SearchGraph(g Graph, fifo bool) Nodes {
	var newNode Node
	var toVisit Nodes
	var out Nodes
	visited := make(map[Node]bool)

	for _, node := range g.nodes {
		if _, ok := visited[node]; !ok {
			toVisit = Nodes{node}
			for len(toVisit) > 0 {
				toVisit, newNode = Pop(toVisit, fifo)
				if _, ok := visited[newNode]; !ok {
					out = append(out, newNode)
					visited[newNode] = true
					newNodes := g.adj[newNode]
					toVisit = append(toVisit, newNodes...)
				}
			}
		}
	}
	return out
}

// SearchFromNode visits the whole connected component of the input node according
// to the adjacency list. If fifo is set to true, it performs a BFS, if it
// is set to false, it performs a DFS.
func SearchFromNode(s Node, adj []Nodes, fifo bool) Nodes {
	var newNode Node
	var out Nodes

	toVisit := Nodes{s}
	visited := make(map[Node]bool)

	for len(toVisit) > 0 {
		toVisit, newNode = Pop(toVisit, fifo)
		if _, ok := visited[newNode]; !ok {
			out = append(out, newNode)
			visited[newNode] = true
			newNodes := adj[newNode]
			toVisit = append(toVisit, newNodes...)
		}
	}
	return out
}

func Pop(s Nodes, fifo bool) (Nodes, Node) {
	var out Node
	if fifo {
		out = s[0]
		s = s[1:]
	} else {
		out = s[len(s)-1]
		s = s[:len(s)-1]
	}
	return s, out
}
