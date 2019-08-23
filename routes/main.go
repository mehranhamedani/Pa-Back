package routes

import (
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
