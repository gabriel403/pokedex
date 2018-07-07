package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"pokedex/internal/file"
	"pokedex/internal/log"
	"pokedex/internal/parser"
	"pokedex/internal/parser/network"
	"pokedex/internal/pokemon"

	"github.com/sirupsen/logrus"
)

var (
	masterList     = pokemon.PokemonList{}
	url            string
	skipMasterList bool
	skipPokemon    bool
)

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	log.Info("Pulling pokemon details and storing in json files")

	flag.StringVar(&url, "url", "http://pokeapi.co/api/v2/pokemon/", "pokemon list url to start from")
	flag.BoolVar(&skipMasterList, "skip-master-list", false, "Skip generating the master list")
	flag.BoolVar(&skipPokemon, "skip-pokemon", false, "Skip processing pokemon")

	flag.Parse()

	err := pokemonList(url)
	if err != nil {
		log.Error(err)
		return
	}

	if skipMasterList {
		log.Infof("Skipping master list")
	} else {
		file.WriteFile(file.OutputDirectory+"pokemon.json", masterList)
	}
}

func pokemonList(url string) error {
	for len(url) > 0 {
		log.Infof("Proccessing url: %s", url)
		res, err := parser.GetPokemonList(url)
		if err != nil {
			log.Error(err)
			return err
		}
		if res.StatusCode != 200 {
			err = fmt.Errorf("unexpected status code %d", res.StatusCode)
			log.Error(err)
			return err
		}

		pokemonList := pokemon.PokemonList{}
		if err := network.UnmarshallResponse(res, &pokemonList); err != nil {
			log.Error(err)
			return err
		}

		if skipPokemon {
			log.Info("Skipping pokemon")
		} else {
			if err := processPokemonList(pokemonList.Results); err != nil {
				log.Error(err)
				return err
			}
		}

		masterList.Results = append(masterList.Results, pokemonList.Results...)
		url = pokemonList.Next
	}

	return nil
}

func processPokemonList(pokemonListItems []pokemon.PokemonListItem) error {
	for _, pokemonListItem := range pokemonListItems {
		log.Infof("Proccessing pokemon: %s", pokemonListItem.Name)

		var res *http.Response
		var err error

		res, err = parser.GetPokemonItem(pokemonListItem.Url)
		log.Debug(res.StatusCode)
		if err != nil {
			log.Error(err)
			return err
		}

		if res.StatusCode == 429 {
			return fmt.Errorf("request limit reached")
		}

		if res.StatusCode != 200 {
			err = fmt.Errorf("unexpected status code")
			log.Error(err)
			return err
		}

		pokemonItem, err := processPokemonListItem(res)
		if err != nil {
			log.Error(err)
			return err
		}

		err = parser.ProcessSprites(&pokemonItem)
		if err != nil {
			log.Error(err)
			return err
		}

		pokemonListItem.Url = "/v1/pokemon/" + strconv.Itoa(pokemonItem.ID)
		file.WriteFile(file.PokemonOutputDirectory+strconv.Itoa(pokemonItem.ID)+".json", pokemonItem)
	}

	return nil
}

func processPokemonListItem(res *http.Response) (pokemon.PokemonItem, error) {
	pokemonItem := pokemon.PokemonItem{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return pokemon.PokemonItem{}, err
	}

	if err := network.UnmarshallBody(body, &pokemonItem); err != nil {
		log.Error(err)
		return pokemon.PokemonItem{}, err
	}

	return pokemonItem, nil
}
