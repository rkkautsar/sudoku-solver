package sudokusolver

import (
	"github.com/rkkautsar/sudoku-solver/sudoku"
)

func GenerateCNFConstraints(s *sudoku.SudokuBoard) CNFInterface {
	var cnf CNFInterface

	shouldUseParallel := s.Size > 4

	cnf = &CNF{
		Board:   s,
		Clauses: make([][]int, 0, s.LenCells()*s.LenValues()*10),
	}

	if shouldUseParallel {
		cnf = &CNFParallel{
			CNF: cnf.(*CNF),
		}
	}

	cnf.setInitialNbVar(s.LenCells() * s.LenValues())
	cnf.initializeLits()

	// var b bytes.Buffer
	// cnf.Print(&b)
	// log.Println(b.String())

	if shouldUseParallel {
		cnf.(*CNFParallel).initWorkers()
	}

	buildCNFCellConstraints(cnf, cnfExactly1)
	buildCNFRangeConstraints(cnf, cnfExactly1)

	if shouldUseParallel {
		cnf.(*CNFParallel).closeAndWait()
	}

	return cnf
}

func (c *CNF) initializeLits() {
	c.lits = make([]int, 0, c.Board.LenCells()*c.Board.LenValues())
	c.litLookup = make([]uint8, c.lookupLen)

	for _, cell := range c.Board.Known {
		row, col, value, size := cell.Row, cell.Col, cell.Value, c.Board.Size

		lit := c.Board.GetLit(row, col, value)
		c.addLit(lit)

		blkIndex := cell.BlockIndex()
		blkRowStart := (blkIndex / size) * size
		blkColStart := (blkIndex % size) * size

		// negatives
		for i := 0; i < size*size; i++ {
			if i+1 != value {
				c.addLit(-c.Board.GetLit(row, col, i+1))
			}
			if i != row {
				c.addLit(-c.Board.GetLit(i, col, value))
			}
			if i != col {
				c.addLit(-c.Board.GetLit(row, i, value))
			}
			blkRow := blkRowStart + i/size
			blkCol := blkColStart + i%size
			if blkRow != row && blkCol != col {
				c.addLit(-c.Board.GetLit(blkRow, blkCol, value))
			}
		}
	}
}

func (c *CNFParallel) initializeLits() {
	c.CNF.initializeLits()
}

func buildCNFCellConstraints(c CNFInterface, builder CNFBuilder) {
	for _, cell := range c.getBoard().Cells() {
		lits := make([]int, 0, c.getBoard().LenValues())
		for val := 1; val <= c.getBoard().LenValues(); val++ {
			lits = append(lits, c.getBoard().GetLit(cell.Row, cell.Col, val))
		}
		c.addFormula(lits, builder)
	}
}

func buildCNFRangeConstraints(
	c CNFInterface,
	builder CNFBuilder,
) {
	size := c.getBoard().Size
	size2 := size * size

	for val := 1; val <= size2; val++ {
		for i := 0; i < size2; i++ {

			blkRowStart := (i / size) * size
			blkColStart := (i % size) * size

			rowLits := make([]int, size2)
			colLits := make([]int, size2)
			blkLits := make([]int, size2)
			for j := 0; j < size2; j++ {
				// row
				rowLits[j] = c.getBoard().GetLit(i, j, val)

				// col
				colLits[j] = c.getBoard().GetLit(j, i, val)

				// block
				blkRow := blkRowStart + j/size
				blkCol := blkColStart + j%size
				blkLits[j] = c.getBoard().GetLit(blkRow, blkCol, val)
			}

			c.addFormula(blkLits, builder)
			c.addFormula(rowLits, builder)
			c.addFormula(colLits, builder)
		}
	}
}
