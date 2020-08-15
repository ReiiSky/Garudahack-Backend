package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1 "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// ProposalForm form
func ProposalForm(c *gin.Context) {
	defer c.Request.Body.Close()

	token := c.Request.Header.Get("Authtoken")
	service := v1.CreateTokenValidator(token)
	user, status := service.Validate()
	if status {
		form := v1.CreateSBOFormService()
		err := helpers.ReadByteAndParse(c.Request.Body, &form.Form)
		if err == nil {
			form.Form.UserID = user.ID.String()
			form.Form.IsActive = false
			form.Insert()
			api.JSONResponse(http.StatusOK, c.Writer, gin.H{
				"message": "proposal inserted",
			})
		}
	} else {
		api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
			"message": "user token not found",
		})
	}
}
