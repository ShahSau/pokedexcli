package main

import (
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("missing location area name")
	}

	locationAreaName := args[0]

	locationName, err := cfg.pokeapiClient.GetDetailsOfLocationArea(locationAreaName)

	if err != nil {
		return err
	}
	fmt.Println("Pokemon in %S:\n", locationName.Name)
	for _, pokemon := range locationName.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
