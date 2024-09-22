package main

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	type ErrorStruct struct {
		Error string `json:"error"`
	}
	err := ErrorStruct{
		Error: message,
	}
	RespondWithJSON(w, code, err)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
