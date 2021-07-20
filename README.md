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

About 24s (2047 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-12               	     688	   1702221 ns/op	  491228 B/op	    3592 allocs/op
BenchmarkSolveHard9x9-12                  	     888	   1339579 ns/op	  602087 B/op	    4089 allocs/op
BenchmarkSolve17clue9x9-12                	    4234	    285438 ns/op	  184120 B/op	     890 allocs/op
BenchmarkSolve25x25-12                    	      74	  15093489 ns/op	 8071064 B/op	   57375 allocs/op
BenchmarkSolve64x64-12                    	       1	1301842971 ns/op	60469976 B/op	  253431 allocs/op
BenchmarkSolve81x81-12                    	       2	 525297125 ns/op	121324236 B/op	  592974 allocs/op
BenchmarkSolve100x100-12                  	       3	 361291717 ns/op	165881909 B/op	  450116 allocs/op
BenchmarkSolve144x144-12                  	       2	 809432180 ns/op	349800128 B/op	  158432 allocs/op
BenchmarkSolve225x225-12                  	       1	3011330858 ns/op	1351737624 B/op	  393876 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12    	     114	  10999670 ns/op	  269628 B/op	    4722 allocs/op
BenchmarkSolveWithCadicalHard9x9-12       	     106	  11112477 ns/op	  304917 B/op	    5762 allocs/op
BenchmarkSolveWithCadicalHard17clue-12    	     154	   6850171 ns/op	  130345 B/op	     668 allocs/op
BenchmarkSolveWithCadical25x25-12         	      27	  42531117 ns/op	 4485805 B/op	   79881 allocs/op
BenchmarkSolveWithCadical64x64-12         	       1	1845246530 ns/op	40840440 B/op	  330760 allocs/op
BenchmarkSolveWithCadical81x81-12         	       1	1488105064 ns/op	91083656 B/op	  809433 allocs/op
BenchmarkSolveWithCadical100x100-12       	       2	 547152644 ns/op	137814748 B/op	  637857 allocs/op
BenchmarkSolveWithCadical144x144-12       	       2	1017842537 ns/op	342127060 B/op	  243101 allocs/op
BenchmarkSolveWithCadical225x225-12       	       1	2978402073 ns/op	1336669584 B/op	  603063 allocs/op
BenchmarkSolveManyHardest110626-12        	       2	 836910400 ns/op	229778124 B/op	 1609714 allocs/op
BenchmarkSolveMany17Clue2k-12             	       2	 974799328 ns/op	600371616 B/op	 3711387 allocs/op
BenchmarkSolveMany17Clue-12               	       1	24084093210 ns/op	14693746808 B/op	90992594 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	65.757s
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
