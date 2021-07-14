package sudokusolver

import (
	"io"
	"runtime"
	"sync"

	"github.com/rkkautsar/sudoku-solver/sudoku"
)

type CNFParallel struct {
	CNFInterface
	*CNF
	workerWg    sync.WaitGroup
	managerWg   sync.WaitGroup
	workChan    chan WorkRequest
	clauseChan  chan [][]int
	literalChan chan NewLiteralRequest
	doneChan    chan bool
	workerCount int
}

func (c *CNFParallel) lookup(lit int) bool {
	return c.CNF.lookup(lit)
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

func (c *CNFParallel) getBoard() *sudoku.SudokuBoard {
	return c.CNF.getBoard()
}

func (c *CNFParallel) getLits() []int {
	return c.CNF.getLits()
}

func (c *CNFParallel) getClauses() [][]int {
	return c.CNF.getClauses()
}

func (c *CNFParallel) requestLiterals(num int) []int {
	resp := make(chan []int)
	c.literalChan <- NewLiteralRequest{num, resp}
	lits := <-resp
	return lits
}

func (c *CNFParallel) initWorkers() {
	var managerWg, workerWg sync.WaitGroup
	c.managerWg = managerWg
	c.workerWg = workerWg
	c.workChan = make(chan WorkRequest)
	c.clauseChan = make(chan [][]int)
	c.literalChan = make(chan NewLiteralRequest)
	c.doneChan = make(chan bool)
	go manager(c)

	c.workerCount = runtime.NumCPU() - 2
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
	close(c.literalChan)
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

type NewLiteralRequest struct {
	num  int
	resp chan []int
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
		case request := <-cnf.literalChan:
			lits := makeRange(cnf.nbVar+1, cnf.nbVar+request.num)
			cnf.nbVar += request.num
			request.resp <- lits
		case <-cnf.doneChan:
			n--
		}
	}
	cnf.managerWg.Done()
}
