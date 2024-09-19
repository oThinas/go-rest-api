package main

import (
	"go-rest-api/routes"
	"net/http"
)

func main() {
	routes.HandleRequest()
	http.ListenAndServe(":8080", nil)
}
