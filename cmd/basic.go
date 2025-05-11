package cmd

import (
	"fmt"

	"gotypeslearn/internal/examples"
	"gotypeslearn/internal/printer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	basicTypes = []string{
		"int",
		"float",
		"bool",
		"string",
		"byte",
		"rune",
		"complex",
	}

	basicCmd = &cobra.Command{
		Use:   "basic [type]",
		Short: "Learn about basic data types in Go",
		Long: `Explore basic data types in the Go programming language.
If no specific type is provided, you'll see an overview of all basic types.
For detailed information about a specific type, add its name as an argument.

Available basic types:
- int (and variants: int8, int16, int32, int64, uint, uint8, etc.)
- float (float32, float64)
- bool
- string
- byte (alias for uint8)
- rune (alias for int32, represents Unicode code points)
- complex (complex64, complex128)`,
		Args: cobra.MaximumNArgs(1),
		Run:  runBasicCmd,
	}
)

func runBasicCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		showBasicTypesOverview()
		return
	}

	typeName := args[0]
	switch typeName {
	case "int":
		showIntType()
	case "float":
		showFloatType()
	case "bool":
		showBoolType()
	case "string":
		showStringType()
	case "byte":
		showByteType()
	case "rune":
		showRuneType()
	case "complex":
		showComplexType()
	default:
		color.Red("Unknown type: %s", typeName)
		fmt.Println("Available basic types:", basicTypes)
	}
}

func showBasicTypesOverview() {
	printer.PrintHeader("Basic Data Types in Go")
	
	printer.PrintSection("Available Basic Types")
	
	for _, t := range basicTypes {
		printer.PrintItem(t)
	}
	
	printer.PrintNote("Run 'gotypeslearn basic TYPE' to learn more about a specific type")
	
	printer.PrintSection("Quick Overview")
	printer.PrintCode(`
// Integer types
var i int = 42         // platform dependent size (32 or 64 bit)
var i8 int8 = 127      // -128 to 127
var i16 int16 = 32767  // -32768 to 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807
var ui uint = 42       // unsigned integer
var ui8 uint8 = 255    // 0 to 255
var ui16 uint16 = 65535

// Float types
var f32 float32 = 3.14159265358979323846
var f64 float64 = 3.14159265358979323846 // more precision

// Boolean type
var b bool = true

// String type
var s string = "Hello, Go!"

// Byte (uint8) and Rune (int32)
var byte byte = 'A'    // ASCII character
var r rune = '世'       // Unicode character

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 1.5 + 3.1i
`)
}

func showIntType() {
	printer.PrintHeader("Integer Types in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Go has several integer types with different sizes and ranges.")
	fmt.Println("The primary types are int and uint (unsigned int), with size dependent on platform (32 or 64 bit).")
	fmt.Println("There are also size-specific variants: int8, int16, int32, int64, uint8, uint16, uint32, uint64.")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.IntExamples)
	
	printer.PrintSection("Common Operations")
	printer.PrintCode(examples.IntOperations)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Counting and indexing")
	fmt.Println("- Storing whole numbers like age, counts, IDs")
	fmt.Println("- Mathematical calculations without decimal points")
	fmt.Println("- Bit manipulation (using uint types)")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use plain 'int' for most purposes unless you have specific memory or range requirements")
	fmt.Println("- Use uint only when you specifically need non-negative numbers or bit operations")
	fmt.Println("- Be cautious about overflow - operations that exceed the type's range wrap around")
}

func showFloatType() {
	printer.PrintHeader("Floating-Point Types in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Go provides two floating-point types for representing decimal numbers:")
	fmt.Println("- float32: Single precision, ~6 decimal digits of precision")
	fmt.Println("- float64: Double precision, ~15 decimal digits of precision (default and recommended)")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.FloatExamples)
	
	printer.PrintSection("Common Operations")
	printer.PrintCode(examples.FloatOperations)
	
	printer.PrintSection("Precision and Comparison")
	printer.PrintCode(examples.FloatPrecision)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Scientific calculations")
	fmt.Println("- Financial applications (though Decimal libraries are often better)")
	fmt.Println("- Graphics and physics simulations")
	fmt.Println("- Anything requiring decimal points or fractional values")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use float64 by default, as it provides better precision")
	fmt.Println("- Never compare floats directly with ==; use a threshold/epsilon approach")
	fmt.Println("- For financial calculations, consider using libraries providing decimal types")
	fmt.Println("- Be aware of potential precision errors in calculations")
}

