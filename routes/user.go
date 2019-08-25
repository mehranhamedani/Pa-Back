package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"pa-back/jsonmodel"
	"pa-back/resources/texts"
	"pa-back/services"
	"pa-back/utilities"
)

func setUserRoutes() {
	router.HandleFunc("/user/registerUser", registerUser).Methods("POST")
	router.HandleFunc("/user/getOTP/{mobileNumber}", getOTP).Methods("GET")
}

func getOTP(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	mobileNumber := params["mobileNumber"]
	otp, error := services.GetOTP(mobileNumber)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, texts.Fa_ABMAS, otp)
}

func registerUser(response http.ResponseWriter, request *http.Request) {
	jsonEncoder := json.NewDecoder(request.Body)
	jsonUser := &jsonmodel.User{}
	err := jsonEncoder.Decode(jsonUser)
	if err != nil {
		log.Println(err.Error())
	}
}
