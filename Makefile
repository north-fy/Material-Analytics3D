.PHONY: build

build:
	go build -v -o ./build/windows/ ./

# БИЛДЫ ПОД РАЗНЫЕ ПЛАТФОРМЫ (кроме WINDOWS)

.PHONY: build-linux

build-linux:
	GOOS=darwin GOARCH=amd64 go build -v -o ./build/linux/ ./

.PHONY: build-macos

build-macos:
	GOOS=linux GOARCH=amd64 go build -v -o ./build/darwin/ ./

# ТЕСТЫ

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: cover-test
cover-test:
	go test -coverprofile=c.out -coverpkg=./... ./...;go tool cover -html=./c.out -o .test/test-coverage.html

# ПО УМОЛЧАНИЮ: GOOS=windows GOARCH=amd64 go build -o ...
.DEFAULT_GOAL := build



