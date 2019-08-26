package jsonmodel

// LoginRequest struct
type LoginRequest struct {
	MobileNumber string `json:"mobileNumber"`
	OTP          string `json:"otp"`
}
