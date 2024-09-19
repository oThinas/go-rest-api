package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
	"net/http"
)

func main() {
	database.Connect()
	r := routes.HandleRequest()
	http.ListenAndServe(":8080", r)
}
