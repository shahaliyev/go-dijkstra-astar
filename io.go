package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Reads the command line: input_file algorithm
func readCommandLine() (string, string) {
	if len(os.Args) != 3 {
		log.Fatalln("Please enter a correct command:\n" +
			"input_path algorithm [dijkstra or astar]\n" +
			"For example: input.txt dijkstra")
	}

	inputPath := os.Args[1]
	algorithm := os.Args[2]

	// error checking
	if (algorithm != "dijkstra" && algorithm != "astar") {
		log.Fatalln("Please enter the correct algorithm: either 'dijkstra' or 'astar'")
	}

	return inputPath, algorithm
}

// Builds graphs by scanning each line of the input file
func buildFromFile(input *os.File) (int, int) {
	var line string
	var start, finish int
	var section = "vertices"

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		// scanning next line
		line = scanner.Text()
		// skipping comments
		if strings.HasPrefix(line, "#") { continue }
		// switching sections whenever we encounter empty line
		if line == "" {
			if section == "vertices" { section = "edges" } else
			if section == "edges" { section = "test" }
			continue
		}

		// splitting comma separated lines into tokens
		tokens := strings.Split(line, ",")

		switch section {
		case "vertices":
			// getting vertex elements
			id := stringToInt(tokens[0])
			cell := stringToInt(tokens[1])
			// building vertex graph
			vertexGraph[id] = cell
			break
		case "edges":
			// getting edge elements
			from := stringToInt(tokens[0])
			to := stringToInt(tokens[1])
			distance := stringToFloat(tokens[2])
			// building undirected edge graph
			edgeGraph[from] = append(edgeGraph[from], Edge{to, distance})
			edgeGraph[to] = append(edgeGraph[to], Edge{from, distance})
			break
		case "test":
			// getting start and destination vertices
			if tokens[0] == "S" { start = stringToInt(tokens[1]) }
			if tokens[0] == "D" { finish = stringToInt(tokens[1]) }
			break
		}
	}
	// error checking
	err := scanner.Err(); check(err)

	return start, finish
}

// Displays the output
func displayOutput(start, finish int){
	if dist[finish] == INFINITY {
		fmt.Print("Impossible to reach the destination")
	} else {
		fmt.Print("Shortest path: ")
		// building path to display
		path := buildPath(start, finish)
		// printing out the reversed path
		printReverse(path)
		fmt.Println()
		fmt.Println("Cost:", math.Round(dist[finish]*100)/100)
		fmt.Println("Visited vertex count:", visitCount)
	}
}
