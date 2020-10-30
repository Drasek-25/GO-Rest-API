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
// UPDATE ROUTE WIPES OUT FIELDS THAT WERNT SENT FROM ORIGINAL
// CREATE HELPERS PACKAGE FOR URL PARAMS AND DECODERS
// PUT POKEMON STRUCT INTO SEPERATE MODEL FILE
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

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: UpdatePokemon")

	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		json.NewEncoder(w).Encode("Missing Key")
		return
	}
	fmt.Println("Searching for: ", keys[0])

	var newPokemon parseJson.Pokemon
	err := json.NewDecoder(r.Body).Decode(&newPokemon)
	if err != nil {
		fmt.Println("Invalid pokemon formatting")
		json.NewEncoder(w).Encode("Invalid pokemon formatting")
		return
	}

	index := -1
	for i := 0; i < len(pokemon); i++ {
		if pokemon[i].Name == keys[0] {
			pokemon[i] = newPokemon
			index = i
			break
		}
	}

	if index == -1 {
		json.NewEncoder(w).Encode("Pokemon Not Found")
		fmt.Println("Search Failed")
	} else {
		json.NewEncoder(w).Encode(pokemon[index])
		fmt.Println("Search Succeded")
	}
}

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: CreatePokemon")

	var newPokemon parseJson.Pokemon
	err := json.NewDecoder(r.Body).Decode(&newPokemon)
	if err != nil {
		fmt.Println("Invalid pokemon formatting")
		json.NewEncoder(w).Encode("Invalid pokemon formatting")
		return
	}

	pokemon = append(pokemon, newPokemon)
	log.Println("CreatePokemon was Succesful")
	json.NewEncoder(w).Encode(pokemon[len(pokemon)-1])
}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: DeletePokemon")

	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		json.NewEncoder(w).Encode("Missing Key")
		return
	}
	fmt.Println("Searching for: ", keys[0])

	var deleted parseJson.Pokemon
	for i := 0; i < len(pokemon); i++ {
		if pokemon[i].Name == keys[0] {
			pokemon[len(pokemon)-1], pokemon[i] = pokemon[i], pokemon[len(pokemon)-1]
			deleted = pokemon[len(pokemon)-1]
			pokemon = pokemon[:len(pokemon)-1]
			break
		}
	}

	if deleted.Name == "" {
		json.NewEncoder(w).Encode("Pokemon Not Found")
		fmt.Println("Search Failed")
	} else {
		json.NewEncoder(w).Encode(deleted)
		fmt.Println("Deletion Succeded")
	}
}

func GetAllPokemon(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: GetPokemonByType")
	json.NewEncoder(w).Encode(pokemon)
}
