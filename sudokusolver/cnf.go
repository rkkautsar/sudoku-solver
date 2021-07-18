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
	lookupTrue(lit int) bool
	requestLiterals(num int) []int
	setInitialNbVar(int)
	getBoard() *sudoku.Board
	getLits() []int
	getClauses() [][]int
	Simplify(SimplifyOptions)
	Print(w io.Writer)
}

type CNF struct {
	CNFInterface
	Board     *sudoku.Board
	Clauses   [][]int
	lits      []int
	litLookup []uint8
	watchers  [][]int
	lookupLen int
	nbVar     int32
}

const (
	unassigned = iota
	assignedTrue
	assignedFalse
)

type CNFBuilder = func(c CNFInterface, lits []int) [][]int

func (c *CNF) addLit(lit int) {
	c.lits = append(c.lits, lit)
	idx := getLitLookupIdx(lit)
	if idx >= len(c.litLookup) {
		return
	}

	if lit < 0 {
		c.litLookup[idx] = assignedFalse
	} else {
		c.litLookup[idx] = assignedTrue
	}
}

func (c *CNF) lookupTrue(lit int) bool {
	idx := getLitLookupIdx(lit)
	if idx >= len(c.litLookup) {
		// log.Println("lookup out of bound", lit, len(c.litLookup))
		return false
	}

	if lit < 0 {
		return c.litLookup[idx] == assignedFalse

	}
	return c.litLookup[idx] == assignedTrue
}

func (c *CNF) requestLiterals(num int) []int {
	lits := makeRange(int(c.nbVar)+1, int(c.nbVar)+num)
	c.nbVar += int32(num)
	return lits
}

func (c *CNF) setInitialNbVar(nbVar int) {
	c.nbVar = int32(nbVar)
	c.lookupLen = nbVar
}

func (c *CNF) addClause(clause []int) {
	c.Clauses = append(c.Clauses, clause)
}

func (c *CNF) addClauses(clauses [][]int) {
	c.Clauses = append(c.Clauses, clauses...)
}

func (c *CNF) addFormula(lits []int, builder CNFBuilder) {
	// log.Println("exactly one", lits)
	formula := builder(c, lits)
	c.addClauses(formula)
}

func (c *CNF) getBoard() *sudoku.Board {
	return c.Board
}

func (c *CNF) getLits() []int {
	return c.lits
}

func (c *CNF) getClauses() [][]int {
	clauses := make([][]int, 0, len(c.Clauses)+len(c.lits))
	for _, l := range c.lits {
		clauses = append(clauses, []int{l})
	}
	clauses = append(clauses, c.Clauses...)
	return clauses
}

func (c *CNF) Print(w io.Writer) {
	clauses := c.getClauses()
	fmt.Fprintf(w, "p cnf %d %d\n", c.nbVar, len(clauses))
	for _, c := range clauses {
		for _, l := range c {
			fmt.Fprint(w, l, " ")
		}
		fmt.Fprintln(w, 0)
	}
}

func getLitLookupIdx(lit int) int {
	if lit < 0 {
		return -lit - 1
	}
	return lit - 1
}
