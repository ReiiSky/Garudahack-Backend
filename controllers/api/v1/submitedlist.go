package v1

import (
	"github.com/gin-gonic/gin"
)

// SubmitedList submited
func SubmitedList(c *gin.Context) {
	/* defer c.Request.Body.Close()

	token := c.Request.Header.Get("Authtoken")
	service := v1.CreateTokenValidator(token)
	user, status := service.Validate()

	if status {
		service := v1s.CreateRequestService()
		err := helpers.ReadByteAndParse(c.Request.Body, &service.Request)

		if err == nil {
			service.Request.UserID = user.ID.String()
			service.Request.ImageLink = "http://3.23.126.114/_nuxt/img/lan_umkm-1.a7cf81e.jpg"
			message, err := service.PlaceRequest()
			if err == nil {
				api.JSONResponse(http.StatusCreated, c.Writer, gin.H{
					"status":  "ok",
					"message": message,
				})
				return
			}
		}
	} else {
		api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
			"message": "user token not found",
		})
	} */
}
