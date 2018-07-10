package file

import (
	"os"
)

var (
	OutputDirectory        = os.Getenv("POKEDEX_ASSETS_DIR")
	PokemonOutputDirectory = OutputDirectory + "pokemon/"
	SpritesOutputDirectory = OutputDirectory + "sprites/"
)

func init() {
	os.MkdirAll(PokemonOutputDirectory, 0755)
	os.MkdirAll(SpritesOutputDirectory, 0755)
}
