package models

import "github.com/Kamva/mgm/v3"

// SBOForm form sbo
type SBOForm struct {
	mgm.DefaultModel `bson:",inline"`
	ProductName      string `bson:"productName" json:"productName"`
	Description      string `bson:"description" json:"description"`
	Contract         string `bson:"contract" json:"contract"`
	UnitPrice        int    `bson:"unitPrice" json:"unitPrice"`
	Stok             int    `bson:"stok" json:"stok"`
	UserID           string `bson:"userid" json:"userid"`
}

// GetCollection return the collection of current user model
func (sbo *SBOForm) GetCollection() *mgm.Collection {
	return mgm.Coll(sbo)
}
