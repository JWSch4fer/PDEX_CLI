package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func CallBackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("Pokemon not found in this area 0_o")
	}
	pokemonName := args[0]

	pokemon_data, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	//probability of catching will be 50%
	threshold := 0.5 * float32(pokemon_data.BaseExperience)
	randNum := rand.Intn(pokemon_data.BaseExperience)
	if randNum < int(threshold) {
		return fmt.Errorf("FAILED TO CATCH %s", pokemonName)
	}
	fmt.Printf("%s WAS CAPTURED >_< \n", pokemonName)
	cfg.caughtPokemon[pokemonName] = pokemon_data
	return nil
}

func CallBackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("Pokemon not found in this area 0_o")
	}
	pokemonName := args[0]

	pokemon_data, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("You haven't caught %s", pokemonName)
	}

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
	return nil
}
