package main

import (
	"com.fullstack.ecommerce/db"
	"context"
	"fmt"
	"net/http"

	"com.fullstack.ecommerce/server"
	"com.fullstack.ecommerce/utils/config"
)

const defaultPort = "8080"

func main() {
	appConfig := config.App()
	db.Init(context.TODO())

	port := appConfig.Port
	if port == "" {
		port = defaultPort
	}

	router := server.Routes()
	fmt.Println("Started server")
	http.ListenAndServe(":"+port, router)

}
