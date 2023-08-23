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
- pretty fast for larger sudokus, for example a 144x144 sudoku can be solved in 4s
- pretty simple code

## Benchmarks

About 5.8s (8474 puzzle/s or 118 ns/puzzle) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks available in `make bench`:

```
↳ make bench
if ! command -v cadical &> /dev/null; then echo "[!] Please install cadical first: $(tput bold)brew install cadical$(tput sgr0)"; exit 1; fi
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: darwin
goarch: arm64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
BenchmarkSolveAiEscargot-10               	    1911	    647132 ns/op	  369585 B/op	    2929 allocs/op
BenchmarkSolveHard9x9-10                  	    2706	    431285 ns/op	  396274 B/op	    3352 allocs/op
BenchmarkSolve17clue9x9-10                	   28108	     47913 ns/op	  107654 B/op	     634 allocs/op
BenchmarkSolve25x25-10                    	     262	   4780719 ns/op	 5455202 B/op	   43100 allocs/op
BenchmarkSolve64x64-10                    	       4	 332938156 ns/op	33512906 B/op	  206324 allocs/op
BenchmarkSolve81x81-10                    	       5	 213718425 ns/op	101065350 B/op	  771714 allocs/op
BenchmarkSolve100x100-10                  	      18	  63925764 ns/op	71169493 B/op	  340853 allocs/op
BenchmarkSolve144x144-10                  	       5	 249753067 ns/op	111240065 B/op	  158053 allocs/op
BenchmarkSolve225x225-10                  	       3	 415970389 ns/op	444639680 B/op	  393479 allocs/op
BenchmarkSolveWithCadicalAiEscargot-10    	     322	   3993171 ns/op	  344782 B/op	    5535 allocs/op
BenchmarkSolveWithCadicalHard9x9-10       	     381	   3329669 ns/op	  444484 B/op	    6674 allocs/op
BenchmarkSolveWithCadicalHard17clue-10    	     598	   2121930 ns/op	  127355 B/op	     720 allocs/op
BenchmarkSolveWithCadical25x25-10         	     100	  12874129 ns/op	 5965580 B/op	  101945 allocs/op
BenchmarkSolveWithCadical64x64-10         	       2	 682339458 ns/op	29716664 B/op	  474691 allocs/op
BenchmarkSolveWithCadical81x81-10         	       4	 321655042 ns/op	110022726 B/op	 1865605 allocs/op
BenchmarkSolveWithCadical100x100-10       	       8	 126807734 ns/op	74976945 B/op	  779323 allocs/op
BenchmarkSolveWithCadical144x144-10       	       4	 256900156 ns/op	111567778 B/op	  181708 allocs/op
BenchmarkSolveWithCadical225x225-10       	       2	 504154167 ns/op	445448628 B/op	  452004 allocs/op
BenchmarkSolveManyHardest110626-10        	       4	 327428104 ns/op	149366518 B/op	 1302273 allocs/op
BenchmarkSolveMany17Clue2k-10             	       5	 236078600 ns/op	365466395 B/op	 2980678 allocs/op
BenchmarkSolveMany17Clue-10               	       1	5851515333 ns/op	8961331456 B/op	73508710 allocs/op
PASS
ok  	github.com/rkkautsar/sudoku-solver/sudokusolver	45.391s
```

## Getting Started

### Dependencies

- Go

### Installing

#### From gobinaries

```sh
curl -sf https://gobinaries.com/rkkautsar/sudoku-solver/cmd/sudokusolver | sh
```

#### From source

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
sudokusolver -solver "cadical -q" < data/sudoku-9-1.txt
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
