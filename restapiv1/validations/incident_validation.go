// Package validations provides input validation logic for incidents.
package validations

import "strings"

// ValidateIncidentTitle checks if the title is non-empty.
func ValidateIncidentTitle(title string) bool {
	return strings.TrimSpace(title) != ""
}
