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

About 20s (2457 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-12               	     717	   1645334 ns/op	  491201 B/op	    3592 allocs/op
BenchmarkSolveHard9x9-12                  	     928	   1275419 ns/op	  602342 B/op	    4089 allocs/op
BenchmarkSolve17clue9x9-12                	    5529	    209680 ns/op	  184120 B/op	     890 allocs/op
BenchmarkSolve25x25-12                    	      86	  12849594 ns/op	 8071041 B/op	   57375 allocs/op
BenchmarkSolve64x64-12                    	       1	1184489944 ns/op	60433064 B/op	  253429 allocs/op
BenchmarkSolve81x81-12                    	       3	 402375792 ns/op	121324656 B/op	  592975 allocs/op
BenchmarkSolve100x100-12                  	       4	 251428482 ns/op	165872564 B/op	  450114 allocs/op
BenchmarkSolve144x144-12                  	       2	 695077576 ns/op	349800012 B/op	  158431 allocs/op
BenchmarkSolve225x225-12                  	       1	2055590642 ns/op	1351737624 B/op	  393876 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12    	     102	  10360182 ns/op	  266746 B/op	    4722 allocs/op
BenchmarkSolveWithCadicalHard9x9-12       	     120	  10183215 ns/op	  304319 B/op	    5762 allocs/op
BenchmarkSolveWithCadicalHard17clue-12    	     169	   6693674 ns/op	  131900 B/op	     669 allocs/op
BenchmarkSolveWithCadical25x25-12         	      22	  57088199 ns/op	 4485766 B/op	   79881 allocs/op
BenchmarkSolveWithCadical64x64-12         	       1	1779014885 ns/op	40840456 B/op	  330760 allocs/op
BenchmarkSolveWithCadical81x81-12         	       1	1423910268 ns/op	91083656 B/op	  809433 allocs/op
BenchmarkSolveWithCadical100x100-12       	       3	 466718087 ns/op	137814626 B/op	  637855 allocs/op
BenchmarkSolveWithCadical144x144-12       	       2	 859802836 ns/op	342126924 B/op	  243100 allocs/op
BenchmarkSolveWithCadical225x225-12       	       1	2095585899 ns/op	1336669680 B/op	  603064 allocs/op
BenchmarkSolveManyHardest110626-12        	       2	 808677660 ns/op	229225656 B/op	 1608967 allocs/op
BenchmarkSolveMany17Clue2k-12             	       2	 845363307 ns/op	597426156 B/op	 3707387 allocs/op
BenchmarkSolveMany17Clue-12               	       1	20653032719 ns/op	14621359112 B/op	90894167 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	60.287s
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
