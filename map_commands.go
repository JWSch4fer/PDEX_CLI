package main

import (
	"errors"
	"fmt"
)

func CallBackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Locations: ")
	for _, area := range resp.Results {
		fmt.Printf("Name: %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return err
}

func CallBackMapb(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("This is the first page...")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Locations: ")
	for _, area := range resp.Results {
		fmt.Printf("Name: %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return err
}
