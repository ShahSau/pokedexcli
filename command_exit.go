package main

import (
	"fmt"
	"os"
)

func exitCommandLine(cfg *config, args ...string) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
