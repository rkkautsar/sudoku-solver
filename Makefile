build: cmd sudoku sudokusolver
	go build -o ./bin/sudokusolver ./cmd/sudokusolver

install: build
	go install ./...

bench: sudokusolver
	go test -run=XXX -benchmem -bench=. ./sudokusolver
