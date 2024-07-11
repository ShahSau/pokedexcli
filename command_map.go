package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	res, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaUrl)

	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, locationArea := range res.Results {
		fmt.Printf(" - %v\n", locationArea.Name)
	}

	cfg.nextLocationAreaUrl = res.Next
	cfg.previousLocationAreaUrl = res.Previous

	return nil
}

func callbackMapBack(cfg *config, args ...string) error {

	if cfg.previousLocationAreaUrl == nil {
		fmt.Println("No previous location area")
		return nil
	}
	res, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousLocationAreaUrl)

	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, locationArea := range res.Results {
		fmt.Printf(" - %v\n", locationArea.Name)
	}

	cfg.nextLocationAreaUrl = res.Next
	cfg.previousLocationAreaUrl = res.Previous

	return nil
}
