package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OTP struct
type OTP struct {
	OTPID      primitive.ObjectID `bson:"_id"`
	UserID     primitive.ObjectID `bson:"userID"`
	OTP        string             `bson:"otp"`
	Status     int                `bson:"status"`
	ExpireTime int64              `bson:"expireTime"`
}
