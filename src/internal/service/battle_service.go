package service

import (
	"pokecat_pokebat/internal/model"
)

type BattleService struct{}

func NewBattleService() *BattleService {
	return &BattleService{}
}

func (bs *BattleService) StartBattle(player1 *model.Player, pokemons1 []*model.CapturedPokemon, player2 *model.Player, pokemons2 []*model.CapturedPokemon) *model.Player {
	if len(pokemons1) == 0 || len(pokemons2) == 0 {
		return nil
	}
	pokemon1 := pokemons1[0]
	pokemon2 := pokemons2[0]

	if pokemon1.Attack > pokemon2.Defense {
		return player1
	}
	return player2
}
