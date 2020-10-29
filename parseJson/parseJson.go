package parseJson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Stats struct {
	Stat      string  `json:"stat"`
	Effort    float64 `json:"effort"`
	Base_stat float64 `json:"base_stat"`
}
type Sprites struct {
	Back_female        string `json:"back_female"`
	Back_shiny_female  string `json:"back_shiny_female"`
	Back_default       string `json:"back_default"`
	Front_female       string `json:"front_female"`
	Front_shiny_female string `json:"front_shiny_female"`
	Back_shiny         string `json:"back_shiny"`
	Front_default      string `json:"front_default"`
	Front_shiny        string `json:"front_shiny"`
}
type Pokemon struct {
	Abilities       []string `json:"abilities"`
	Stats           []Stats  `json:"stats"`
	Name            string   `json:"name"`
	Weight          float64  `json:"weight"`
	Moves           []string `json:"moves"`
	Sprites         Sprites  `json:"sprites"`
	Height          float64  `json:"height"`
	Id              float64  `json:"id"`
	Base_experience float64  `json:"base_experience"`
	Types           []string `json:"types"`
}
type JsonObject struct {
	Pokemons []Pokemon `json:"Pokemons"`
}

func ParseJson() []Pokemon {
	content, err := ioutil.ReadFile("data/pokemons.json")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	var Pokemons JsonObject
	err2 := json.Unmarshal([]byte(text), &Pokemons)
	if err2 != nil {
		fmt.Println("error: ", err2)
	}
	if Pokemons.Pokemons == nil {
		fmt.Println("Error Unmarshaling Json")
	} else {
		fmt.Println("Pokemons Json Unmarshaled Succesfully")
	}
	return Pokemons.Pokemons
}
