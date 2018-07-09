package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"pokedex/internal/log"
	"pokedex/internal/pokemon"
)

const (
	OutputDirectory        = "json/"
	PokemonOutputDirectory = OutputDirectory + "pokemon/"

	serviceName = "pokedex"
)

var (
	version = "dev"

	router *chi.Mux
)

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	log.Infof("Preparing to run %s (%s)", serviceName, version)

	log.Debug("Initiating routes")
	routes()
	listen := fmt.Sprintf(":%s", "8080")
	log.Debug("Listening on", listen)

	var err error

	if len(os.Getenv("POKEDEX_HTTP_TLS_CERTIFICATE")) > 0 && len(os.Getenv("POKEDEX_HTTP_TLS_KEY")) > 0 {
		log.Info("Listen on TLS")
		err = http.ListenAndServeTLS(listen, os.Getenv("POKEDEX_HTTP_TLS_CERTIFICATE"), os.Getenv("POKEDEX_HTTP_TLS_KEY"), router)
	} else {
		err = http.ListenAndServe(listen, router)
	}

	if err != nil {
		log.Fatal(err, "Listen and serve error")
	}
}

func routes() {
	router = chi.NewRouter()
	registerRoutes(router)
}

func registerRoutes(router *chi.Mux) {
	router.Route("/v1", func(r chi.Router) {
		r.Use(telemetryMiddleware)
		r.Route("/pokemon", func(r chi.Router) {
			r.Get("/", getPokemonList)
			r.Route("/{pokemonID}", func(r chi.Router) {
				r.Use(pokemonCTX)
				r.Get("/", getPokemon)
			})
		})
		r.Route("/assets", func(r chi.Router) {
			r.Route("/{fileName}", func(r chi.Router) {
				r.Use(assetCTX)
				r.Get("/", getAsset)
			})
		})
	})
}

func assetCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileName := chi.URLParam(r, "fileName")

		asset, err := pokemon.GetAsset(fileName)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "asset", asset)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func pokemonCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokemonID, err := ParamInt(r, "pokemonID")
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		pokemon, err := pokemon.GetPokemon(pokemonID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "pokemon", pokemon)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ParamInt(r *http.Request, key string) (int, error) {
	val, err := strconv.Atoi(chi.URLParam(r, key))
	if err != nil {
		return 0, err
	}
	return val, nil
}

func telemetryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Debugf("http_request %s %v", r.Method, r.RequestURI)
	})
}
