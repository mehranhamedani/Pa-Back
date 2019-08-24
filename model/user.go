package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	UserID       primitive.ObjectID `bson:"_id"`
	FirstName    string             `bson:"firstName"`
	LastName     string             `bson:"lastName"`
	Email        string             `bson:"email"`
	UserName     string             `bson:"userName"`
	MobileNumber string             `bson:"mobileNumber"`
	Password     string             `bson:"password"`
}
