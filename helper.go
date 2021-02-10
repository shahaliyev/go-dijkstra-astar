/* HELPER FUNCTIONS */

package main

import (
	"fmt"
	"os"
	"strconv"
)

// Panics in case of errors
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Implementation of memset from C++
// Fills the whole array with a specific number
func memsetInt(arr []int, v int) {
	for i := range arr {
		arr[i] = v
	}
}

func memsetFloat(arr []float64, v int) {
	for i := range arr {
		arr[i] = float64(v)
	}
}

// Resets values
func reset() {
	memsetFloat(dist, INFINITY)
	memsetInt(parent, -1)
	visitCount = 0
}

// String to integer / float64 converter
// In order to not check for errors each time
func stringToInt(str string) (val int) {
	val, err := strconv.Atoi(str)
	check(err)

	return val
}

func stringToFloat(str string) (val float64) {
	val, err := strconv.ParseFloat(str, 64)
	check(err)

	return val
}

// Reads the input file for vertices and edges
func readInput(filePath string) *os.File {
	input, err := os.Open(filePath)
	check(err)

	return input
}

// Prints array in reverse
func printReverse(path []int) {
	for i := len(path)-1; i >= 0; i-- {
		fmt.Print(path[i], " ")
	}
}

// Checks if arrays are equal (for debugging purposes)
func isEqual(a, b []int) bool {
	if len(a) != len(b) { return false }
	for i, j := range a {
		if j != b[i] { return false }
	}
	return true
}
