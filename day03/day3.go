package day03

import (
	"log"
	"strconv"

	"github.com/bennyd312/adventofcode2023/utilities"
)

type symbol struct {
	row int
	col int
}

var numbers_map map[byte]struct{} = map[byte]struct{}{
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
	'0': {},
}

func getAdjacentCoords(x, y, maxRow, maxCol int) [][2]int {
	potential := [][2]int{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
	}

	var valid [][2]int

	for i := range potential {
		row := potential[i][0]
		col := potential[i][1]
		if -1 < row && row < maxRow && -1 < col && col < maxCol {
			valid = append(valid, potential[i])
		}
	}

	return valid
}

func getSchemaSymbols(lines []string, day03b ...bool) (map[[2]int]byte, []symbol) {
	var schema map[[2]int]byte = map[[2]int]byte{}
	symbols := []symbol{}
	for row := range lines {
		for col := range lines[0] {
			if lines[row][col] != '.' {
				if _, found := numbers_map[lines[row][col]]; found {
					schema[[2]int{row, col}] = lines[row][col]
				} else {
					if day03b[0] && lines[row][col] == '*' {
						symbols = append(symbols, symbol{row, col})
					} else {
						symbols = append(symbols, symbol{row, col})
					}
				}
			}
		}
	}
	return schema, symbols
}

func evaluateSchemaSimple(lines []string) int {

	score := 0
	schema, symbols := getSchemaSymbols(lines)

	for i := range symbols {
		adjacent := getAdjacentCoords(symbols[i].row, symbols[i].col, len(lines), len(lines[0]))
		for j := range adjacent {
			row := adjacent[j][0]
			col := adjacent[j][1]
			if _, found := schema[[2]int{row, col}]; found {
				val, start, end := getNum(lines[row], col)
				score += val
				for i := start; i < end+1; i++ {
					delete(schema, [2]int{row, i})
				}
			}
		}
	}

	return score
}

func mergeCoords(coords [][2]int) [][2]int {
	merged := [][2]int{}
	var rows map[int][]int = map[int][]int{}
	for i := range coords {
		row := coords[i][0]
		col := coords[i][1]

		rows[row] = append(rows[row], col)
	}

	for k, v := range rows {

		switch len(v) {
		case 3:
			merged = append(merged, [2]int{k, v[0]})
		case 2:
			if min(v[0], v[1])+1 != max(v[0], v[1]) {
				merged = append(merged, [2]int{k, v[0]})
			}
			merged = append(merged, [2]int{k, v[1]})
		default:
			merged = append(merged, [2]int{k, v[0]})
		}
	}
	return merged
}

func evaluateSchemaGears(lines []string) int {
	score := 0

	schema, symbols := getSchemaSymbols(lines, true)

	for i := range symbols {
		symbol := symbols[i]
		adjacent := getAdjacentCoords(symbol.row, symbol.col, len(lines), len(lines[0]))
		adjValid := [][2]int{}
		for j := range adjacent {
			row := adjacent[j][0]
			col := adjacent[j][1]
			if _, found := schema[[2]int{row, col}]; found {
				adjValid = append(adjValid, [2]int{row, col})
			}

		}
		parts := mergeCoords(adjValid)
		if len(parts) == 2 {
			row_1 := parts[0][0]
			col_1 := parts[0][1]
			row_2 := parts[1][0]
			col_2 := parts[1][1]
			val_1, _, _ := getNum(lines[row_1], col_1)
			val_2, _, _ := getNum(lines[row_2], col_2)
			score += val_1 * val_2
		}
	}
	return score
}

func getNum(line string, x int) (int, int, int) {
	start, end := x, x
	for {
		if start == 0 {
			break
		}
		if _, found := numbers_map[line[start-1]]; found {
			start--
		} else {
			break
		}
	}
	for {
		if end == len(line)-1 {
			break
		}
		if _, found := numbers_map[line[end+1]]; found {
			end++
		} else {
			break
		}
	}
	val, err := strconv.Atoi(line[start : end+1])
	if err != nil {
		log.Fatal(err)
	}
	return val, start, end
}

