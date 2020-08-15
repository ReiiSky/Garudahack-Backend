package route

import (
	v1 "github.com/Satssuki/Go-Service-Boilerplate/controllers/api/v1"
	"github.com/Satssuki/Go-Service-Boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRouter Register each route here
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.AppliAllCORS)

	apig := router.Group("/api")
	{
		v1g := apig.Group("/v1")
		{
			v1g.POST("/signup", v1.InsertUser)
			v1g.POST("/signin", v1.FindUser)
			v1g.POST("/com/request", v1.PlaceRequest)
			v1g.GET("/gen/marketplace", v1.ProductRequestList)
			v1g.POST("/sbo/proposals/form", v1.ProposalForm)
			v1g.POST("/sbo/proposals", v1.ProposalList)
		}
	}
	return router
}
