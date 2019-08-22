package routes

import (
	"net/http"
)

func setUserRoutes() {
	router.HandleFunc("", registerUser).Methods("POST")
}

func registerUser(response http.ResponseWriter, request *http.Request) {

}
