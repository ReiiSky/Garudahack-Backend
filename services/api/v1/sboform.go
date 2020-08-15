package v1

import (
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"go.mongodb.org/mongo-driver/bson"
)

// SBOFormService form sbo
type SBOFormService struct {
	Form models.SBOForm
}

// CreateSBOFormService just a shorthand create object from struct
func CreateSBOFormService() SBOFormService {
	return SBOFormService{models.SBOForm{}}
}

// Insert SBO form
func (service *SBOFormService) Insert() error {
	return service.Form.GetCollection().Create(&service.Form)
}

// Get SBO form
func (service *SBOFormService) Get(userid string) []models.SBOForm {
	forms := []models.SBOForm{}
	service.Form.GetCollection().SimpleFind(&forms, bson.M{
		"userid": userid,
	})
	return forms
}
