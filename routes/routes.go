package routes

import (
	"go-rest-api/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.GetAllPersonalities)

	return r
}
