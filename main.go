package main

import (
	"fmt"
	"net/http"
)

func main() {
	HandleRequest()
	http.ListenAndServe(":8080", nil)
}

func HandleRequest() {
	http.HandleFunc("/", Home)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}
