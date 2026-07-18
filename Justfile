#!/usr/bin/env just --justfile

# Default recipe to list available commands
default:
    @just --list

# Build the commit-increment binary
build:
    go build -o commit-increment .

# Run tests
test:
    go test -v ./...

# Run tests with coverage
coverage:
    go test -v -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
    rm -f commit-increment coverage.out coverage.html

# Run the binary with example usage
run *ARGS:
    go run . {{ARGS}}

# Format code
fmt:
    go fmt ./...

# Lint code
lint:
    golangci-lint run ./...

# Install dependencies
deps:
    go mod download
    go mod tidy

# Build and run tests
all: build test

# Help for building
help:
    @echo "Available recipes:"
    @echo "  build     - Build the commit-increment binary"
    @echo "  test      - Run tests"
    @echo "  coverage  - Generate coverage report"
    @echo "  clean     - Remove build artifacts"
    @echo "  run       - Run the binary directly"
    @echo "  fmt       - Format code"
    @echo "  lint      - Lint code"
    @echo "  deps      - Download and tidy dependencies"
    @echo "  all       - Build and run tests"
