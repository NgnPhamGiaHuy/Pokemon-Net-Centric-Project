package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"pokecat_pokebat/controller"
	"pokecat_pokebat/internal/model"
	"pokecat_pokebat/internal/service"
	"strings"
	"time"
)

func loadPokedexData(filename string) ([]model.Pokemon, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var pokedexData []model.Pokemon
	err = json.Unmarshal(data, &pokedexData)
	if err != nil {
		return nil, err
	}

	return pokedexData, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	pokedexFile := "data/pokedex.json"
	pokedexData, err := loadPokedexData(pokedexFile)
	if err != nil {
		fmt.Printf("Failed to load Pokedex data: %v\n", err)
		return
	}

	worldID := rand.Intn(1000)
	worldController := controller.NewWorldController(worldID)

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
		return
	}
	playerPokemonDataFile := filepath.Join(workingDir, "data", "player_pokemon_list.json")
	playerService := service.NewPlayerService()
	playerPokemonMap := playerService.LoadPlayerList(playerPokemonDataFile)

	for playerName, pokemons := range playerPokemonMap {
		player := playerService.CreatePlayer(playerName, model.Position{X: rand.Intn(1000), Y: rand.Intn(1000)})
		player.Pokemons = pokemons
		worldController.AddPlayer(player)
	}

	go worldController.SpawnPokemons(pokedexData, 50)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter player name and direction (e.g., 'Ash up') or 'auto' to auto-move: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		if len(parts) == 1 && parts[0] == "auto" {
			fmt.Print("Enter player name for auto-move: ")
			playerName, _ := reader.ReadString('\n')
			playerName = strings.TrimSpace(playerName)

			var player *model.Player
			for _, p := range playerService.Players {
				if p.Name == playerName {
					player = p
					break
				}
			}

			if player == nil {
				fmt.Printf("Player %s not found.\n", playerName)
				continue
			}

			fmt.Println("Starting automatic movement...")
			pokemon := playerService.AutoMovePlayer(player, worldController.WorldService)
			if pokemon != nil {
				fmt.Printf("Player %s found a Pokémon: %s. Do you want to catch it? (yes/no): ", player.Name, pokemon.Name)
				answer, _ := reader.ReadString('\n')
				answer = strings.TrimSpace(answer)
				if strings.ToLower(answer) == "yes" {
					playerService.CatchPokemon(player, pokemon)
					fmt.Printf("Player %s caught the Pokémon: %s\n", player.Name, pokemon.Name)
				} else {
					fmt.Printf("Player %s did not catch the Pokémon: %s\n", player.Name, pokemon.Name)
				}
			}
		} else if len(parts) == 2 {
			playerName := parts[0]
			direction := parts[1]

			var player *model.Player
			for _, p := range playerService.Players {
				if p.Name == playerName {
					player = p
					break
				}
			}

			if player == nil {
				fmt.Printf("Player %s not found.\n", playerName)
				continue
			}

			if direction == "up" || direction == "down" || direction == "left" || direction == "right" {
				playerService.MovePlayer(player, direction, worldController.WorldService)
				fmt.Printf("Player %s moved %s to position %+v\n", player.Name, direction, player.Position)
			} else {
				fmt.Println("Invalid direction. Please enter 'up', 'down', 'left', or 'right'.")
			}
		} else {
			fmt.Println("Invalid input. Please enter in the format 'PlayerName Direction' or 'auto'.")
		}
	}
}
