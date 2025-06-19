// Package dal contains data access logic for incidents using pgx.
package dal

import (
	"database/sql"
	"itsm-itom-ai-assistant/restapiv1/models/dto"
	"time"
)

// InsertIncident inserts a new incident into the database.
func InsertIncident(db *sql.DB, incident dto.IncidentDTO) error {
	_, err := db.Exec(`INSERT INTO incidents (title, description, affected_service, ai_severity, ai_category, created_at) VALUES ($1, $2, $3, $4, $5, $6)`, incident.Title, incident.Description, incident.AffectedService, incident.AISeverity, incident.AICategory, time.Now().UTC())
	return err
}
