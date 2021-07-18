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

About 54s (910 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-12               	     697	   1687623 ns/op	  491796 B/op	    3593 allocs/op
BenchmarkSolveHard9x9-12                  	     906	   1314003 ns/op	  609422 B/op	    4249 allocs/op
BenchmarkSolve17clue9x9-12                	    1238	    952341 ns/op	  905815 B/op	    5778 allocs/op
BenchmarkSolve25x25-12                    	      64	  16936977 ns/op	 8465780 B/op	   64830 allocs/op
BenchmarkSolve64x64-12                    	       1	2113378361 ns/op	72375544 B/op	  311863 allocs/op
BenchmarkSolve81x81-12                    	       3	 342890040 ns/op	134011754 B/op	  670371 allocs/op
BenchmarkSolve100x100-12                  	       5	 222726956 ns/op	173762041 B/op	  568440 allocs/op
BenchmarkSolve144x144-12                  	       2	 649767466 ns/op	540966428 B/op	 1775294 allocs/op
BenchmarkSolve225x225-12                  	       1	2741436484 ns/op	1436309272 B/op	  979385 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12    	      63	  20319456 ns/op	  268032 B/op	    4722 allocs/op
BenchmarkSolveWithCadicalHard9x9-12       	      92	  13955196 ns/op	  313934 B/op	    6030 allocs/op
BenchmarkSolveWithCadicalHard17clue-12    	     100	  10950169 ns/op	  477029 B/op	    8500 allocs/op
BenchmarkSolveWithCadical25x25-12         	      25	  47552508 ns/op	 4840385 B/op	   90033 allocs/op
BenchmarkSolveWithCadical64x64-12         	       1	2148136893 ns/op	43759336 B/op	  409028 allocs/op
BenchmarkSolveWithCadical81x81-12         	       2	 701421391 ns/op	95692404 B/op	  924666 allocs/op
BenchmarkSolveWithCadical100x100-12       	       2	 515719512 ns/op	144273756 B/op	  811772 allocs/op
BenchmarkSolveWithCadical144x144-12       	       1	1504072771 ns/op	430382712 B/op	 2491800 allocs/op
BenchmarkSolveWithCadical225x225-12       	       1	3114840298 ns/op	1386768752 B/op	 1561386 allocs/op
BenchmarkSolveManyHardest110626-12        	       1	2528963576 ns/op	17054176 B/op	   65788 allocs/op
BenchmarkSolveMany17Clue2k-12             	       1	1171240693 ns/op	11874632 B/op	   69286 allocs/op
BenchmarkSolveMany17Clue-12               	       1	54531775049 ns/op	177058296 B/op	 1182925 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	92.457s
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
