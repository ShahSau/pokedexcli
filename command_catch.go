package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("missing pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)

	const threshold = 50

	if randNum > threshold {
		return fmt.Errorf("%s got away", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Println("You caught the pokemon!")

	return nil
}
