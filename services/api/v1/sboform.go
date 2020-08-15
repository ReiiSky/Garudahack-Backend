package v1

import "github.com/Satssuki/Go-Service-Boilerplate/models"

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
