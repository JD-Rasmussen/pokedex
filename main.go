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
			err := commands[cleanInput[0]].callback()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
			//fmt.Print("Your command was: " + cleanInput[0] + "\n")

		}
	}

}
