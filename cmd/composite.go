package cmd

import (
	"fmt"

	"gotypeslearn/internal/examples"
	"gotypeslearn/internal/printer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	compositeTypes = []string{
		"array",
		"slice",
		"map",
		"struct",
		"pointer",
		"function",
		"channel",
	}

	compositeCmd = &cobra.Command{
		Use:   "composite [type]",
		Short: "Learn about composite data types in Go",
		Long: `Explore composite data types in the Go programming language.
If no specific type is provided, you'll see an overview of all composite types.
For detailed information about a specific type, add its name as an argument.

Available composite types:
- array (fixed-size sequence of elements)
- slice (dynamic-size sequence of elements)
- map (key-value pairs)
- struct (custom data type with named fields)
- pointer (reference to memory address)
- function (first-class functions)
- channel (communication between goroutines)`,
		Args: cobra.MaximumNArgs(1),
		Run:  runCompositeCmd,
	}
)

func runCompositeCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		showCompositeTypesOverview()
		return
	}

	typeName := args[0]
	switch typeName {
	case "array":
		showArrayType()
	case "slice":
		showSliceType()
	case "map":
		showMapType()
	case "struct":
		showStructType()
	case "pointer":
		showPointerType()
	case "function":
		showFunctionType()
	case "channel":
		showChannelType()
	default:
		color.Red("Unknown type: %s", typeName)
		fmt.Println("Available composite types:", compositeTypes)
	}
}

func showCompositeTypesOverview() {
	printer.PrintHeader("Composite Data Types in Go")
	
	printer.PrintSection("Available Composite Types")
	
	for _, t := range compositeTypes {
		printer.PrintItem(t)
	}
	
	printer.PrintNote("Run 'gotypeslearn composite TYPE' to learn more about a specific type")
	
	printer.PrintSection("Quick Overview")
	printer.PrintCode(`
// Arrays - fixed-size collection of elements
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slices - dynamic-size collection (view into array)
var slice []int = []int{1, 2, 3, 4}
slice = append(slice, 5) // dynamic growth

// Maps - key-value pairs
var m map[string]int = map[string]int{
    "one": 1,
    "two": 2,
}

// Structs - custom data types with fields
type Person struct {
    Name string
    Age  int
}
p := Person{Name: "Alice", Age: 30}

// Pointers - references to memory addresses
var ptr *int
num := 42
ptr = &num    // ptr points to num's memory

// Functions - first-class citizens
add := func(a, b int) int {
    return a + b
}
result := add(5, 3)

// Channels - communication between goroutines
ch := make(chan int)
go func() { ch <- 42 }() // send
value := <-ch            // receive
`)
}

func showArrayType() {
	printer.PrintHeader("Arrays in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Arrays in Go are fixed-size sequences of elements of the same type.")
	fmt.Println("The size is part of the array's type, which means [5]int and [10]int are different types.")
	
	printer.PrintSection("Declaration and Initialization")
	printer.PrintCode(examples.ArrayExamples)
	
	printer.PrintSection("Operations")
	printer.PrintCode(examples.ArrayOperations)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- When you need a fixed collection of items")
	fmt.Println("- Performance-critical code where size is known in advance")
	fmt.Println("- Backing storage for slices")
	fmt.Println("- Memory-efficient representation when size is known")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Arrays in Go are values, not references; they are copied when assigned or passed to functions")
	fmt.Println("- Use arrays when you know the exact size needed and it won't change")
	fmt.Println("- For most cases, slices are more flexible and commonly used instead of arrays")
	fmt.Println("- Arrays are rarely used directly in Go - slices are the typical choice")
}

func showSliceType() {
	printer.PrintHeader("Slices in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Slices are dynamic, flexible views into arrays. They have three components:")
	fmt.Println("- A pointer to an underlying array")
	fmt.Println("- A length (the number of elements it contains)")
	fmt.Println("- A capacity (the maximum number of elements it can hold before reallocation)")
	
	printer.PrintSection("Declaration and Initialization")
	printer.PrintCode(examples.SliceExamples)
	
	printer.PrintSection("Operations")
	printer.PrintCode(examples.SliceOperations)
	
	printer.PrintSection("Slicing and Appending")
	printer.PrintCode(examples.SliceManipulation)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Dynamic collections of elements")
	fmt.Println("- Building lists incrementally")
	fmt.Println("- Working with unknown amounts of data")
	fmt.Println("- Processing collections of similar items")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Slices are references to arrays; changes to a slice modify the underlying array")
	fmt.Println("- Use make() to pre-allocate capacity for better performance when the size is known")
	fmt.Println("- Be careful with large slices in functions - they can keep large arrays in memory")
	fmt.Println("- Remember that multiple slices can share the same underlying array")
}

