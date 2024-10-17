package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartREPL() {
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
			fmt.Printf("This is not an available command: %v", InputCommand)
			AvailableCommands["help"].callback()
			continue
		}

		err := command.callback()
		if err != nil {
			_ = fmt.Errorf("Error in callback function for  %v: %w", command.name, err)
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
	callback    func() error
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
	}
}
