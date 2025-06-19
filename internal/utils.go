// Package internal provides internal utilities, including AI logic for incident triage.
package internal

import "strings"

// AISuggestion holds the AI-determined severity and category.
type AISuggestion struct {
	Severity string
	Category string
}

// AnalyzeIncident runs basic AI logic to assign severity and category based on keywords.
func DeriveAISeverityAndCategory(title, description string) AISuggestion {
	text := strings.ToLower(title + " " + description)
	var severity, category string

	switch {
	case strings.Contains(text, "outage"):
		severity = "critical"
		category = "infrastructure"
	case strings.Contains(text, "fail") || strings.Contains(text, "error"):
		severity = "high"
		category = "application"
	default:
		severity = "normal"
		category = "general"
	}
	return AISuggestion{Severity: severity, Category: category}
}
