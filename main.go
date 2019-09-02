package main

import (
	"fmt"
	"net/http"
	"pa-back/config"
	_ "pa-back/repository"
	"pa-back/routes"

	"github.com/rs/cors"
)

func main() {
	router := routes.GetRouter()
	fmt.Println("listen on " + config.JSONConfig.Host + ":" + config.JSONConfig.Port)
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.JSONConfig.Host, config.JSONConfig.Port),
		handler)
	if err != nil {
		fmt.Println(err.Error())
	}
}
