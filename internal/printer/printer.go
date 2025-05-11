package printer

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	headerStyle  = color.New(color.FgHiCyan, color.Bold)
	sectionStyle = color.New(color.FgHiYellow, color.Bold)
	itemStyle    = color.New(color.FgHiGreen)
	noteStyle    = color.New(color.FgHiMagenta, color.Italic)
	codeStyle    = color.New(color.FgHiWhite)
)

// PrintHeader prints a main topic header
func PrintHeader(text string) {
	fmt.Println()
	headerStyle.Println("==================================================")
	headerStyle.Printf("  %s\n", text)
	headerStyle.Println("==================================================")
	fmt.Println()
}

// PrintSection prints a section header
func PrintSection(text string) {
	fmt.Println()
	sectionStyle.Printf("▶ %s\n", text)
	fmt.Println(strings.Repeat("-", len(text)+2))
}

// PrintItem prints a bulleted item
func PrintItem(text string) {
	itemStyle.Printf(" • %s\n", text)
}

// PrintNote prints a highlighted note
func PrintNote(text string) {
	fmt.Println()
	noteStyle.Printf("Note: %s\n", text)
	fmt.Println()
}

// PrintCode prints a code block with syntax highlighting
func PrintCode(code string) {
	fmt.Println()
	
	// Trim leading and trailing newlines
	code = strings.Trim(code, "\n")
	
	// Print the code block
	fmt.Println("```go")
	codeStyle.Println(code)
	fmt.Println("```")
	fmt.Println()
}