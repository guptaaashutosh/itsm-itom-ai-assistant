// Package response defines response payloads for the API.
package response

import "time"

// IncidentResponse represents the outgoing payload for a created incident.
type IncidentResponse struct {
	ID              string    `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	Description     string    `json:"description,omitempty"`
	AffectedService string    `json:"affected_service"`
	AISeverity      string    `json:"ai_severity"`
	AICategory      string    `json:"ai_category"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}
