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

	if shouldUseParallel {
		cnf.(*CNFParallel).initWorkers()
	}

	getCNFCellConstraints(cnf, cnfExactly1)
	getCNFRangeConstraints(cnf, s.Rows(), cnfExactly1)
	getCNFRangeConstraints(cnf, s.Columns(), cnfExactly1)
	getCNFRangeConstraints(cnf, s.Blocks(), cnfExactly1)

	if shouldUseParallel {
		cnf.(*CNFParallel).closeAndWait()
	}

	return cnf
}

func (c *CNF) initializeLits() {
	c.lits = make([]int, 0, len(c.Board.Known)*c.Board.LenValues()*4)
	c.litLookup = make([]uint8, c.lookupLen)

	for _, cell := range c.Board.Known {
		c.addLit(c.Board.GetLit(cell.Row, cell.Col, cell.Value))

		ranges := [][]*sudoku.Cell{
			c.Board.Row(cell.Row),
			c.Board.Column(cell.Col),
			c.Board.Block(cell.BlockIndex()),
		}

		// negatives
		for _, val := range c.Board.Values() {
			if val != cell.Value {
				c.addLit(-c.Board.GetLit(cell.Row, cell.Col, val))
			}
		}

		for _, r := range ranges {
			for _, i := range r {
				if i.Row == cell.Row && i.Col == cell.Col {
					continue
				}
				c.addLit(-c.Board.GetLit(i.Row, i.Col, cell.Value))
			}
		}
	}
}

func (c *CNFParallel) initializeLits() {
	c.CNF.initializeLits()
}

func getCNFCellConstraints(c CNFInterface, builder CNFBuilder) {
	for _, cell := range c.getBoard().Cells() {
		lits := make([]int, 0, c.getBoard().LenValues())
		for val := 1; val <= c.getBoard().LenValues(); val++ {
			lits = append(lits, c.getBoard().GetLit(cell.Row, cell.Col, val))
		}
		c.addFormula(lits, builder)
	}
}

func getCNFRangeConstraints(
	c CNFInterface,
	list [][]*sudoku.Cell,
	builder CNFBuilder,
) {
	for _, cells := range list {
		for _, val := range c.getBoard().Values() {
			lits := make([]int, 0, len(list))
			for _, cell := range cells {
				lits = append(lits, c.getBoard().GetLit(cell.Row, cell.Col, val))
			}
			c.addFormula(lits, builder)
		}
	}
}
