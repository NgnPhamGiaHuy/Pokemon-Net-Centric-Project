package models

type Pokemon struct {
	No        int    `json:"no"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Exp       int    `json:"exp"`
	HP        int    `json:"hp"`
	Attack    int    `json:"attack"`
	Defense   int    `json:"defense"`
	SpAttack  int    `json:"sp_attack"`
	SpDefense int    `json:"sp_defense"`
	Speed     int    `json:"speed"`
	TotalEvs  int    `json:"total_evs"`
	Element   string `json:"element"` // Assuming each Pokemon has a primary element
}

type Player struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Pokemons []Pokemon `json:"pokemons"`
	Active   int       `json:"active"`
}

type Battle struct {
	Player1 Player `json:"player1"`
	Player2 Player `json:"player2"`
	Turn    int    `json:"turn"`
}

type BattleRequest struct {
	Player1 Player `json:"player1"`
	Player2 Player `json:"player2"`
}

type InteractionRequest struct {
	Battle           Battle `json:"battle"`
	PlayerID         int    `json:"player_id"`
	Action           string `json:"action"`
	NewActivePokemon int    `json:"new_active_pokemon,omitempty"`
}
