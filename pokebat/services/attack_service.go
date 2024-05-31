package services

import (
	"errors"
	"pokebat/models"
)

func HandlePlayerInteraction(request models.InteractionRequest) (models.Battle, error) {
	battle := request.Battle

	attacker, defender := GetActivePokemons(battle)

	switch request.Action {
	case "attack":
		damage := PerformAttack(attacker, defender)
		defender.HP -= damage
		if defender.HP <= 0 {
			HandlePokemonFaint(&battle, request.PlayerID)
		}
		// Switch turns
		battle.Turn = 3 - battle.Turn
	case "switch":
		if request.PlayerID == 1 {
			battle.Player1.Active = request.NewActivePokemon
		} else {
			battle.Player2.Active = request.NewActivePokemon
		}
		// Turn ends after switching
		battle.Turn = 3 - battle.Turn
	case "surrender":
		HandleSurrender(&battle, request.PlayerID)
	default:
		return battle, errors.New("invalid action")
	}

	return battle, nil
}

func GetActivePokemons(battle models.Battle) (*models.Pokemon, *models.Pokemon) {
	var attacker, defender *models.Pokemon

	if battle.Turn == 1 {
		attacker = &battle.Player1.Pokemons[battle.Player1.Active]
		defender = &battle.Player2.Pokemons[battle.Player2.Active]
	} else {
		attacker = &battle.Player2.Pokemons[battle.Player2.Active]
		defender = &battle.Player1.Pokemons[battle.Player1.Active]
	}

	return attacker, defender
}

func HandlePokemonFaint(battle *models.Battle, playerID int) {
	if playerID == 1 {
		battle.Player1.Active++
		if battle.Player1.Active >= len(battle.Player1.Pokemons) {
			HandleSurrender(battle, playerID)
		}
	} else {
		battle.Player2.Active++
		if battle.Player2.Active >= len(battle.Player2.Pokemons) {
			HandleSurrender(battle, playerID)
		}
	}
}

func HandleSurrender(battle *models.Battle, playerID int) {
	if playerID == 1 {
		// Player 2 wins
		distributeExp(&battle.Player2, battle.Player1)
	} else {
		// Player 1 wins
		distributeExp(&battle.Player1, battle.Player2)
	}
}

func distributeExp(winner *models.Player, loser models.Player) {
	totalExp := 0
	for _, p := range loser.Pokemons {
		totalExp += p.Exp
	}
	expShare := totalExp / 3
	for i := range winner.Pokemons {
		winner.Pokemons[i].Exp += expShare
	}
}
