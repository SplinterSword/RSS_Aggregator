package main

import (
	"net/http"

	"github.com/SplinterSword/RSS_Aggregator/internal/database"
)

func (cfg *apiConfig) handleGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, http.StatusOK, user)
}
