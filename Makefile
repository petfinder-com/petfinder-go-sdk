.PHONY: all build

all: test build

unit_test:
	go test -v ./... -cover

test: unit_test