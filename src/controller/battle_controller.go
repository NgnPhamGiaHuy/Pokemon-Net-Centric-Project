package controller

import (
	"fmt"
	"math/rand"
	"pokecat_pokebat/internal/model"
	"pokecat_pokebat/internal/service"
)

var elementalMultiplier = map[string]float64{
	// Define elemental multipliers for each type
	"Normal": 1.0,
	"Fire":   1.5,
	"Water":  0.8,
	// Add more types as needed
}

type BattleController struct {
	BattleService *service.BattleService
}

func NewBattleController() *BattleController {
	return &BattleController{
		BattleService: service.NewBattleService(),
	}
}

func (bc *BattleController) StartBattle(player1, player2 *model.Player) *model.Player {
	// Set active pokemons
	player1.ActivePokemon = player1.Pokemons[0]
	player2.ActivePokemon = player2.Pokemons[0]

	// Determine who goes first based on speed
	firstPlayer, secondPlayer := determineFirstPlayer(player1, player2)

	// Main battle loop
	for {
		// First player's turn
		bc.performTurn(firstPlayer, secondPlayer)
		if bc.isBattleOver(firstPlayer, secondPlayer) {
			return firstPlayer
		}

		// Second player's turn
		bc.performTurn(secondPlayer, firstPlayer)
		if bc.isBattleOver(secondPlayer, firstPlayer) {
			return secondPlayer
		}
	}
}

func determineFirstPlayer(player1, player2 *model.Player) (firstPlayer, secondPlayer *model.Player) {
	if player1.ActivePokemon.Speed > player2.ActivePokemon.Speed {
		return player1, player2
	} else if player2.ActivePokemon.Speed > player1.ActivePokemon.Speed {
		return player2, player1
	} else {
		// If speeds are equal, randomly choose first player
		if rand.Intn(2) == 0 {
			return player1, player2
		}
		return player2, player1
	}
}

func (bc *BattleController) performTurn(attacker, defender *model.Player) {
	// Switch pokemon if active pokemon fainted
	if attacker.ActivePokemon.HP <= 0 {
		attacker.ActivePokemon = bc.chooseActivePokemon(attacker)
		return // End turn after switching
	}

	// Check for status conditions (e.g., Poison, Burn) and apply damage
	// Implement forced switching mechanics if applicable
	// Introduce surrender check

	// Perform attack
	bc.performAttack(attacker, defender)
}

func (bc *BattleController) performAttack(attacker, defender *model.Player) {
	// Randomly choose between normal attack and special attack
	isSpecialAttack := rand.Intn(2) == 0

	var damage int
	if isSpecialAttack {
		damage = calculateSpecialDamage(attacker.ActivePokemon, defender.ActivePokemon)
	} else {
		damage = calculateNormalDamage(attacker.ActivePokemon, defender.ActivePokemon)
	}

	// Ensure damage is positive
	if damage < 0 {
		damage = 0
	}

	// Apply damage to the defender's active pokemon
	defender.ActivePokemon.HP -= damage

	// Log the attack details
	fmt.Printf("%s's %s attacks %s's %s for %d damage!\n",
		attacker.Name, attacker.ActivePokemon.Name,
		defender.Name, defender.ActivePokemon.Name, damage)
}

func calculateNormalDamage(attacker, defender *model.CapturedPokemon) int {
	return attacker.Attack - defender.Defense
}

func calculateSpecialDamage(attacker, defender *model.CapturedPokemon) int {
	// Apply elemental multiplier
	elementalMultiplier := elementalMultiplier[attacker.Type]
	damage := int(float64(attacker.SpAttack)*elementalMultiplier - float64(defender.SpDefense))
	if damage < 0 {
		damage = 0
	}
	return damage
}

func (bc *BattleController) chooseActivePokemon(player *model.Player) *model.CapturedPokemon {
	// Logic to choose the next active pokemon, for simplicity, let's choose the next available pokemon
	for _, pokemon := range player.Pokemons {
		if pokemon.HP > 0 {
			return pokemon
		}
	}
	// If all pokemons are fainted, return nil
	return nil
}

func (bc *BattleController) isBattleOver(player, opponent *model.Player) bool {
	// Check if all opponent's pokemons are fainted
	for _, pokemon := range opponent.Pokemons {
		if pokemon.HP > 0 {
			return false // Battle is not over
		}
	}

	// Check if player has surrendered
	if player.Surrendered {
		return true // Battle is over
	}

	return true // All opponent's pokemons are fainted
}

func (bc *BattleController) LogBattle(player1, player2, winner *model.Player) {
	fmt.Printf("Battle between %s and %s\n", player1.Name, player2.Name)
	fmt.Printf("Winner: %s\n", winner.Name)
	// Add more detailed battle logging here
}
