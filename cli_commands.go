package main

import (
	"fmt"
	"os"
)

func CallBackHelp(cfg *config) error {
	fmt.Println("Pokedex command line interface options:")
	fmt.Println("")
	for _, cmd := range GetCLICommands() {
		fmt.Printf("%s  | %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func CallBackExit(cfg *config) error {
	os.Exit(0)
	return nil
}
