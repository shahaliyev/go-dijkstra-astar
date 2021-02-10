package main

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
