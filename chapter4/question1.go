package chapter4

import "fmt"

/*
Route Between Nodes:
Given a directed graph, design an algorithm to find out
whether there is a route between two nodes.
*/

type node struct {
	name     string
	adjecent []*node
	state    nodeState
}

type nodeState int

const (
	unvisited nodeState = iota
	visiting
	visited
)

func hasRouteBetweenNodes(n1, n2 *node) bool {
	if n1 == n2 {
		return true
	}
	fmt.Printf("path %s -> %s \n", n1.name, n2.name)
	// use slice for queueing
	var q []*node

	// uses BFS on the first node
	n1.state = visiting
	q = append(q, n1)

	var current *node
	for len(q) > 0 {
		// dequeue
		current = q[0]
		q = q[1:]
		fmt.Printf("arrived at %s \n", current.name)
		if current != nil {
			for _, adj := range current.adjecent {

				if adj.state == unvisited {
					fmt.Printf("visiting %s\n", adj.name)
					if adj == n2 {
						return true
					}
					adj.state = visiting
					q = append(q, adj)
				}
			}
			current.state = visited
		}
	}
	return false
}
