package sudokusolver

import "github.com/rkkautsar/sudoku-solver-2/sudoku"

func GenerateCNFConstraints(s *sudoku.SudokuBoard) *CNF {
	cnf := CNF{
		Board:   s,
		Clauses: make([][]int, 0, s.LenCells()*s.LenValues()*10),
	}

	cnf.nbVar = s.LenCells() * s.LenValues()
	cnf.lookupLen = cnf.nbVar
	cnf.generateLitLookup()

	for _, k := range cnf.lits {
		cnf.addClause([]int{k})
	}

	cnf.addClauses(cnf.getCNFCellConstraints(cnfExactly1))
	cnf.addClauses(cnf.getCNFRangeConstraints(s.Rows(), cnfExactly1))
	cnf.addClauses(cnf.getCNFRangeConstraints(s.Columns(), cnfExactly1))
	cnf.addClauses(cnf.getCNFRangeConstraints(s.Blocks(), cnfExactly1))

	return &cnf
}

func (c *CNF) generateLitLookup() {
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

func (c *CNF) getCNFCellConstraints(builder CNFBuilder) [][]int {
	formula := make([][]int, 0, c.Board.LenCells()*c.Board.LenCells()/2)

	for _, cell := range c.Board.Cells() {
		lits := make([]int, 0, c.Board.LenValues())
		for val := 1; val <= c.Board.LenValues(); val++ {
			lits = append(lits, c.Board.GetLit(cell.Row, cell.Col, val))
		}
		formula = append(formula, builder(c, lits)...)
	}

	return formula
}

func (c *CNF) getCNFRangeConstraints(
	list [][]*sudoku.Cell,
	builder CNFBuilder,
) [][]int {
	formula := make([][]int, 0, len(list)*len(list)/2)
	for _, cells := range list {
		for _, val := range c.Board.Values() {
			lits := make([]int, 0, len(list))
			for _, cell := range cells {
				lits = append(lits, c.Board.GetLit(cell.Row, cell.Col, val))
			}
			formula = append(formula, builder(c, lits)...)
		}
	}
	return formula
}
