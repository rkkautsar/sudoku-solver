package sudokusolver

import (
	"io"

	"github.com/irifrance/gini"
	"github.com/irifrance/gini/z"
	"github.com/rkkautsar/sudoku-solver/sudoku"
)

type CNFInterface interface {
	addLit(lit int)
	addClause(clause []int)
	addFormula(lits []int, builder CNFBuilder)
	requestLiterals(num uint32) []int
	getBoard() *sudoku.Board
	Print(w io.Writer)
}

type CNF struct {
	CNFInterface
	g     *gini.Gini
	Board *sudoku.Board
	nbVar uint32
}

type CNFBuilder = func(c CNFInterface, lits []int)

func (c *CNF) addLit(lit int) {
	c.g.Add(z.Dimacs2Lit(lit))
	c.g.Add(0)
}

func (c *CNF) requestLiterals(num uint32) []int {
	lits := makeRange(c.nbVar+1, c.nbVar+num)
	c.nbVar += num
	return lits
}

func (c *CNF) addClause(clause []int) {
	for _, lit := range clause {
		c.g.Add(z.Dimacs2Lit(lit))
	}
	c.g.Add(0)
}

func (c *CNF) addFormula(lits []int, builder CNFBuilder) {
	builder(c, lits)
}

func (c *CNF) getBoard() *sudoku.Board {
	return c.Board
}

func (c *CNF) Print(w io.Writer) {
	c.g.Write(w)
}
