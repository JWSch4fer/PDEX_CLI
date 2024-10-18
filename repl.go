package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartREPL(cfg *config) {
	for {
		fmt.Println("Hello, please type something...")
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

		AvailableCommands := GetCLICommands()

		command, ok := AvailableCommands[InputCommand]
		if ok == false {
			fmt.Printf("\n>>>This is not an available command: %v\n", InputCommand)
			AvailableCommands["help"].callback(cfg)
			continue
		}

		err := command.callback(cfg)
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
	callback    func(*config) error
}

func GetCLICommands() map[string]cli_commands {
	return map[string]cli_commands{
		"help": {
			name:        "help",
			description: "print the available commands",
			callback:    CallBackHelp,
		},
		"exit": {
			name:        "exit",
			description: "shut down the Pokedex",
			callback:    CallBackExit,
		},
		"map": {
			name:        "map ",
			description: "Show 20 available locations",
			callback:    CallBackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 locations",
			callback:    CallBackMapb,
		},
	}
}
