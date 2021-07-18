package sudoku

import (
	"math"
)

type Board struct {
	Size       int
	Size2      int
	Candidates []bool // lit
	Lookup     []int  // idx

	NumCandidates int
	lit_cLit      []int // lit -> compressed lit, 1-indexed
	cLit_lit      []int // compressed lit -> lit, 1-indexed
}

func New(size int) *Board {
	size2 := size * size
	candidates := make([]bool, size2*size2*size2+1)
	board := &Board{
		Size:   size,
		Size2:  size2,
		Lookup: make([]int, size2*size2),

		NumCandidates: len(candidates) - 1,
	}

	for i := 1; i < len(candidates); i++ {
		candidates[i] = true
	}

	board.Candidates = candidates

	return board
}

func (b *Board) SetValue(row, col, val int, truth bool) {
	if !truth {
		lit := b.Lit(row, col, val)
		prev := b.Candidates[lit]
		b.Candidates[lit] = false
		if prev {
			b.NumCandidates--
		}
		return
	}

	b.Lookup[b.Idx(row, col)] = val
	blkRStart := b.Size * (row / b.Size)
	blkCStart := b.Size * (col / b.Size)

	for i := 0; i < b.Size2; i++ {
		if i+1 != val {
			b.SetValue(row, col, i+1, false)
		}
		if i != row {
			b.SetValue(i, col, val, false)
		}
		if i != col {
			b.SetValue(row, i, val, false)
		}
		blkR := blkRStart + i/b.Size
		blkC := blkCStart + i%b.Size
		if blkR != row && blkC != col {
			b.SetValue(blkR, blkC, val, false)
		}
	}
}

func (b *Board) InitCompressedLits() {
	b.lit_cLit = make([]int, len(b.Candidates)+1)
	b.cLit_lit = make([]int, b.NumCandidates+1)
	j := 1
	for i := 1; i < len(b.Candidates); i++ {
		if !b.Candidates[i] {
			continue
		}
		b.lit_cLit[i] = j
		b.cLit_lit[j] = i
		j++
	}
}

// from model of compressed lits
func (b *Board) SolveWithModel(model []bool) {
	// for i := 1; i < len(b.Candidates); i++ {
	//   if !model[i-1] {
	//     continue
	//   }
	//   lit := b.cLit_lit[i]
	//   b.Lookup[(lit-1)/b.Size2] = 1 + (lit-1)%b.Size2
	// }
	for i := 0; i < min(len(model), len(b.cLit_lit)-1); i++ {
		if !model[i] {
			continue
		}
		lit := b.cLit_lit[i+1]
		// log.Println(i+1, lit, (lit-1)/b.Size2, 1+(lit-1)%b.Size2)
		b.Lookup[(lit-1)/b.Size2] = 1 + (lit-1)%b.Size2
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 1-indexed
func (b *Board) Lit(row, col, val int) int {
	return 1 + b.Idx(row, col)*b.Size2 + (val - 1)
}

// 1-indexed
func (b *Board) CLit(row, col, val int) int {
	return b.lit_cLit[b.Lit(row, col, val)]
}

// 0-indexed
func (b *Board) Idx(row, col int) int {
	return row*b.Size2 + col
}

func getSize(size2 int) int {
	size := int(math.Sqrt(float64(size2)))
	if size2 != size*size {
		panic("Size is not a square")
	}
	return size
}
