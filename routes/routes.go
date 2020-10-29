package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Drasek-25/Rest_API/parseJson"
)

var pokemon = parseJson.ParseJson()

func AllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(pokemon)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
