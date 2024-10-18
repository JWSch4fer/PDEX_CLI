package main

import (
	"PDEX_CLI/internal/pokeapi"
	"fmt"
	"log"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	// StartREPL()
}
