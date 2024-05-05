package main

import (
	"time"
)

func main() {
	pokeClient := NewClient(5*time.Second, time.Minute*5)
	cfg := &Config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]Pokemon{},
	}
	start(cfg)
}
