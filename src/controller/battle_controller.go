package controller

import (
	"fmt"
	"math/rand"
	"pokecat_pokebat/internal/model"
	"pokecat_pokebat/internal/service"
)

var elementalMultiplier = map[string]float64{
	"Normal": 1.0,
	"Fire":   1.5,
	"Water":  0.8,
	// Add more types as needed
}

type BattleController struct {
	BattleService *service.BattleService
	Logs          []string
}

func NewBattleController() *BattleController {
	return &BattleController{
		BattleService: service.NewBattleService(),
		Logs:          []string{},
	}
}

func (bc *BattleController) StartBattle(player1, player2 *model.Player) *model.Player {
	bc.LogBattleStart(player1, player2)

	player1.ActivePokemon = player1.Pokemons[0]
	player2.ActivePokemon = player2.Pokemons[0]

	firstPlayer, secondPlayer := determineFirstPlayer(player1, player2)
	bc.LogTurnOrder(firstPlayer, secondPlayer)

	for {
		if firstPlayer.ActivePokemon != nil {
			bc.performTurn(firstPlayer, secondPlayer)
			if bc.isBattleOver(firstPlayer, secondPlayer) {
				bc.LogBattleEnd(firstPlayer, secondPlayer)
				return firstPlayer
			}
		}

		if secondPlayer.ActivePokemon != nil {
			bc.performTurn(secondPlayer, firstPlayer)
			if bc.isBattleOver(secondPlayer, firstPlayer) {
				bc.LogBattleEnd(secondPlayer, firstPlayer)
				return secondPlayer
			}
		}
	}
}

func determineFirstPlayer(player1, player2 *model.Player) (firstPlayer, secondPlayer *model.Player) {
	if player1.ActivePokemon.Speed > player2.ActivePokemon.Speed {
		return player1, player2
	} else if player2.ActivePokemon.Speed > player1.ActivePokemon.Speed {
		return player2, player1
	} else {
		if rand.Intn(2) == 0 {
			return player1, player2
		}
		return player2, player1
	}
}

func (bc *BattleController) performTurn(attacker, defender *model.Player) {
	if attacker.ActivePokemon.HP <= 0 {
		attacker.ActivePokemon = bc.chooseActivePokemon(attacker)
		if attacker.ActivePokemon != nil {
			bc.LogSwitch(attacker)
		}
		return
	}

	bc.performAttack(attacker, defender)

	if defender.ActivePokemon.HP <= 0 {
		bc.LogFaint(defender)
		if bc.isBattleOver(attacker, defender) {
			return
		}
		defender.ActivePokemon = bc.chooseActivePokemon(defender)
		if defender.ActivePokemon != nil {
			bc.LogSwitch(defender)
		}
	}
}

func (bc *BattleController) performAttack(attacker, defender *model.Player) {
	isSpecialAttack := rand.Intn(2) == 0
	var damage int
	var attackType string

	if isSpecialAttack {
		damage = calculateSpecialDamage(attacker.ActivePokemon, defender.ActivePokemon)
		attackType = "special attack"
	} else {
		damage = calculateNormalDamage(attacker.ActivePokemon, defender.ActivePokemon)
		attackType = "normal attack"
	}

	if damage < 0 {
		damage = 0
	}

	defender.ActivePokemon.HP -= damage

	bc.LogAttack(attacker, defender, damage, attackType)
}

func calculateNormalDamage(attacker, defender *model.CapturedPokemon) int {
	return attacker.Attack - defender.Defense
}

func calculateSpecialDamage(attacker, defender *model.CapturedPokemon) int {
	maxMultiplier := 1.0
	for _, atkType := range attacker.Type {
		if multiplier, exists := elementalMultiplier[atkType]; exists {
			if multiplier > maxMultiplier {
				maxMultiplier = multiplier
			}
		}
	}
	damage := int(float64(attacker.SpAttack)*maxMultiplier - float64(defender.SpDefense))
	if damage < 0 {
		damage = 0
	}
	return damage
}

func (bc *BattleController) chooseActivePokemon(player *model.Player) *model.CapturedPokemon {
	for _, pokemon := range player.Pokemons {
		if pokemon.HP > 0 {
			return pokemon
		}
	}
	return nil
}

func (bc *BattleController) isBattleOver(player, opponent *model.Player) bool {
	for _, pokemon := range opponent.Pokemons {
		if pokemon.HP > 0 {
			return false
		}
	}
	return true
}

func (bc *BattleController) LogBattleStart(player1, player2 *model.Player) {
	bc.Logs = append(bc.Logs, fmt.Sprintf("Battle started between %s and %s", player1.Name, player2.Name))
}

func (bc *BattleController) LogTurnOrder(firstPlayer, secondPlayer *model.Player) {
	bc.Logs = append(bc.Logs, fmt.Sprintf("%s's %s will go first", firstPlayer.Name, firstPlayer.ActivePokemon.Name))
}

func (bc *BattleController) LogSwitch(player *model.Player) {
	bc.Logs = append(bc.Logs, fmt.Sprintf("%s switched to %s", player.Name, player.ActivePokemon.Name))
}

func (bc *BattleController) LogAttack(attacker, defender *model.Player, damage int, attackType string) {
	bc.Logs = append(bc.Logs, fmt.Sprintf("%s's %s used a %s and dealt %d damage to %s's %s",
		attacker.Name, attacker.ActivePokemon.Name, attackType, damage, defender.Name, defender.ActivePokemon.Name))
}

func (bc *BattleController) LogFaint(player *model.Player) {
	bc.Logs = append(bc.Logs, fmt.Sprintf("%s's %s has fainted", player.Name, player.ActivePokemon.Name))
}

func (bc *BattleController) LogBattleEnd(winner, loser *model.Player) {
	bc.Logs = append(bc.Logs, fmt.Sprintf("Battle ended. %s won against %s", winner.Name, loser.Name))
	fmt.Println("Battle Log:")
	for _, log := range bc.Logs {
		fmt.Println(log)
	}
}
