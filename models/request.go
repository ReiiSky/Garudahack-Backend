package models

import (
	"github.com/Kamva/mgm/v3"
)

// Request struct
type Request struct {
	mgm.DefaultModel `bson:",inline"`
	ProductName      string `bson:"productName" json:"productName"`
	Description      string `bson:"description" json:"description"`
	TotalStock       int    `bson:"totalStock" json:"totalStock"`
	UserID           string `bson:"userid" json:"userid"`
	SBOID            string `bson:"sboid" json:"sboid"`
	ImageLink        string `bson:"imageLink" json:"imageLink"`
}

// GetCollection return the collection of current request model
func (request *Request) GetCollection() *mgm.Collection {
	return mgm.Coll(request)
}
