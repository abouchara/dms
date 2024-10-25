// cmd/root.go
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "depview",
    Short: "A tool to manage dependencies within IT infrastructure and applications.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to Dependency Tool CLI!")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
