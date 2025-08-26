# Go - The Complete Course

This repository contains course materials, my notes, and my source code for the Udemy course:

[Go - The Complete Course](https://www.udemy.com/course/go-the-complete-guide)

Tracking:

* Course started: 8/25/2025
* Course completed: In progress

## Project Structure

```
go-complete-course/
├── lessons/                 # All lessons organized by topic
│   ├── 01-first-app/       # Lesson 1: First Go Application
│   │   └── first-app.go
│   └── ...                 # Future lessons
├── bin/                    # Compiled binaries (ignored by git)
├── .vscode/               # VS Code configuration
│   └── launch.json        # Debug configurations
├── Makefile               # Build automation
├── build.sh               # Alternative build script
├── go.mod                 # Go module definition
└── README.md              # This file
```

## Quick Start

### Using Make (Recommended)

```bash
# List all available lessons
make list

# Run a specific lesson
make run LESSON=01-first-app

# Build all lessons
make build

# Build specific lesson
make build-lesson LESSON=01-first-app

# Create new lesson
make new LESSON=04-arrays

# Clean binaries
make clean

# Show help
make help
```

### Using Build Script

```bash
# Build all lessons
./build.sh

# Build specific lesson
./build.sh lesson 01-first-app

# Run specific lesson
./build.sh run 01-first-app

# List lessons
./build.sh list

# Show help
./build.sh help
```

### Using Go Commands Directly

```bash
# Run a lesson
cd lessons/01-first-app && go run first-app.go

# Build a lesson
go build -o bin/01-first-app lessons/01-first-app/first-app.go
```

### Using VS Code

1. Open any lesson file (e.g., `lessons/01-first-app/first-app.go`)
2. Press `F5` to debug
3. Choose "Launch Current Lesson" or specific lesson configuration

## Adding New Lessons

### Method 1: Using Make
```bash
make new LESSON=05-structs
```

### Method 2: Manual Creation
1. Create new directory: `lessons/XX-topic-name/`
2. Add `topic-name.go` file with your code
3. Update VS Code launch.json if needed (optional)

## Available Lessons

* **01-first-app**: Your first Go application

## Development Workflow

1. **Create lesson**: `make new LESSON=XX-topic`
2. **Edit code**: Open `lessons/XX-topic/XX-topic.go`
3. **Test quickly**: `make run LESSON=XX-topic`
4. **Debug**: Use VS Code F5 or `make build-lesson LESSON=XX-topic`
5. **Build all**: `make build` before committing

## Tips

* Use `make help` to see all available commands
* Binaries are automatically placed in `bin/` folder
* VS Code configurations are pre-configured for easy debugging
* Each lesson is self-contained with its own `.go` file
