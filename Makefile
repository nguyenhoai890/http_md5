.PHONY: build test

build:
	go build -o build/myhttp main.go

test:
	go test -race -v .