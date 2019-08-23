package model

// OTP struct
type OTP struct {
	UserID     string `bson:"userID"`
	OTP        string `bson:"otp"`
	Status     int    `bson:"status"`
	ExpireTime int64  `bson:"expireTime"`
}
