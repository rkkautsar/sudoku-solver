build: cmd sudoku sudokusolver
	go build -o ./bin/sudokusolver ./cmd/sudokusolver

install: build
	go install ./...
