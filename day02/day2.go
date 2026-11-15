package day02

import (
	"log"
	"strconv"
	"strings"

	"github.com/bennyd312/adventofcode2023/utilities"
)

func evaluateResultA(set [3]int) int {
	red := 13
	green := 14
	blue := 15

	if !(set[0] < red && set[1] < green && set[2] < blue) {
		return 0
	}
	return 1
}

func scoringResultB(sets [][3]int) int {
	r, g, b := 0, 0, 0

	for i := range sets {
		r = max(r, sets[i][0])
		g = max(g, sets[i][1])
		b = max(b, sets[i][2])
	}
	return r * g * b
}

func parseSet(set string) [3]int {

	r, g, b := 0, 0, 0

	substrings := strings.Split(set, ", ")

	for i := range substrings {
		roll := strings.Split(substrings[i], " ")

		switch color := roll[1]; color {
		case "red":
			num, err := strconv.Atoi(roll[0])
			if err != nil {
				log.Fatal(err)
			}
			r = num
		case "green":
			num, err := strconv.Atoi(roll[0])
			if err != nil {
				log.Fatal(err)
			}
			g = num
		case "blue":
			num, err := strconv.Atoi(roll[0])
			if err != nil {
				log.Fatal(err)
			}
			b = num

		}

	}
	return [3]int{r, g, b}
}

func parseGame(line string) (int, [][3]int) {

	var clean_sets [][3]int
	split_str := strings.Split(line, ": ")
	//first, get game id
	game_id, err := strconv.Atoi(strings.Replace(split_str[0], "Game ", "", -1))
	if err != nil {
		log.Fatal(err)
	}
	sets := strings.Split(split_str[1], "; ")
	for i := range sets {
		set := parseSet(sets[i])
		clean_sets = append(clean_sets, set)
	}
	return game_id, clean_sets
}

func evaluateSetA(game_id int, sets [][3]int) int {
	for i := range sets {
		if valid := evaluateResultA(sets[i]); valid == 0 {
			return 0
		}
	}
	return game_id
}

func evaluateSetB(sets [][3]int) int {
	return scoringResultB(sets)
}

func Day02A(filepath string) int {
	solution := 0

	lines := utilities.Lines_to_list(filepath)

	for i := range lines {
		game_id, sets := parseGame(lines[i])
		solution += evaluateSetA(game_id, sets)
	}

	return solution
}

func Day02B(filepath string) int {
	solution := 0

	lines := utilities.Lines_to_list(filepath)
	for i := range lines {
		_, sets := parseGame(lines[i])
		solution += evaluateSetB(sets)
	}

	return solution
}
