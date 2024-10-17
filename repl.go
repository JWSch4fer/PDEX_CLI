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

		fmt.Printf("I got your message:\n%v\n", clean_text)
	}
}

func CleanInput(s string) []string {
	split_string := strings.Fields(strings.ToLower(s))

	return split_string
}
