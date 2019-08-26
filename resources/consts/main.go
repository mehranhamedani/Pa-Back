package consts

import "time"

const (
	OTPLifeTime   int64         = 30000
	OTPMinLength  int           = 5
	OTPSalt       string        = "OTPSalt"
	JWTLifeMinute time.Duration = 5
	JWTKey        string        = "JWTKey"
)
