.PHONY: build test lint clean all

# Default binary name
BINARY=gotypeslearn

# Version information
VERSION?=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"

# Default target
all: clean test build

# Build the application
build:
	@echo "Building $(BINARY)..."
	go build $(LDFLAGS) -o $(BINARY)
	@echo "Done!"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run linting
lint:
	@echo "Running linters..."
	go fmt ./...
	go vet ./...
	@which golint > /dev/null && golint ./... || echo "golint not installed"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY)
	go clean

# Install the binary to GOPATH/bin
install: build
	@echo "Installing..."
	go install

# Cross-compilation for various platforms
cross-build:
	@echo "Cross-compiling for multiple platforms..."
	# Windows
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-windows-amd64.exe
	# Linux
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-linux-amd64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY)-linux-arm64
	# macOS
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY)-darwin-arm64
	@echo "Cross-compilation complete."

# Run the application
run: build
	./$(BINARY)

# Help target
help:
	@echo "Available targets:"
	@echo "  all         - Run clean, test, and build"
	@echo "  build       - Build the application"
	@echo "  test        - Run tests"
	@echo "  lint        - Run linters (fmt, vet, lint)"
	@echo "  clean       - Clean build artifacts"
	@echo "  install     - Install to GOPATH/bin"
	@echo "  cross-build - Build for multiple platforms"
	@echo "  run         - Build and run the application"