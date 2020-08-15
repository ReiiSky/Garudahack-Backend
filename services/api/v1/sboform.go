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

func AggregatePropsReq(token string) []models.AggregateProps {``
	validator := CreateTokenValidator(token)
	user, _ := validator.Validate()
	sboService := CreateSBOFormService()
	requestService := CreateRequestService()

	filter := bson.M{"userid": user.ID.String()}
	sboS := []models.SBOForm{}
	sboService.Form.GetCollection().SimpleFind(&sboS, filter)

	reqS := []models.Request{}
	requestService.Request.GetCollection().SimpleFind(&reqS, filter)
	results := []models.AggregateProps{}
	for _, x := range sboS {
		for _, y := range reqS {
			if x.UserID == y.UserID && !x.IsActive && len(y.SBOID) != 0 {
				user := CreateUserService()
				user.FindByID(y.SBOID)
				object := models.AggregateProps{
					ProductName: x.ProductName,
					NamaUMKM:    user.User.Name,
					Desc:        y.Description,
					Stock:       y.TotalStock,
					HargaSatuan: x.UnitPrice,
					ID:          x.ID.String(),
				}
				results = append(results, object)
			}
		}
	}
	return results
}
