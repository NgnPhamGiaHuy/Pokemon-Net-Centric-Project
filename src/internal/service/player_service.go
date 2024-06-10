package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"pokecat_pokebat/internal/model"
	"time"
)

type PlayerService struct {
	Players []*model.Player
}

func NewPlayerService() *PlayerService {
	return &PlayerService{Players: []*model.Player{}}
}

func (ps *PlayerService) CreatePlayer(name string, initialPosition model.Position) *model.Player {
	player := &model.Player{
		Name:     name,
		Pokemons: []*model.CapturedPokemon{},
		Position: initialPosition,
	}
	ps.Players = append(ps.Players, player)
	return player
}

func (ps *PlayerService) LoadCapturedPokemons(player *model.Player, playerPokemonDataFile string) {
	data, err := ioutil.ReadFile(playerPokemonDataFile)
	if err != nil {
		log.Fatalf("Failed to read player Pokemon data file: %v", err)
	}

	var playerPokemons []*model.CapturedPokemon

	err = json.Unmarshal(data, &playerPokemons)
	if err != nil {
		log.Fatalf("Failed to unmarshal player Pokemon data: %v", err)
	}

	player.Pokemons = append(player.Pokemons, playerPokemons...)
}

func (ps *PlayerService) CatchPokemon(player *model.Player, pokemon *model.Pokemon) {
	capturedPokemon := &model.CapturedPokemon{
		No:             pokemon.No,
		Image:          pokemon.Image,
		Name:           pokemon.Name,
		Type:           pokemon.Type,
		Level:          1,
		AccumulatedExp: 0,
		EV:             0,
		HP:             pokemon.HP,
		Attack:         pokemon.Attack,
		Defense:        pokemon.Defense,
		SpAttack:       pokemon.SpAttack,
		SpDefense:      pokemon.SpDefense,
		Speed:          pokemon.Speed,
		TotalEvs:       pokemon.TotalEvs,
	}
	player.Pokemons = append(player.Pokemons, capturedPokemon)

	ps.SavePlayerPokemons(player)
}

func (ps *PlayerService) SavePlayerPokemons(player *model.Player) {
	// Create a map to save all players' Pokémon
	playerPokemonMap := make(map[string][]*model.CapturedPokemon)
	for _, p := range ps.Players {
		playerPokemonMap[p.Name] = p.Pokemons
	}

	data, err := json.MarshalIndent(playerPokemonMap, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal player Pokémon data: %v", err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	playerPokemonDataFile := filepath.Join(workingDir, "data", "player_pokemon_list.json")
	err = ioutil.WriteFile(playerPokemonDataFile, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write player Pokémon data file: %v", err)
	}
}

func (ps *PlayerService) MovePlayer(player *model.Player, direction string, worldService *WorldService) {
	switch direction {
	case "up":
		ps.MoveUp(player)
	case "down":
		ps.MoveDown(player)
	case "left":
		ps.MoveLeft(player)
	case "right":
		ps.MoveRight(player)
	}

	pos := player.Position
	if worldService.HasPokemon(pos) {
		pokemon := worldService.CapturePokemon(pos)
		ps.CatchPokemon(player, pokemon)
	}
}

func (ps *PlayerService) AutoMovePlayer(player *model.Player, worldService *WorldService) *model.Pokemon {
	directions := []func(*model.Player){
		ps.MoveUp,
		ps.MoveDown,
		ps.MoveLeft,
		ps.MoveRight,
	}
	directionNames := []string{"up", "down", "left", "right"}

	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			moveIndex := rand.Intn(len(directions))
			move := directions[moveIndex]
			direction := directionNames[moveIndex]
			move(player)
			fmt.Printf("Player %s moved %s to position %+v\n", player.Name, direction, player.Position)
			pos := player.Position
			if worldService.HasPokemon(pos) {
				return worldService.CapturePokemon(pos)
			}
		}
	}
}

func (ps *PlayerService) LoadPlayerList(filename string) map[string][]*model.CapturedPokemon {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read player data file: %v", err)
	}

	playerPokemonMap := make(map[string][]*model.CapturedPokemon)
	err = json.Unmarshal(data, &playerPokemonMap)
	if err != nil {
		log.Fatalf("Failed to unmarshal player data: %v", err)
	}

	return playerPokemonMap
}

func (ps *PlayerService) MoveUp(player *model.Player) {
	player.Position.Y++
}

func (ps *PlayerService) MoveDown(player *model.Player) {
	player.Position.Y--
}

func (ps *PlayerService) MoveLeft(player *model.Player) {
	player.Position.X--
}

func (ps *PlayerService) MoveRight(player *model.Player) {
	player.Position.X++
}
