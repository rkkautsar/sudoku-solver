package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLit(t *testing.T) {
	s := SudokuBoard{Size: 2}

	assert.Equal(t, 1, s.GetLit(0, 0, 1))
	assert.Equal(t, 4, s.GetLit(0, 0, 4))

	cell := s.NewCellFromLit(1)
	assert.Equal(t, 0, cell.Row)
	assert.Equal(t, 0, cell.Col)
	assert.Equal(t, 1, cell.Value)

	cell = s.NewCellFromLit(4)
	assert.Equal(t, 0, cell.Row)
	assert.Equal(t, 0, cell.Col)
	assert.Equal(t, 4, cell.Value)

	assert.Equal(t, 64, s.GetLit(3, 3, 4))
	cell = s.NewCellFromLit(64)
	assert.Equal(t, 3, cell.Row)
	assert.Equal(t, 3, cell.Col)
	assert.Equal(t, 4, cell.Value)
}
