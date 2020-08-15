package v1

import (
	"github.com/Kamva/mgm/v3"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"go.mongodb.org/mongo-driver/bson"
)

// TokenValidationService service for validating the token in db
type TokenValidationService struct {
	Token string
}

// CreateTokenValidator just for validator
func CreateTokenValidator(token string) TokenValidationService {
	return TokenValidationService{
		Token: token,
	}
}

// Validate func for validate the token
func (service *TokenValidationService) Validate() (models.User, bool) {
	user := CreateUserService()
	result := user.User.GetCollection().FindOne(mgm.Ctx(), bson.M{
		"token": service.Token,
	})
	result.Decode(&user.User)
	if user.User.Token == service.Token {
		return user.User, true
	}
	return user.User, false
}
