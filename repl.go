package main

import (
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
