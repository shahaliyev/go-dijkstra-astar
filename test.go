package main

import (
	"fmt"
	"math"
)

// Nested loop wrapper for printing all values
func printAll(testSize int, f func(start, finish int)) {
	for i := 0; i < testSize; i++ {
		for j := 0; j < SIZE; j++ {
			f(i , j)
		}
	}
}

// Output will be the same for dijkstra
func testShortestPath(start, finish int) {
	findShortestPath(start, finish, "astar")
	path := buildPath(start, finish)

	fmt.Print(start, " -> ", finish, ": ")
	fmt.Print("Shortest path: ")
	fmt.Print("[ "); printReverse(path); fmt.Print("] ")
	fmt.Println("Cost:", math.Round(dist[finish]*100)/100)
}

// Prints the number of vertices visited by both algorithms
var visitD, visitA int
var visitCnt, visitDiff, visitDiffSum int
func testVisitCount(start, finish int) {
	findShortestPath(start, finish, "dijkstra")
	visitD = visitCount

	findShortestPath(start, finish, "astar")
	visitA = visitCount

	// calculating difference, sum, and the number of operations
	visitDiff = visitD - visitA
	visitDiffSum += visitDiff
	visitCnt++

	fmt.Print(start, " -> ", finish, ": ")
	fmt.Println("A* visited", visitDiff, "vertices less")
}

// Checks if there are different paths between algorithms
// Ideally, diffCount should be zero
var pathA, pathD []int
var diffCount int
func testDiff(start, finish int){
	findShortestPath(start, finish, "dijkstra")
	pathD = buildPath(start, finish)

	findShortestPath(start, finish, "astar")
	pathA = buildPath(start, finish)

	// increasing the counter if paths are not the same
	if !isEqual(pathD, pathA) { diffCount++ }
}

// Test cases (uncomment/modify any)
func test(){
	//testShortestPath(4, 47)
	//printAll(SIZE, testShortestPath)

	//printAll(SIZE, testVisitCount); fmt.Println()
	//fmt.Println("Average:", visitDiffSum/visitCnt, "(rounded down)")

	//printAll(SIZE, testDiff); fmt.Println(diffCount, "paths are suboptimal")
}

