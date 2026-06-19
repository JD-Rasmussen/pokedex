package main

import (
	"strings"
)

func cleanInput(input string) []string {

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return strings.Fields(input)

}
