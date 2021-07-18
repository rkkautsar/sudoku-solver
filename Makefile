build: cmd sudoku sudokusolver
	go build -o ./bin/sudokusolver ./cmd/sudokusolver

install: build
	go install ./...

test: sudoku sudokusolver
	go test -cover ./...

bench: sudokusolver
	go test -run=XXX -benchmem -bench=. ./sudokusolver

benchprofile:
	go test -run=XXX -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -bench=. ./sudokusolver

profile: build
	./bin/sudokusolver -cpuprofile=cpu.prof -memprofile=mem.prof ${ARGS}
