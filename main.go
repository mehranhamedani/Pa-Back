package main

import (
	"fmt"
	"net/http"
	"pa-back/config"
	_ "pa-back/repository"
	"pa-back/routes"
)

func main() {
	router := routes.GetRouter()
	fmt.Println("listen on " + config.JSONConfig.Host + ":" + config.JSONConfig.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.JSONConfig.Host, config.JSONConfig.Port), router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
