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
- fast, but not as fast as specialized solvers (0.6ms for ai-escargot, naïve backtracking is around 30ms)
- pretty fast for larger sudokus (with custom state of the art SAT solver like cadical, 144x144 took 20s, may be better if using parallel SAT solvers)
- pretty simple code

## Benchmarks

About 51.5s (954 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-12               	    1711	    663113 ns/op	  603690 B/op	    3984 allocs/op
BenchmarkSolveHard9x9-12                  	     656	   1819127 ns/op	  745202 B/op	    5069 allocs/op
BenchmarkSolve17clue9x9-12                	     918	   1242395 ns/op	 1213969 B/op	    7963 allocs/op
BenchmarkSolve25x25-12                    	      43	  23817584 ns/op	15700733 B/op	  109126 allocs/op
BenchmarkSolve64x64-12                    	       3	 405321750 ns/op	99750616 B/op	  498922 allocs/op
BenchmarkSolve81x81-12                    	       3	 467102880 ns/op	241131421 B/op	 1335135 allocs/op
BenchmarkSolve100x100-12                  	       4	 303379450 ns/op	232606390 B/op	  917229 allocs/op
BenchmarkSolve144x144-12                  	       2	 884218184 ns/op	751492956 B/op	 3344856 allocs/op
BenchmarkSolve225x225-12                  	       1	2679304741 ns/op	1457409928 B/op	 1205887 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12    	      58	  18298219 ns/op	  303392 B/op	    5747 allocs/op
BenchmarkSolveWithCadicalHard9x9-12       	     114	  11335283 ns/op	  431605 B/op	    7209 allocs/op
BenchmarkSolveWithCadicalHard17clue-12    	      93	  12866351 ns/op	  734494 B/op	   12246 allocs/op
BenchmarkSolveWithCadical25x25-12         	      15	  77647577 ns/op	 9598772 B/op	  176605 allocs/op
BenchmarkSolveWithCadical64x64-12         	       1	2314926730 ns/op	59734048 B/op	  827781 allocs/op
BenchmarkSolveWithCadical81x81-12         	       1	1364799701 ns/op	149652664 B/op	 2407136 allocs/op
BenchmarkSolveWithCadical100x100-12       	       2	 744538528 ns/op	171159232 B/op	 1508925 allocs/op
BenchmarkSolveWithCadical144x144-12       	       1	2212441629 ns/op	556941648 B/op	 5890258 allocs/op
BenchmarkSolveWithCadical225x225-12       	       1	2814346249 ns/op	1404577904 B/op	 2020374 allocs/op
BenchmarkSolveManyHardest110626-12        	       1	2680093046 ns/op	29963776 B/op	   92572 allocs/op
BenchmarkSolveMany17Clue2k-12             	       1	1508717751 ns/op	18968024 B/op	   97994 allocs/op
BenchmarkSolveMany17Clue-12               	       1	67656855355 ns/op	194538384 B/op	 1224765 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	111.019s
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
