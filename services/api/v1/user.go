package v1

import (
	"errors"

	"github.com/Kamva/mgm/v3"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/services/api/validation"
	"go.mongodb.org/mongo-driver/bson"
)

// UserService struct that wrapper the user model api
type UserService struct {
	User models.User
}

// CreateUserService just a shorthand create object from struct
func CreateUserService() UserService {
	return UserService{models.User{}}
}

// Insert implementation of function in base interface
func (user *UserService) Insert() (string, error) {
	err := validation.ValidateUser(&user.User)
	var message string = "Users created"
	if err == nil {
		_ = err
		currentUser := &user.User
		count, Err := currentUser.GetCollection().CountDocuments(mgm.Ctx(), bson.M{
			"name":  user.User.Name,
			"email": user.User.Email,
		})
		err = Err
		if count > 0 {
			message = "Users already defined"
			err = errors.New(message)
		} else {
			err = currentUser.GetCollection().Create(currentUser)
		}
	}
	return message, err
}
