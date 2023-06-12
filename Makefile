run: build
	@./bin/alejandro

build:
	@go build -o bin/alejandro

test:
	@go test -v ./...

all: run

.PHONY: run build test all