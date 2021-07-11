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

About 94s to solve the benchmark here (fastest is 0.2s :eyes:): https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveWithGophersatAiEscargot-12           	     412	   2732315 ns/op	 1530785 B/op	   14615 allocs/op
BenchmarkSolveWithGophersatHard9x9-12              	     423	   2918257 ns/op	 1618744 B/op	   15580 allocs/op
BenchmarkSolveWithGophersat17clue9x9-12            	     709	   1675453 ns/op	 1522718 B/op	   15970 allocs/op
BenchmarkSolveWithGophersat25x25-12                	      16	 268754379 ns/op	86939063 B/op	  206440 allocs/op
BenchmarkSolveWithGophersat64x64-12                	       1	2582235586 ns/op	1532554568 B/op	 2585080 allocs/op
BenchmarkSolveWithGophersat81x81-12                	       1	11991080500 ns/op	4707919632 B/op	 5038647 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12             	     104	  12126437 ns/op	  896867 B/op	   13649 allocs/op
BenchmarkSolveWithCadicalHard9x9-12                	      86	  13233215 ns/op	  919754 B/op	   14223 allocs/op
BenchmarkSolveWithCadicalHard17clue-12             	     122	   9536498 ns/op	  963679 B/op	   15207 allocs/op
BenchmarkSolveWithCadical25x25-12                  	       8	 161531516 ns/op	31091319 B/op	  237504 allocs/op
BenchmarkSolveWithCadical64x64-12                  	       1	1617848709 ns/op	1099744112 B/op	 3539319 allocs/op
BenchmarkSolveWithCadical81x81-12                  	       1	3015859903 ns/op	2475459944 B/op	 7040312 allocs/op
BenchmarkSolveWithCadical144x144-12                	       1	23423450732 ns/op	22835249400 B/op	38713340 allocs/op
BenchmarkSolveManyWithGophersatHardest110626-12    	       1	1634277900 ns/op	671532304 B/op	 5957585 allocs/op
BenchmarkSolveManyWithGophersat17Clue-12           	       1	94037820983 ns/op	78070546592 B/op	786381350 allocs/op
PASS
```

## Getting Started

### Dependencies

- Go

### Installing

```sh
go get github.com/rkkautsar/sudoku-solver
go install github.com/rkkautsar/sudoku-solver
```

### Executing program

```sh
./sudokusolver -help
./sudokusolver -cnf < data/sudoku-9-1.txt
./sudokusolver -solve < data/sudoku-9-1.txt
./sudokusolver -solve -many < data/sudoku.many.2k.txt

# brew install cadical
./sudokusilver -solver "cadical -q" < data/sudoku-9-1.txt
```

## License

This project is licensed under the MIT License - see the LICENSE.md file for details

## Acknowledgments

Inspiration, code snippets, etc.

- [gophersat](https://github.com/crillab/gophersat)
- Kwon, Gihwon, and Himanshu Jain. "Optimized CNF encoding for sudoku puzzles." Proc. 13th International Conference on Logic for Programming Artificial Intelligence and Reasoning (LPAR2006). 2006. ([PDF](http://www.cs.cmu.edu/~hjain/papers/sudoku-as-SAT.pdf))
- Klieber, Will, and Gihwon Kwon. "Efficient CNF encoding for selecting 1 from n objects." Proc. International Workshop on Constraints in Formal Verification. 2007. ([PDF](https://www.cs.cmu.edu/~wklieber/papers/2007_efficient-cnf-encoding-for-selecting-1.pdf))
- Nguyen, Van-Hau, and Son T. Mai. "A new method to encode the at-most-one constraint into SAT." Proceedings of the Sixth International Symposium on Information and Communication Technology. 2015. ([PDF](https://www.researchgate.net/profile/Van-Hau-Nguyen/publication/301455290_A_New_Method_to_Encode_the_At-Most-One_Constraint_into_SAT/links/5d2bfbaba6fdcc2462e0e269/A-New-Method-to-Encode-the-At-Most-One-Constraint-into-SAT.pdf))