func showBoolType() {
	printer.PrintHeader("Boolean Type in Go")
	
	printer.PrintSection("Description")
	fmt.Println("The bool type represents boolean values. It can be either true or false.")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.BoolExamples)
	
	printer.PrintSection("Common Operations")
	printer.PrintCode(examples.BoolOperations)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Conditional statements (if, for, switch)")
	fmt.Println("- State flags (is enabled, is valid, etc.)")
	fmt.Println("- Control flow decisions")
	fmt.Println("- Return values indicating success/failure")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use descriptive variable names for booleans, often starting with 'is', 'has', 'should', etc.")
	fmt.Println("- Prefer positive boolean names to avoid double negatives in conditions")
	fmt.Println("- The zero value is false")
}

func showStringType() {
	printer.PrintHeader("String Type in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Strings in Go are immutable sequences of bytes, typically representing text.")
	fmt.Println("They're UTF-8 encoded by default, making them excellent for multilingual text.")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.StringExamples)
	
	printer.PrintSection("String Operations")
	printer.PrintCode(examples.StringOperations)
	
	printer.PrintSection("Runes and Bytes")
	printer.PrintCode(examples.StringRunes)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Text processing and manipulation")
	fmt.Println("- User input and output")
	fmt.Println("- File paths, URLs, and identifiers")
	fmt.Println("- Configuration values and settings")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Strings are immutable - operations create new strings")
	fmt.Println("- Use string builder for efficient string concatenation")
	fmt.Println("- Process strings as runes when working with international text")
	fmt.Println("- Raw string literals (using backticks) help with multi-line strings or regex patterns")
}

func showByteType() {
	printer.PrintHeader("Byte Type in Go")
	
	printer.PrintSection("Description")
	fmt.Println("The byte type is an alias for uint8, representing a single 8-bit unsigned integer.")
	fmt.Println("It's commonly used to represent ASCII characters or binary data.")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.ByteExamples)
	
	printer.PrintSection("Working with Bytes")
	printer.PrintCode(examples.ByteOperations)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Binary data processing")
	fmt.Println("- File I/O operations")
	fmt.Println("- Network programming")
	fmt.Println("- Cryptography and hashing")
	fmt.Println("- ASCII character processing")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use []byte for efficient string manipulation in performance-critical code")
	fmt.Println("- Convert to string when interfacing with text-based APIs")
	fmt.Println("- Remember that byte can only represent values 0-255")
	fmt.Println("- Use byte for ASCII, but use rune for non-ASCII characters")
}

func showRuneType() {
	printer.PrintHeader("Rune Type in Go")
	
	printer.PrintSection("Description")
	fmt.Println("The rune type is an alias for int32, representing a Unicode code point.")
	fmt.Println("It's used to represent individual characters, including international characters.")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.RuneExamples)
	
	printer.PrintSection("Working with Runes")
	printer.PrintCode(examples.RuneOperations)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- International text processing")
	fmt.Println("- Character-by-character string manipulation")
	fmt.Println("- Unicode handling")
	fmt.Println("- Text validation and filtering")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use runes when processing text that may contain non-ASCII characters")
	fmt.Println("- Remember that a rune represents a Unicode code point, not a byte")
	fmt.Println("- When iterating over strings, range provides runes, not bytes")
	fmt.Println("- A single rune may be encoded as multiple bytes in a string")
}

func showComplexType() {
	printer.PrintHeader("Complex Number Types in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Go provides two complex number types:")
	fmt.Println("- complex64: Complex numbers with float32 real and imaginary parts")
	fmt.Println("- complex128: Complex numbers with float64 real and imaginary parts")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.ComplexExamples)
	
	printer.PrintSection("Operations")
	printer.PrintCode(examples.ComplexOperations)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Signal processing")
	fmt.Println("- Electrical engineering calculations")
	fmt.Println("- Physics simulations")
	fmt.Println("- Mathematical computations")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use complex128 for most applications, as it offers better precision")
	fmt.Println("- Extract real and imaginary parts using real() and imag() functions")
	fmt.Println("- The 'math/cmplx' package provides additional functions for complex number manipulations")
}
