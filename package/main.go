package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cfg := &Config{} // single config for pagination

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("Goodbye!")
			return
		}

		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if input == "" {
			continue
		}

		words := cleanInput(input)
		commandName := words[0]

		if cmd, ok := commands[commandName]; ok {
			cmd.callback(cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
