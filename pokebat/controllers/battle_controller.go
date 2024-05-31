package controllers

import (
	"encoding/json"
	"net/http"
	"pokebat/internal/services"
	"pokebat/models"
)

func BattleStartController(w http.ResponseWriter, r *http.Request) {
	var battleRequest models.BattleRequest

	if err := json.NewDecoder(r.Body).Decode(&battleRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	battle := services.StartBattle(battleRequest.Player1, battleRequest.Player2)
	json.NewEncoder(w).Encode(battle)
}