func evaluateSchema(lines []string) int {
	score := 0

	var schema map[[2]int]byte = map[[2]int]byte{}
	symbols := []symbol{}
	for row := range lines {
		for col := range lines[0] {
			if lines[row][col] != '.' {
				if _, found := numbers_map[lines[row][col]]; !found {
					symbols = append(symbols, symbol{row, col})
				} else {
					schema[[2]int{row, col}] = lines[row][col]
				}
			}
		}
	}
	for i := range symbols {
		symbol := symbols[i]
		//horizontal pass
		for col := range symbol.col + 1 {
			if _, found := schema[[2]int{symbol.row, symbol.col - col}]; found {
				value, start, end := getNum(lines[symbol.row], symbol.col-col)
				score += value
				for i := start; i < end+1; i++ {
					delete(schema, [2]int{symbol.row, i})
				}
				break
			}
		}
		for col := range len(lines[0]) - symbol.col {
			if _, found := schema[[2]int{symbol.row, symbol.col + col}]; found {
				value, start, end := getNum(lines[symbol.row], symbol.col+col)
				score += value
				for i := start; i < end+1; i++ {
					delete(schema, [2]int{symbol.row, i})
				}
				break
			}
		}

		//vertical pass
		for row := range symbol.row {
			if _, found := schema[[2]int{symbol.row - row, symbol.col}]; found {
				value, start, end := getNum(lines[symbol.row-row], symbol.col)
				score += value
				for i := start; i < end+1; i++ {
					delete(schema, [2]int{symbol.row - row, i})
				}
				break
			}
		}
		for row := range len(lines) - symbol.row {
			if _, found := schema[[2]int{symbol.row + row, symbol.col}]; found {
				value, start, end := getNum(lines[symbol.row+row], symbol.col)
				score += value
				for i := start; i < end+1; i++ {
					delete(schema, [2]int{symbol.row + row, i})
				}
				break
			}
		}
		//diagonal pass
		counter := 0
		maxCol := len(lines[0])
		maxRow := len(lines)
		for {
			if (symbol.col-counter == 0) || (symbol.row-counter == 0) {
				break
			}
			if _, found := schema[[2]int{symbol.row - counter, symbol.col - counter}]; found {
				value, x_start, x_end := getNum(lines[symbol.row-counter], symbol.col-counter)
				score += value
				for i := x_start; i < x_end+1; i++ {
					delete(schema, [2]int{symbol.row - counter, i})
				}
				break
			}
			counter++
		}
		counter = 0
		for {
			if (symbol.col+counter == maxCol-1) || (symbol.row+counter == maxRow-1) {
				break
			}
			if _, found := schema[[2]int{symbol.row + counter, symbol.col + counter}]; found {
				value, x_start, x_end := getNum(lines[symbol.row+counter], symbol.col+counter)
				score += value
				for i := x_start; i < x_end+1; i++ {
					delete(schema, [2]int{symbol.row + counter, i})
				}
				break
			}
			counter++
		}
		counter = 0
		//antidiagonal pass
		for {
			if (symbol.row-counter == 0) || (symbol.col+counter == maxCol-1) {
				break
			}
			if _, found := schema[[2]int{symbol.row - counter, symbol.col + counter}]; found {
				value, x_start, x_end := getNum(lines[symbol.row-counter], symbol.col+counter)
				score += value
				for i := x_start; i < x_end+1; i++ {
					delete(schema, [2]int{symbol.row - counter, i})
				}
				break
			}
			counter++
		}
		counter = 0
		for {
			if (symbol.row+counter == maxRow) || (symbol.col-counter == 0) {
				break
			}
			if _, found := schema[[2]int{symbol.row + counter, symbol.col - counter}]; found {
				value, x_start, x_end := getNum(lines[symbol.row+counter], symbol.col-counter)
				score += value
				for i := x_start; i < x_end+1; i++ {
					delete(schema, [2]int{symbol.row + counter, i})
				}
				break
			}
			counter++
		}
	}
	return score
}

func Day03A(filepath string) int {
	lines := utilities.Lines_to_list(filepath)
	solution := evaluateSchemaSimple(lines)

	return solution
}

func Day03B(filepath string) int {
	lines := utilities.Lines_to_list(filepath)
	solution := evaluateSchemaGears(lines)

	return solution
}
