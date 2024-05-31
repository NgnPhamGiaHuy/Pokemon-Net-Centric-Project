package controllers

import (
	"encoding/json"
	"net/http"
	"pokebat/internal/services"
	"pokebat/models"
)

func PlayerInteractionController(w http.ResponseWriter, r *http.Request) {
	var interactionRequest models.InteractionRequest

	if err := json.NewDecoder(r.Body).Decode(&interactionRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := services.HandlePlayerInteraction(interactionRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
