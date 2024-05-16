package services

import (
	"encoding/json"
)

// Pokemon represents a pokemon
type Pokemon struct {
	No         int    `json:"no"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Exp        int    `json:"exp"`
	HP         int    `json:"hp"`
	Attack     int    `json:"attack"`
	Defense    int    `json:"defense"`
	SpAttack   int    `json:"sp_attack"`
	SpDefense  int    `json:"sp_defense"`
	Speed      int    `json:"speed"`
	TotalEVs   int    `json:"total_evs"`
	CurrentEVs int    `json:"current_evs"`
	Captured   bool   `json:"captured"`
}

// NewPokemon creates a new pokemon
func NewPokemon(no int, name string, image string) Pokemon {
	return Pokemon{
		No:         no,
		Image:      image,
		Name:       name,
		Exp:        0,
		HP:         0,
		Attack:     0,
		Defense:    0,
		SpAttack:   0,
		SpDefense:  0,
		Speed:      0,
		TotalEVs:   0,
		CurrentEVs: 0,
		Captured:   false,
	}
}

// Spawn creates a new pokemon from a JSON string
func Spawn(jsonData string) (*Pokemon, error) {
	// Parse the JSON data
	var pokemon Pokemon
	err := json.Unmarshal([]byte(jsonData), &pokemon)
	if err != nil {
		return nil, err
	}

	// Create a new pokemon with the loaded data
	newPokemon := NewPokemon(pokemon.No, pokemon.Name, pokemon.Image)

	return &newPokemon, nil
}
