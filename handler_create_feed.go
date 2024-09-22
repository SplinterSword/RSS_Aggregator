package main

import (
	"net/http"

	"github.com/SplinterSword/RSS_Aggregator/internal/database"
)

func (cfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	var params parameters
	err := getParameters(&params, r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to read parameters")
		return
	}

	feed, err := cfg.client.CreateFeed(user, params.Name, params.URL)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	RespondWithJSON(w, http.StatusOK, feed)
}
