package cmd

import (
	"fmt"

	"gotypeslearn/internal/examples"
	"gotypeslearn/internal/printer"

	"github.com/spf13/cobra"
)

var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "Learn about interfaces in Go",
	Long: `Interfaces in Go define behavior, not data. 
They are a powerful way to achieve polymorphism and decouple code.
This command explains interfaces, their usage, and best practices.`,
	Run: func(cmd *cobra.Command, args []string) {
		showInterfaces()
	},
}

func showInterfaces() {
	printer.PrintHeader("Interfaces in Go")

	printer.PrintSection("Description")
	fmt.Println("Interfaces define behavior through method signatures. A type implements an interface")
	fmt.Println("by implementing its methods, without explicit declaration. This is called")
	fmt.Println("\"implicit implementation\" and is a key feature of Go's design.")

	printer.PrintSection("Basic Interface Definition")
	printer.PrintCode(examples.InterfaceDefinitionExample)

	printer.PrintSection("Implementing Interfaces")
	printer.PrintCode(examples.InterfaceImplementation)

	printer.PrintSection("Interface Values")
	printer.PrintCode(examples.InterfaceValues)

	printer.PrintSection("Empty Interface")
	printer.PrintCode(examples.EmptyInterface)

	printer.PrintSection("Type Assertions and Type Switches")
	printer.PrintCode(examples.TypeAssertions)

	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Dependency injection")
	fmt.Println("- Testing with mocks")
	fmt.Println("- Polymorphism and abstraction")
	fmt.Println("- Creating flexible APIs")
	fmt.Println("- Composition of behavior")

	printer.PrintSection("Common Standard Library Interfaces")
	printer.PrintCode(`
// io.Reader - implemented by types that can read bytes
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer - implemented by types that can write bytes
type Writer interface {
    Write(p []byte) (n int, err error)
}

// fmt.Stringer - implemented by types with a string representation
type Stringer interface {
    String() string
}

// sort.Interface - implemented by types that can be sorted
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// error - the built-in error interface
type error interface {
    Error() string
}
`)

	printer.PrintSection("Best Practices")
	fmt.Println("- Keep interfaces small (often just one or two methods)")
	fmt.Println("- Define interfaces where they're used, not implemented")
	fmt.Println("- Accept interfaces, return concrete types")
	fmt.Println("- Use the empty interface (interface{}) sparingly")
	fmt.Println("- Interfaces are about behavior, not data")
}
