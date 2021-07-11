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
- fast but not as fast as specialized solvers (3.7ms for ai-escargot)
- pretty fast for larger sudokus (with custom state of the art SAT solver like cadical, 144x144 took 22s, may be better if using parallel solvers)
- pretty simple code

## Benchmarks

About 120s (100s before using goroutine) to solve the benchmark here (fastest is 0.2s :eyes:): https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Other benchmarks:

```
> go test -run=XXX -bench=./sudokusolver -benchmem
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveWithGophersatAiEscargot-12    	     318	   3713396 ns/op	 1653520 B/op	   15804 allocs/op
BenchmarkSolveWithGophersatHard9x9-12       	     268	   4340095 ns/op	 1770123 B/op	   16784 allocs/op
BenchmarkSolveWithGophersat17clue9x9-12     	     519	   2243898 ns/op	 1594683 B/op	   16876 allocs/op
BenchmarkSolveWithGophersat25x25-12         	       6	 203096073 ns/op	83865096 B/op	  205988 allocs/op
BenchmarkSolveWithGophersat64x64-12         	       1	6170177136 ns/op	2195560920 B/op	 2596360 allocs/op
BenchmarkSolveWithGophersat81x81-12         	       1	13232220182 ns/op	5017162712 B/op	 5039706 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12      	      98	  10866806 ns/op	  963540 B/op	   14770 allocs/op
BenchmarkSolveWithCadicalHard9x9-12         	      97	  12835094 ns/op	  981650 B/op	   15279 allocs/op
BenchmarkSolveWithCadicalHard17clue-12      	     100	  10040787 ns/op	 1018491 B/op	   16095 allocs/op
BenchmarkSolveWitCadical25x25-12            	       7	 160545458 ns/op	31089931 B/op	  237506 allocs/op
BenchmarkSolveWitCadical64x64-12            	       1	1750162007 ns/op	1099779552 B/op	 3539349 allocs/op
BenchmarkSolveWitCadical81x81-12            	       1	2946934571 ns/op	2475453808 B/op	 7040300 allocs/op
BenchmarkSolveWitCadical144x144-12          	       1	26018999229 ns/op	22835248400 B/op	38713298 allocs/op
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
./sudokusilver -solve < data/sudoku-9-1.txt

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
