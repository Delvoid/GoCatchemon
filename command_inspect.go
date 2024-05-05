package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon>")
	}

	name := args[0]
	caughtPokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("Pokemon not caught")
	}

	fmt.Println("Name:", caughtPokemon.Name)
	fmt.Println("Height:", caughtPokemon.Height)
	fmt.Println("Weight:", caughtPokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range caughtPokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range caughtPokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}
