# Todo-Finder

A command-line tool written in Go that scans your codebase for TODO and FIXME comments. This tool helps developers track their pending tasks and fixes across their project files.

## Features

- Recursively scans directories for Go files
- Identifies TODO and FIXME comments
- Shows file location and line numbers
- Simple and fast execution
- Supports both `-v` and `--version` flags

## Installation

### Prerequisites

- Go 1.11 or higher (for Go modules support)
- Git (optional, for version control)

### Building from Source

1. Clone the repository:

```sh
git clone https://github.com/thatbeautifuldream/todo-finder.git --depth 1
```

2. Navigate to the project directory:

```sh
cd todo-finder
```

3. Build the project:

```sh
go build
```

This will create an executable file named `todo-finder` in the current directory.

## Using `go install`

You can directly install the tool using `go install`:

```sh
go install github.com/thatbeautifuldream/todo-finder@latest
```

Then you can use the tool like this:

```sh
todo-finder -dir=./path/to/your/project
```

## Run the tool

```sh
go run .
```

- Or with a specific directory:

```sh
go run . -dir=./path/to/your/project
```
