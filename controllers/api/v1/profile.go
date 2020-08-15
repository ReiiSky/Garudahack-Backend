package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1 "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// Profile get profile
func Profile(c *gin.Context) {
	token := c.Request.Header.Get("Authtoken")
	service := v1.CreateTokenValidator(token)
	user, status := service.Validate()
	if status {
		api.JSONResponse(http.StatusOK, c.Writer, gin.H{
			"profile": user,
		})
	} else {
		api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
			"message": "user token not found",
		})
	}
}
