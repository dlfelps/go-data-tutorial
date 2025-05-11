# GoTypesLearn

GoTypesLearn is an interactive CLI tool designed to help developers learn about Go's data types from the standard library. It provides comprehensive, colorful explanations and examples of basic and composite types, interfaces, and type conversions.

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/gotypeslearn)](https://goreportcard.com/report/github.com/yourusername/gotypeslearn)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

![GoTypesLearn Demo](./assets/demo.gif)

## Features

- 📚 Comprehensive coverage of Go's type system
- 🌈 Colorful command-line output with syntax highlighting
- 🧩 Organized by type categories for easy navigation
- 📝 Practical examples and best practices
- 💡 Tips for avoiding common pitfalls

## Installation

### Using Go Install

```bash
go install github.com/yourusername/gotypeslearn@latest
```

### From Source

```bash
git clone https://github.com/yourusername/gotypeslearn.git
cd gotypeslearn
go build -o gotypeslearn
```

## Usage

```bash
# Show general help and available commands
gotypeslearn

# Learn about basic data types (overview)
gotypeslearn basic

# Learn about a specific basic type
gotypeslearn basic int
gotypeslearn basic float
gotypeslearn basic string
gotypeslearn basic bool
gotypeslearn basic byte
gotypeslearn basic rune
gotypeslearn basic complex

# Learn about composite types (overview)
gotypeslearn composite

# Learn about specific composite types
gotypeslearn composite array
gotypeslearn composite slice
gotypeslearn composite map
gotypeslearn composite struct
gotypeslearn composite pointer
gotypeslearn composite function
gotypeslearn composite channel

# Learn about interfaces
gotypeslearn interface

# Learn about type conversion
gotypeslearn conversion
```

## Type Categories

### Basic Types
- `int` (and variants: int8, int16, int32, int64, uint, uint8, etc.)
- `float` (float32, float64)
- `bool`
- `string`
- `byte` (alias for uint8)
- `rune` (alias for int32, represents Unicode code points)
- `complex` (complex64, complex128)

### Composite Types
- `array` (fixed-size collection of elements)
- `slice` (dynamic-size sequence of elements)
- `map` (key-value pairs)
- `struct` (custom data type with named fields)
- `pointer` (reference to memory address)
- `function` (first-class functions)
- `channel` (communication between goroutines)

### Other Topics
- `interface` (type definitions based on behavior)
- `conversion` (converting between types)

## Project Structure

- `cmd/`: Command implementations
- `internal/examples/`: Code examples for each type
- `internal/printer/`: Formatting utilities for console output
- `main.go`: Application entry point

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.