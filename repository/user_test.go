package repository

import (
	"testing"
)

func TestGetOrCreateUserByMobileNumber(t *testing.T) {
	user, error := GetOrCreateUserByMobileNumber("09126333305")
	if error != nil {
		t.Error("", user)
	}
}
