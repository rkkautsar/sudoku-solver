# SAT-based Sudoku Solver

A rewrite of a side project: https://github.com/rkkautsar/sudoku-solver-nim

## Description

A general n &times; n Sudoku solver that works by encoding the constraints to SAT and use SAT solver to solve it.

Featuring:

- accepts n &times; n sudoku input (either multiline or one line for 9x9 sudoku)
- can print out the CNF encoding only
- has built-in SAT solver (gophersat) or can use custom SAT solver
- bimander encoding for at-most-one
- parallel CNF encoding
- fast but not as fast as specialized solvers (2.7ms for ai-escargot, naive backtracking is around 30ms)
- pretty fast for larger sudokus (with custom state of the art SAT solver like cadical, 144x144 took 20s, may be better if using parallel SAT solvers)
- pretty simple code

## Benchmarks

About 68s (1.4ms per puzzle) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s but done in different machine. We should be able to achieve the same number (or even faster?) if we can reuse the SAT solver instead (AFAIK currently gophersat doesn't support that).

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveWithGophersatAiEscargot-12           	     440	   2619218 ns/op	 1492665 B/op	   13764 allocs/op
BenchmarkSolveWithGophersatHard9x9-12              	     445	   2664543 ns/op	 1581811 B/op	   14747 allocs/op
BenchmarkSolveWithGophersat17clue9x9-12            	     800	   1497033 ns/op	 1486668 B/op	   15200 allocs/op
BenchmarkSolveWithGophersat25x25-12                	      13	 214647199 ns/op	89722513 B/op	  171241 allocs/op
BenchmarkSolveWithGophersat64x64-12                	       1	1461878184 ns/op	1406503392 B/op	 1786650 allocs/op
BenchmarkSolveWithGophersat81x81-12                	       1	11567383716 ns/op	4623310624 B/op	 3386890 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12             	      86	  11679865 ns/op	  868394 B/op	   12884 allocs/op
BenchmarkSolveWithCadicalHard9x9-12                	      85	  13694774 ns/op	  893992 B/op	   13482 allocs/op
BenchmarkSolveWithCadicalHard17clue-12             	     110	  10720756 ns/op	  939978 B/op	   14526 allocs/op
BenchmarkSolveWithCadical25x25-12                  	       8	 142845386 ns/op	30053894 B/op	  202405 allocs/op
BenchmarkSolveWithCadical64x64-12                  	       1	1102019556 ns/op	1076979416 B/op	 2745999 allocs/op
BenchmarkSolveWithCadical81x81-12                  	       1	2413306978 ns/op	2425167816 B/op	 5392853 allocs/op
BenchmarkSolveWithCadical144x144-12                	       1	17188514020 ns/op	22526224632 B/op	28845185 allocs/op
BenchmarkSolveManyWithGophersatHardest110626-12    	       2	 554761716 ns/op	379259636 B/op	 5097771 allocs/op
BenchmarkSolveManyWithGophersat17Clue-12           	       1	68379438066 ns/op	49290168648 B/op	648892469 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	118.915s
```

## Getting Started

### Dependencies

- Go

### Installing

```sh
go get github.com/rkkautsar/sudoku-solver/cmd/sudokusolver
```

### Executing program

```sh
sudokusolver -help
sudokusolver -cnf < data/sudoku-9-1.txt
sudokusolver -solve < data/sudoku-9-1.txt
sudokusolver -solve -many < data/sudoku.many.17clue.txt

# brew install cadical
sudokusilver -solver "cadical -q" < data/sudoku-9-1.txt
```

## License

This project is licensed under the MIT License - see the LICENSE.md file for details

## Acknowledgments

Inspiration, code snippets, etc.

- [Gophersat](https://github.com/crillab/gophersat)
- Kwon, Gihwon, and Himanshu Jain. "Optimized CNF encoding for sudoku puzzles." Proc. 13th International Conference on Logic for Programming Artificial Intelligence and Reasoning (LPAR2006). 2006. ([PDF](http://www.cs.cmu.edu/~hjain/papers/sudoku-as-SAT.pdf))
- Klieber, Will, and Gihwon Kwon. "Efficient CNF encoding for selecting 1 from n objects." Proc. International Workshop on Constraints in Formal Verification. 2007. ([PDF](https://www.cs.cmu.edu/~wklieber/papers/2007_efficient-cnf-encoding-for-selecting-1.pdf))
- Nguyen, Van-Hau, and Son T. Mai. "A new method to encode the at-most-one constraint into SAT." Proceedings of the Sixth International Symposium on Information and Communication Technology. 2015. ([PDF](https://www.researchgate.net/profile/Van-Hau-Nguyen/publication/301455290_A_New_Method_to_Encode_the_At-Most-One_Constraint_into_SAT/links/5d2bfbaba6fdcc2462e0e269/A-New-Method-to-Encode-the-At-Most-One-Constraint-into-SAT.pdf))
