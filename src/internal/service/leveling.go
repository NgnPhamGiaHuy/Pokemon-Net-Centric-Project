package service

import (
	"pokecat_pokebat/internal/model"
)

func CapturePokemon(pokedex *model.Pokemon, level int, evs string) model.CapturedPokemon {
	return model.CapturedPokemon{
		No:        pokedex.No,
		Image:     pokedex.Image,
		Name:      pokedex.Name,
		Exp:       0,
		Type:      pokedex.Type,
		Level:     level,
		EVs:       evs,
		HP:        pokedex.HP,
		Attack:    pokedex.Attack,
		Defense:   pokedex.Defense,
		SpAttack:  pokedex.SpAttack,
		SpDefense: pokedex.SpDefense,
		Speed:     pokedex.Speed,
		TotalEvs:  pokedex.TotalEvs,
	}
}
