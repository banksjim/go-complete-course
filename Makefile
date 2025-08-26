# Go project Makefile

# Variables
BINARY_NAME=first-app
MODULE_NAME=go-complete-course
BIN_DIR=bin

# Default target
all: build

# Build single file
build-file:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) first-app.go

# Build entire module
build-module:
	@echo "Building $(MODULE_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(MODULE_NAME) .

# Build both
build: build-file build-module

# Clean binaries
clean:
	@echo "Cleaning binaries..."
	@rm -rf $(BIN_DIR)

# Run the binary
run:
	@./$(BIN_DIR)/$(BINARY_NAME)

# Run with go run (no binary creation)
dev:
	go run first-app.go

.PHONY: all build build-file build-module clean run dev
