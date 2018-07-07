package parser

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"pokedex/internal/file"
	"pokedex/internal/log"
	"pokedex/internal/parser/network"
	"pokedex/internal/pokemon"
)

func GetPokemonList(url string) (*http.Response, error) {
	req, err := network.CreateGETRequest(url)
	if err != nil {
		return &http.Response{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		logErrorBody(res)
		return &http.Response{}, err
	}

	return res, nil
}

func GetPokemonItem(url string) (*http.Response, error) {
	req, err := network.CreateGETRequest(url)
	if err != nil {
		return &http.Response{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
		logErrorBody(res)
		return &http.Response{}, err
	}

	return res, nil
}

func logErrorBody(r *http.Response) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return
	}
	log.Error(string(body))
}

func ProcessSprites(pokemonItem *pokemon.PokemonItem) error {
	if len(pokemonItem.Sprites.FrontDefault) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_front_default.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.FrontDefault); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.FrontDefault = "/v1/assets/" + fileName
	}

	if len(pokemonItem.Sprites.BackDefault) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_back_default.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.BackDefault); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.BackDefault = "/v1/assets/" + fileName
	}

	if len(pokemonItem.Sprites.FrontFemale) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_front_female.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.FrontFemale); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.FrontFemale = "/v1/assets/" + fileName
	}

	if len(pokemonItem.Sprites.BackFemale) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_back_female.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.BackFemale); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.BackFemale = "/v1/assets/" + fileName
	}

	if len(pokemonItem.Sprites.FrontShiny) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_front_shiny.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.FrontShiny); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.FrontShiny = "/v1/assets/" + fileName
	}

	if len(pokemonItem.Sprites.BackShiny) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_back_shiny.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.BackShiny); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.BackShiny = "/v1/assets/" + fileName
	}

	if len(pokemonItem.Sprites.FrontShinyFemale) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_front_shiny_female.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.FrontShinyFemale); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.FrontShinyFemale = "/v1/assets/" + fileName
	}

	if len(pokemonItem.Sprites.BackShinyFemale) > 0 {
		fileName := strconv.Itoa(pokemonItem.ID) + "_back_shiny_female.png"
		if err := file.DownloadFile(file.SpritesOutputDirectory+fileName, pokemonItem.Sprites.BackShinyFemale); err != nil {
			log.Error(err)
			return err
		}

		pokemonItem.Sprites.BackShinyFemale = "/v1/assets/" + fileName
	}

	return nil
}
