package main

import (
	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/route"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	models.SetConfig()
	models.DatabasePing()
	route.SetupRouter().Run(helpers.GetPORT())
}
