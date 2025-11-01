package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationAreaResponse maps the API response
type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

// fetchLocationAreas fetches a page from PokeAPI
func fetchLocationAreas(url string) (*LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data LocationAreaResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// commandMap shows the next 20 locations
func commandMap(cfg *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	data, err := fetchLocationAreas(url)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil
}

// commandMapBack shows the previous 20 locations
func commandMapBack(cfg *Config, args []string) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	data, err := fetchLocationAreas(*cfg.Previous)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil
}

// commandExplore shows all Pokémon in a location area
func commandExplore(cfg *Config, args []string) error {
	if len(args) < 1 {
		fmt.Println("usage: explore <location-area-name>")
		return nil
	}

	location := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	data, err := getFromAPI(cfg, url) // <-- uses cache if available
	if err != nil {
		return err
	}

	var resp struct {
		PokemonEncounters []struct {
			Pokemon struct {
				Name string `json:"name"`
			} `json:"pokemon"`
		} `json:"pokemon_encounters"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}

	fmt.Printf("Pokémon in %s:\n", location)
	for _, p := range resp.PokemonEncounters {
		fmt.Println(" -", p.Pokemon.Name)
	}

	return nil
}

// getFromAPI fetches data from API or cache
func getFromAPI(cfg *Config, url string) ([]byte, error) {
	// check cache
	if data, ok := cfg.cache.Get(url); ok {
		return data, nil
	}

	// get from API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// store in cache
	cfg.cache.Add(url, body)
	return body, nil
}
