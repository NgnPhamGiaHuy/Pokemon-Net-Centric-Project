package main

import (
	"PokeGo/models"
	"PokeGo/services"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if _, err := os.Stat("data/pokedex.json"); os.IsNotExist(err) {
		fmt.Println("pokedex.json not found, scraping data...")
		err := services.ScrapePokedex()
		if err != nil {
			fmt.Println("Error scraping pokedex data:", err)
			return
		}
	}

	pokedex, err := services.LoadPokedex("data/pokedex.json")
	if err != nil {
		fmt.Println("Error loading pokedex:", err)
		return
	}

	playerPokemons, err := services.LoadPlayerPokemonList("data/player_pokemons.json")
	if err != nil {
		fmt.Println("Error loading player's pokemon list:", err)
		playerPokemons = &models.PlayerPokemonList{}
	}

	ev := 0.5 + rand.Float64()*0.5
	if len(*pokedex) > 0 {
		newPokemon := services.CapturePokemon(&(*pokedex)[0], 1, ev)
		playerPokemons.Pokemons = append(playerPokemons.Pokemons, newPokemon)

		err = services.SavePlayerPokemonList("data/player_pokemons.json", playerPokemons)
		if err != nil {
			fmt.Println("Error saving player's pokemon list:", err)
		}
	} else {
		fmt.Println("Pokedex is empty.")
	}

	if len(playerPokemons.Pokemons) > 0 {
		playerPokemons.Pokemons[0].AccumulatedExp += 50
		services.LevelUpPokemon(&playerPokemons.Pokemons[0])

		err = services.SavePlayerPokemonList("data/player_pokemons.json", playerPokemons)
		if err != nil {
			fmt.Println("Error saving player's pokemon list:", err)
		} else {
			fmt.Println("Player's pokemon list updated and saved successfully.")
		}
	}
}
