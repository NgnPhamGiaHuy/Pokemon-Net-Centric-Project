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
}
