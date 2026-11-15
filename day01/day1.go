package day01

import (
	"github.com/bennyd312/adventofcode2023/utilities"
)

var numbers map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getOrdinarySolution(line string) int {
	line_len := len(line)
	var calibration int
	ptr := 0

	//front pass
	for true {
		if line[ptr] > 47 && line[ptr] < 58 {
			calibration += int(line[ptr]-'0') * 10
			ptr = 0
			break
		} else {
			ptr++
		}
	}
	//back pass
	for true {
		if line[line_len-1-ptr] > 47 && line[line_len-1-ptr] < 58 {
			calibration += int(line[line_len-1-ptr] - '0')
			break
		} else {
			ptr++
		}
	}
	return calibration
}

func getCombinedSolution(line string) int {
	n := len(line)
	calibration := 0
	//front pass
	for i := range line {
		if line[i] > 47 && line[i] < 58 {
			calibration += int(line[i]-'0') * 10
			break
		}
		if i > 1 {
			if num := numbers[line[i-2:i+1]]; num != 0 {
				calibration += num * 10
				break
			}
		}
		if i > 2 {
			if num := numbers[line[i-3:i+1]]; num != 0 {
				calibration += num * 10
				break
			}
		}
		if i > 3 {
			if num := numbers[line[i-4:i+1]]; num != 0 {
				calibration += num * 10
				break
			}
		}
	}

	for i := range line {
		if line[n-1-i] > 47 && line[n-1-i] < 58 {
			calibration += int(line[n-1-i] - '0')
			break
		}
		if i > 1 {
			if num := numbers[line[n-1-i:n-1-i+3]]; num != 0 {
				calibration += num
				break
			}
		}
		if i > 2 {
			if num := numbers[line[n-1-i:n-1-i+4]]; num != 0 {
				calibration += num
				break
			}
		}
		if i > 3 {
			if num := numbers[line[n-1-i:n-1-i+5]]; num != 0 {
				calibration += num
				break
			}
		}
	}
	return calibration
}

func Day01A(filepath string) int {
	solution := 0

	lines := utilities.Lines_to_list(filepath)
	n := len(lines)
	for i := range n {
		solution += getOrdinarySolution(lines[i])
	}
	return solution
}

func Day01B(filepath string) int {
	solution := 0

	lines := utilities.Lines_to_list(filepath)

	for i := range lines {
		line := lines[i]
		if len(line) < 2 {
			solution += getOrdinarySolution(line)
		} else {
			solution += getCombinedSolution(line)
		}
	}
	return solution
}
