package dto

import "time"

// IncidentDTO is used to transfer incident data between layers.
type IncidentDTO struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	AffectedService string    `json:"affected_service"`
	AISeverity      string    `json:"ai_severity"`
	AICategory      string    `json:"ai_category"`
	CreatedAt       time.Time `json:"created_at"`
}
