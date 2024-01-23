package main

import (
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("teste")

	r := router.NewAPIRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
