package model

type CapturedPokemon struct {
	No             int     `json:"no"`
	Image          string  `json:"image"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Level          int     `json:"level"`
	AccumulatedExp int     `json:"accumulated_exp"`
	EV             float64 `json:"ev"`
	HP             int     `json:"hp"`
	Attack         int     `json:"attack"`
	Defense        int     `json:"defense"`
	SpAttack       int     `json:"sp_attack"`
	SpDefense      int     `json:"sp_defense"`
	Speed          int     `json:"speed"`
	TotalEvs       int     `json:"total_evs"`
}
