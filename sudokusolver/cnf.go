package sudokusolver

import (
	"fmt"
	"io"

	"github.com/rkkautsar/sudoku-solver/sudoku"
)

type CNFInterface interface {
	addLit(lit int)
	addClause(clause []int)
	addClauses(clauses [][]int)
	addFormula(lits []int, builder CNFBuilder)
	initializeLits()
	lookup(lit int) bool
	requestLiterals(num int) []int
	setInitialNbVar(int)
	getBoard() *sudoku.SudokuBoard
	getLits() []int
	getClauses() [][]int
	Print(w io.Writer)
}

type CNF struct {
	CNFInterface
	Board     *sudoku.SudokuBoard
	Clauses   [][]int
	lits      []int
	litLookup []uint8
	lookupLen int
	nbVar     int
}

type CNFBuilder = func(c CNFInterface, lits []int) [][]int

func (c *CNF) addLit(lit int) {
	c._addLit(lit)
	c.addClause([]int{lit})
}

func (c *CNF) _addLit(lit int) {
	c.lits = append(c.lits, lit)
	c.addClause([]int{lit})
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

func (c *CNF) requestLiterals(num int) []int {
	lits := makeRange(c.nbVar+1, c.nbVar+num)
	c.nbVar += num
	return lits
}

func (c *CNF) setInitialNbVar(nbVar int) {
	c.nbVar = nbVar
	c.lookupLen = nbVar
}

func (c *CNF) addClause(clause []int) {
	c.Clauses = append(c.Clauses, clause)
}

func (c *CNF) addClauses(clauses [][]int) {
	c.Clauses = append(c.Clauses, clauses...)
}

func (c *CNF) addFormula(lits []int, builder CNFBuilder) {
	formula := builder(c, lits)
	c.addClauses(formula)
}

func (c *CNF) getBoard() *sudoku.SudokuBoard {
	return c.Board
}

func (c *CNF) getLits() []int {
	return c.lits
}

func (c *CNF) getClauses() [][]int {
	return c.Clauses
}

func (c *CNF) Print(w io.Writer) {
	fmt.Fprintf(w, "p cnf %d %d\n", c.nbVar, len(c.Clauses))
	for _, c := range c.Clauses {
		for _, l := range c {
			fmt.Fprint(w, l)
		}
		fmt.Fprintln(w, 0)
	}
}
