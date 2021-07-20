package sudokusolver

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	// g.Write(os.Stdout)
	status := g.Solve()

	if status < 0 {
		ms := g.Why([]z.Lit{})
		log.Println(ms)
		panic("UNSAT")
	}
	model := make([]bool, board.NumCandidates)
	for i := 1; i <= len(model); i++ {
		model[i-1] = g.Value(z.Dimacs2Lit(i))
	}
	// log.Println(model)
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
	board := base.Board

	for scanner.Scan() {
		input := scanner.Text()
		board.ReplaceWithSingleRowString(input, false)
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
	// log.Println("new")
	// giniAddConstraints(g, base.getClauses())
	// log.Println("constraints")
	// board := base.Board

	// giniSolve(g, board)
	board := sudoku.New(3)

	// actLits := make([]z.Lit, 0, 81)

	for scanner.Scan() {
		// log.Println("start")
		input := scanner.Text()
		board.ReplaceWithSingleRowString(input, true)
		// board.BasicSolve()
		// board.NumCandidates = 729
		// // SolveWithGini(board)
		// // cnf := &CNF{Board: board, nbVar: base.nbVar}
		// actLits = actLits[:0]
		// for i := 1; i <= 729; i++ {
		// 	if board.Lookup[(i-1)/9] == ((i-1)%9)+1 {
		// 		// log.Println("new:", i)
		// 		g.Add(z.Dimacs2Lit(i))
		// 		m := g.Activate()
		// 		// fmt.Println("activation", i, m.Dimacs())
		// 		actLits = append(actLits, m)
		// 	}
		// 	// else if !board.Candidates[i] {
		// 	// 	// log.Println("new:", -i)
		// 	// 	g.Add(z.Dimacs2Lit(-i))
		// 	// 	m := g.Activate()
		// 	// 	actLits = append(actLits, m)
		// 	// 	// fmt.Println("activation", -i, m.Dimacs())
		// 	// }
		// }
		// // log.Println("assume")
		// // log.Println(actLits)
		// g.Assume(actLits...)
		// giniSolve(g, board)
		// // log.Println("solve")
		// board.PrintOneLine(writer)
		// for _, m := range actLits {
		// 	g.Deactivate(m)
		// }
		// log.Println("end")

		g := gini.New()
		board.ReplaceWithSingleRowString(input, false)
		cnf := GenerateCNFConstraints(board)
		giniAddConstraints(g, cnf.getClauses())
		giniSolve(g, board)
		board.PrintOneLine(writer)
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
