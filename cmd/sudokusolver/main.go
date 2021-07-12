package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	"github.com/rkkautsar/sudoku-solver/sudoku"
	"github.com/rkkautsar/sudoku-solver/sudokusolver"
)

var (
	isCNFMode    bool
	isSolveMode  bool
	isManyMode   bool
	cpuprofile   string
	customSolver string
)

func init() {
	flag.BoolVar(&isCNFMode, "cnf", false, "Generate CNF")
	flag.BoolVar(&isSolveMode, "solve", true, "Solve with SAT solver")
	flag.BoolVar(&isManyMode, "many", false, "Solve many one-line 9x9 sudoku w/ gophersat")
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

	mode := "solve"
	if isCNFMode {
		mode = "cnf"
	}
	if !isCNFMode && customSolver != "gophersat" {
		mode = "custom"
	}

	if isManyMode {
		solveMany()
	} else {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		input := string(bytes)
		solve(mode, input)
	}
}

func solve(mode, input string) {
	board := sudoku.NewFromString(input)

	if mode == "cnf" {
		cnf := sudokusolver.GenerateCNFConstraints(&board)
		cnf.Print(os.Stdout)
		return
	}

	if mode == "solve" {
		sudokusolver.SolveWithGophersat(&board)
	}

	if mode == "custom" {
		sudokusolver.SolveWithCustomSolver(&board, customSolver)
	}

	board.Print()
}

// only support gophersat since otherwise it has the overhead of spawning subproc
func solveMany() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	base := sudokusolver.GetBase9x9Clauses()

	for scanner.Scan() {
		input := scanner.Text()
		board := sudoku.NewFromString(input)
		// sudokusolver.SolveWithGophersat(&board)
		sudokusolver.SolveWithGophersatAndBase(&board, base)
		board.PrintOneLine(writer)
	}
	writer.Flush()
}
