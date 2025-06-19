// Package controllers implements HTTP handler functions for incidents.
package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"itsm-itom-ai-assistant/constants"
	"itsm-itom-ai-assistant/internal"
	"itsm-itom-ai-assistant/repository"
	"itsm-itom-ai-assistant/restapiv1/models/dto"
	"itsm-itom-ai-assistant/restapiv1/models/request"
	"itsm-itom-ai-assistant/restapiv1/models/response"
)

// CreateIncident handles POST /incidents requests.
func CreateIncident(c *gin.Context) {

	var req request.IncidentRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Apply AI logic
	aiSuggestion := internal.DeriveAISeverityAndCategory(req.Title, req.Description)

	incident := dto.IncidentDTO{
		Title:           req.Title,
		Description:     req.Description,
		AffectedService: req.AffectedService,
		AISeverity:      aiSuggestion.Severity,
		AICategory:      aiSuggestion.Category,
		CreatedAt:       time.Now(),
	}

	db, ok := c.Request.Context().Value(constants.DB_CONN_CTX_KEY).(*sql.DB)
	if !ok || db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not available"})
		return
	}
	incidentRepo := repository.NewIncidentRepository(db)

	err := incidentRepo.CreateIncident(incident)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.MsgInternalError})
		return
	}

	// Respond
	c.JSON(http.StatusOK, response.IncidentResponse{
		Title:           incident.Title,
		Description:     incident.Description,
		AffectedService: incident.AffectedService,
		AISeverity:      incident.AISeverity,
		AICategory:      incident.AICategory,
		CreatedAt:       incident.CreatedAt,
	})
}
