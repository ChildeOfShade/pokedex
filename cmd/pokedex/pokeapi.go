package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Config struct for pagination
type Config struct {
	Next     *string
	Previous *string
}

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
func commandMap(cfg *Config) error {
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
func commandMapBack(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
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
