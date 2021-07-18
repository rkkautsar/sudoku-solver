package sudokusolver

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/crillab/gophersat/explain"
	"github.com/crillab/gophersat/solver"
	"github.com/irifrance/gini"
	"github.com/irifrance/gini/z"
	"github.com/rkkautsar/sudoku-solver/sudoku"
)

func SolveWithGophersat(board *sudoku.Board) {
	cnf := GenerateCNFConstraints(board)
	pb := solver.ParseSlice(cnf.getClauses())
	solvePbWithGophersat(board, pb)
}

func SolveWithGini(board *sudoku.Board) {
	g := gini.New()
	cnf := GenerateCNFConstraints(board)
	giniAddConstraints(g, cnf.getClauses())
	giniSolve(g, board)
}

func giniAddConstraints(g *gini.Gini, clauses [][]int) {
	for _, clause := range clauses {
		// log.Println("add clause", clause)
		for _, lit := range clause {
			g.Add(z.Dimacs2Lit(lit))
			// log.Println("add lit", lit)
		}
		g.Add(0)
	}
}

func giniSolve(g *gini.Gini, board *sudoku.Board) {
	g.Solve()
	model := make([]bool, board.NumCandidates)
	for i := 1; i <= len(model); i++ {
		model[i-1] = g.Value(z.Dimacs2Lit(i))
	}
	board.SolveWithModel(model)
}

func solvePbWithGophersat(board *sudoku.Board, pb *solver.Problem) {
	s := solver.New(pb)
	status := s.Solve()

	if status == solver.Unsat {
		ExplainUnsat(pb)
		return
	}

	board.SolveWithModel(s.Model())
}

func SolveWithCustomSolver(board *sudoku.Board, solver string) {
	solverArgs := strings.Split(solver, " ")
	cmd := exec.Command(solverArgs[0], solverArgs[1:]...)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	reader := bufio.NewScanner(stdout)
	writer := bufio.NewWriter(stdin)

	cmd.Start()
	defer cmd.Wait()
	cnf := GenerateCNFConstraints(board)
	cnf.Print(writer)
	writer.Flush()
	stdin.Close()

	model := make([]bool, board.NumCandidates)

	for reader.Scan() {
		line := reader.Text()

		if strings.HasPrefix(line, "s UNSATISFIABLE") {
			fmt.Println("UNSAT")
			return
		}

		if len(line) < 1 || !strings.HasPrefix(line, "v") {
			continue
		}

		values := strings.Split(line, " ")[1:]
		for _, val := range values {
			parsed, _ := strconv.Atoi(val)
			polarity := parsed > 0

			if parsed < 0 {
				parsed = -parsed
			}

			if parsed > 0 && parsed < len(model) {
				model[parsed-1] = polarity
			}
		}
	}

	board.SolveWithModel(model)
}

func ExplainUnsat(pb *solver.Problem) {
	fmt.Println("UNSAT")
	cnf := pb.CNF()

	unsatPb, err := explain.ParseCNF(strings.NewReader(cnf))
	if err != nil {
		panic(err)
	}

	mus, err := unsatPb.MUSDeletion()
	if err != nil {
		panic(err)
	}
	musCnf := mus.CNF()
	// Sort clauses so as to always have the same output
	lines := strings.Split(musCnf, "\n")
	sort.Sort(sort.StringSlice(lines[1:]))
	musCnf = strings.Join(lines, "\n")
	fmt.Println(musCnf)
}

// only support gophersat since otherwise it has the overhead of spawning subproc
func SolveManyGophersat(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	base := GetBase9x9Clauses()

	for scanner.Scan() {
		input := scanner.Text()
		board := sudoku.NewFromString(input)
		SolveWithGophersatAndBase(board, base)
		board.PrintOneLine(writer)
	}
	writer.Flush()
}

func SolveManyGini(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	// base := GetBase9x9Clauses()
	// g := gini.New()
	// // log.Println("new")
	// giniAddConstraints(g, base.getClauses())
	// // log.Println("constraints")

	for scanner.Scan() {
		// log.Println("start")
		input := scanner.Text()
		board := sudoku.NewFromString(input)
		SolveWithGini(board)
		// cnf := &CNF{Board: board, nbVar: base.nbVar}
		// actLits := make([]z.Lit, len(cnf.lits))
		// for i, l := range cnf.lits {
		// 	g.Add(z.Dimacs2Lit(l))
		// 	actLits[i] = g.Activate()
		// }
		// // log.Println("assume")
		// g.Assume(actLits...)
		// giniSolve(g, board)
		// // log.Println("solve")
		board.PrintOneLine(writer)
		// for _, m := range actLits {
		// 	g.Deactivate(m)
		// }
		// log.Println("end")
	}
	writer.Flush()
}

func GetBase9x9Clauses() *CNF {
	board := sudoku.New(3)
	cnf := GenerateCNFConstraints(board)
	return cnf.(*CNF)
}

func SolveWithGophersatAndBase(board *sudoku.Board, base *CNF) {
	clauses := make([][]int, len(base.Clauses))
	copy(clauses, base.Clauses)

	cnf := &CNF{Board: board, Clauses: clauses, nbVar: base.nbVar}
	cnf.lits = append(cnf.lits, base.lits...)
	cnf.Simplify(SimplifyOptions{disablePureLiteralElimination: false})
	pb := solver.ParseSlice(cnf.Clauses)
	solvePbWithGophersat(board, pb)
}
