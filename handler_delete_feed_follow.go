package main

import (
	"net/http"

	"github.com/SplinterSword/RSS_Aggregator/internal/database"
)

func (cfg *apiConfig) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowId := r.PathValue("feedFollowID")
	if feedFollowId == "" {
		RespondWithError(w, http.StatusBadRequest, "Missing feedFollowID")
		return
	}

	err := cfg.client.DeleteFeedFollows(user, feedFollowId)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "OK")
}
