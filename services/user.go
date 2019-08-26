package services

import (
	"pa-back/model"
	"pa-back/repository"
)

// GetOTP func
func GetOTP(mobileNumber string) (string, error) {
	userModel, error := repository.GetOrCreateUserByMobileNumber(mobileNumber)
	otpModel, error := repository.GetOrCreateActiveOTP(userModel.UserID)
	return otpModel.OTP, error
}

func Login(mobileNumber string, otp string) (*model.User, error) {
	userModel, error := repository.GetUserByMobileNumber(mobileNumber)
	otpModel, error := repository.GetOTPByUserIDAndOTP(userModel.UserID, otp)
	if !otpModel.OTPID.IsZero() {
		error = repository.UseOTP(otpModel.OTPID)
	}
	return userModel, error
}
