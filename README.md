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
- pretty fast for larger sudokus (with custom state of the art SAT solver like cadical, 144x144 took 23s, may be better if using parallel SAT solvers)
- pretty simple code

## Benchmarks

About 8.5s (0.17ms per puzzle) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s but done in different machine, so seems we have similar performance. (TODO: can do better if we can reuse gophersat solver instead)

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveWithGophersatAiEscargot-12           	     438	   2615063 ns/op	 1524794 B/op	   13989 allocs/op
BenchmarkSolveWithGophersatHard9x9-12              	     440	   2687688 ns/op	 1614202 B/op	   14966 allocs/op
BenchmarkSolveWithGophersat17clue9x9-12            	     774	   1572928 ns/op	 1518762 B/op	   15405 allocs/op
BenchmarkSolveWithGophersat25x25-12                	      12	 138194101 ns/op	68188720 B/op	  194649 allocs/op
BenchmarkSolveWithGophersat64x64-12                	       1	1722969355 ns/op	1392623656 B/op	 2406550 allocs/op
BenchmarkSolveWithGophersat81x81-12                	       1	30294197463 ns/op	7817848024 B/op	 4702085 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12             	      97	  10944531 ns/op	  892244 B/op	   13110 allocs/op
BenchmarkSolveWithCadicalHard9x9-12                	      84	  14291798 ns/op	  917637 B/op	   13702 allocs/op
BenchmarkSolveWithCadicalHard17clue-12             	      97	  10767028 ns/op	  963013 B/op	   14731 allocs/op
BenchmarkSolveWithCadical25x25-12                  	       7	 166361657 ns/op	31037640 B/op	  228389 allocs/op
BenchmarkSolveWithCadical64x64-12                  	       1	1644846296 ns/op	1098798216 B/op	 3366567 allocs/op
BenchmarkSolveWithCadical81x81-12                  	       1	3219062370 ns/op	2473455232 B/op	 6688904 allocs/op
BenchmarkSolveWithCadical144x144-12                	       1	23011183541 ns/op	22821517000 B/op	36672972 allocs/op
BenchmarkSolveManyWithGophersatHardest110626-12    	      15	  71453908 ns/op	136387861 B/op	  659023 allocs/op
BenchmarkSolveManyWithGophersat17Clue-12           	       1	8558544307 ns/op	17239184152 B/op	67466798 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	83.036s
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
