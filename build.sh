#!/bin/bash

# Go build script - outputs binaries to bin folder

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Variables
BIN_DIR="bin"
BINARY_NAME="first-app"
MODULE_NAME="go-complete-course"

echo -e "${YELLOW}Building Go binaries...${NC}"

# Create bin directory if it doesn't exist
mkdir -p "$BIN_DIR"

# Build single file
echo -e "${YELLOW}Building $BINARY_NAME...${NC}"
if go build -o "$BIN_DIR/$BINARY_NAME" first-app.go; then
    echo -e "${GREEN}✓ Successfully built $BIN_DIR/$BINARY_NAME${NC}"
else
    echo -e "${RED}✗ Failed to build $BINARY_NAME${NC}"
    exit 1
fi

# Build entire module
echo -e "${YELLOW}Building $MODULE_NAME...${NC}"
if go build -o "$BIN_DIR/$MODULE_NAME" .; then
    echo -e "${GREEN}✓ Successfully built $BIN_DIR/$MODULE_NAME${NC}"
else
    echo -e "${RED}✗ Failed to build $MODULE_NAME${NC}"
    exit 1
fi

echo -e "${GREEN}All binaries built successfully!${NC}"
echo -e "${YELLOW}Binaries location: $(pwd)/$BIN_DIR/${NC}"
ls -la "$BIN_DIR/"
