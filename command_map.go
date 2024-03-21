package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locationsResp.Next
	cfg.prev = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.prev == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prev)
	if err != nil {
		return err
	}

	cfg.next = locationResp.Next
	cfg.prev = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
