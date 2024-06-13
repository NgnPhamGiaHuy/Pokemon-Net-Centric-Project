package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"pokecat_pokebat/controller"
	"pokecat_pokebat/internal/model"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//playerService := service.NewPlayerService()
	battleController := controller.NewBattleController()

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
		return
	}

	playerPokemonList, err := loadPlayerPokemonList(filepath.Join(workingDir, "data/player_pokemon_list.json"))
	if err != nil {
		fmt.Printf("Failed to load player pokemon list: %v\n", err)
		return
	}

	if len(playerPokemonList) < 2 {
		fmt.Println("Insufficient players to start a battle.")
		return
	}

	var eligiblePlayers []*model.Player
	for playerName, pokemons := range playerPokemonList {
		if len(pokemons) >= 3 {
			player := playerService.CreatePlayer(playerName, model.Position{X: 0, Y: 0})
			player.Pokemons = pokemons
			eligiblePlayers = append(eligiblePlayers, player)
		}
	}

	if len(eligiblePlayers) < 2 {
		fmt.Println("Insufficient players with at least 3 pokemons to start a battle.")
		return
	}

	player1 := eligiblePlayers[rand.Intn(len(eligiblePlayers))]
	var player2 *model.Player
	for {
		player2 = eligiblePlayers[rand.Intn(len(eligiblePlayers))]
		if player1 != player2 {
			break
		}
	}

	player1.Pokemons = choosePokemon(player1.Name, player1.Pokemons)
	player2.Pokemons = choosePokemon(player2.Name, player2.Pokemons)

	battleController.StartBattle(player1, player2)
}

func loadPlayerPokemonList(filePath string) (map[string][]*model.CapturedPokemon, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var playerPokemonList map[string][]*model.CapturedPokemon
	if err := json.Unmarshal(data, &playerPokemonList); err != nil {
		return nil, err
	}

	return playerPokemonList, nil
}

func choosePokemon(playerName string, pokemons []*model.CapturedPokemon) []*model.CapturedPokemon {
	fmt.Printf("%s, choose your 3 Pokémon by entering the numbers (separated by space):\n", playerName)
	for i, pokemon := range pokemons {
		fmt.Printf("%d: %s (HP: %d, Speed: %d)\n", i+1, pokemon.Name, pokemon.HP, pokemon.Speed)
	}

	var choices [3]int
	for {
		fmt.Print("Enter the numbers of your choices: ")
		_, err := fmt.Scanf("%d %d %d", &choices[0], &choices[1], &choices[2])
		if err == nil && isValidChoice(choices[:], len(pokemons)) {
			break
		}
		fmt.Println("Invalid choice. Please enter again.")
	}

	var selectedPokemons []*model.CapturedPokemon
	for _, choice := range choices {
		selectedPokemons = append(selectedPokemons, pokemons[choice-1])
	}

	fmt.Printf("%s has chosen:\n", playerName)
	for _, pokemon := range selectedPokemons {
		fmt.Printf("%s (HP: %d, Speed: %d)\n", pokemon.Name, pokemon.HP, pokemon.Speed)
	}

	return selectedPokemons
}

func isValidChoice(choices []int, max int) bool {
	seen := make(map[int]bool)
	for _, choice := range choices {
		if choice < 1 || choice > max || seen[choice] {
			return false
		}
		seen[choice] = true
	}
	return true
}
