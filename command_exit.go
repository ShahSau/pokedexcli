package main

import (
	"fmt"
	"os"
)

func exitCommandLine() error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
