package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1 "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// ProposalList list
func ProposalList(c *gin.Context) {
	token := c.Request.Header.Get("Authtoken")
	service := v1.CreateTokenValidator(token)
	user, status := service.Validate()
	if status {
		service := v1.CreateSBOFormService()
		list := service.Get(user.ID.String())
		api.JSONResponse(http.StatusOK, c.Writer, gin.H{
			"list": list,
		})
	} else {
		api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
			"message": "user token not found",
		})
	}
}
