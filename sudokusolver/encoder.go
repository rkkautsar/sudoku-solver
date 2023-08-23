package sudokusolver

import (
	"github.com/irifrance/gini"
	"github.com/rkkautsar/sudoku-solver/sudoku"
)

func GenerateCNFConstraints(s *sudoku.Board, g *gini.Gini) CNFInterface {
	var cnf CNFInterface

	s.InitCompressedLits()
	cnf = &CNF{
		Board: s,
		g:     g,
		nbVar: uint32(s.NumCandidates),
	}
	initializeLits(cnf)
	buildCNFCellConstraints(cnf, cnfExactly1)
	buildCNFRangeConstraints(cnf, cnfExactly1)

	return cnf
}

func initializeLits(c CNFInterface) {
	b := c.getBoard()
	for i := 0; i < len(b.Lookup); i++ {
		if b.Lookup[i] != 0 {
			c.addLit(b.CLit(i/b.Size2, i%b.Size2, b.Lookup[i]))
		}
	}
}

func buildCNFCellConstraints(cnf CNFInterface, builder CNFBuilder) {
	b := cnf.getBoard()
	idx := 0
	for r := 0; r < b.Size2; r++ {
		for c := 0; c < b.Size2; c++ {
			if b.Lookup[idx] != 0 {
				idx++
				continue
			}
			idx++
			lits := make([]int, b.Size2)
			for v := 1; v <= b.Size2; v++ {
				lits[v-1] = b.CLit(r, c, v)
			}
			cnf.addFormula(filterZero(lits), builder)
		}
	}
}

func buildCNFRangeConstraints(
	c CNFInterface,
	builder CNFBuilder,
) {
	b := c.getBoard()
	size := b.Size
	size2 := b.Size2

	for v := 1; v <= size2; v++ {
		for i := 0; i < size2; i++ {

			blkRowStart := (i / size) * size
			blkColStart := (i % size) * size
			rowLits := make([]int, size2)
			colLits := make([]int, size2)
			blkLits := make([]int, size2)
			for j := 0; j < size2; j++ {
				// block
				blkRow := blkRowStart + j/size
				blkCol := blkColStart + j%size
				blkLits[j] = b.CLit(blkRow, blkCol, v)

				// row
				rowLits[j] = b.CLit(i, j, v)

				// col
				colLits[j] = b.CLit(j, i, v)
			}

			c.addFormula(filterZero(blkLits), builder)
			c.addFormula(filterZero(rowLits), builder)
			c.addFormula(filterZero(colLits), builder)
		}
	}
}

func filterZero(slice []int) []int {
	newSlice := slice[:0]
	for _, x := range slice {
		if x != 0 {
			newSlice = append(newSlice, x)
		}
	}
	return newSlice
}
