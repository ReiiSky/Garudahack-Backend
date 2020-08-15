package v1

import (
	"errors"

	"github.com/Kamva/mgm/v3"
	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
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
	var message string = "Users created"
	currentUser := &user.User
	count, err := currentUser.GetCollection().CountDocuments(mgm.Ctx(), bson.M{
		"name":  user.User.Name,
		"email": user.User.Email,
	})

	if count > 0 {
		message = "Users already defined"
		err = errors.New(message)
	} else {
		err = currentUser.GetCollection().Create(currentUser)
	}
	return message, err
}

// FindUserAndUpdateToken implementation of function in base interface
func (user *UserService) FindUserAndUpdateToken() (string, error) {
	currentUser := &user.User
	filter := bson.M{
		"email":    user.User.Email,
		"password": user.User.Password,
	}
	generatedID := helpers.GenerateID(16)
	result := currentUser.GetCollection().FindOne(mgm.Ctx(), filter)
	userDecoded := models.User{}
	result.Decode(&userDecoded)
	if userDecoded.Email == user.User.Email && userDecoded.Password == user.User.Password {
		userDecoded.Token = generatedID
		object, err := currentUser.GetCollection().UpdateOne(mgm.Ctx(), filter, bson.M{
			"$set": bson.M{
				"token": generatedID,
			},
		})
		return generatedID, err
	}
	return "", errors.New("Could not find an user.")
}
