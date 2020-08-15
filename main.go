package main

import (
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	v1 "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	models.SetConfig()
	models.DatabasePing()
	// v1.AggregatePropsReq("8ED0207A25817D70")
	route.SetupRouter().Run(helpers.GetPORT())
}
