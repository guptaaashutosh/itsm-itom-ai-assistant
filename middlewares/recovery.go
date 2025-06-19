// Package middlewares provides HTTP middleware for the application, including recovery.
package middlewares

import (
	"github.com/gin-gonic/gin"
)

// RegisterRecovery attaches Gin's recovery middleware to the router.
func RegisterRecovery(r *gin.Engine) {
	r.Use(gin.Recovery())
}
