package main

import (
	"fmt"
)

func CallBackPokedex(cfg *config, args ...string) error {
	fmt.Println("--------------Showing all PokeDex info--------------")
	for _, pokemon_data := range cfg.caughtPokemon {
		fmt.Printf("Name: %s\n", pokemon_data.Name)
		fmt.Printf("Height: %v\n", pokemon_data.Height)
		fmt.Printf("Weight: %v\n", pokemon_data.Weight)
		for _, stat := range pokemon_data.Stats {
			fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		for _, typ := range pokemon_data.Types {
			fmt.Printf(" - Type: %s", typ.Type.Name)
		}
		fmt.Println("")
	}
	fmt.Println("---------------------------------------------------")
	return nil
}
