package utilities

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Lines_to_list(path string) []string {
	var lines = []string{}

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadSolution(filepath string, solution int) {
	fmt.Printf("For input %s the solution is %d", filepath, solution)
}
