package cmd

import (
	"fmt"

	"gotypeslearn/internal/examples"
	"gotypeslearn/internal/printer"

	"github.com/spf13/cobra"
)

var conversionCmd = &cobra.Command{
	Use:   "conversion",
	Short: "Learn about type conversion in Go",
	Long: `Learn about explicit and implicit type conversion in Go.
This command explains how to convert between different data types,
type assertions, and common conversion patterns.`,
	Run: func(cmd *cobra.Command, args []string) {
		showTypeConversion()
	},
}

func showTypeConversion() {
	printer.PrintHeader("Type Conversion in Go")
	
	printer.PrintSection("Basic Type Conversion")
	fmt.Println("Go requires explicit type conversion between most types, even related ones.")
	fmt.Println("This is to prevent subtle bugs from implicit conversions.")
	printer.PrintCode(examples.BasicConversion)
	
	printer.PrintSection("Numeric Type Conversion")
	printer.PrintCode(examples.NumericConversion)
	
	printer.PrintSection("String Conversions")
	printer.PrintCode(examples.StringConversion)
	
	printer.PrintSection("Interface Type Assertions")
	printer.PrintCode(examples.InterfaceConversion)
	
	printer.PrintSection("Type Conversions with Structs")
	printer.PrintCode(examples.StructConversion)
	
	printer.PrintSection("Using the strconv Package")
	printer.PrintCode(examples.StrconvExample)
	
	printer.PrintSection("Common Pitfalls")
	printer.PrintCode(examples.ConversionPitfalls)
	
	printer.PrintSection("Best Practices")
	fmt.Println("- Always check for potential overflow when converting between numeric types")
	fmt.Println("- Handle errors when converting strings to numbers")
	fmt.Println("- Use type assertions with the two-value form to safely check interfaces")
	fmt.Println("- Be aware of precision loss when converting between float types")
	fmt.Println("- Remember that conversions create new values, they don't modify the original")
}
