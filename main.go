package main

import (
	"flag"
	"log"

	"github.com/bennyd312/adventofcode2023/day03"
	"github.com/bennyd312/adventofcode2023/utilities"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "path/to/file")
	flag.Parse()
	if len(path) == 0 {
		log.Fatal("Path not supplied")
	}

	solution := day03.Day03B(path)
	utilities.ReadSolution(path, solution)
}
