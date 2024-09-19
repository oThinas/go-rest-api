package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middlewares"

	"github.com/gorilla/mux"
)

func HandleRequest() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.Use(middlewares.SetContentType)

	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")

	return r
}
