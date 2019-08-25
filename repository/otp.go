package repository

import (
	"context"
	"pa-back/model"
	"pa-back/resources/consts"
	"pa-back/resources/enums"
	"pa-back/resources/texts"
	"pa-back/utilities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetOrCreateActiveOTP func
func GetOrCreateActiveOTP(userID primitive.ObjectID) (*model.OTP, error) {
	otpModel, error := GetActiveOTP(userID)
	if otpModel.OTPID.IsZero() {
		otpModel, error = CreateOTP(userID)
	}
	return otpModel, error
}

func GetActiveOTP(userID primitive.ObjectID) (*model.OTP, error) {
	var error error
	otpModel := model.OTP{}
	otpCollection := db.Collection(enums.DBCollectionNameOTPs.ToString())
	utc := time.Now().UTC().Unix()
	result := otpCollection.FindOne(context.Background(), bson.M{"userID": userID, "status": enums.OTPStatusEnable, "expireTime": bson.M{"$gt": utc}})
	if result.Err() == nil {
		error = result.Decode(&otpModel)
	} else if result.Err().Error() != texts.En_MNDIR {
		error = result.Err()
	}

	return &otpModel, error
}

func CreateOTP(userID primitive.ObjectID) (*model.OTP, error) {
	var error error
	otpModel := &model.OTP{}
	otp := utilities.GenerateNewOTP()
	otpCollection := db.Collection(enums.DBCollectionNameOTPs.ToString())
	expireTime := time.Now().UTC().Unix() + consts.OTPLifeTime
	result, error := otpCollection.InsertOne(context.Background(), bson.M{"userID": userID, "otp": otp, "status": enums.OTPStatusEnable, "expireTime": expireTime})
	if error == nil {
		otpModel, error = GetOTPByID(result.InsertedID.(primitive.ObjectID))
	}
	return otpModel, error
}

func GetOTPByID(otpID primitive.ObjectID) (*model.OTP, error) {
	var error error
	otpModel := model.OTP{}
	otpCollection := db.Collection(enums.DBCollectionNameOTPs.ToString())
	result := otpCollection.FindOne(context.Background(), bson.M{"_id": otpID})
	if result.Err() != nil && result.Err().Error() != texts.En_MNDIR {
		error = result.Err()
	}
	error = result.Decode(&otpModel)
	return &otpModel, error
}
