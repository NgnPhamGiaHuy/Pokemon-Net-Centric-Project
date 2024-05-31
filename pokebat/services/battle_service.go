package services

import (
	"math/rand"
	"pokebat/models"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StartBattle(player1, player2 models.Player) models.Battle {
	return models.Battle{
		Player1: player1,
		Player2: player2,
		Turn:    DetermineFirstMove(player1, player2),
	}
}

func DetermineFirstMove(p1, p2 models.Player) int {
	p1Speed := p1.Pokemons[p1.Active].Speed
	p2Speed := p2.Pokemons[p2.Active].Speed

	if p1Speed > p2Speed {
		return 1
	} else if p1Speed < p2Speed {
		return 2
	} else {
		return rand.Intn(2) + 1
	}
}
