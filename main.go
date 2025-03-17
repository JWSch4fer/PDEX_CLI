package main

import (
	"PDEX_CLI/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 15),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	var show_list_of_commands bool = true
	StartREPL(&cfg, show_list_of_commands)
}
