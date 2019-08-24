package enums

// OTPStatus type
type OTPStatus int

// ToInt func
func (otpStatus OTPStatus) ToInt() int {
	result := 0
	switch otpStatus {
	case OTPStatusEnable:
		result = 1
	}
	return result
}

const (
	// OTPStatusEnable enable
	OTPStatusEnable OTPStatus = 1
)
