.Default_Goal := all
PROJECT_NAME := "gitgo"
BINARY_NAME := "gitgo"

.PHONY : all clean test test-verbose run build help

all: clean build test
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME) || true

test: build 
	go test -v ./...


build:
	go build -o $(BINARY_NAME) main.go
	chmod +x $(BINARY_NAME)