package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var playerName string

func main() {
	fmt.Println("Welcome to PokeBat!")
	login()
}

func login() {
	var name string
	fmt.Print("Enter your name to login: ")
	fmt.Scanln(&name)

	data := map[string]string{"name": name}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error logging in:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Login successful!")
		playerName = name
		showMainMenu()
	} else {
		fmt.Println("Failed to login:", resp.Status)
	}
}

func showMainMenu() {
	for {
		var choice int
		fmt.Println("1. Join PokeBat")
		fmt.Println("2. Surrender")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			joinPokeBat()
		case 2:
			surrender()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func joinPokeBat() {
	var opponentName string
	fmt.Print("Enter the name of your opponent: ")
	fmt.Scanln(&opponentName)

	playerPokemons := fetchPlayerPokemons(playerName)
	fmt.Println("Your Pokémon list:")
	for _, pokemon := range playerPokemons {
		fmt.Println(pokemon.Name)
	}

	var selectedPokemons []string
	for i := 0; i < 3; i++ {
		var pokemonName string
		fmt.Printf("Enter the name of Pokémon %d: ", i+1)
		fmt.Scanln(&pokemonName)
		selectedPokemons = append(selectedPokemons, pokemonName)
	}

	data := map[string]interface{}{
		"player1":          playerName,
		"player2":          opponentName,
		"selectedPokemons": selectedPokemons,
	}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post("http://localhost:8080/start-battle", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error starting battle:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result struct {
			Winner           string   `json:"winner"`
			Logs             []string `json:"logs"`
			OpponentPokemons []string `json:"opponent_pokemons"`
		}
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Printf("Battle finished! Winner: %s\n", result.Winner)
		fmt.Println("Battle Log:")
		for _, log := range result.Logs {
			fmt.Println(log)
		}
		fmt.Println("Opponent's Pokémon list:")
		for _, pokemon := range result.OpponentPokemons {
			fmt.Println(pokemon)
		}
	} else {
		fmt.Println("Failed to start battle:", resp.Status)
	}
}

func fetchPlayerPokemons(playerName string) []Pokemon {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/player-pokemons?name=%s", playerName))
	if err != nil {
		fmt.Println("Error fetching player pokemons:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var pokemons []Pokemon
		json.NewDecoder(resp.Body).Decode(&pokemons)
		return pokemons
	} else {
		fmt.Println("Failed to fetch player pokemons:", resp.Status)
		return nil
	}
}

func surrender() {
	data := map[string]string{"player": playerName}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post("http://localhost:8080/surrender", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error surrendering:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result struct {
			Winner string   `json:"winner"`
			Logs   []string `json:"logs"`
		}
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Printf("You surrendered! Winner: %s\n", result.Winner)
		fmt.Println("Battle Log:")
		for _, log := range result.Logs {
			fmt.Println(log)
		}
	} else {
		fmt.Println("Failed to surrender:", resp.Status)
	}
}

type Pokemon struct {
	Name string `json:"name"`
}
