package sudokusolver

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"github.com/irifrance/gini"
	"github.com/irifrance/gini/z"
	"github.com/rkkautsar/sudoku-solver/sudoku"
)

func SolveWithGini(board *sudoku.Board) {
	board.BasicSolve()
	g := gini.NewVc(2*board.NumCandidates, 3*board.NumCandidates)
	GenerateCNFConstraints(board, g)
	giniSolve(g, board)
}

func giniSolve(g *gini.Gini, board *sudoku.Board) {
	status := g.Solve()

	if status < 0 {
		ms := g.Why([]z.Lit{})
		fmt.Println(ms)
		panic("UNSAT")
	}
	model := make([]bool, board.NumCandidates)
	for i := 1; i <= len(model); i++ {
		model[i-1] = g.Value(z.Dimacs2Lit(i))
	}
	// log.Println(model)
	board.SolveWithModel(model)
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
	board.BasicSolve()
	g := gini.NewVc(2*board.NumCandidates, 3*board.NumCandidates)
	cnf := GenerateCNFConstraints(board, g)
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

func SolveManyGini(in io.Reader, out io.Writer) {
	shouldPrintPuzzle := false

	if out == nil {
		out = io.Discard
	} else {
		shouldPrintPuzzle = true
	}

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	board := sudoku.New(3)

	for scanner.Scan() {
		input := scanner.Text()
		if shouldPrintPuzzle {
			writer.WriteString(input + ",")
		}
		board.ReplaceWithSingleRowString(input, false)
		SolveWithGini(board)
		board.PrintOneLine(writer)
	}
	writer.Flush()
}
