// Package routes defines and registers API routes for the application.
package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"itsm-itom-ai-assistant/middlewares"
	"itsm-itom-ai-assistant/restapiv1/controllers"
)

// SetupRouter initializes the Gin engine, registers middleware and routes.
func InitializeRoutes(db *sql.DB) *gin.Engine {

	r := gin.Default()
	middlewares.RegisterRecovery(r)

	r.Use(middlewares.SetDBConnection(db))

	// swagger:route POST /incidents incidents createIncident
	// Create a new incident.
	// Responses:
	//   200: IncidentResponse
	r.POST("/incidents", controllers.CreateIncident)

	return r
}
