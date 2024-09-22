package main

import "net/http"

func (cfg *apiConfig) handleMakeUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	var params parameters
	err := getParameters(&params, r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to read parameters")
		return
	}

	user, err := cfg.client.MakeUser(params.Name)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	RespondWithJSON(w, http.StatusOK, user)
}
