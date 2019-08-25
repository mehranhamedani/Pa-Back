package services

import "pa-back/repository"

// GetOTP func
func GetOTP(mobileNumber string) (string, error) {
	userModel, error := repository.GetOrCreateUserByMobileNumber(mobileNumber)
	otpModel, error := repository.GetOrCreateActiveOTP(userModel.UserID)
	return otpModel.OTP, error
}
