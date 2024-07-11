package main

import (
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return fmt.Errorf("you havent caught this %s yet", pokemonName)
	}

	fmt.Printf("Pokemon: %s\n", pokemon.Name)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("%s ", t.Type.Name)
	}

	for _, a := range pokemon.Abilities {
		fmt.Printf("\nAbility: %s\n", a.Ability.Name)
	}

	for _, s := range pokemon.Stats {
		fmt.Printf(" %s: %v\n", s.Stat.Name, s.BaseStat)
	}

	return nil
}
