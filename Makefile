build: cmd sudoku sudokusolver
	go build -o ./bin/sudokusolver ./cmd/sudokusolver

install: build
	go install ./...

test: sudoku sudokusolver
	go test -cover ./...

bench: sudokusolver
	go test -run=XXX -benchmem -bench=. ./sudokusolver

profile_many: build
	./bin/sudokusolver -many -cpuprofile=cpu.prof -memprofile=mem.prof
