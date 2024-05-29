package main

import (
	"log"
	"net/http"
	"pokebat/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/battle/start", controllers.BattleStartController).Methods("POST")
	router.HandleFunc("/battle/interaction", controllers.PlayerInteractionController).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
