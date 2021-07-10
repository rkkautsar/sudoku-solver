package sudokusolver

import (
	"fmt"
	"io"

	"github.com/rkkautsar/sudoku-solver-2/sudoku"
)

type CNF struct {
	Board     *sudoku.SudokuBoard
	Clauses   [][]int
	lits      []int
	litLookup []uint8
	lookupLen int
	nbVar     int
}

type CNFBuilder = func(c *CNF, lits []int) [][]int

func (c *CNF) addLit(lit int) {
	c.lits = append(c.lits, lit)
	if lit < 0 && -lit <= c.lookupLen {
		c.litLookup[-lit-1] = 2
	} else if lit > 0 && lit <= c.lookupLen {
		c.litLookup[lit-1] = 1
	}
}

func (c *CNF) lookup(lit int) bool {
	if lit < 0 && -lit <= c.lookupLen {
		return c.litLookup[-lit-1] == 2
	} else if lit > 0 && lit <= c.lookupLen {
		return c.litLookup[lit-1] == 1
	}

	return false
}

func (c *CNF) addClause(clause []int) {
	c.Clauses = append(c.Clauses, clause)
}

func (c *CNF) addClauses(clauses [][]int) {
	c.Clauses = append(c.Clauses, clauses...)
}

func (c *CNF) Print(w io.Writer) {
	fmt.Fprintf(w, "p cnf %d %d\n", c.nbVar, len(c.Clauses))
	for _, c := range c.Clauses {
		for _, l := range c {
			fmt.Fprintf(w, "%d ", l)
		}
		fmt.Fprintln(w, 0)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
