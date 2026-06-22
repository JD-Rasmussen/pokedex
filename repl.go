package main

import (
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	os.Exit(0)
	return nil
}
