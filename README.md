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
- fast, but not as fast as specialized solvers (1ms for ai-escargot, naïve backtracking is around 30ms)
- pretty fast for larger sudokus, for example a 144x144 sudoku can be solved in 4s
- pretty simple code

## Benchmarks

About 10.7s (4593 puzzle/s) to solve the benchmark of [49k 17-clue 9x9 sudoku](data/sudoku.many.17clue.txt) from here: https://codegolf.stackexchange.com/questions/190727/the-fastest-sudoku-solver

Fastest on this benchmark is [tdoku](https://www.github.com/t-dillon/tdoku) which took 0.2s to complete :rocket:. Other SAT-based solver with minisat took 11.7s.

Other benchmarks available in `make bench`:

```
> make bench
go test -run=XXX -benchmem -bench=. ./sudokusolver
goos: linux
goarch: amd64
pkg: github.com/rkkautsar/sudoku-solver/sudokusolver
cpu: AMD Ryzen 5 3600X 6-Core Processor
BenchmarkSolveAiEscargot-12                         1215           1023191 ns/op          493599 B/op       3499 allocs/op
BenchmarkSolveHard9x9-12                            1546            762088 ns/op          595014 B/op       3997 allocs/op
BenchmarkSolve17clue9x9-12                         13698             85393 ns/op          185647 B/op        568 allocs/op
BenchmarkSolve25x25-12                               128           9396040 ns/op         8907000 B/op      62807 allocs/op
BenchmarkSolve64x64-12                                 2         676392630 ns/op        66929668 B/op     296221 allocs/op
BenchmarkSolve81x81-12                                 3         346799245 ns/op        192092229 B/op    963324 allocs/op
BenchmarkSolve100x100-12                               7         143129774 ns/op        171782198 B/op    472035 allocs/op
BenchmarkSolve144x144-12                               3         407516616 ns/op        349730328 B/op     75485 allocs/op
BenchmarkSolve225x225-12                               1        1156088001 ns/op        1351336632 B/op   191375 allocs/op
BenchmarkSolveWithCadicalAiEscargot-12               468           2450860 ns/op          265146 B/op       4617 allocs/op
BenchmarkSolveWithCadicalHard9x9-12                  603           2116695 ns/op          299252 B/op       5659 allocs/op
BenchmarkSolveWithCadicalHard17clue-12              1292            915425 ns/op          126206 B/op        336 allocs/op
BenchmarkSolveWithCadical25x25-12                     66          18147812 ns/op         5027100 B/op      92772 allocs/op
BenchmarkSolveWithCadical64x64-12                      1        1073545781 ns/op        45360520 B/op     440751 allocs/op
BenchmarkSolveWithCadical81x81-12                      2         590886469 ns/op        123323796 B/op   1693362 allocs/op
BenchmarkSolveWithCadical100x100-12                    5         232460068 ns/op        141966470 B/op    717556 allocs/op
BenchmarkSolveWithCadical144x144-12                    3         402889437 ns/op        342061024 B/op    160144 allocs/op
BenchmarkSolveWithCadical225x225-12                    1        1170865434 ns/op        1336467432 B/op   400554 allocs/op
BenchmarkSolveManyHardest110626-12                     3         489099958 ns/op        225490930 B/op   1574077 allocs/op
BenchmarkSolveMany17Clue2k-12                          3         429496144 ns/op        590409965 B/op   3235638 allocs/op
BenchmarkSolveMany17Clue-12                            1        10773818004 ns/op       14442420704 B/op        79414952 allocs/op
PASS
ok      github.com/rkkautsar/sudoku-solver/sudokusolver 48.061s
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
