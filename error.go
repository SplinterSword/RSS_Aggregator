package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
