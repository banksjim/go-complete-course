#!/bin/bash

# Go Complete Course Build Script - supports multiple lessons

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Variables
BIN_DIR="bin"
LESSONS_DIR="lessons"

echo -e "${BLUE}Go Complete Course - Build Script${NC}"

# Function to build all lessons
build_all() {
    echo -e "${YELLOW}Building all lessons...${NC}"
    
    # Create bin directory if it doesn't exist
    mkdir -p "$BIN_DIR"
    
    # Find all lesson directories
    for lesson_dir in "$LESSONS_DIR"/*; do
        if [ -d "$lesson_dir" ]; then
            lesson_name=$(basename "$lesson_dir")
            echo -e "${YELLOW}Building $lesson_name...${NC}"
            
            if go build -o "$BIN_DIR/$lesson_name" "$lesson_dir"/*.go; then
                echo -e "${GREEN}✓ Successfully built $lesson_name${NC}"
            else
                echo -e "${RED}✗ Failed to build $lesson_name${NC}"
                exit 1
            fi
        fi
    done
    
    echo -e "${GREEN}All lessons built successfully!${NC}"
    echo -e "${BLUE}Binaries location: $(pwd)/$BIN_DIR/${NC}"
    ls -la "$BIN_DIR/"
}

# Function to build specific lesson
build_lesson() {
    local lesson=$1
    if [ -z "$lesson" ]; then
        echo -e "${RED}Error: Please specify a lesson name${NC}"
        echo -e "${YELLOW}Usage: $0 lesson <lesson-name>${NC}"
        echo -e "${YELLOW}Example: $0 lesson 01-hello-world${NC}"
        exit 1
    fi
    
    lesson_path="$LESSONS_DIR/$lesson"
    if [ ! -d "$lesson_path" ]; then
        echo -e "${RED}Error: Lesson '$lesson' not found${NC}"
        echo -e "${YELLOW}Available lessons:${NC}"
        ls -1 "$LESSONS_DIR" 2>/dev/null || echo -e "${RED}No lessons found${NC}"
        exit 1
    fi
    
    echo -e "${YELLOW}Building lesson $lesson...${NC}"
    mkdir -p "$BIN_DIR"
    
    if go build -o "$BIN_DIR/$lesson" "$lesson_path"/*.go; then
        echo -e "${GREEN}✓ Successfully built $lesson${NC}"
        echo -e "${BLUE}Binary: $BIN_DIR/$lesson${NC}"
    else
        echo -e "${RED}✗ Failed to build $lesson${NC}"
        exit 1
    fi
}

# Function to run a lesson
run_lesson() {
    local lesson=$1
    if [ -z "$lesson" ]; then
        echo -e "${RED}Error: Please specify a lesson name${NC}"
        echo -e "${YELLOW}Usage: $0 run <lesson-name>${NC}"
        exit 1
    fi
    
    lesson_path="$LESSONS_DIR/$lesson"
    if [ ! -d "$lesson_path" ]; then
        echo -e "${RED}Error: Lesson '$lesson' not found${NC}"
        exit 1
    fi
    
    echo -e "${BLUE}Running lesson $lesson...${NC}"
    echo -e "${YELLOW}--- Output ---${NC}"
    cd "$lesson_path" && go run *.go
    echo -e "${YELLOW}--- End ---${NC}"
}

# Function to list lessons
list_lessons() {
    echo -e "${BLUE}Available lessons:${NC}"
    if [ -d "$LESSONS_DIR" ]; then
        for lesson_dir in "$LESSONS_DIR"/*; do
            if [ -d "$lesson_dir" ]; then
                lesson_name=$(basename "$lesson_dir")
                echo -e "${GREEN}  $lesson_name${NC}"
            fi
        done
    else
        echo -e "${RED}No lessons directory found${NC}"
    fi
}

# Main script logic
case "${1:-all}" in
    "all"|"")
        build_all
        ;;
    "lesson")
        build_lesson "$2"
        ;;
    "run")
        run_lesson "$2"
        ;;
    "list")
        list_lessons
        ;;
    "clean")
        echo -e "${YELLOW}Cleaning binaries...${NC}"
        rm -rf "$BIN_DIR"
        echo -e "${GREEN}✓ Cleaned${NC}"
        ;;
    "help")
        echo -e "${BLUE}Usage: $0 [command] [options]${NC}"
        echo -e "${GREEN}Commands:${NC}"
        echo -e "  all (default)     Build all lessons"
        echo -e "  lesson <name>     Build specific lesson"
        echo -e "  run <name>        Run specific lesson"
        echo -e "  list              List available lessons"
        echo -e "  clean             Remove all binaries"
        echo -e "  help              Show this help"
        echo
        echo -e "${YELLOW}Examples:${NC}"
        echo -e "  $0                    # Build all lessons"
        echo -e "  $0 lesson 01-hello-world  # Build specific lesson"
        echo -e "  $0 run 02-variables       # Run specific lesson"
        echo -e "  $0 list                   # List all lessons"
        ;;
    *)
        echo -e "${RED}Error: Unknown command '$1'${NC}"
        echo -e "${YELLOW}Run '$0 help' for usage information${NC}"
        exit 1
        ;;
esac
