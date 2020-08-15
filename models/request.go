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
}

// GetCollection return the collection of current request model
func (request *Request) GetCollection() *mgm.Collection {
	return mgm.Coll(request)
}
