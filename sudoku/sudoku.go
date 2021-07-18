package sudoku

import (
	"math"
)

type Board struct {
	Size        int
	Known       []*Cell
	KnownLookup []int
}

func (s *Board) LenValues() int {
	return s.Size * s.Size
}

func (s *Board) LenRows() int {
	return s.LenValues()
}

func (s *Board) LenCols() int {
	return s.LenValues()
}

func (s *Board) LenCells() int {
	return s.LenValues() * s.LenValues()
}

func (s *Board) LenBlocks() int {
	return s.LenValues()
}

func (s *Board) Values() []int {
	values := make([]int, s.LenValues())
	for i := 1; i <= s.LenValues(); i++ {
		values[i-1] = i
	}
	return values
}

func (s *Board) Rows() [][]*Cell {
	rows := make([][]*Cell, s.LenRows())
	for rowIndex := 0; rowIndex < s.LenRows(); rowIndex++ {
		rows[rowIndex] = s.Row(rowIndex)
	}
	return rows
}

func (s *Board) Columns() [][]*Cell {
	cols := make([][]*Cell, s.LenCols())
	for colIndex := 0; colIndex < s.LenCols(); colIndex++ {
		cols[colIndex] = s.Column(colIndex)
	}
	return cols
}

func (s *Board) Blocks() [][]*Cell {
	blocks := make([][]*Cell, s.LenBlocks())
	for blkIndex := 0; blkIndex < s.LenBlocks(); blkIndex++ {
		blocks[blkIndex] = s.Block(blkIndex)
	}
	return blocks
}

func (s *Board) Cells() []*Cell {
	cells := make([]*Cell, s.LenCells())
	for _, row := range s.Rows() {
		for _, cell := range row {
			cells[cell.Index()] = cell
		}
	}
	return cells
}

func (s *Board) Row(rowIndex int) []*Cell {
	row := make([]*Cell, s.LenCols())
	for colIndex := 0; colIndex < s.LenCols(); colIndex++ {
		cell := s.NewCell(rowIndex, colIndex)
		idx := cell.Index()
		if s.KnownLookup != nil && idx < s.LenCells() && s.KnownLookup[idx] != 0 {
			cell.Value = s.KnownLookup[idx]
		}
		row[colIndex] = cell
	}
	return row
}

func (s *Board) Column(colIndex int) []*Cell {
	col := make([]*Cell, s.LenRows())
	for rowIndex := 0; rowIndex < s.LenRows(); rowIndex++ {
		cell := s.NewCell(rowIndex, colIndex)
		col[rowIndex] = cell
	}
	return col
}

// example indexing if size = 2
// 0 0 1 1
// 0 0 1 1
// 2 2 3 3
// 2 2 3 3
func (s *Board) Block(blkIndex int) []*Cell {
	rowStart := (blkIndex / s.Size) * s.Size
	colStart := (blkIndex % s.Size) * s.Size
	block := make([]*Cell, s.LenValues())

	idx := 0
	for rowIndex := rowStart; rowIndex < rowStart+s.Size; rowIndex++ {
		for colIndex := colStart; colIndex < colStart+s.Size; colIndex++ {
			cell := s.NewCell(rowIndex, colIndex)
			block[idx] = cell
			idx++
		}
	}
	return block
}

func (s *Board) NewCell(row int, col int) *Cell {
	return &Cell{Row: row, Col: col, size: s.Size}
}

func (s *Board) NewCellFromLit(lit int) *Cell {
	lit -= 1
	val := 1 + (lit % s.LenValues())
	lit /= s.LenValues()
	col := lit % s.LenCols()
	lit /= s.LenCols()
	row := lit
	return &Cell{Row: row, Col: col, Value: val, size: s.Size}
}

func (s *Board) GetLit(row int, col int, val int) int {
	if val <= 0 {
		panic("Value should not be <= 0")
	}

	cell := Cell{Row: row, Col: col, Value: val, size: s.Size}
	return cell.toInt()
}

func (s *Board) generateKnownLookup() {
	s.KnownLookup = make([]int, s.LenCells())
	for _, cell := range s.Known {
		if cell.Index() >= s.LenCells() {
			continue
		}
		s.KnownLookup[cell.Index()] = cell.Value
	}
}

func (s *Board) SolveWithModel(model []bool) {
	s.Known = make([]*Cell, 0, s.LenCells())

	for lit, val := range model {
		if !val {
			continue
		}

		s.Known = append(s.Known, s.NewCellFromLit(lit+1))
	}

	s.generateKnownLookup()
}

type Cell struct {
	Row   int // 0-based
	Col   int // 0-based
	Value int
	size  int
}

func (cell *Cell) Index() int {
	size2 := cell.size * cell.size
	return (cell.Row*size2 + cell.Col)
}

func (cell *Cell) toInt() int {
	size2 := cell.size * cell.size
	if cell.Value < 0 {
		panic("cell value < 0")
	}
	return 1 + cell.Row*(size2*size2) + cell.Col*size2 + (cell.Value - 1)
}

func (cell *Cell) BlockIndex() int {
	return (cell.Row/cell.size)*cell.size + (cell.Col / cell.size)
}

func getSize(size2 int) int {
	size := int(math.Sqrt(float64(size2)))
	if size2 != size*size {
		panic("Size is not a square")
	}
	return size
}
