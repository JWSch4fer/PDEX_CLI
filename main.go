package main

import (
	"PDEX_CLI/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 15),
	}
	StartREPL(&cfg)
}
