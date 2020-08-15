package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// InsertUser sample controller to perform insert user function
func InsertUser(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateUserService()
	err := helpers.ReadByteAndParse(c.Request.Body, &service.User)

	if err == nil {
		message, err := service.Insert()
		if err == nil {
			api.JSONResponse(http.StatusCreated, c.Writer, gin.H{
				"status":  "ok",
				"message": message,
			})
			return
		} else {
			if message != "Users created" {
				api.JSONResponse(409, c.Writer, gin.H{
					"status":  "ok",
					"message": message,
				})
				return
			}
		}
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"status":  "failure",
		"message": err,
	})
}
