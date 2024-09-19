package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	database.Connect()
	r := routes.HandleRequest()
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:5173"}))(r))
}
