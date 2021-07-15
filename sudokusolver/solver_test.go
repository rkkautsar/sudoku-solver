package sudokusolver_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/rkkautsar/sudoku-solver/sudoku"
	"github.com/rkkautsar/sudoku-solver/sudokusolver"
	"github.com/stretchr/testify/assert"
)

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

func TestSolveWithGophersatAiEscargot(t *testing.T) {
	solution := gophersatSolveOneLiner(aiEscargot[0])
	assert.Equal(t, aiEscargot[1], solution)
}

func TestSolveWithGophersatHard1(t *testing.T) {
	solution := gophersatSolveOneLiner(hard1[0])
	assert.Equal(t, hard1[1], solution)
}

func TestSolveWithGophersatHard17clue(t *testing.T) {
	solution := gophersatSolveOneLiner(hard17clue[0])
	assert.Equal(t, hard17clue[1], solution)
}
func BenchmarkSolveWithGophersatAiEscargot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gophersatSolveOneLiner(aiEscargot[0])
	}
}

func BenchmarkSolveWithGophersatHard9x9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gophersatSolveOneLiner(hard1[0])
	}
}

func BenchmarkSolveWithGophersat17clue9x9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gophersatSolveOneLiner(hard17clue[0])
	}
}

func BenchmarkSolveWithGophersat25x25(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-25-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gophersatSolveOneLiner(input)
	}
}

// func BenchmarkSolveWithGophersat64x64(b *testing.B) {
// 	bytes, _ := ioutil.ReadFile("../data/sudoku-64-1.txt")
// 	input := string(bytes)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		gophersatSolveOneLiner(input)
// 	}
// }

// func BenchmarkSolveWithGophersat81x81(b *testing.B) {
// 	bytes, _ := ioutil.ReadFile("../data/sudoku-81-1.txt")
// 	input := string(bytes)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		gophersatSolveOneLiner(input)
// 	}
// }

func BenchmarkSolveWithCadicalAiEscargot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(aiEscargot[0], "cadical -q")
	}
}

func BenchmarkSolveWithCadicalHard9x9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(hard1[0], "cadical -q")
	}
}

func BenchmarkSolveWithCadicalHard17clue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(hard17clue[0], "cadical -q")
	}
}

func BenchmarkSolveWithCadical25x25(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-25-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, "cadical -q")
	}
}

func BenchmarkSolveWithCadical64x64(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-64-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, "cadical -q")
	}
}

func BenchmarkSolveWithCadical81x81(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-81-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, "cadical -q")
	}
}

func BenchmarkSolveWithCadical144x144(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-144-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, "cadical -q")
	}
}

func BenchmarkSolveWithCadical225x225(b *testing.B) {
	bytes, _ := ioutil.ReadFile("../data/sudoku-225-1.txt")
	input := string(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customSolveOneLiner(input, "cadical -q")
	}
}

func BenchmarkSolveManyWithGophersatHardest110626(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveManyWithGophersat("../data/sudoku.many.hardest110626.txt")
	}
}

func BenchmarkSolveManyWithGophersat17Clue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveManyWithGophersat("../data/sudoku.many.17clue.txt")
	}
}

func gophersatSolveOneLiner(input string) string {
	board := sudoku.NewFromString(input)
	sudokusolver.SolveWithGophersat(&board)
	var b bytes.Buffer
	board.PrintOneLine(&b)
	return strings.TrimSpace(b.String())
}

func customSolveOneLiner(input, solver string) string {
	board := sudoku.NewFromString(input)
	sudokusolver.SolveWithCustomSolver(&board, solver)
	var b bytes.Buffer
	board.PrintOneLine(&b)
	return strings.TrimSpace(b.String())
}

func solveManyWithGophersat(inputFile string) {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	base := sudokusolver.GetBase9x9Clauses()
	for scanner.Scan() {
		input := scanner.Text()
		board := sudoku.NewFromString(input)
		sudokusolver.SolveWithGophersatAndBase(&board, base)
		var b bytes.Buffer
		board.PrintOneLine(&b)
	}
}
