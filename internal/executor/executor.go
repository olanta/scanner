package executor

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/olanta/olanta/scanner/internal/models"
	"github.com/olanta/olanta/scanner/internal/scanner"
	"github.com/olanta/olanta/scanner/internal/utils"
)

func ExecuteScan(language, path string) []models.Issue {
	var issues []models.Issue

	switch language {
	case "java":
		s := scanner.NewJavaScanner()
		issues = s.Scan(path)
	case "python":
		s := scanner.NewPythonScanner()
		issues = s.Scan(path)
	default:
		fmt.Printf("Unsupported language: %s\n", language)
	}

	return issues
}

func GenerateHTMLReport(issues []models.Issue, path string) {
	report := utils.CreateHTMLReport(issues)
	reportPath := filepath.Join(path, "scan_report_"+time.Now().Format("20060102150405")+".html")

	file, err := os.Create(reportPath)
	if err != nil {
		fmt.Printf("Error creating report file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(report)
	if err != nil {
		fmt.Printf("Error writing to report file: %v\n", err)
		return
	}

	fmt.Printf("Report generated at: %s\n", reportPath)
}
