package main

import (
	"encoding/json"
	"net/http"
)

func getParameters(params any, r *http.Request) error {
	Decoder := json.NewDecoder(r.Body)
	err := Decoder.Decode(&params)
	if err != nil {
		return err
	}
	return nil
}
