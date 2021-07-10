package sudokusolver

import "github.com/rkkautsar/sudoku-solver-2/sudoku"

func GenerateCNFConstraints(s *sudoku.SudokuBoard) *CNF {
	cnf := CNF{
		Board:   s,
		Clauses: [][]int{},
	}

	cnf.generateLitLookup()

	cnf.nbVar = s.LenCells() * s.LenValues()

	for k := range cnf.LitLookup {
		cnf.addClause([]int{k})
	}

	cnf.addClauses(cnf.getCNFCellConstraints(cnfExactly1))
	cnf.addClauses(cnf.getCNFRangeConstraints(s.Rows(), cnfExactly1))
	cnf.addClauses(cnf.getCNFRangeConstraints(s.Columns(), cnfExactly1))
	cnf.addClauses(cnf.getCNFRangeConstraints(s.Blocks(), cnfExactly1))

	return &cnf
}

func (c *CNF) generateLitLookup() {
	c.LitLookup = make(map[int]bool, len(c.Board.Known)*c.Board.Size*c.Board.Size)

	for _, cell := range c.Board.Known {
		c.LitLookup[c.Board.GetLit(cell.Row, cell.Col, cell.Value)] = true

		ranges := [][]*sudoku.Cell{
			c.Board.Row(cell.Row),
			c.Board.Column(cell.Col),
			c.Board.Block(cell.BlockIndex()),
		}

		// negatives
		for _, val := range c.Board.Values() {
			if val != cell.Value {
				c.LitLookup[-c.Board.GetLit(cell.Row, cell.Col, val)] = true
			}
		}

		for _, r := range ranges {
			for _, i := range r {
				if i.Row == cell.Row && i.Col == cell.Col {
					continue
				}
				c.LitLookup[-c.Board.GetLit(i.Row, i.Col, cell.Value)] = true
			}
		}
	}
}

func (c *CNF) getCNFCellConstraints(builder CNFBuilder) [][]int {
	formula := [][]int{}

	for _, cell := range c.Board.Cells() {
		lits := []int{}
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
	formula := [][]int{}
	for _, cells := range list {
		for _, val := range c.Board.Values() {
			lits := []int{}
			for _, cell := range cells {

				lits = append(lits, c.Board.GetLit(cell.Row, cell.Col, val))
			}
			formula = append(formula, builder(c, lits)...)
		}
	}
	return formula
}
