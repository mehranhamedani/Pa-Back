package repository

import (
	"context"
	"pa-back/model"
	"pa-back/resources/enums"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetOrCreateActiveOTP func
func GetOrCreateActiveOTP(userID primitive.ObjectID) (*model.OTP, error) {
	otpCollection := db.Collection("otps")
	a := time.Now().UTC()
	otpCollection.FindOne(context.Background(), bson.M{"userID": userID, "status": enums.OTPStatusEnable})
}
