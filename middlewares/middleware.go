package middlewares

import (
	"context"
	"database/sql"
	"itsm-itom-ai-assistant/constants"

	"github.com/gin-gonic/gin"
)

func SetDBConnection(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constants.DB_CONN_CTX_KEY, dbConn)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
