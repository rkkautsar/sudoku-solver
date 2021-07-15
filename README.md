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

About 221s (222 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s. We should be able to achieve the same number (or even faster?) if we can reuse the SAT solver instead (AFAIK currently gophersat doesn't support that).

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveWithGophersatAiEscargot-12           	     685	   1620889 ns/op	 1508844 B/op	   17462 allocs/op
BenchmarkSolveWithGophersatHard9x9-12              	     348	   3495819 ns/op	 1858280 B/op	   20632 allocs/op
BenchmarkSolveWithGophersat17clue9x9-12            	     676	   1720905 ns/op	 1617252 B/op	   19611 allocs/op
BenchmarkSolveWithGophersat25x25-12                	      27	  41956661 ns/op	32261674 B/op	  296832 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12             	     160	   6955917 ns/op	  604066 B/op	   15825 allocs/op
BenchmarkSolveWithCadicalHard9x9-12                	     166	   7018180 ns/op	  666724 B/op	   17827 allocs/op
BenchmarkSolveWithCadicalHard17clue-12             	     158	   7659411 ns/op	  792406 B/op	   21299 allocs/op
BenchmarkSolveWithCadical25x25-12                  	      21	  54921948 ns/op	12117852 B/op	  325142 allocs/op
BenchmarkSolveWithCadical64x64-12                  	       2	 688249084 ns/op	162273080 B/op	 4206661 allocs/op
BenchmarkSolveWithCadical81x81-12                  	       1	1340205275 ns/op	320514408 B/op	 8254170 allocs/op
BenchmarkSolveWithCadical144x144-12                	       1	6522589842 ns/op	1639835504 B/op	40348928 allocs/op
BenchmarkSolveWithCadical225x225-12                	       1	31404501130 ns/op	6905996256 B/op	174948946 allocs/op
BenchmarkSolveManyWithGophersatHardest110626-12    	       1	3296745097 ns/op	1939708792 B/op	22988226 allocs/op
BenchmarkSolveManyWithGophersat17Clue-12           	       1	221472069615 ns/op	191509349544 B/op	2687855662 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	279.498s
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
