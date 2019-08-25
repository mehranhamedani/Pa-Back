package repository

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetOrCreateActiveOTP(t *testing.T) {
	userId, error := primitive.ObjectIDFromHex("5d5e3111be61d14c806c7b0d")
	user, error := GetOrCreateActiveOTP(userId)
	if error != nil {
		t.Error("", user)
	}
}
