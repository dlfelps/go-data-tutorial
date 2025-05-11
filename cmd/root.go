package cmd

import (
        "fmt"

        "github.com/fatih/color"
        "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
        Use:   "gotypeslearn",
        Short: "Learn Golang data types interactively",
        Long: `A command-line tool that teaches you about different data types in Go.
Navigate through basic and composite data types, see examples, and learn
about type conversion and practical applications.`,
        Run: func(cmd *cobra.Command, args []string) {
                showWelcomeMessage()
                cmd.Help()
        },
}

// Execute runs the root command
func Execute() error {
        return rootCmd.Execute()
}

func init() {
        rootCmd.AddCommand(basicCmd)
        rootCmd.AddCommand(compositeCmd)
        rootCmd.AddCommand(interfaceCmd)
        rootCmd.AddCommand(conversionCmd)
}

func showWelcomeMessage() {
        title := color.New(color.FgHiCyan, color.Bold)
        title.Println("\n==== Welcome to Go Types Interactive Learning ====")
        
        info := color.New(color.FgHiWhite)
        info.Println("This tool will help you learn about Golang's data types.")
        info.Println("Use the commands below to explore different categories of types.")
        
        usage := color.New(color.FgHiYellow)
        usage.Println("\nUsage Examples:")
        examples := color.New(color.FgHiGreen)
        examples.Println("  gotypeslearn basic       - Learn about basic data types")
        examples.Println("  gotypeslearn composite   - Learn about composite data types")
        examples.Println("  gotypeslearn interface   - Learn about interfaces")
        examples.Println("  gotypeslearn conversion  - Learn about type conversion\n")
        
        tip := color.New(color.FgHiMagenta, color.Italic)
        tip.Println("Tip: Add --help to any command for more information")
        fmt.Println()
}
