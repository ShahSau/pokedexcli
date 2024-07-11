package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter text: ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		commands := getCommands()

		cmd, ok := commands[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()

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
	callback    func()
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "display help menu",
			callback:    callbackHelp,
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
