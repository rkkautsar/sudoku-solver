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
- fast, but not as fast as specialized solvers (1.8ms for ai-escargot, naïve backtracking is around 30ms)
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
BenchmarkSolveAiEscargot-12               	     716	   1806322 ns/op	  490891 B/op	    3592 allocs/op
BenchmarkSolveHard9x9-12                  	     912	   1285606 ns/op	  609571 B/op	    4249 allocs/op
BenchmarkSolve17clue9x9-12                	    1263	    934811 ns/op	  905213 B/op	    5778 allocs/op
BenchmarkSolve25x25-12                    	      64	  16796403 ns/op	 8465768 B/op	   64830 allocs/op
BenchmarkSolve64x64-12                    	       1	2054103411 ns/op	72338888 B/op	  311863 allocs/op
BenchmarkSolve81x81-12                    	       3	 343304668 ns/op	134011749 B/op	  670371 allocs/op
BenchmarkSolve100x100-12                  	       5	 228876627 ns/op	173762048 B/op	  568441 allocs/op
BenchmarkSolve144x144-12                  	       2	 654317958 ns/op	540966420 B/op	 1775294 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12    	      78	  13209624 ns/op	  270865 B/op	    4723 allocs/op
BenchmarkSolveWithCadicalHard9x9-12       	     120	  10153001 ns/op	  310502 B/op	    6029 allocs/op
BenchmarkSolveWithCadicalHard17clue-12    	     100	  10160937 ns/op	  477742 B/op	    8500 allocs/op
BenchmarkSolveWithCadical25x25-12         	      24	  46477649 ns/op	 4841759 B/op	   90033 allocs/op
BenchmarkSolveWithCadical64x64-12         	       1	1834994077 ns/op	43759320 B/op	  409028 allocs/op
BenchmarkSolveWithCadical64x64Hard-12     	       1	66145999146 ns/op	50348200 B/op	  574721 allocs/op
BenchmarkSolveWithCadical81x81-12         	       2	 658297041 ns/op	95692532 B/op	  924668 allocs/op
BenchmarkSolveWithCadical100x100-12       	       3	 501373854 ns/op	144273909 B/op	  811773 allocs/op
BenchmarkSolveWithCadical144x144-12       	       1	1429326864 ns/op	430382712 B/op	 2491800 allocs/op
BenchmarkSolveManyHardest110626-12        	       2	 817315026 ns/op	232346776 B/op	 1638300 allocs/op
BenchmarkSolveMany17Clue2k-12             	       1	2133558634 ns/op	1782471736 B/op	11993486 allocs/op
BenchmarkSolveMany17Clue-12               	       1	51558690517 ns/op	44412021688 B/op	298044564 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	151.958s
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
