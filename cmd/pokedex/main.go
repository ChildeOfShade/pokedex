package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "pokedex/internal/pokecache"
    "time"
)

type Config struct {
    cache    *pokecache.Cache
    Next     *string
    Previous *string
}

func main() {
    cfg := &Config{
        cache: pokecache.NewCache(5 * time.Second),
    }

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
            err := cmd.callback(cfg, words[1:])
            if err != nil {
                fmt.Println("Error:", err)
            }
        } else {
            fmt.Println("Unknown command")
        }
    }
}