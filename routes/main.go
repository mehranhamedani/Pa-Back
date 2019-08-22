package routes

import (
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

func getRoutes() *mux.Router {
	return router
}
