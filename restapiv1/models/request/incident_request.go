package request

// IncidentRequest represents the expected payload for POST /incidents.
// swagger:model IncidentRequest
type IncidentRequest struct {
	Title           string `json:"title" binding:"required"`
	Description     string `json:"description" binding:"required"`
	AffectedService string `json:"affected_service" binding:"required"`
}
