package main

import (
	"go-rest-api/routes"
	"net/http"
)

func main() {
	r := routes.HandleRequest()
	http.ListenAndServe(":8080", r)
}
