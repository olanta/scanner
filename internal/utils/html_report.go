package utils

import (
	"fmt"
	"strings"

	"github.com/olanta/olanta/scanner/internal/models"
)

func CreateHTMLReport(issues []models.Issue) string {
	var builder strings.Builder

	builder.WriteString("<html><head><title>Scan Report</title></head><body>")
	builder.WriteString("<h1>Scan Report</h1>")
	builder.WriteString("<table border='1'><tr><th>Description</th><th>Severity</th><th>File</th><th>Line</th><th>Column</th></tr>")

	for _, issue := range issues {
		builder.WriteString("<tr>")
		builder.WriteString(fmt.Sprintf("<td>%s</td>", issue.Description))
		builder.WriteString(fmt.Sprintf("<td>%s</td>", issue.Severity))
		builder.WriteString(fmt.Sprintf("<td>%s</td>", issue.File))
		builder.WriteString(fmt.Sprintf("<td>%d</td>", issue.Line))
		builder.WriteString(fmt.Sprintf("<td>%d</td>", issue.Column))
		builder.WriteString("</tr>")
	}

	builder.WriteString("</table>")
	builder.WriteString("</body></html>")

	return builder.String()
}
