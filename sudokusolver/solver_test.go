package sudokusolver_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/rkkautsar/sudoku-solver/sudoku"
	"github.com/rkkautsar/sudoku-solver/sudokusolver"
	"github.com/stretchr/testify/assert"
)

const CUSTOM_SOLVER = "cadical -q"

var aiEscargot = [...]string{
	"100007090030020008009600500005300900010080002600004000300000010041000007007000300",
	"162857493534129678789643521475312986913586742628794135356478219241935867897261354",
}

var hard1 = [...]string{
	"........8..3...4...9..2..6.....79.......612...6.5.2.7...8...5...1.....2.4.5.....3",
	"621943758783615492594728361142879635357461289869532174238197546916354827475286913",
}

var hard17clue = [...]string{
	"000000010400000000020000000000050407008000300001090000300400200050100000000806000",
	"693784512487512936125963874932651487568247391741398625319475268856129743274836159",
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestSolveAiEscargot(t *testing.T) {
	solution := solveOneLiner(aiEscargot[0])
	assert.Equal(t, aiEscargot[1], solution)
}

func TestSolveHard1(t *testing.T) {
	solution := solveOneLiner(hard1[0])
	assert.Equal(t, hard1[1], solution)
}

func TestSolveHard17clue(t *testing.T) {
	solution := solveOneLiner(hard17clue[0])
	assert.Equal(t, hard17clue[1], solution)
}
func BenchmarkSolveAiEscargot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveOneLiner(aiEscargot[0])
	}
}

func BenchmarkSolveHard9x9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveOneLiner(hard1[0])
	}
}

func BenchmarkSolve17clue9x9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveOneLiner(hard17clue[0])
	}
}

func BenchmarkSolve25x25(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-25-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOneLiner(input)
	}
}

func BenchmarkSolve64x64(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-64-2.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOneLiner(input)
	}
}

// TLE (>11m)
// func BenchmarkSolve64x64Hard(b *testing.B) {
// 	bytes, _ := ioutil.ReadFile("../data/sudoku-64-1.txt")
// 	input := string(bytes)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		solveOneLiner(input)
// 	}
// }

func BenchmarkSolve81x81(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-81-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOneLiner(input)
	}
}

func BenchmarkSolve100x100(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-100-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOneLiner(input)
	}
}

func BenchmarkSolve144x144(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-144-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOneLiner(input)
	}
}

func BenchmarkSolve225x225(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-225-2.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOneLiner(input)
	}
}

func BenchmarkSolveWithCadicalAiEscargot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(aiEscargot[0], CUSTOM_SOLVER)
	}
}

func BenchmarkSolveWithCadicalHard9x9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(hard1[0], CUSTOM_SOLVER)
	}
}

func BenchmarkSolveWithCadicalHard17clue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(hard17clue[0], CUSTOM_SOLVER)
	}
}

func BenchmarkSolveWithCadical25x25(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-25-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, CUSTOM_SOLVER)
	}
}

func BenchmarkSolveWithCadical64x64(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-64-2.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, CUSTOM_SOLVER)
	}
}

// func BenchmarkSolveWithCadical64x64Hard(b *testing.B) {
// 	bytes, _ := ioutil.ReadFile("../data/sudoku-64-1.txt")
// 	input := string(bytes)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		customSolveOneLiner(input, CUSTOM_SOLVER)
// 	}
// }

func BenchmarkSolveWithCadical81x81(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-81-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, CUSTOM_SOLVER)
	}
}

func BenchmarkSolveWithCadical100x100(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-100-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, CUSTOM_SOLVER)
	}
}

func BenchmarkSolveWithCadical144x144(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-144-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, CUSTOM_SOLVER)
	}
}

func BenchmarkSolveWithCadical225x225(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-225-2.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, CUSTOM_SOLVER)
	}
}

func BenchmarkSolveManyHardest110626(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveMany("../data/sudoku.many.hardest110626.txt")
	}
}

func BenchmarkSolveMany17Clue2k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveMany("../data/sudoku.many.17clue.2k.txt")
	}
}

func BenchmarkSolveMany17Clue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveMany("../data/sudoku.many.17clue.txt")
	}
}

func solveOneLiner(input string) string {
	board := sudoku.NewFromString(input)
	// sudokusolver.Solve(board)
	sudokusolver.SolveWithGini(board)
	var b bytes.Buffer
	board.PrintOneLine(&b)
	return strings.TrimSpace(b.String())
}

func customSolveOneLiner(input, solver string) string {
	board := sudoku.NewFromString(input)
	sudokusolver.SolveWithCustomSolver(board, solver)
	var b bytes.Buffer
	board.PrintOneLine(&b)
	return strings.TrimSpace(b.String())
}

func solveMany(inputFile string) {
	solveManyWithGini(inputFile)
}

func solveManyWithGophersat(inputFile string) {
	file, _ := os.Open(inputFile)
	sudokusolver.SolveManyGophersat(file, ioutil.Discard)
}

func solveManyWithGini(inputFile string) {
	file, _ := os.Open(inputFile)
	sudokusolver.SolveManyGini(file, ioutil.Discard)
}
