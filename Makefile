.PHONY: build

build:
	go build -v ./

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: cover-test
cover-test:
	go test -coverprofile=c.out -coverpkg=./... ./...;go tool cover -html=./c.out -o .test/test-coverage.html


.DEFAULT_GOAL := build



