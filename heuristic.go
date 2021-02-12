package main

import (
	"math"
)

// Given a cell number, finds its location relative to the vertex zero
func location(cellNumber int) (int, int) {
	var locX, locY int

	if (cellNumber / 10 != 0) {
		locX = cellNumber % 10
		locY = cellNumber / 10
	} else {
		locX = cellNumber
	}

	return locX, locY
}

// Calculates euclidean distance
func findEuclid(start, finish float64) float64 {
	return math.Sqrt(start * start + finish * finish)
}

// Heuristic for finding the euclidean distance between start and end vertices
func heuristic(start, finish int) float64 {
	var startLocX, startLocY int
	var finishLocX, finishLocY int
	var distX, distY float64
	var euclidean float64

	// normalizes euclidean distance to achieve better efficiency
	// for the test case, any number above 0.33 leads to suboptimal solutions
	var normalize = 330.0 / CELLSIZE

	// getting cell numbers from graph
	startCell := vertexGraph[start]
	finishCell := vertexGraph[finish]

	// if it is the same cell, euclidean = 0
	if start == finish { return euclidean }

	// if are in the same cell, approximating the distance
	if startCell == finishCell { return CELLSIZE / 10.0 }

	// finding cell locations relative to vertex zero
	startLocX, startLocY = location(startCell)
	finishLocX, finishLocY = location(finishCell)

	// calculating distance along x and y axes
	distX = math.Abs(float64(finishLocX - startLocX))
	distY = math.Abs(float64(finishLocY - startLocY))

	// calculating euclidean distance
	euclidean = findEuclid(distX, distY) * CELLSIZE * normalize

	return euclidean
}
