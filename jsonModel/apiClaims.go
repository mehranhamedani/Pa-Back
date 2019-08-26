package jsonmodel

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// APIClaims struct
type APIClaims struct {
	UserID       primitive.ObjectID `bson:"_id"`
	FirstName    string             `bson:"firstName"`
	LastName     string             `bson:"lastName"`
	Email        string             `bson:"email"`
	UserName     string             `bson:"userName"`
	MobileNumber string             `bson:"mobileNumber"`
	jwt.StandardClaims
}
