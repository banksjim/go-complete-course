# Go Complete Course Makefile
# Supports multiple lessons and programs

# Variables
BIN_DIR := bin
LESSONS_DIR := lessons
GO_FILES := $(shell find $(LESSONS_DIR) -name "*.go" -type f)
LESSON_DIRS := $(shell find $(LESSONS_DIR) -mindepth 1 -maxdepth 1 -type d)

# Colors for output
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[1;33m
BLUE := \033[0;34m
NC := \033[0m # No Color

# Default target
.PHONY: all
all: help

# Help target
.PHONY: help
help:
	@echo -e "$(BLUE)Go Complete Course - Available Commands:$(NC)"
	@echo -e "$(GREEN)  make list$(NC)          - List all available lessons"
	@echo -e "$(GREEN)  make build$(NC)         - Build all lessons"
	@echo -e "$(GREEN)  make build-lesson$(NC)  - Build specific lesson (e.g., make build-lesson LESSON=01-hello-world)"
	@echo -e "$(GREEN)  make run$(NC)           - Run specific lesson (e.g., make run LESSON=01-hello-world)"
	@echo -e "$(GREEN)  make run-all$(NC)       - Run all lessons sequentially"
	@echo -e "$(GREEN)  make clean$(NC)         - Remove all binaries"
	@echo -e "$(GREEN)  make test$(NC)          - Run tests (if any)"
	@echo -e "$(GREEN)  make new$(NC)           - Create new lesson (e.g., make new LESSON=04-arrays)"

# List all lessons
.PHONY: list
list:
	@echo -e "$(BLUE)Available Lessons:$(NC)"
	@for dir in $(LESSON_DIRS); do \
		lesson=$$(basename $$dir); \
		echo -e "$(GREEN)  $$lesson$(NC)"; \
	done

# Build all lessons
.PHONY: build
build:
	@echo -e "$(YELLOW)Building all lessons...$(NC)"
	@mkdir -p $(BIN_DIR)
	@for dir in $(LESSON_DIRS); do \
		lesson=$$(basename $$dir); \
		echo -e "$(YELLOW)Building $$lesson...$(NC)"; \
		if go build -o $(BIN_DIR)/$$lesson $$dir/*.go; then \
			echo -e "$(GREEN)✓ Successfully built $$lesson$(NC)"; \
		else \
			echo -e "$(RED)✗ Failed to build $$lesson$(NC)"; \
		fi; \
	done

# Build specific lesson
.PHONY: build-lesson
build-lesson:
ifndef LESSON
	@echo -e "$(RED)Error: LESSON not specified. Usage: make build-lesson LESSON=01-hello-world$(NC)"
	@exit 1
endif
	@echo -e "$(YELLOW)Building lesson $(LESSON)...$(NC)"
	@mkdir -p $(BIN_DIR)
	@if [ -d "$(LESSONS_DIR)/$(LESSON)" ]; then \
		if go build -o $(BIN_DIR)/$(LESSON) $(LESSONS_DIR)/$(LESSON)/*.go; then \
			echo -e "$(GREEN)✓ Successfully built $(LESSON)$(NC)"; \
		else \
			echo -e "$(RED)✗ Failed to build $(LESSON)$(NC)"; \
		fi; \
	else \
		echo -e "$(RED)Error: Lesson $(LESSON) not found$(NC)"; \
		exit 1; \
	fi

# Run specific lesson
.PHONY: run
run:
ifndef LESSON
	@echo -e "$(RED)Error: LESSON not specified. Usage: make run LESSON=01-hello-world$(NC)"
	@exit 1
endif
	@echo -e "$(BLUE)Running lesson $(LESSON)...$(NC)"
	@if [ -d "$(LESSONS_DIR)/$(LESSON)" ]; then \
		echo -e "$(YELLOW)--- Output ---$(NC)"; \
		cd $(LESSONS_DIR)/$(LESSON) && go run *.go; \
		echo -e "$(YELLOW)--- End ---$(NC)"; \
	else \
		echo -e "$(RED)Error: Lesson $(LESSON) not found$(NC)"; \
		exit 1; \
	fi

# Run all lessons
.PHONY: run-all
run-all:
	@echo -e "$(BLUE)Running all lessons...$(NC)"
	@for dir in $(LESSON_DIRS); do \
		lesson=$$(basename $$dir); \
		echo -e "$(BLUE)--- Running $$lesson ---$(NC)"; \
		cd $$dir && go run *.go; \
		echo -e "$(YELLOW)--- End of $$lesson ---$(NC)"; \
		echo; \
	done

# Clean binaries
.PHONY: clean
clean:
	@echo -e "$(YELLOW)Cleaning binaries...$(NC)"
	@rm -rf $(BIN_DIR)
	@echo -e "$(GREEN)✓ Cleaned$(NC)"

# Create new lesson
.PHONY: new
new:
ifndef LESSON
	@echo -e "$(RED)Error: LESSON not specified. Usage: make new LESSON=04-arrays$(NC)"
	@exit 1
endif
	@echo -e "$(YELLOW)Creating new lesson $(LESSON)...$(NC)"
	@if [ -d "$(LESSONS_DIR)/$(LESSON)" ]; then \
		echo -e "$(RED)Error: Lesson $(LESSON) already exists$(NC)"; \
		exit 1; \
	else \
		mkdir -p $(LESSONS_DIR)/$(LESSON); \
		echo 'package main\n\nimport "fmt"\n\nfunc main() {\n\tfmt.Println("Lesson: $(LESSON)")\n\t// Add your code here\n}' > $(LESSONS_DIR)/$(LESSON)/$(LESSON).go; \
		echo -e "$(GREEN)✓ Created lesson $(LESSON)$(NC)"; \
		echo -e "$(BLUE)Edit: $(LESSONS_DIR)/$(LESSON)/$(LESSON).go$(NC)"; \
	fi

# Run tests (if any)
.PHONY: test
test:
	@echo -e "$(YELLOW)Running tests...$(NC)"
	@go test ./...

# Development - run with file watching (requires entr)
.PHONY: watch
watch:
ifndef LESSON
	@echo -e "$(RED)Error: LESSON not specified. Usage: make watch LESSON=01-hello-world$(NC)"
	@exit 1
endif
	@echo -e "$(BLUE)Watching lesson $(LESSON) for changes...$(NC)"
	@if command -v entr >/dev/null 2>&1; then \
		find $(LESSONS_DIR)/$(LESSON) -name "*.go" | entr -c make run LESSON=$(LESSON); \
	else \
		echo -e "$(RED)Error: entr not installed. Install with: sudo apt install entr$(NC)"; \
	fi
