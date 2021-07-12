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

About 70s (1.4ms per puzzle) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s but done in different machine. We should be able to achieve the same number (or even faster?) if we can reuse the SAT solver instead (AFAIK currently gophersat doesn't support that).

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveWithGophersatAiEscargot-12           	     427	   2671008 ns/op	 1517566 B/op	   13972 allocs/op
BenchmarkSolveWithGophersatHard9x9-12              	     416	   2831866 ns/op	 1606058 B/op	   14949 allocs/op
BenchmarkSolveWithGophersat17clue9x9-12            	     798	   1542301 ns/op	 1510787 B/op	   15388 allocs/op
BenchmarkSolveWithGophersat25x25-12                	      13	 198769497 ns/op	81853725 B/op	  171645 allocs/op
BenchmarkSolveWithGophersat64x64-12                	       1	1945220966 ns/op	1432926776 B/op	 1800925 allocs/op
BenchmarkSolveWithGophersat81x81-12                	       1	11398970081 ns/op	4513232568 B/op	 3411549 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12             	     109	  11186665 ns/op	  889477 B/op	   13093 allocs/op
BenchmarkSolveWithCadicalHard9x9-12                	      88	  12784883 ns/op	  913035 B/op	   13684 allocs/op
BenchmarkSolveWithCadicalHard17clue-12             	     122	  10207471 ns/op	  957309 B/op	   14713 allocs/op
BenchmarkSolveWithCadical25x25-12                  	       9	 134338770 ns/op	30209425 B/op	  203807 allocs/op
BenchmarkSolveWithCadical64x64-12                  	       1	1362732578 ns/op	1079056248 B/op	 2759366 allocs/op
BenchmarkSolveWithCadical81x81-12                  	       1	2338041475 ns/op	2432113984 B/op	 5418335 allocs/op
BenchmarkSolveWithCadical144x144-12                	       1	19602138983 ns/op	22571851160 B/op	28943003 allocs/op
BenchmarkSolveManyWithGophersatHardest110626-12    	       2	 579917183 ns/op	385855932 B/op	 5132818 allocs/op
BenchmarkSolveManyWithGophersat17Clue-12           	       1	69941989808 ns/op	50119247888 B/op	652776439 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	122.987s
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
