package services

import "PokeGo/models"

func LevelUpPokemon(pokemon *models.CapturedPokemon) {
	requiredExp := 2 * pokemon.Level * pokemon.Level
	for pokemon.AccumulatedExp >= requiredExp {
		pokemon.Level++
		requiredExp = 2 * pokemon.Level * pokemon.Level

		pokemon.HP = int(float64(pokemon.HP) * (1 + pokemon.EV))
		pokemon.Attack = int(float64(pokemon.Attack) * (1 + pokemon.EV))
		pokemon.Defense = int(float64(pokemon.Defense) * (1 + pokemon.EV))
		pokemon.SpAttack = int(float64(pokemon.SpAttack) * (1 + pokemon.EV))
		pokemon.SpDefense = int(float64(pokemon.SpDefense) * (1 + pokemon.EV))
		pokemon.Speed = int(float64(pokemon.Speed) * (1 + pokemon.EV))
	}
}

func CapturePokemon(pokedex *models.Pokemon, level int, ev float64) models.CapturedPokemon {
	return models.CapturedPokemon{
		No:             pokedex.No,
		Image:          pokedex.Image,
		Name:           pokedex.Name,
		Type:           "Unknown",
		Level:          level,
		AccumulatedExp: 0,
		EV:             ev,
		HP:             pokedex.HP,
		Attack:         pokedex.Attack,
		Defense:        pokedex.Defense,
		SpAttack:       pokedex.SpAttack,
		SpDefense:      pokedex.SpDefense,
		Speed:          pokedex.Speed,
		TotalEvs:       pokedex.TotalEvs,
	}
}
