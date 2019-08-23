package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"pa-back/jsonmodel"
)

func setUserRoutes() {
	router.HandleFunc("", registerUser).Methods("POST")
}

func getOTP(mobileNumber string) string, error{
	
}

func registerUser(response http.ResponseWriter, request *http.Request) {
	jsonEncoder := json.NewDecoder(request.Body)
	jsonUser := &jsonmodel.User{}
	err := jsonEncoder.Decode(jsonUser)
	if err != nil {
		log.Println(err.Error())
	}
}
