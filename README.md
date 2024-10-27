# Go Interpreter

A custom interpreter built in Go. 

## Improvements

### Error Handling
- **Current Issue**: Error messages lack specific context like filenames and line numbers, making debugging more challenging.
- **Proposed Improvement**: Incorporate `io.Reader` to parse files more flexibly and include filename information in error outputs, enhancing error reporting for easier debugging.

### Unicode Support
- **Current Issue**: Limited to ASCII characters due to a byte-based input handling method in `lexer.go`.
- **Proposed Improvement**: Switch to `rune`-based input handling to fully support Unicode, enabling the interpreter to process non-ASCII characters smoothly.

### Float Support
- **Current Issue**: The interpreter does not recognize floating-point numbers, limiting its arithmetic capabilities.
- **Proposed Improvement**: Update the lexer to detect and parse floating-point literals, enabling operations on both integer and decimal values.

## Getting Started

### Prerequisites
- **Go**: Install Go 1.18 or higher. [Get Go here](https://golang.org/dl/).

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-interpreter.git
   cd go-interpreter

2. Build the project
   ```bash
   go build -o interpreter main.go

3. Run in interactive mode, uses a REPL
   ```
   ./interpreter
