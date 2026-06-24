package main

import (
	"pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type locationResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type Config struct {
	Next     string
	Previous string
	Cache    *internal.Cache
}
