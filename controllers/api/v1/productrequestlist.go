package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// ProductRequestList ...
func ProductRequestList(c *gin.Context) {
	service := v1s.CreateRequestService()
	result, err := service.ProductRequestList()
	if err == nil {
		api.JSONResponse(http.StatusOK, c.Writer, gin.H{
			"RequestList": result,
		})
		return
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"status":  "failure",
		"message": err,
	})
}
