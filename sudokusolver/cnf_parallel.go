package sudokusolver

import (
	"io"
	"sync"
	"sync/atomic"

	"github.com/rkkautsar/sudoku-solver/sudoku"
)

type CNFParallel struct {
	CNFInterface
	*CNF
	workerWg    sync.WaitGroup
	managerWg   sync.WaitGroup
	workChan    chan WorkRequest
	clauseChan  chan [][]int
	doneChan    chan bool
	workerCount int
}

func (c *CNFParallel) clauseLen() int {
	return c.CNF.clauseLen()
}

func (c *CNFParallel) lookupTrue(lit int) bool {
	return c.CNF.lookupTrue(lit)
}

func (c *CNFParallel) addLit(lit int) {
	c.CNF.addLit(lit)
}

func (c *CNFParallel) addClause(clause []int) {
	c.clauseChan <- [][]int{clause}
}

func (c *CNFParallel) addClauses(clauses [][]int) {
	c.clauseChan <- clauses
}

func (c *CNFParallel) addFormula(lits []int, builder CNFBuilder) {
	c.workChan <- WorkRequest{lits, builder}
}

func (c *CNFParallel) setInitialNbVar(nbVar int) {
	c.CNF.setInitialNbVar(nbVar)
}

func (c *CNFParallel) getBoard() *sudoku.Board {
	return c.CNF.getBoard()
}

func (c *CNFParallel) getLits() []int {
	return c.CNF.getLits()
}

func (c *CNFParallel) getClauses() [][]int {
	return c.CNF.getClauses()
}

func (c *CNFParallel) requestLiterals(num int) []int {
	newNbVar := atomic.AddInt32(&c.nbVar, int32(num))
	return makeRange(int(newNbVar)-num+1, int(newNbVar))
}

func (c *CNFParallel) initWorkers() {
	var managerWg, workerWg sync.WaitGroup
	c.managerWg = managerWg
	c.workerWg = workerWg
	c.workChan = make(chan WorkRequest, c.workerCount)
	c.clauseChan = make(chan [][]int, 100)
	c.doneChan = make(chan bool, c.workerCount)
	go manager(c)

	c.workerCount = 2
	if c.workerCount <= 0 {
		c.workerCount = 1
	}
	for i := 0; i < c.workerCount; i++ {
		go worker(c)
	}
}

func (c *CNFParallel) closeAndWait() {
	close(c.workChan)
	c.workerWg.Wait()
	close(c.clauseChan)
	c.managerWg.Wait()
}

func (c *CNFParallel) Print(w io.Writer) {
	c.CNF.Print(w)
}

type WorkRequest struct {
	lits    []int
	builder CNFBuilder
}

func worker(cnf *CNFParallel) {
	cnf.workerWg.Add(1)
	for instruction := range cnf.workChan {
		formula := instruction.builder(cnf, instruction.lits)
		cnf.clauseChan <- formula
	}
	cnf.doneChan <- true
	cnf.workerWg.Done()
}

func manager(cnf *CNFParallel) {
	cnf.managerWg.Add(1)
	for n := cnf.workerCount; n > 0; {
		select {
		case clauses := <-cnf.clauseChan:
			cnf.Clauses = append(cnf.Clauses, clauses...)
		case <-cnf.doneChan:
			n--
		}
	}
	cnf.managerWg.Done()
}
