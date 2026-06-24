package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		"map": {
			name:        "map",
			description: "List map of the Pokedex",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous page of the Pokedex",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore pokemon in location",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Printf("Welcome to the Pokedex!\n" + "Usage: \n" + "\n" + "help: Displays a help message \n" + "exit: Exit the Pokedex")
	return nil
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if cfg.Next != "" {
		url = cfg.Next
	}
	if val, exists := cfg.Cache.Get(url); exists {
		// Use cached data
		locations := locationResponse{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached map data: %v", err)
		}

		// Update next and previous in config
		cfg.Next = locations.Next
		cfg.Previous = locations.Previous

		// Print the location
		for _, location := range locations.Results {
			fmt.Println(location.Name)

		}
		return nil
	}
	data, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch map data: %v", err)
	}
	defer data.Body.Close()
	if data.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch map data: received status code %d", data.StatusCode)
	}
	dat, err := io.ReadAll(data.Body)
	if err != nil {
		return fmt.Errorf("failed to read map data: %v", err)
	}
	locations := locationResponse{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return fmt.Errorf("failed to unmarshal map data: %v", err)
	}
	//update next and previous in config
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	//print the location
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	if val, exists := cfg.Cache.Get(cfg.Previous); exists {
		// Use cached data
		locations := locationResponse{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached map data: %v", err)
		}

		// Update next and previous in config
		cfg.Next = locations.Next
		cfg.Previous = locations.Previous

		// Print the location
		for _, location := range locations.Results {
			fmt.Println(location.Name)

		}
		return nil
	}
	data, err := http.Get(cfg.Previous)
	if err != nil {
		return fmt.Errorf("failed to fetch map data: %v", err)
	}
	defer data.Body.Close()
	if data.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch map data: received status code %d", data.StatusCode)
	}
	dat, err := io.ReadAll(data.Body)
	if err != nil {
		return fmt.Errorf("failed to read map data: %v", err)
	}
	locations := locationResponse{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return fmt.Errorf("failed to unmarshal map data: %v", err)
	}

	//update next and previous in config
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	//print the location
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.ExploreLocation == "" {
		fmt.Println("No location selected. Please provide a location to explore.")
		return nil
	}
	if val, exists := cfg.Cache.Get(url + cfg.ExploreLocation); exists {
		// Use cached data
		pokemon := pokemonResponse{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached pokemon data: %v", err)
		}
	} else {
		data, err := http.Get(url + cfg.ExploreLocation)
		if err != nil {
			return fmt.Errorf("failed to find pokemon data: %v", err)
		}
		defer data.Body.Close()
		if data.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to find pokemon data: received status code %d", data.StatusCode)
		}
		dat, err := io.ReadAll(data.Body)
		if err != nil {
			return fmt.Errorf("failed to read pokemon data: %v", err)
		}
		pokemon := pokemonResponse{}
		err = json.Unmarshal(dat, &pokemon)
		if err != nil {
			return fmt.Errorf("failed to unmarshal pokemon data: %v", err)
		}
		fmt.Printf("Exploring %s... \n"+"Found Pokemon: \n", cfg.ExploreLocation)
		for _, encounter := range pokemon.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
	}
	return nil
}
