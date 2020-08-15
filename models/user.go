package models

import (
	"github.com/Kamva/mgm/v3"
)

// User just dummy struct
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name" json:"name"`
	Email            string `bson:"email" json:"email"`
	Password         string `bson:"password" json:"password"`
	Role             string `bson:"role" json:"role"`
	OrganizationName string `bson:"organizationName" json:"organizationName"`
	Token            string `bson:"token"`
}

// GetCollection return the collection of current user model
func (user *User) GetCollection() *mgm.Collection {
	return mgm.Coll(user)
}