func showMapType() {
	printer.PrintHeader("Maps in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Maps are unordered collections of key-value pairs. They provide fast lookups,")
	fmt.Println("insertions, and deletions based on keys. Maps are reference types in Go.")
	
	printer.PrintSection("Declaration and Initialization")
	printer.PrintCode(examples.MapExamples)
	
	printer.PrintSection("Operations")
	printer.PrintCode(examples.MapOperations)
	
	printer.PrintSection("Checking for Keys")
	printer.PrintCode(examples.MapKeyCheck)
	
	printer.PrintSection("Iterating Over Maps")
	printer.PrintCode(examples.MapIteration)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Lookups and dictionaries")
	fmt.Println("- Caching values")
	fmt.Println("- Counting occurrences (frequency maps)")
	fmt.Println("- Deduplication")
	fmt.Println("- Representing relationships between data")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Maps must be initialized before use (with make() or literal)")
	fmt.Println("- Map iteration order is not guaranteed - it's intentionally randomized")
	fmt.Println("- Maps are not safe for concurrent use; use sync.Map or mutexes for concurrency")
	fmt.Println("- Valid map keys must be comparable (can use == operator)")
	fmt.Println("- Slices, maps, and functions cannot be used as map keys")
}

func showStructType() {
	printer.PrintHeader("Structs in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Structs are composite data types that group together variables under a single name.")
	fmt.Println("They allow you to create custom data types with named fields of different types.")
	
	printer.PrintSection("Declaration and Initialization")
	printer.PrintCode(examples.StructExamples)
	
	printer.PrintSection("Accessing and Modifying Fields")
	printer.PrintCode(examples.StructOperations)
	
	printer.PrintSection("Methods and Receivers")
	printer.PrintCode(examples.StructMethods)
	
	printer.PrintSection("Embedding and Composition")
	printer.PrintCode(examples.StructEmbedding)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Representing real-world entities (User, Product, etc.)")
	fmt.Println("- Custom data structures")
	fmt.Println("- Request/response models")
	fmt.Println("- Configuration objects")
	fmt.Println("- Implementing object-oriented patterns")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use capitalized field names for exported fields (visible outside package)")
	fmt.Println("- Use tags for metadata (often used with encoding packages or ORM)")
	fmt.Println("- Choose value or pointer receivers based on your needs")
	fmt.Println("- Use composition over inheritance through embedding")
	fmt.Println("- Consider using constructor functions for complex structs")
}

func showPointerType() {
	printer.PrintHeader("Pointers in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Pointers in Go hold the memory address of a value. They allow you to indirectly reference")
	fmt.Println("and modify values, enabling efficient memory usage and pass-by-reference semantics.")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.PointerExamples)
	
	printer.PrintSection("Pointers to Structs")
	printer.PrintCode(examples.StructPointers)
	
	printer.PrintSection("Common Patterns")
	printer.PrintCode(examples.PointerPatterns)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Modifying arguments in functions")
	fmt.Println("- Avoiding copies of large data structures")
	fmt.Println("- Implementing data structures like linked lists, trees")
	fmt.Println("- Working with receiver methods that modify state")
	fmt.Println("- Representing optional values (nil pointers)")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use pointers when you need to modify a variable in a function")
	fmt.Println("- Use pointers for large structs to avoid copying")
	fmt.Println("- Be cautious about nil pointers - always check before dereferencing")
	fmt.Println("- Go has no pointer arithmetic (unlike C/C++)")
	fmt.Println("- The zero value of a pointer is nil")
}

func showFunctionType() {
	printer.PrintHeader("Functions in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Functions in Go are first-class citizens, which means they can be assigned to variables,")
	fmt.Println("passed as arguments, and returned from other functions.")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.FunctionExamples)
	
	printer.PrintSection("Function Types and Values")
	printer.PrintCode(examples.FunctionValues)
	
	printer.PrintSection("Higher-Order Functions")
	printer.PrintCode(examples.HigherOrderFunctions)
	
	printer.PrintSection("Closures")
	printer.PrintCode(examples.Closures)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Callbacks and event handlers")
	fmt.Println("- Strategy patterns")
	fmt.Println("- Dependency injection")
	fmt.Println("- Middleware chains")
	fmt.Println("- Functional programming patterns")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Use named return values for improved readability")
	fmt.Println("- Keep functions focused on a single task")
	fmt.Println("- Handle errors explicitly")
	fmt.Println("- Use variadic functions for flexible argument lists")
	fmt.Println("- Use defer for cleanup operations")
}

func showChannelType() {
	printer.PrintHeader("Channels in Go")
	
	printer.PrintSection("Description")
	fmt.Println("Channels are communication mechanisms that allow goroutines to synchronize and")
	fmt.Println("exchange data. They implement the principle: \"Don't communicate by sharing memory;")
	fmt.Println("share memory by communicating.\"")
	
	printer.PrintSection("Declaration and Usage")
	printer.PrintCode(examples.ChannelBasics)
	
	printer.PrintSection("Buffered Channels")
	printer.PrintCode(examples.BufferedChannels)
	
	printer.PrintSection("Directionality")
	printer.PrintCode(examples.ChannelDirections)
	
	printer.PrintSection("Select Statement")
	printer.PrintCode(examples.ChannelSelect)
	
	printer.PrintSection("Practical Use Cases")
	fmt.Println("- Synchronization between goroutines")
	fmt.Println("- Worker pools and job distribution")
	fmt.Println("- Rate limiting")
	fmt.Println("- Fan-out/fan-in patterns")
	fmt.Println("- Quit signals and timeouts")
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Close channels from the sender side, not the receiver")
	fmt.Println("- Use buffered channels when you know the exact number of items to send")
	fmt.Println("- Use directional channel types to clarify intent")
	fmt.Println("- Always handle potential blocking operations")
	fmt.Println("- Remember that sending to a closed channel causes a panic")
}
