package services

import (
	"PokeGo/models"
	"encoding/json"
	"os"
)

func LoadPlayerPokemonList(filename string) (*models.PlayerPokemonList, error) {
	var list models.PlayerPokemonList
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &list, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &list)
	return &list, err
}

func SavePlayerPokemonList(filename string, list *models.PlayerPokemonList) error {
	file, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}
