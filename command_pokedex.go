package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	fmt.Println("Pokemon in Pokedex:")

	for _, p := range cfg.caughtPokemon {
		fmt.Printf("Pokemon: %s\n", p.Name)
		fmt.Printf("Base Experience: %d\n", p.BaseExperience)
		fmt.Printf("Height: %d\n", p.Height)
		fmt.Printf("Weight: %d\n", p.Weight)
		fmt.Printf("Types: ")
		for _, t := range p.Types {
			fmt.Printf("%s ", t.Type.Name)
		}
		for _, a := range p.Abilities {
			fmt.Printf("\nAbility: %s\n", a.Ability.Name)
		}

		for _, s := range p.Stats {
			fmt.Printf(" %s: %v\n", s.Stat.Name, s.BaseStat)
		}
	}

	return nil
}
