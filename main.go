package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()
	cfg := &Config{}

	for {
		//	scanner.Scan()

		fmt.Print("pokedex > ")
		for scanner.Scan() {
			input := scanner.Text()

			if input == "" {
				fmt.Print("pokedex > ")
				continue
			}
			cleanInput := cleanInput(input)
			cmd, ok := commands[cleanInput[0]]
			if !ok {
				fmt.Print("Your command was: " + cleanInput[0] + "\n")
				continue
			}
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		}
	}

}
