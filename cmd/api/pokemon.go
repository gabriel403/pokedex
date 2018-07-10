package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"pokedex/internal/log"
	"pokedex/internal/pokemon"
)

func getPokemonList(w http.ResponseWriter, r *http.Request) {
	pokemonList, err := pokemon.GetPokemonkList()
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	result, err := json.MarshalIndent(pokemonList, "", "  ")
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon := r.Context().Value("pokemon")
	result, err := json.MarshalIndent(pokemon, "", "  ")
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}

func getAsset(w http.ResponseWriter, r *http.Request) {
	asset := r.Context().Value("asset").(*bytes.Buffer)

	w.Header().Set("Content-Type", http.DetectContentType(asset))
	w.WriteHeader(200)
	w.Write(asset)
}
