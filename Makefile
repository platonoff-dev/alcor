.SUFFIXES: .peg .go

.peg.go:
	peg -switch -inline -strict -output scanner/scanner.go $<

all: grammar.go