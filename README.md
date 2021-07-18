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

About 51.9s (946 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks on MacBook Pro (15-inch, 2019):

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-12               	     640	   1841238 ns/op	 1473245 B/op	   14840 allocs/op
BenchmarkSolveHard9x9-12                  	     502	   2329957 ns/op	 1490162 B/op	   15318 allocs/op
BenchmarkSolve17clue9x9-12                	     630	   1893870 ns/op	 1565765 B/op	   16695 allocs/op
BenchmarkSolve25x25-12                    	      26	  39358552 ns/op	25523239 B/op	  216653 allocs/op
BenchmarkSolve64x64-12                    	       1	1883747768 ns/op	406859344 B/op	 3087570 allocs/op
BenchmarkSolve81x81-12                    	       1	2074903980 ns/op	814669112 B/op	 5846540 allocs/op
BenchmarkSolve100x100-12                  	       1	1942511736 ns/op	1503319000 B/op	 9391110 allocs/op
BenchmarkSolve144x144-12                  	       1	5225423238 ns/op	3787917824 B/op	24658830 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12    	      54	  23870606 ns/op	  816679 B/op	   17116 allocs/op
BenchmarkSolveWithCadicalHard9x9-12       	      67	  23467677 ns/op	  833513 B/op	   17868 allocs/op
BenchmarkSolveWithCadicalHard17clue-12    	      61	  18024857 ns/op	  886772 B/op	   19762 allocs/op
BenchmarkSolveWithCadical25x25-12         	      12	  94144301 ns/op	15625046 B/op	  282364 allocs/op
BenchmarkSolveWithCadical64x64-12         	       1	1955876941 ns/op	278617648 B/op	 4163922 allocs/op
BenchmarkSolveWithCadical81x81-12         	       1	2441800155 ns/op	557864976 B/op	 8025225 allocs/op
BenchmarkSolveWithCadical100x100-12       	       1	3267921925 ns/op	1007263048 B/op	13333760 allocs/op
BenchmarkSolveWithCadical144x144-12       	       1	8815720086 ns/op	2814243288 B/op	36203749 allocs/op
BenchmarkSolveManyHardest110626-12        	       1	1800657236 ns/op	34435800 B/op	  383468 allocs/op
BenchmarkSolveMany17Clue2k-12             	       1	1569262306 ns/op	87826232 B/op	 1454889 allocs/op
BenchmarkSolveMany17Clue-12               	       1	51935033960 ns/op	2183699544 B/op	35075295 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	94.932s
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
