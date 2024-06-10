package service

import (
	"encoding/json"
	"os"
	"pokecat_pokebat/internal/model"
)

func LoadPlayerPokemonList(filename string) (*model.PlayerPokemonList, error) {
	var list model.PlayerPokemonList
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

func SavePlayerPokemonList(filename string, list *model.PlayerPokemonList) error {
	file, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}
