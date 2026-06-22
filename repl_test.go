package main

import (
	"testing"
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
	if len(commands) != 2 {
		t.Errorf("getCommands() returned %d commands, expected 2", len(commands))
	}
	if _, ok := commands["exit"]; !ok {
		t.Error("getCommands() did not return 'exit' command")
	}
	if _, ok := commands["help"]; !ok {
		t.Error("getCommands() did not return 'help' command")
	}
}

func TestCommandHelp(t *testing.T) {
	err := commandHelp()
	if err != nil {
		t.Errorf("commandHelp() returned an error: %v", err)
	}
}
