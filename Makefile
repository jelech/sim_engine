# SimEngine Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Project parameters
BINARY_NAME=sim_engine
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PATH=./cmd/sim_engine

# Build parameters
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
COMMIT ?= $(shell git rev-parse HEAD)

LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.commit=$(COMMIT)"

.PHONY: all build clean test coverage lint fmt vet deps help install-tools

# Default target
all: test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PATH)

# Build for production
build-prod:
	@echo "Building $(BINARY_NAME) for production..."
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) $(LDFLAGS) -a -installsuffix cgo -o $(BINARY_NAME) $(MAIN_PATH)

# Cross-platform builds
build-linux:
	@echo "Building for Linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_UNIX) $(MAIN_PATH)

build-windows:
	@echo "Building for Windows..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME).exe $(MAIN_PATH)

build-darwin:
	@echo "Building for macOS..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME)_darwin $(MAIN_PATH)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with race detection
test-race:
	@echo "Running tests with race detection..."
	$(GOTEST) -race -v ./...

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Generate test coverage
coverage:
	@echo "Generating test coverage..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
	@echo "Running linter..."
	$(GOLINT) run

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .
	goimports -w .

# Run go vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

# Update dependencies
deps:
	@echo "Updating dependencies..."
	$(GOMOD) tidy
	$(GOMOD) verify

# Install development tools
install-tools:
	@echo "Installing development tools..."
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOGET) golang.org/x/tools/cmd/goimports@latest
	$(GOGET) golang.org/x/tools/cmd/godoc@latest

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_NAME).exe
	rm -f $(BINARY_NAME)_darwin
	rm -f coverage.out
	rm -f coverage.html

# Run all checks (formatting, linting, vetting, testing)
check: fmt vet lint test

# Generate documentation
docs:
	@echo "Generating documentation..."
	godoc -http=:6060
	@echo "Documentation server started at http://localhost:6060"

# Run examples
run-examples:
	@echo "Running basic example..."
	$(GOCMD) run examples/basic/main.go

# Development server (auto-reload)
dev:
	@echo "Starting development server..."
	air

# Install the binary
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GOCMD) install $(LDFLAGS) $(MAIN_PATH)

# Create release archives
release: clean build-linux build-windows build-darwin
	@echo "Creating release archives..."
	mkdir -p dist
	tar -czf dist/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz $(BINARY_UNIX) README.md LICENSE
	zip -r dist/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BINARY_NAME).exe README.md LICENSE
	tar -czf dist/$(BINARY_NAME)-$(VERSION)-darwin-amd64.tar.gz $(BINARY_NAME)_darwin README.md LICENSE

# Docker targets
docker-build:
	@echo "Building Docker image..."
	docker build -t sim_engine:$(VERSION) .

docker-run:
	@echo "Running Docker container..."
	docker run --rm -it sim_engine:$(VERSION)

# Help target
help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  build-prod    - Build for production"
	@echo "  build-linux   - Build for Linux"
	@echo "  build-windows - Build for Windows"
	@echo "  build-darwin  - Build for macOS"
	@echo "  test          - Run tests"
	@echo "  test-race     - Run tests with race detection"
	@echo "  bench         - Run benchmarks"
	@echo "  coverage      - Generate test coverage report"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  deps          - Update dependencies"
	@echo "  install-tools - Install development tools"
	@echo "  clean         - Clean build artifacts"
	@echo "  check         - Run all checks"
	@echo "  docs          - Start documentation server"
	@echo "  run-examples  - Run example code"
	@echo "  dev           - Start development server"
	@echo "  install       - Install the binary"
	@echo "  release       - Create release archives"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container"
	@echo "  help          - Show this help message"
