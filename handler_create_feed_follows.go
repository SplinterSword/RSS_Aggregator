package main

import (
	"net/http"

	"github.com/SplinterSword/RSS_Aggregator/internal/database"
)

func (cfg *apiConfig) handleCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	params := parameters{}
	err := getParameters(&params, r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to read parameters")
		return
	}

	feedFollow, err := cfg.client.CreateFeedFollows(user, params.FeedID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, feedFollow)
}
