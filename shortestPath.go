package main

// Builds path array for the shortest path
func buildPath(start, finish int) []int {
	var path []int
	var vertex = finish

	// backtracking and creating a (reversed) path from parent array
	for vertex != start {
		path = append(path, vertex)
		vertex = parent[vertex]
	}
	path = append(path, start)

	return path
}

// Finds the shortest path
func findShortestPath(start, finish int, algorithm string) {
	// resetting previous values
	reset()

	// declaring costs
	var fCost float64
	var gCost float64
	var hCost float64

	// declaring other variables
	var currentVertex, to int
	var edge Edge

	// initializing priority queue
	queue := NewPriorityQueue()
	queue.Push(Edge{start, 0})
	dist[start] = 0

	// while queue is not empty
	for queue.Length() > 0 {
		// increasing visit count
		visitCount++

		// getting the current edge and vertex
		edge = queue.Front().(Edge)
		queue.Pop() // dequeue
		currentVertex = edge.to

		// terminating search when reaching goal
		if currentVertex == finish { break }

		// iterating through all the vertices that we can reach from current vertex
		for i := 0; i < len(edgeGraph[currentVertex]); i++ {
			// getting the next connected vertex
			to = edgeGraph[currentVertex][i].to
			// calculating gCost
			gCost = dist[currentVertex] + edgeGraph[currentVertex][i].distance

			// relaxing the edge if needed
			if gCost < dist[to] {
				// updating distance if shorter path is found
				dist[to] = gCost
				// setting heuristic in case of A*
				if algorithm == "astar" { hCost = heuristic(to, finish) }
				// calculating fCost (hCost is zero in case of dijkstra)
				fCost = gCost + hCost
				// enqueue
				queue.Push(Edge{to, fCost})
				// memorizing the parent vertex for building path
				parent[to] = currentVertex
			}
		}
	}
}
