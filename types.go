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
	Next          string
	Previous      string
	Cache         *internal.Cache
	SecondCommand string
	Pokedex       map[string]pokemonData
}

type pokemonResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type pokemonData struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
