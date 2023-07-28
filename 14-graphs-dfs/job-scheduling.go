package main

// ScheduleJob takes a DAG on a collection of Nodes and returns a total order (scheduling)
// so that whenever a Node is scheduled, any jobs pointing to it have already been scheduled
// (i.e. a linear extension of the DAG)
func ScheduleJob(nodes Nodes, adj []Nodes) Nodes {
	var newNode Node
	var toVisit Nodes
	var currentStack Nodes
	var out Nodes
	visited := make(map[Node]bool)

	for _, node := range nodes {
		if _, ok := visited[node]; !ok {
			toVisit = Nodes{node}
			for len(toVisit) > 0 {
				toVisit, newNode = Pop(toVisit, false)
				if _, ok := visited[newNode]; !ok {
					currentStack = append(currentStack, newNode)
					visited[newNode] = true
					newNodes := adj[newNode]
					toVisit = append(toVisit, newNodes...)
				}
			}
			reverse(currentStack)
			out = append(out, currentStack...)
			// don't forget to empty the current stack...!
			currentStack = Nodes{}
		}
	}
	reverse(out)
	return out
}

func reverse[T any](out []T) {
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
}
