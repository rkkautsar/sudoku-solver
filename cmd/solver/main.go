package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"

	"github.com/crillab/gophersat/explain"
	"github.com/crillab/gophersat/solver"
	"github.com/rkkautsar/sudoku-solver-2/sudoku"
	"github.com/rkkautsar/sudoku-solver-2/sudokusolver"
)

var (
	isCNFMode    bool
	isSolveMode  bool
	isManyMode   bool
	cpuprofile   string
	customSolver string
)

func init() {
	flag.BoolVar(&isCNFMode, "cnf", true, "Generate CNF")
	flag.BoolVar(&isSolveMode, "solve", false, "Solve with SAT solver")
	flag.BoolVar(&isManyMode, "many", false, "Solve many one-line 9x9 sudoku")
	flag.StringVar(&customSolver, "solver", "gophersat", "Solve with specified SAT solver [implies -solve if set]")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "Write CPU profile to a file")
	flag.Parse()
}

func main() {
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	mode := "cnf"
	if isSolveMode || !isCNFMode {
		mode = "solve"
	}
	if customSolver != "gophersat" {
		mode = "custom"
	}

	if isManyMode {
		solveMany(mode)
	} else {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		input := string(bytes)
		solve(mode, input)
	}
}

func solve(mode, input string) {
	board := sudoku.NewFromString(input)

	var cnf *sudokusolver.CNF

	if mode != "custom" {
		cnf = sudokusolver.GenerateCNFConstraints(&board)
	}

	if mode == "cnf" {
		cnf.Print(os.Stdout)
		return
	}

	if mode == "solve" {
		solveWithGophersat(&board, cnf)
	}

	if mode == "custom" {
		solveWithCustomSolver(&board, customSolver)
	}

	board.Print()
}

func solveMany(mode string) {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		input := scanner.Text()
		board := sudoku.NewFromString(input)

		if mode == "solve" {
			cnf := sudokusolver.GenerateCNFConstraints(&board)
			solveWithGophersat(&board, cnf)
		}

		if mode == "custom" {
			solveWithCustomSolver(&board, customSolver)
		}

		board.PrintOneLine(writer)
	}
	writer.Flush()
}

func solveWithGophersat(board *sudoku.SudokuBoard, cnf *sudokusolver.CNF) {
	pb := solver.ParseSlice(cnf.Clauses)
	s := solver.New(pb)
	status := s.Solve()

	if status == solver.Unsat {
		explainUnsat(pb)
		return
	}

	board.SolveWithModel(s.Model())
}

func solveWithCustomSolver(board *sudoku.SudokuBoard, solver string) {
	solverArgs := strings.Split(solver, " ")
	cmd := exec.Command(solverArgs[0], solverArgs[1:]...)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	reader := bufio.NewScanner(stdout)
	writer := bufio.NewWriter(stdin)

	cmd.Start()
	defer cmd.Wait()
	cnf := sudokusolver.GenerateCNFConstraints(board)
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

func explainUnsat(pb *solver.Problem) {
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
