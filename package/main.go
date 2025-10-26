package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		// Get user input and clean it
		words := cleanInput(scanner.Text())

		// Ignore empty lines
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cmd, exists := commands[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		// Execute the command
		if err := cmd.callback(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
