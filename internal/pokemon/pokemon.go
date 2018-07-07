package pokemon

import (
	"bytes"
	"encoding/json"
	"strconv"

	"pokedex/internal/file"
	"pokedex/internal/log"
)

type PokemonListItem struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type PokemonList struct {
	Count    int               `json:"count"`
	Previous string            `json:"previous"`
	Next     string            `json:"next"`
	Results  []PokemonListItem `json:"results"`
}

type PokemonItem struct {
	ID             int    `json:"id"`
	Order          int    `json:"order"`
	Name           string `json:"name"`
	Weight         int    `json:"weight"`
	Height         int    `json:"height"`
	BaseExperience int    `json:"base_experience"`
	Forms          []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"forms"`
	Abilities []struct {
		Slot    int  `json:"slot"`
		Hidden  bool `json:"is_hidden"`
		Ability struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		Effort   int `json:"effort"`
		BaseStat int `json:"base_stat"`
		Stat     struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Sprites struct {
		FrontDefault     string `json:"front_default"`
		BackDefault      string `json:"back_default"`
		FrontFemale      string `json:"front_female"`
		BackFemale       string `json:"back_female"`
		FrontShiny       string `json:"front_shiny"`
		BackShiny        string `json:"back_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		BackShinyFemale  string `json:"back_shiny_female"`
	} `json:"sprites"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetPokemonkList() (PokemonList, error) {
	pokemonList := PokemonList{}

	if err := file.ReadFile(file.OutputDirectory+"pokemon.json", &pokemonList); err != nil {
		log.Error(err)
		return PokemonList{}, err
	}

	return pokemonList, nil
}

func GetPokemon(id int) (interface{}, error) {
	pokemon := struct {
		json.RawMessage
	}{}

	if err := file.ReadFile(file.PokemonOutputDirectory+strconv.Itoa(id)+".json", &pokemon); err != nil {
		log.Error(err)
		return struct{}{}, err
	}

	return pokemon, nil
}

func GetAsset(fileName string) (*bytes.Buffer, error) {
	return file.StreamFileOut(file.SpritesOutputDirectory + fileName)
}
