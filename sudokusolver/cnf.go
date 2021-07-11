package sudokusolver

import (
	"fmt"
	"io"
	"runtime"
	"sync"

	"github.com/rkkautsar/sudoku-solver/sudoku"
)

type CNF struct {
	Board       *sudoku.SudokuBoard
	Clauses     [][]int
	lits        []int
	litLookup   []uint8
	lookupLen   int
	nbVar       int
	workerWg    sync.WaitGroup
	managerWg   sync.WaitGroup
	workChan    chan WorkRequest
	clauseChan  chan [][]int
	literalChan chan NewLiteralRequest
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
	c.clauseChan <- [][]int{clause}
}

func (c *CNF) addClauses(clauses [][]int) {
	c.clauseChan <- clauses
}

func (c *CNF) addFormula(lits []int, builder CNFBuilder) {
	c.workChan <- WorkRequest{lits, builder}
}

func (c *CNF) requestLiterals(num int) []int {
	resp := make(chan []int)
	c.literalChan <- NewLiteralRequest{num, resp}
	lits := <-resp
	return lits
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

func (c *CNF) initWorkers() {
	var managerWg, workerWg sync.WaitGroup
	c.managerWg = managerWg
	c.workerWg = workerWg
	c.workChan = make(chan WorkRequest)
	c.clauseChan = make(chan [][]int)
	c.literalChan = make(chan NewLiteralRequest)
	go clauseManager(c)
	go newLiteralManager(c)

	workerCount := runtime.NumCPU() - 2
	if workerCount <= 0 {
		workerCount = 1
	}
	for i := 0; i < workerCount; i++ {
		go worker(c)
	}
}

func (c *CNF) closeAndWait() {
	close(c.workChan)
	c.workerWg.Wait()
	close(c.literalChan)
	close(c.clauseChan)
	c.managerWg.Wait()
}

type WorkRequest struct {
	lits    []int
	builder CNFBuilder
}

type NewLiteralRequest struct {
	num  int
	resp chan []int
}

func worker(cnf *CNF) {
	cnf.workerWg.Add(1)
	for instruction := range cnf.workChan {
		formula := instruction.builder(cnf, instruction.lits)
		cnf.clauseChan <- formula
	}
	cnf.workerWg.Done()
}

func clauseManager(cnf *CNF) {
	cnf.managerWg.Add(1)
	for clauses := range cnf.clauseChan {
		cnf.Clauses = append(cnf.Clauses, clauses...)
	}
	cnf.managerWg.Done()
}

func newLiteralManager(cnf *CNF) {
	cnf.managerWg.Add(1)
	for request := range cnf.literalChan {
		lits := makeRange(cnf.nbVar+1, cnf.nbVar+request.num)
		cnf.nbVar += request.num
		request.resp <- lits
	}
	cnf.managerWg.Done()
}
