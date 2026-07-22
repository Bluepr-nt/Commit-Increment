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

# Build the GitHub org runner image
runner-build:
    docker build -f Dockerfile.runner -t gh-org-runner .

# Run the GitHub org runner container (requires GH_PAT env vars)
runner-run name='gh-org-runner-1' labels='docker,linux,x64' group='Default':
    test -n "$GH_PAT" || (echo "GH_PAT is required" && exit 1)
    docker run -d \
        --name {{name}} \
        -e GH_PAT="$GH_PAT" \
        -e RUNNER_NAME="{{name}}" \
        -e RUNNER_LABELS="{{labels}}" \
        -e RUNNER_GROUP="{{group}}" \
        gh-org-runner

# Run the GitHub org runner container with dlv debug port exposed (attach VS Code to localhost:40000)
runner-debug name='gh-org-runner-debug' labels='docker,linux,x64' group='Default':
    test -n "$GH_PAT" || (echo "GH_PAT is required" && exit 1)
    docker run \
        --name {{name}} \
        -e GH_PAT="$GH_PAT" \
        -e RUNNER_NAME="{{name}}" \
        -e RUNNER_LABELS="{{labels}}" \
        -e RUNNER_GROUP="{{group}}" \
        -e DEBUG_DLV=true \
        -p 40000:40000 \
        gh-org-runner

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
    @echo "  runner-build - Build the GitHub org runner Docker image"
    @echo "  runner-run   - Run the GitHub org runner container"
