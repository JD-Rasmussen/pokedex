package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//	scanner.Scan()

		fmt.Print("pokedex > ")
		for scanner.Scan() {
			input := scanner.Text()
			if input == "quit" {
				fmt.Println("Exiting...")
				os.Exit(0)
			}
			if input == "" {
				fmt.Print("pokedex > ")
				continue
			}
			cleanInput := cleanInput(input)
			fmt.Print("Your command was: " + cleanInput[0] + "\n")

		}
	}

}
