package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(input string) []string {

	if input == "" {
		return []string{}
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return strings.Fields(input)

}
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show help message",
			callback:    commandHelp,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n" + "Usage: \n" + "\n" + "help: Displays a help message \n" + "exit: Exit the Pokedex")
	return nil
}
