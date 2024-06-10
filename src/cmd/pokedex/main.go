package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"pokecat_pokebat/internal/model"
	"pokecat_pokebat/internal/service"
	"strconv"
	"strings"
	"time"
)

const minPokemonPerUser = 3

var realWorldNames = []string{"Alice", "Bob", "Charlie", "David", "Emma", "Frank", "Grace", "Henry", "Ivy", "Jack"}

func main() {
	rand.Seed(time.Now().UnixNano())

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}

	if _, err := os.Stat("data/pokedex.json"); os.IsNotExist(err) {
		fmt.Println("pokedex.json not found, scraping data...")
		err := service.ScrapePokedex()
		if err != nil {
			fmt.Println("Error scraping pokedex data:", err)
			return
		}
	}

	pokedex, err := service.LoadPokedex("data/pokedex.json")
	if err != nil {
		fmt.Println("Error loading pokedex:", err)
		return
	}

	fmt.Println("Choose generation method:")
	fmt.Println("1. Random")
	fmt.Println("2. Manual Input")
	fmt.Print("Enter your choice: ")
	reader := bufio.NewReader(os.Stdin)
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Invalid choice. Exiting...")
		return
	}

	var playerPokemonLists map[string]*model.PlayerPokemonList

	switch choice {
	case 1:
		playerPokemonLists = generateRandomPlayerPokemonLists(pokedex)
	case 2:
		playerPokemonLists = generateManualPlayerPokemonLists(pokedex)
	}

	err = savePlayerPokemonLists("data/player_pokemon_list.json", playerPokemonLists)
	if err != nil {
		fmt.Println("Error saving player's pokemon lists:", err)
	} else {
		fmt.Println("Player's pokemon lists saved successfully.")
	}
}

func generateRandomPlayerPokemonLists(pokedex *[]model.Pokemon) map[string]*model.PlayerPokemonList {
	playerPokemonLists := make(map[string]*model.PlayerPokemonList)

	fmt.Print("Enter the number of players to generate: ")
	reader := bufio.NewReader(os.Stdin)
	numPlayersStr, _ := reader.ReadString('\n')
	numPlayersStr = strings.TrimSpace(numPlayersStr)
	numPlayers, err := strconv.Atoi(numPlayersStr)
	if err != nil || numPlayers <= 0 {
		fmt.Println("Invalid number of players. Exiting...")
		os.Exit(1)
	}

	fmt.Print("Enter the minimum number of pokemons per player: ")
	minPokemonsStr, _ := reader.ReadString('\n')
	minPokemonsStr = strings.TrimSpace(minPokemonsStr)
	minPokemons, err := strconv.Atoi(minPokemonsStr)
	if err != nil || minPokemons <= 0 {
		fmt.Println("Invalid minimum number of pokemons. Exiting...")
		os.Exit(1)
	}

	fmt.Print("Enter the maximum number of pokemons per player: ")
	maxPokemonsStr, _ := reader.ReadString('\n')
	maxPokemonsStr = strings.TrimSpace(maxPokemonsStr)
	maxPokemons, err := strconv.Atoi(maxPokemonsStr)
	if err != nil || maxPokemons <= 0 || maxPokemons < minPokemons {
		fmt.Println("Invalid maximum number of pokemons. Exiting...")
		os.Exit(1)
	}

	for i := 1; i <= numPlayers; i++ {
		playerName := getRandomRealWorldName()
		numPokemons := rand.Intn(maxPokemons-minPokemons+1) + minPokemons
		playerPokemonLists[playerName] = generatePokemonList(pokedex, numPokemons)
	}

	return playerPokemonLists
}

func generateManualPlayerPokemonLists(pokedex *[]model.Pokemon) map[string]*model.PlayerPokemonList {
	playerPokemonLists := make(map[string]*model.PlayerPokemonList)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of players: ")
	numPlayersStr, _ := reader.ReadString('\n')
	numPlayersStr = strings.TrimSpace(numPlayersStr)
	numPlayers, err := strconv.Atoi(numPlayersStr)
	if err != nil || numPlayers <= 0 {
		fmt.Println("Invalid number of players. Exiting...")
		os.Exit(1)
	}

	for i := 1; i <= numPlayers; i++ {
		fmt.Printf("Enter name for player %d: ", i)
		playerName, _ := reader.ReadString('\n')
		playerName = strings.TrimSpace(playerName)

		numPokemons := minPokemonPerUser
		playerPokemonLists[playerName] = generatePokemonList(pokedex, numPokemons)
	}

	return playerPokemonLists
}

func generatePokemonList(pokedex *[]model.Pokemon, numPokemon int) *model.PlayerPokemonList {
	playerPokemons := model.PlayerPokemonList{} // Changed from pointer to value

	for i := 0; i < numPokemon; i++ {
		ev := 0.5 + rand.Float64()*0.5
		if len(*pokedex) > 0 {
			newPokemon := service.CapturePokemon(&(*pokedex)[rand.Intn(len(*pokedex))], 1, ev)
			playerPokemons = append(playerPokemons, newPokemon) // Append directly to the slice
		}
	}

	return &playerPokemons // Return pointer to the slice
}

func savePlayerPokemonLists(filename string, playerPokemonLists map[string]*model.PlayerPokemonList) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encodedPlayerPokemonLists := make(map[string][]model.CapturedPokemon)

	for playerName, pokemonList := range playerPokemonLists {
		encodedPlayerPokemonLists[playerName] = *pokemonList
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(encodedPlayerPokemonLists)
	if err != nil {
		return err
	}

	return nil
}

func getRandomRealWorldName() string {
	return realWorldNames[rand.Intn(len(realWorldNames))]
}
