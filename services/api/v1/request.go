package v1

import (
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"go.mongodb.org/mongo-driver/bson"
)

// RequestService struct that wrapper the request model api
type RequestService struct {
	Request models.Request
}

// CreateRequestService just a shorthand create object from struct
func CreateRequestService() RequestService {
	return RequestService{models.Request{}}
}

// PlaceRequest ....
func (request *RequestService) PlaceRequest() (string, error) {
	var message string = "Request created"
	currentRequest := &request.Request
	err := currentRequest.GetCollection().Create(currentRequest)
	return message, err
}

// ProductRequestList ....
func (request *RequestService) ProductRequestList() ([]models.Request, error) {
	currentRequest := &request.Request
	result := []models.Request{}
	err := currentRequest.GetCollection().SimpleFind(&result, bson.M{})
	return result, err
}

// RequestList ....
func (request *RequestService) RequestList() ([]models.Request, error) {
	currentRequest := &request.Request
	result := []models.Request{}
	err := currentRequest.GetCollection().SimpleFind(&result, bson.M{"productName": "Susu Murni "})
	return result, err
}
