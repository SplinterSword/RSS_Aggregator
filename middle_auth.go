package main

import (
	"net/http"

	"github.com/SplinterSword/RSS_Aggregator/internal/auth"
	"github.com/SplinterSword/RSS_Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetKey(&r.Header)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Couldn't get API key")
			return
		}

		user, err := cfg.client.GetUserByApiKey(api_key)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		handler(w, r, user)
	}
}
