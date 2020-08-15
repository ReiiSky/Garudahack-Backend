package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// PlaceRequest ....
func PlaceRequest(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateRequestService()
	err := helpers.ReadByteAndParse(c.Request.Body, &service.Request)

	if err == nil {
		message, err := service.PlaceRequest()
		if err == nil {
			api.JSONResponse(http.StatusCreated, c.Writer, gin.H{
				"status":  "ok",
				"message": message,
			})
			return
		}
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"status":  "failure",
		"message": err,
	})
}

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
