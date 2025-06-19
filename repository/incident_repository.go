// Package repository provides the implementation of the IncidentRepository interface to interact with the DAL.
package repository

import (
	"database/sql"
	"itsm-itom-ai-assistant/dal"
	"itsm-itom-ai-assistant/restapiv1/models/dto"
)

// IncidentRepository defines methods to interact with incidents in the DB.
type IncidentRepository interface {
	CreateIncident(incident dto.IncidentDTO) error
}

// incidentRepositoryImpl implements IncidentRepository.
type incidentRepositoryImpl struct {
	db *sql.DB
}

// NewIncidentRepository returns a new IncidentRepository.
func NewIncidentRepository(db *sql.DB) IncidentRepository {
	return &incidentRepositoryImpl{db: db}
}

// CreateIncident stores a new incident using the DAL.
func (r *incidentRepositoryImpl) CreateIncident(incident dto.IncidentDTO) error {
	return dal.InsertIncident(r.db, incident)
}
