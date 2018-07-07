package file

import (
	"os"
)

const (
	OutputDirectory        = "assets/"
	PokemonOutputDirectory = OutputDirectory + "pokemon/"
	SpritesOutputDirectory = OutputDirectory + "sprites/"
)

func init() {
	os.MkdirAll(PokemonOutputDirectory, 0755)
	os.MkdirAll(SpritesOutputDirectory, 0755)
}
