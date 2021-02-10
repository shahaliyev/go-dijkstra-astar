# Shortest Path implementation in Go (Dijkstra & A*)

Program reads the graph and takes two inputs: start and end points, and calculates the shortest path from start to end.  

## Installation

Go should be installed beforehand https://golang.org/doc/install and GOPATH should be set.

Then either download .zip file of this package or go get from command line to GOPATH:

```
$ go get github.com/shahaliyev/go-dijkstra-astar
```

## Application

Given a random graph. The “nodes” in a graph are located on a 10x10 2D "chess board", consisting of 100 squares (from 0 to 99). Each square is itself of size 1000 by 1000 units.

### Input

Input is a **input_name.txt** file with three sections: 

**Vertex section** includes the node data, which is **vertex ID**, and the **“square”** it falls on. This is given in this format:

```
1,72 // Vertex #1 falls in “square” 72
2,69 // Vertex #2 falls in “square” 69
```

**Edge section** includes the set of edges, which are **fromVertex**, **toVertex**, and the **distance** between them:

```
1,3,349.22 // Vertices 1 and 3 are 349.22 distance apart, respectively.
```
**Test section** includes start and destination values for testing:

```
S,0 // start vertex ID
D,99 // destination vertex ID
```

### How to run

If the installation part is complete, it remains to run the command line with the following arguments:

* input_file path
* algorithm choice [dijkstra or astar]

For example:

```
$ go run go-dijkstra-astar input/graph.txt astar
```

Before running, make sure that the package is located in GOPATH and the input path is specified correctly.

### Package

The main package consists of the following files:

* main.go
* heuristic.go
* helper.go
* io.go
* edge.go
* queue.go
* test.go

The core functions are implemented in **main.go**, ```findShortestPath``` being the most important one, as the logic for both algorithms  is located in it.

```
func findShortestPath(start, finish int, algorithm string) { ... }
```

It takes either **“dijkstra”** or **“astar”** as its final argument, when the only difference is the calculation of the heuristic function and fCost.

Other files include: 

* **heuristic.go** which implements the heuristic function and finds euclidean distance between the starting and ending cells.
* **helper.go** includes helper functions (memset, printReverse, etc) to ease computation in other files.
* **io.go** functions take care of input and output operations.
* **edge.go** implements the Edge type with  to int and  distance float64 fields.
* **queue.go** uses the priority queue implementation library from https://programmer.help/blogs/simple-priority-queue-implemented-by-golang.html
* **test.go** implements the function for displaying all shortest path combinations. The result of the output files is achieved with its help.

### Heuristic
Heuristic function calculates the **euclidean distance** between vertices. After which, further normalization allows us to achieve better efficiency of A star over Dijkstra without overestimation.

### Input & Output
There are two folders in the repository: input and output. Input is **graph.txt** file. Program’s output is saved in **out.txt** with all the shortest path combinations given in the below format:

```
0 -> 17: Shortest path: [ 0 31 17 ] Cost: 7026
```

Meaning, the shortest path from 0 to 17 is 0 -> 31 -> 17 with the cost of 7026

**Cost** is the shortest distance between the start and end points. It’s the same for both Dijkstra and A*, as even if heuristic increases f(v), it does not affect the distance between vertices.

**comparison.txt** stores by how many vertices A* was more efficient than Dijkstra. Functions in **test.go** further visualizes different aspects of the program. 