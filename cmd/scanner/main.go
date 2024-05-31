package main

import (
	"fmt"
	"os"

	"github.com/olanta/olanta/scanner/cmd/scanner/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "scanner",
		Short: "Olanta scanner",
	}

	rootCmd.AddCommand(commands.NewScanCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
