package main

import (
	"fmt"
	"net/http"
	"pa-back/config"
	_ "pa-back/repository"
)

func main() {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.JSONConfig.Host, config.JSONConfig.Port), )
	if err != nil {
		fmt.Println(error.Error())
	}
}
