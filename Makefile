# Makefile for g4n - Cross-platform math library
# Supports Windows and Linux builds with version injection

# Variables
BINARY_NAME=g4o
GO_FILES=$(shell find . -name "*.go" -type f)
BUILD_DIR=build
INCLUDE_DIR=include

# Version information extraction
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Extract version from branch name (format: feature/x.y)
VERSION := $(shell echo "$(GIT_BRANCH)" | sed 's/.*\///g')

# Build flags with version injection
VERSION_LDFLAGS := -X 'github.com/ifabos/g4o/internal/version.Version=$(VERSION)' \
	-X 'github.com/ifabos/g4o/internal/version.GitCommit=$(GIT_COMMIT)' \
	-X 'github.com/ifabos/g4o/internal/version.GitBranch=$(GIT_BRANCH)' \
	-X 'github.com/ifabos/g4o/internal/version.BuildTime=$(BUILD_TIME)'

# Go build flags for dynamic libraries
CGO_ENABLED=1

# Windows specific settings
WINDOWS_LDFLAGS=-ldflags "-s -w $(VERSION_LDFLAGS)"
WINDOWS_BUILDMODE=-buildmode=c-shared
WINDOWS_OUTPUT=$(BUILD_DIR)/$(BINARY_NAME).dll

# Linux specific settings  
LINUX_LDFLAGS=-ldflags "-s -w $(VERSION_LDFLAGS)"
LINUX_BUILDMODE=-buildmode=c-shared
LINUX_OUTPUT=$(BUILD_DIR)/lib$(BINARY_NAME).so

# Default target
.PHONY: all
all: clean linux windows

# Create build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Build for Linux
.PHONY: linux
linux: $(BUILD_DIR)
	@echo "Building for Linux..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 \
		go build $(LINUX_BUILDMODE) $(LINUX_LDFLAGS) \
		-o $(LINUX_OUTPUT) ./cmd/library
	@echo "Linux build completed: $(LINUX_OUTPUT)"

# Build for Windows (cross-compile from Linux)
.PHONY: windows
windows: $(BUILD_DIR)
	@echo "Building for Windows..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=amd64 \
		CC=x86_64-w64-mingw32-gcc \
		go build $(WINDOWS_BUILDMODE) $(WINDOWS_LDFLAGS) \
		-o $(WINDOWS_OUTPUT) ./cmd/library
	@echo "Windows build completed: $(WINDOWS_OUTPUT)"

# Build for current platform only
.PHONY: local
local: $(BUILD_DIR)
	@echo "Building for current platform..."
	CGO_ENABLED=$(CGO_ENABLED) \
		go build $(LINUX_BUILDMODE) $(LINUX_LDFLAGS) \
		-o $(BUILD_DIR)/lib$(BINARY_NAME)_local.so ./cmd/library
	@echo "Local build completed: $(BUILD_DIR)/lib$(BINARY_NAME)_local.so"

# Test the Go modules
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...
	@echo "Tests completed."

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)
	@echo "Clean completed."

# Install dependencies (mainly for cross-compilation)
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@echo "For Windows cross-compilation, install mingw-w64:"
	@echo "  Ubuntu/Debian: sudo apt-get install gcc-mingw-w64"
	@echo "  CentOS/RHEL: sudo yum install mingw64-gcc"
	@echo "  Arch: sudo pacman -S mingw-w64-gcc"

# Package the library with headers
.PHONY: package
package: all
	@echo "Creating package..."
	mkdir -p $(BUILD_DIR)/package/lib
	mkdir -p $(BUILD_DIR)/package/include
	cp $(BUILD_DIR)/*.dll $(BUILD_DIR)/package/lib/ 2>/dev/null || true
	cp $(BUILD_DIR)/*.so $(BUILD_DIR)/package/lib/ 2>/dev/null || true
	cp $(INCLUDE_DIR)/*.h $(BUILD_DIR)/package/include/
	@echo "Package created in $(BUILD_DIR)/package/"

# Development server (for testing)
.PHONY: dev
dev:
	@echo "Building development version..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME)_dev ./cmd/library

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all      - Build for both Linux and Windows"
	@echo "  linux    - Build for Linux (.so)"
	@echo "  windows  - Build for Windows (.dll)"
	@echo "  local    - Build for current platform"
	@echo "  test     - Run Go tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  deps     - Show dependency installation instructions"
	@echo "  package  - Package library with headers"
	@echo "  dev      - Build development version"
	@echo "  help     - Show this help"

# Show version information extracted from git
.PHONY: version
version:
	@echo "Version Information Extraction:"
	@echo "  Current branch: $(GIT_BRANCH)"
	@echo "  Extracted version: $(VERSION)"
	@echo "  Git commit: $(GIT_COMMIT)"
	@echo "  Build time: $(BUILD_TIME)"
	@echo ""
	@if [ "$(VERSION)" = "unknown" ]; then \
		echo "⚠️  Warning: Could not extract version from branch name."; \
		echo "   Expected format: feature/x.y (e.g., feature/1.0)"; \
		echo "   Current branch: $(GIT_BRANCH)"; \
	else \
		echo "✅ Version successfully extracted from branch name"; \
	fi

# Show build information
.PHONY: info
info:
	@echo "Build Information:"
	@echo "  Go version: $(shell go version)"
	@echo "  Build directory: $(BUILD_DIR)"
	@echo "  Windows output: $(WINDOWS_OUTPUT)"
	@echo "  Linux output: $(LINUX_OUTPUT)"
	@echo ""
	@echo "Version Information:"
	@echo "  Git branch: $(GIT_BRANCH)"
	@echo "  Extracted version: $(VERSION)"
	@echo "  Git commit: $(GIT_COMMIT)"
	@echo "  Build time: $(BUILD_TIME)"
