# Chronus Makefile

.PHONY: build test lint clean example fmt vet install-tools help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOVET=$(GOCMD) vet

# Project variables
BINARY_NAME=chronus
CLI_BINARY=chronus-cli
EXAMPLE_DIR=./example
CLI_DIR=./cmd/chronus
EXAMPLE_BINARY=$(EXAMPLE_DIR)/example

# Default target
all: test build

# Build the library
build:
	@echo "Building library..."
	$(GOBUILD) -v ./...

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -func=coverage.out

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

# Vet code
vet:
	@echo "Vetting code..."
	$(GOVET) ./...

# Lint code (requires golangci-lint)
lint: install-golangci-lint
	@echo "Linting code..."
	golangci-lint run

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(EXAMPLE_BINARY)
	rm -f $(CLI_DIR)/$(CLI_BINARY)
	rm -f coverage.out coverage.html

# Build and run example
example:
	@echo "Building and running example..."
	cd $(EXAMPLE_DIR) && $(GOBUILD) -o example .
	cd $(EXAMPLE_DIR) && ./example

# Build example only
build-example:
	@echo "Building example..."
	cd $(EXAMPLE_DIR) && $(GOBUILD) -o example .

# Build CLI tool
build-cli:
	@echo "Building CLI tool..."
	cd $(CLI_DIR) && $(GOBUILD) -o $(CLI_BINARY) .

# Install CLI tool
install-cli: build-cli
	@echo "Installing CLI tool..."
	cd $(CLI_DIR) && go install .

# Install development tools
install-tools: install-golangci-lint

# Install golangci-lint
install-golangci-lint:
	@echo "Installing golangci-lint..."
	@which golangci-lint > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download

# Verify dependencies
verify:
	@echo "Verifying dependencies..."
	$(GOMOD) verify

# Check for vulnerabilities
vuln-check:
	@echo "Checking for vulnerabilities..."
	$(GOCMD) list -json -m all | nancy sleuth

# Pre-commit checks
pre-commit: fmt vet lint test

# Release preparation
release-check: pre-commit
	@echo "Release check completed successfully!"

# Help
help:
	@echo "Available targets:"
	@echo "  build         - Build the library"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  fmt           - Format code"
	@echo "  vet           - Vet code"
	@echo "  lint          - Lint code (requires golangci-lint)"
	@echo "  clean         - Clean build artifacts"
	@echo "  example       - Build and run example"
	@echo "  build-example - Build example only"
	@echo "  build-cli     - Build CLI tool"
	@echo "  install-cli   - Install CLI tool"
	@echo "  install-tools - Install development tools"
	@echo "  tidy          - Tidy dependencies"
	@echo "  deps          - Download dependencies"
	@echo "  verify        - Verify dependencies"
	@echo "  vuln-check    - Check for vulnerabilities"
	@echo "  pre-commit    - Run pre-commit checks"
	@echo "  release-check - Check if ready for release"
	@echo "  help          - Show this help"
