package main

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	type Readiness struct {
		Status string `json:"status"`
	}

	readiness := Readiness{
		Status: "OK",
	}
	RespondWithJSON(w, http.StatusOK, readiness)
}
