# SAT-based Sudoku Solver

A rewrite of a side project: https://github.com/rkkautsar/sudoku-solver-nim

## Description

A general n &times; n Sudoku solver that works by encoding the constraints to SAT and use SAT solver to solve it.

Featuring:

- accepts n &times; n sudoku input (either multiline or one line for 9x9 sudoku)
- can print out the CNF encoding only
- has built-in SAT solver (gini) or can use custom SAT solver
- bimander encoding for at-most-one
- parallel CNF encoding
- fast, but not as fast as specialized solvers (0.6ms for ai-escargot, naïve backtracking is around 30ms)
- pretty fast for larger sudokus, for example a 144x144 sudoku can be solved in 4s
- pretty simple code

## Benchmarks

About 7.3s (6733 puzzle/s or 0.15 ms/puzzle) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks available in `make bench`:

```
↳ make bench

go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: arm64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-10               	    1672	    681784 ns/op	  494076 B/op	    3499 allocs/op
BenchmarkSolveHard9x9-10                  	    2335	    529959 ns/op	  595289 B/op	    3997 allocs/op
BenchmarkSolve17clue9x9-10                	   19891	     51179 ns/op	  185785 B/op	     568 allocs/op
BenchmarkSolve25x25-10                    	     192	   5740740 ns/op	 8906262 B/op	   62807 allocs/op
BenchmarkSolve64x64-10                    	       3	 421615486 ns/op	66795490 B/op	  294591 allocs/op
BenchmarkSolve81x81-10                    	       5	 214429175 ns/op	192092593 B/op	  963851 allocs/op
BenchmarkSolve100x100-10                  	      14	  78694101 ns/op	171781423 B/op	  472035 allocs/op
BenchmarkSolve144x144-10                  	       4	 258500177 ns/op	349729978 B/op	   75485 allocs/op
BenchmarkSolve225x225-10                  	       2	 548826021 ns/op	1351335860 B/op	  191375 allocs/op
BenchmarkSolveWithCadicalAiEscargot-10    	     267	   4271864 ns/op	  262297 B/op	    4631 allocs/op
BenchmarkSolveWithCadicalHard9x9-10       	     325	   3630465 ns/op	  298358 B/op	    5673 allocs/op
BenchmarkSolveWithCadicalHard17clue-10    	     475	   2576579 ns/op	  125342 B/op	     350 allocs/op
BenchmarkSolveWithCadical25x25-10         	      96	  24191828 ns/op	 5024244 B/op	   92785 allocs/op
BenchmarkSolveWithCadical64x64-10         	       2	 703740396 ns/op	45394548 B/op	  440766 allocs/op
BenchmarkSolveWithCadical81x81-10         	       4	 337242250 ns/op	123322466 B/op	 1693375 allocs/op
BenchmarkSolveWithCadical100x100-10       	       7	 148710054 ns/op	141965497 B/op	  717570 allocs/op
BenchmarkSolveWithCadical144x144-10       	       4	 304871969 ns/op	342056742 B/op	  160152 allocs/op
BenchmarkSolveWithCadical225x225-10       	       2	 657442375 ns/op	1336465844 B/op	  400567 allocs/op
BenchmarkSolveManyHardest110626-10        	       3	 358715625 ns/op	225478709 B/op	 1574099 allocs/op
BenchmarkSolveMany17Clue2k-10             	       4	 306551583 ns/op	590373882 B/op	 3235654 allocs/op
BenchmarkSolveMany17Clue-10               	       1	7351459250 ns/op	14441545904 B/op	79415491 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	47.154s
```

## Getting Started

### Dependencies

- Go

### Installing

#### From gobinaries

```sh
curl -sf https://gobinaries.com/rkkautsar/sudoku-solver/cmd/sudokusolver | sh
```

#### From source

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
sudokusolver -solver "cadical -q" < data/sudoku-9-1.txt
```

## License

This project is licensed under the MIT License - see the LICENSE.md file for details

## Acknowledgments

Inspiration, code snippets, etc.

- [gini](https://github.com/irifrance/gini)
- [tdoku](https://github.com/t-dillon/tdoku)
- Kwon, Gihwon, and Himanshu Jain. "Optimized CNF encoding for sudoku puzzles." Proc. 13th International Conference on Logic for Programming Artificial Intelligence and Reasoning (LPAR2006). 2006. ([PDF](http://www.cs.cmu.edu/~hjain/papers/sudoku-as-SAT.pdf))
- Klieber, Will, and Gihwon Kwon. "Efficient CNF encoding for selecting 1 from n objects." Proc. International Workshop on Constraints in Formal Verification. 2007. ([PDF](https://www.cs.cmu.edu/~wklieber/papers/2007_efficient-cnf-encoding-for-selecting-1.pdf))
- Nguyen, Van-Hau, and Son T. Mai. "A new method to encode the at-most-one constraint into SAT." Proceedings of the Sixth International Symposium on Information and Communication Technology. 2015. ([PDF](https://www.researchgate.net/profile/Van-Hau-Nguyen/publication/301455290_A_New_Method_to_Encode_the_At-Most-One_Constraint_into_SAT/links/5d2bfbaba6fdcc2462e0e269/A-New-Method-to-Encode-the-At-Most-One-Constraint-into-SAT.pdf))
