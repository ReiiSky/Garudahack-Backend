package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1 "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// RequestList ....
func RequestList(c *gin.Context) {
	defer c.Request.Body.Close()

	token := c.Request.Header.Get("Authtoken")
	services := v1.CreateTokenValidator(token)
	user, status := services.Validate()

	if status {
		service := v1s.CreateRequestService()
		result, err := service.RequestList(user.ID.String())
		if err == nil {
			api.JSONResponse(http.StatusOK, c.Writer, gin.H{
				"RequestList": result,
			})
			return
		}
	} else {
		api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
			"message": "user token not found",
		})
	}

}
