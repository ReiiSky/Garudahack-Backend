package v1

import (
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
func (service *TokenValidationService) Validate() *mongo.SingleResult {
	user := CreateUserService()
	result := user.User.GetCollection().FindOne(mgm.Ctx(), bson.M{
		"token": service.Token,
	})
	return result
}
