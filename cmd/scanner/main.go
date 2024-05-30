package main

import (
	"fmt"
	"os"

	"github.com/olanta/olanta/scanner/internal/executor"
	"github.com/olanta/olanta/scanner/internal/submitter"
	"github.com/spf13/cobra"
)

func main() {
	var coreURL string
	var dryRun bool

	rootCmd := &cobra.Command{
		Use:   "scanner",
		Short: "Olanta scanner",
	}

	scanCmd := &cobra.Command{
		Use:   "scan [language] [path]",
		Short: "Scan the specified path for code smells and bugs.",
		Long: `Scan the specified path for code smells and bugs.
Available languages: java, python.

Examples:
  # Scan a Java project
  scanner scan java /path/to/java/project

  # Scan a Python project
  scanner scan python /path/to/python/project

  # Run a dry-run scan and generate a local HTML report
  scanner scan java /path/to/java/project --dry-run
`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			language := args[0]
			path := args[1]

			issues := executor.ExecuteScan(language, path)
			if dryRun {
				executor.GenerateHTMLReport(issues, path)
				fmt.Println("Dry-run completed. Report generated.")
			} else {
				submitter.SubmitOrPrintIssues(coreURL, issues)
			}
		},
	}

	scanCmd.Flags().StringVar(&coreURL, "core-url", "localhost:8080", "URL of the Olanta core server")
	scanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Run the scan in dry-run mode and generate a local HTML report")
	rootCmd.AddCommand(scanCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
