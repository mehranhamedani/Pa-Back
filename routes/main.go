package routes

import (
	"pa-back/jsonmodel"
	"pa-back/model"
	"pa-back/resources/consts"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
	setUserRoutes()
}

// GetRouter function
func GetRouter() *mux.Router {
	return router
}

func getToken(userModel *model.User) (string, error) {
	expirationTime := time.Now().Add(consts.JWTLifeMinute * time.Minute)
	apiClaims := jsonmodel.APIClaims{
		UserID:       userModel.UserID,
		FirstName:    userModel.FirstName,
		LastName:     userModel.LastName,
		Email:        userModel.Email,
		UserName:     userModel.UserName,
		MobileNumber: userModel.MobileNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, apiClaims)
	jwtKey := []byte(consts.JWTKey)
	return token.SignedString(jwtKey)
}
