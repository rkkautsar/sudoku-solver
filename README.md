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
- pretty fast for larger sudokus (with custom state of the art SAT solver like cadical, 144x144 took 20s, may be better if using parallel SAT solvers)
- pretty simple code

## Benchmarks

About 17.8s (2761 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-12               	     709	   1666217 ns/op	  493419 B/op	    3500 allocs/op
BenchmarkSolveHard9x9-12                  	     921	   1286598 ns/op	  603805 B/op	    4001 allocs/op
BenchmarkSolve17clue9x9-12                	    7153	    163063 ns/op	  184390 B/op	     570 allocs/op
BenchmarkSolve25x25-12                    	      80	  12950373 ns/op	 8083937 B/op	   56287 allocs/op
BenchmarkSolve64x64-12                    	       1	1220353582 ns/op	60514968 B/op	  242670 allocs/op
BenchmarkSolve81x81-12                    	       3	 399683702 ns/op	121419346 B/op	  576312 allocs/op
BenchmarkSolve100x100-12                  	       5	 206062494 ns/op	165985179 B/op	  422056 allocs/op
BenchmarkSolve144x144-12                  	       2	 697157792 ns/op	349824684 B/op	   75491 allocs/op
BenchmarkSolve225x225-12                  	       1	2142547654 ns/op	1351756120 B/op	  191380 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12    	     104	  10387281 ns/op	  272787 B/op	    4631 allocs/op
BenchmarkSolveWithCadicalHard9x9-12       	     122	   9572524 ns/op	  308229 B/op	    5674 allocs/op
BenchmarkSolveWithCadicalHard17clue-12    	     178	   6628977 ns/op	  131383 B/op	     348 allocs/op
BenchmarkSolveWithCadical25x25-12         	      25	  41032223 ns/op	 4497729 B/op	   78792 allocs/op
BenchmarkSolveWithCadical64x64-12         	       1	1714517015 ns/op	40885528 B/op	  320000 allocs/op
BenchmarkSolveWithCadical81x81-12         	       1	1365326700 ns/op	91180048 B/op	  792776 allocs/op
BenchmarkSolveWithCadical100x100-12       	       3	 412547854 ns/op	137918040 B/op	  609797 allocs/op
BenchmarkSolveWithCadical144x144-12       	       2	 877587856 ns/op	342151780 B/op	  160161 allocs/op
BenchmarkSolveWithCadical225x225-12       	       1	2278051818 ns/op	1336688080 B/op	  400567 allocs/op
BenchmarkSolveManyHardest110626-12        	       2	 837861886 ns/op	228404280 B/op	 1574401 allocs/op
BenchmarkSolveMany17Clue2k-12             	       2	 704271392 ns/op	590739820 B/op	 3235576 allocs/op
BenchmarkSolveMany17Clue-12               	       1	17838533350 ns/op	14457925264 B/op	79414142 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	57.400s
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

- [gini](https://github.com/irifrance/gini)
- [tdoku](https://github.com/t-dillon/tdoku)
- Kwon, Gihwon, and Himanshu Jain. "Optimized CNF encoding for sudoku puzzles." Proc. 13th International Conference on Logic for Programming Artificial Intelligence and Reasoning (LPAR2006). 2006. ([PDF](http://www.cs.cmu.edu/~hjain/papers/sudoku-as-SAT.pdf))
- Klieber, Will, and Gihwon Kwon. "Efficient CNF encoding for selecting 1 from n objects." Proc. International Workshop on Constraints in Formal Verification. 2007. ([PDF](https://www.cs.cmu.edu/~wklieber/papers/2007_efficient-cnf-encoding-for-selecting-1.pdf))
- Nguyen, Van-Hau, and Son T. Mai. "A new method to encode the at-most-one constraint into SAT." Proceedings of the Sixth International Symposium on Information and Communication Technology. 2015. ([PDF](https://www.researchgate.net/profile/Van-Hau-Nguyen/publication/301455290_A_New_Method_to_Encode_the_At-Most-One_Constraint_into_SAT/links/5d2bfbaba6fdcc2462e0e269/A-New-Method-to-Encode-the-At-Most-One-Constraint-into-SAT.pdf))
