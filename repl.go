package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartREPL(cfg *config, show_list_of_commands bool) {
	for {
		fmt.Println("==================================================================================")
		fmt.Println("Hello, please type something...")
		fmt.Println("==================================================================================")
		AvailableCommands := GetCLICommands()

		if show_list_of_commands {
			AvailableCommands["help"].callback(cfg, "")
			show_list_of_commands = false
		}
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		//Read user input and parse to lowercase slice
		text := scanner.Text()
		clean_text := CleanInput(text)

		// If enter is pressed we just continue
		if len(clean_text) == 0 {
			continue
		}

		InputCommand := clean_text[0]
		args := []string{}
		if len(clean_text) > 1 {
			args = clean_text[1:]
		}

		command, ok := AvailableCommands[InputCommand]
		if ok == false {
			fmt.Printf("\n>>>This is not an available command: %v\n", InputCommand)
			AvailableCommands["help"].callback(cfg, args...)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error in callback function for  %v: %v\n", command.name, err)
		}

	}
}

func CleanInput(s string) []string {
	split_string := strings.Fields(strings.ToLower(s))

	return split_string
}

type cli_commands struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func GetCLICommands() map[string]cli_commands {
	return map[string]cli_commands{
		"help": {
			name:        "help                 ",
			description: "print the available commands",
			callback:    CallBackHelp,
		},
		"exit": {
			name:        "exit                 ",
			description: "shut down the Pokedex",
			callback:    CallBackExit,
		},
		"map": {
			name:        "map                  ",
			description: "Show 20 available locations",
			callback:    CallBackMap,
		},
		"mapb": {
			name:        "mapb                 ",
			description: "Show previous 20 locations",
			callback:    CallBackMapb,
		},
		"explore": {
			name:        "explore {area_name}  ",
			description: "show all pokemon encounters in this area",
			callback:    CallBackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}  ",
			description: "try to catch a pokemon",
			callback:    CallBackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "inspect pokemon you've caught",
			callback:    CallBackInspect,
		},
		"pokedex": {
			name:        "pokedex               ",
			description: "inspect all pokemon you've caught",
			callback:    CallBackPokedex,
		},
	}
}
