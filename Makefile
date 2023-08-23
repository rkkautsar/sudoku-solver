build: cmd sudoku sudokusolver
	go build -o ./bin/sudokusolver ./cmd/sudokusolver

install: build
	go install ./...

test: sudoku sudokusolver
	go test -cover ./...

bench: sudokusolver
	if ! command -v cadical &> /dev/null; then echo "[!] Please install cadical first: $$(tput bold)brew install cadical$$(tput sgr0)"; exit 1; fi
	go test -run=XXX -benchmem -bench=. ./sudokusolver

benchprofile:
	go test -run=XXX -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -bench=. ./sudokusolver

profile: build
	./bin/sudokusolver -cpuprofile=cpu.prof -memprofile=mem.prof ${ARGS}
