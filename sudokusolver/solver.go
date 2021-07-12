package sudokusolver

import (
	"bufio"
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/crillab/gophersat/explain"
	"github.com/crillab/gophersat/solver"
	"github.com/rkkautsar/sudoku-solver/sudoku"
)

func SolveWithGophersat(board *sudoku.SudokuBoard) {
	cnf := GenerateCNFConstraints(board)
	solveCNFwithGophersat(cnf)
}

func solveCNFwithGophersat(cnf CNFInterface) {
	pb := solver.ParseSlice(cnf.getClauses())
	s := solver.New(pb)
	status := s.Solve()

	if status == solver.Unsat {
		ExplainUnsat(pb)
		return
	}

	cnf.getBoard().SolveWithModel(s.Model())
}

func SolveWithCustomSolver(board *sudoku.SudokuBoard, solver string) {
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

	model := make([]bool, board.LenCells()*board.LenValues())

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

func GetBase9x9Clauses() [][]int {
	board := sudoku.SudokuBoard{Size: 3, Known: []sudoku.Cell{}}
	cnf := GenerateCNFConstraints(&board)
	return cnf.getClauses()
}

func SolveWithGophersatAndBaseClauses(board *sudoku.SudokuBoard, clauses [][]int) {
	cnf := &CNF{Board: board, Clauses: clauses}
	cnf.generateLitLookup()
	for _, lit := range cnf.getLits() {
		cnf.addClause([]int{lit})
	}
	solveCNFwithGophersat(cnf)
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
