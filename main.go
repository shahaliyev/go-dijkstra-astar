package main

// Global constants
const SIZE = 100
const CELLSIZE = 1000
const INFINITY = 0xf3f3f3f

// Global arrays
var edgeGraph = make([][]Edge, SIZE)
var vertexGraph = make([]int, SIZE)
var dist = make([]float64, SIZE)
var parent = make([]int, SIZE)

// Global variables
var visitCount int

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
	// initializing heuristic cost to zero
	var hCost float64
	// initializing priority queue
	queue := NewPriorityQueue()
	queue.Push(Edge{start, 0})
	dist[start] = 0

	// implementing dijkstra / A*
	for queue.Length() > 0 {
		// increasing visit count
		visitCount++

		// getting the current edge and vertex
		edge := queue.Front().(Edge)
		queue.Pop() // dequeue
		currentVertex := edge.to

		// terminating search when reaching goal
		if currentVertex == finish { break }

		// iterating through all the vertices that we can reach from current vertex
		for i := 0; i < len(edgeGraph[currentVertex]); i++ {
			// getting the next connected vertex
			to := edgeGraph[currentVertex][i].to
			// calculating gCost
			gCost := dist[currentVertex] + edgeGraph[currentVertex][i].distance

			// relaxing the edge if needed
			if gCost < dist[to] {
				// updating distance if shorter path is found
				dist[to] = gCost
				// setting heuristic in case of A*
				if algorithm == "astar" { hCost = heuristic(to, finish) }
				// calculating fCost (hCost is zero in case of dijkstra)
				fCost := gCost + hCost
				// enqueue
				queue.Push(Edge{to, fCost})
				// memorizing the parent vertex for building path
				parent[to] = currentVertex
			}
		}
	}
}

func main() {
	// reading values from command line
	inputPath, algorithm := readCommandLine()

	// reading input file from path
	input := readInput(inputPath); defer input.Close()

	// building graphs & getting test values from input
	start, finish := buildFromFile(input)

	// finding shortest path
	findShortestPath(start, finish, algorithm)

	// displaying output
	displayOutput(start, finish)

	// testing (see test.go)
	test()
}
