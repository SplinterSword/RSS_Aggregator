package main

import "net/http"

func (cfg *apiConfig) handleGetAllFeedFollows(w http.ResponseWriter, r *http.Request) {
	feedFollows, err := cfg.client.GetAllFeedFollows()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, feedFollows)
}
