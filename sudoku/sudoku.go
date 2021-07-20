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
		Size:       size,
		Size2:      size2,
		Lookup:     make([]int, size2*size2),
		Candidates: candidates,

		NumCandidates: len(candidates) - 1,
	}

	for i := 1; i < len(candidates); i++ {
		candidates[i] = true
	}

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

func (b *Board) BasicSolve() {
	restart := true
	for restart {
		restart = false
		restart = restart || b.NakedSingles()
		restart = restart || b.HiddenSingles()
	}
}

func (b *Board) NakedSingles() bool {
	restart := false
	for r := 0; r < b.Size2; r++ {
		for c := 0; c < b.Size2; c++ {
			if b.Lookup[b.Idx(r, c)] != 0 {
				continue
			}

			// naked singles
			last := 0
			for v := 1; v <= b.Size2; v++ {
				if !b.Candidates[b.Lit(r, c, v)] {
					continue
				}
				if last != 0 {
					last = 0
					break
				}
				last = v
			}
			if last != 0 {
				b.SetValue(r, c, last, true)
				restart = true
			}
		}
	}
	return restart
}

func (b *Board) HiddenSingles() bool {
	restart := false
	for v := 1; v <= b.Size2; v++ {
		for i := 0; i < b.Size2; i++ {
			var (
				count [3]int
				last  [3]int
			)
			blkRStart := (i / b.Size) * b.Size
			blkCStart := (i % b.Size) * b.Size

			for j := 0; j < b.Size2; j++ {
				// row
				if b.Candidates[b.Lit(i, j, v)] && b.Lookup[b.Idx(i, j)] != v {
					count[0] += 1
					last[0] = j
				}
				// col
				if b.Candidates[b.Lit(j, i, v)] && b.Lookup[b.Idx(j, i)] != v {
					count[1] += 1
					last[1] = j
				}
				// block
				blkR := blkRStart + j/b.Size
				blkC := blkCStart + j%b.Size
				if b.Candidates[b.Lit(blkR, blkC, v)] && b.Lookup[b.Idx(blkR, blkC)] != v {
					count[2] += 1
					last[2] = j
				}
			}

			if count[0] == 1 {
				b.SetValue(i, last[0], v, true)
				restart = true
			}
			if count[1] == 1 {
				b.SetValue(last[1], i, v, true)
				restart = true
			}
			if count[2] == 1 {
				blkR := blkRStart + last[2]/b.Size
				blkC := blkCStart + last[2]%b.Size
				b.SetValue(blkR, blkC, v, true)
				restart = true
			}
		}
	}

	return restart
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
