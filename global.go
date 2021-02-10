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
