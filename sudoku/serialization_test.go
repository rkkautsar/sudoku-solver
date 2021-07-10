package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFromString(t *testing.T) {
	exampleSudoku := `
	0 0 3 0 2 0 6 0 0
	9 0 0 3 0 5 0 0 1
	0 0 1 8 0 6 4 0 0
	0 0 8 1 0 2 9 0 0
	7 0 0 0 0 0 0 0 8
	0 0 6 7 0 8 2 0 0
	0 0 2 6 0 9 5 0 0
	8 0 0 2 0 3 0 0 9
	0 0 5 0 1 0 3 0 0`

	sudokuBoard := NewFromString(exampleSudoku)

	assert.Equal(t, 3, sudokuBoard.Size)
	assert.Equal(t, 32, len(sudokuBoard.Known))
}

func TestParseFromSingleRowString(t *testing.T) {
	exampleSudoku := "........8..3...4...9..2..6.....79.......612...6.5.2.7...8...5...1.....2.4.5.....3"
	sudokuBoard := NewFromString(exampleSudoku)

	assert.Equal(t, 3, sudokuBoard.Size)
	assert.Equal(t, 22, len(sudokuBoard.Known))
}
