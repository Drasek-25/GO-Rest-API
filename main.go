package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Drasek-25/Rest_API/routes"
	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", routes.HomePage)
	r.HandleFunc("/articles", routes.AllArticles).Methods("GET")

	fmt.Println("Server Running at 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	handleRequests()
}
