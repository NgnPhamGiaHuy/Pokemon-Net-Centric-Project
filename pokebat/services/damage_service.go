package services

import (
	"math/rand"
	"pokebat/models"
)

func NormalAttack(attacker, defender *models.Pokemon) int {
	damage := attacker.Attack - defender.Defense
	if damage < 0 {
		damage = 0
	}
	return damage
}

func SpecialAttack(attacker, defender *models.Pokemon, elementalMultiplier float64) int {
	damage := int(float64(attacker.SpAttack)*elementalMultiplier) - defender.SpDefense
	if damage < 0 {
		damage = 0
	}
	return damage
}

func CalculateElementalMultiplier(attackerElement, defenderElement string) float64 {
	// Implement elemental effectiveness logic here
	// Placeholder for example
	return 1.0
}

func PerformAttack(attacker, defender *models.Pokemon) int {
	if rand.Intn(2) == 0 {
		return NormalAttack(attacker, defender)
	} else {
		elementalMultiplier := CalculateElementalMultiplier(attacker.Element, defender.Element)
		return SpecialAttack(attacker, defender, elementalMultiplier)
	}
}
