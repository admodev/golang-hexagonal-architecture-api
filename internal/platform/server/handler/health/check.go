package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheckHandler returns an HTTP handler to perform API health checks
func HealthCheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Server is ok!")
	}
}
