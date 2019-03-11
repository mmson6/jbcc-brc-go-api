package rmwebservice

import (
	"github.com/gorilla/mux"

	"github.com/jbcc/brc-api/internal/webmiddleware"
)

func AddRoutes(r *mux.Router) {
	// GET requests to the read model
	rGets := r.Methods("GET").Subrouter()
	rGets.Use(webmiddleware.ResponseExpiresImmediately)

	// Health GET requests
	rGets.HandleFunc("/health", GetHealth)
	rGets.HandleFunc("/records/{user_id}", GetUserRecords)
	rGets.HandleFunc("/leaderboard", GetLeaderboard)
}
