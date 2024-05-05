package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {

	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s \n", name)
	}

	return nil
}
