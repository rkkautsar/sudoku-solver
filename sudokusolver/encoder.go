package sudokusolver

import (
	"github.com/rkkautsar/sudoku-solver/sudoku"
)

func GenerateCNFConstraints(s *sudoku.Board) CNFInterface {
	var cnf CNFInterface

	shouldUseParallel := s.Size > 4

	cnf = &CNF{
		Board:   s,
		Clauses: make([][]int, 0, s.LenCells()*s.LenValues()*5),
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

	buildCNFRangeConstraints(cnf, cnfExactly1)
	buildCNFCellConstraints(cnf, cnfExactly1)
	// buildCNFRangeConstraints2(cnf, cnf.getBoard().Rows(), cnfExactly1)
	// buildCNFRangeConstraints2(cnf, cnf.getBoard().Columns(), cnfExactly1)
	// buildCNFRangeConstraints2(cnf, cnf.getBoard().Blocks(), cnfExactly1)

	if shouldUseParallel {
		cnf.(*CNFParallel).closeAndWait()
	}

	if s.Size > 6 {
		cnf.Simplify(SimplifyOptions{})
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

func buildCNFRangeConstraints2(
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

func buildCNFRangeConstraints(
	c CNFInterface,
	builder CNFBuilder,
) {
	size := c.getBoard().Size
	size2 := size * size

	triadAux := c.requestLiterals(2 * size * size2 * size2)
	getTriadAuxIdx := func(i, j, v, isCol int) int {
		return (v-1)*size*size2 + i*size + j + isCol*(size*size2*size2)
	}

	for v := 1; v <= size2; v++ {
		for i := 0; i < size2; i++ {
			vblkTriads := make([]int, size)
			hblkTriads := make([]int, size)
			blkRowStart := (i / size) * size
			blkColStart := (i % size) * size

			// row triad
			c.addFormula(triadAux[getTriadAuxIdx(i, 0, v, 0):getTriadAuxIdx(i, size, v, 0)], builder)
			// col triad
			c.addFormula(triadAux[getTriadAuxIdx(i, 0, v, 1):getTriadAuxIdx(i, size, v, 1)], builder)

			for j := 0; j < size; j++ {

				hblkTriads[j] = triadAux[getTriadAuxIdx(blkRowStart+j, i%size, v, 0)]
				vblkTriads[j] = triadAux[getTriadAuxIdx(blkColStart+j, i/size, v, 1)]
			}

			// block triads
			c.addFormula(hblkTriads, builder)
			c.addFormula(vblkTriads, builder)

			for j := 0; j < size; j++ {
				rowTriadLits := make([]int, size+1)
				colTriadLits := make([]int, size+1)
				for k := 0; k < size; k++ {
					rowTriadLits[k] = c.getBoard().GetLit(i, j*size+k, v)
					colTriadLits[k] = c.getBoard().GetLit(j*size+k, i, v)
				}
				rowTriadLits[size] = -triadAux[getTriadAuxIdx(i, j, v, 0)]
				colTriadLits[size] = -triadAux[getTriadAuxIdx(i, j, v, 1)]
				c.addFormula(rowTriadLits, builder)
				c.addFormula(colTriadLits, builder)
			}

			// rowLits := make([]int, size2)
			// colLits := make([]int, size2)
			// blkLits := make([]int, size2)
			// for j := 0; j < size2; j++ {
			// 	// row
			// 	rowLits[j] = c.getBoard().GetLit(i, j, val)

			// 	// col
			// 	colLits[j] = c.getBoard().GetLit(j, i, val)

			// 	// block
			// 	blkRow := blkRowStart + j/size
			// 	blkCol := blkColStart + j%size
			// 	blkLits[j] = c.getBoard().GetLit(blkRow, blkCol, val)
			// }

			// c.addFormula(blkLits, builder)
			// c.addFormula(rowLits, builder)
			// c.addFormula(colLits, builder)
		}
	}
}
