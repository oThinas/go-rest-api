package routes

import (
	"go-rest-api/controllers"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
}
