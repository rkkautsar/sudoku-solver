package sudoku

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

/*
Parse newline and space separated sudoku problem
0 0 3 ...
9 0 0 ...
0 0 1 ...
...
*/
func NewFromString(input string) SudokuBoard {
	input = strings.Trim(input, " \n\t")
	space := regexp.MustCompile(`  +`)
	input = space.ReplaceAllString(input, " ")

	// standard 9x9 single row
	if strings.Index(input, "\n") == -1 {
		return NewFromSingleRowString(input)
	}

	rows := strings.Split(input, "\n")
	size2 := len(rows)
	cells := make([][]int, size2)

	for i, row := range rows {
		row = strings.Trim(row, " \t")
		// row = space.ReplaceAllString(row, " ")
		cols := strings.Split(row, " ")

		cells[i] = make([]int, size2)
		for j, col := range cols {
			if parsed, err := strconv.Atoi(col); err == nil {
				cells[i][j] = parsed
			}
		}
	}

	return NewFromArray(cells)
}

func NewFromSingleRowString(input string) SudokuBoard {
	size2 := 9
	size := 3
	known := []Cell{}

	for i, c := range input {
		if c != '0' && c != '.' {
			known = append(known, Cell{
				Row:   i / size2,
				Col:   i % size2,
				Value: int(c - '0'),
				size:  size,
			})
		}
	}

	board := SudokuBoard{
		Known: known,
		Size:  size,
	}

	return board
}

func NewFromArray(cells [][]int) SudokuBoard {
	size2 := len(cells)
	size := getSize(size2)

	known := []Cell{}

	for r, row := range cells {
		for c, val := range row {
			if val < 1 || val > size2 {
				continue
			}

			known = append(known, Cell{
				Row:   r,
				Col:   c,
				Value: val,
				size:  size,
			})

		}
	}

	board := SudokuBoard{
		Known: known,
		Size:  size,
	}

	board.generateKnownLookup()

	return board
}

func (s *SudokuBoard) Print() {
	charLen := int(math.Floor(math.Log10(float64(s.LenCells()))))
	formatter := fmt.Sprintf("%%%dd", charLen)

	for _, row := range s.Rows() {
		formatted := make([]string, s.LenCols())
		for i, cell := range row {
			formatted[i] = fmt.Sprintf(formatter, cell.Value)
		}
		fmt.Println(strings.Join(formatted, " "))
	}
}

func (s *SudokuBoard) PrintOneLine(w io.Writer) {
	for _, cell := range s.Cells() {
		fmt.Fprint(w, cell.Value)
	}
	fmt.Fprintln(w)
}
