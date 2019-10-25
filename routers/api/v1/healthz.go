package v1

import (
	"kusnandartoni/starter/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck func for health check
func HealthCheck(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	appG.Response(http.StatusOK, "Ok", map[string]interface{}{
		"Status": "Health",
	})
}
