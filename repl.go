package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter command: ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}

		commands := getCommands()

		cmd, ok := commands[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg, args...)

		if err != nil {
			fmt.Println("Error:", err)
		}

		// switch commandName {
		// case "help":
		// 	callbackHelp()

		// case "exit":
		// 	fmt.Println("Exiting...")
		// 	os.Exit(0)

		// default:
		// 	fmt.Println("Unknown command")
		// }

	}
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "display help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "display location areas",
			callback:    callbackMap,
		},
		"mapback": {
			name:        "back",
			description: "display previous location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore {location area name}",
			description: "List all the pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "Catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon name}",
			description: "Vie information about the caught pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the caught pokemon",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    exitCommandLine,
		},
	}
}

func cleanInput(str string) []string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	words := strings.Fields(str)
	return words
}
