package main

import (
	"backend/src/config"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.SetupEnvironment()

	fmt.Println(config.Port)
	fmt.Println(config.ServerAddress)

	r := router.NewAPIRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
