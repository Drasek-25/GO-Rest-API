package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Drasek-25/Rest_API/parseJson"
)

// ADD HTTP STATUS CODES
var pokemon = parseJson.ParseJson()

func GetPokemonByType(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: GetPokemonByType")
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		json.NewEncoder(w).Encode("Missing Key")
		return
	}
	fmt.Println("Searching for: ", keys[0])

	var result []string

	for i := 0; i < len(pokemon); i++ {
		for j := 0; j < len(pokemon[i].Types); j++ {
			if strings.ToLower(pokemon[i].Types[j]) == strings.ToLower(keys[0]) {
				result = append(result, pokemon[i].Name)
				break
			}
		}

	}
	if result == nil {
		json.NewEncoder(w).Encode("Type Not Found")
		fmt.Println("Search Failed")
	} else {
		json.NewEncoder(w).Encode(result)
		fmt.Println("Search Succeded")
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: HomePage")
}

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: CreatePokemon")
	var newPokemon parseJson.Pokemon
	err := json.NewDecoder(r.Body).Decode(&newPokemon)
	if err != nil {
		fmt.Println("Invalid formatting")
		json.NewEncoder(w).Encode("Invalid formatting")
		return
	}

	pokemon = append(pokemon, newPokemon)
	log.Println("CreatePokemon was Succesful")
	json.NewEncoder(w).Encode(pokemon[len(pokemon)-1])
}

func GetAllPokemon(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: GetPokemonByType")
	json.NewEncoder(w).Encode(pokemon)
}
