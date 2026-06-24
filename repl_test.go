package main

import (
	"pokedex/internal"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander Bulbasaur PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d items, expected %d", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}
}

func TestGetCommands(t *testing.T) {
	commands := getCommands()
	if len(commands) != 4 {
		t.Errorf("getCommands() returned %d commands, expected 4", len(commands))
	}
	if _, ok := commands["exit"]; !ok {
		t.Error("getCommands() did not return 'exit' command")
	}
	if _, ok := commands["help"]; !ok {
		t.Error("getCommands() did not return 'help' command")
	}
	if _, ok := commands["map"]; !ok {
		t.Error("getCommands() did not return 'map' command")
	}
	if _, ok := commands["mapb"]; !ok {
		t.Error("getCommands() did not return 'mapb' command")
	}
}

func TestCommandHelp(t *testing.T) {
	err := commandHelp(&Config{})
	if err != nil {
		t.Errorf("commandHelp() returned an error: %v", err)
	}
}

func TestCommandMap(t *testing.T) {
	cfg := &Config{
		Cache: internal.NewCache(5 * time.Second),
	}
	err := commandMap(cfg)
	if err != nil {
		t.Errorf("commandMap() returned an error: %v", err)
	}
}

func TestCommandMapb(t *testing.T) {
	cfg := &Config{
		Cache: internal.NewCache(5 * time.Second),
	}
	err := commandMap(cfg)
	if err != nil {
		t.Errorf("commandMapback() returned an error: %v", err)
	}
}
