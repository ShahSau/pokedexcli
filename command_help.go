package main

import (
	"fmt"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Available commands:")
	// fmt.Println("help - display this message")
	// fmt.Println("exit - exit the program")
	for _, cmd := range getCommands() {
		fmt.Printf("%s - %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
