package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SplinterSword/RSS_Aggregator/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	client database.DBConfig
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	db := database.CreateConnection()

	cfg := apiConfig{
		client: db,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", handleReadiness)
	mux.HandleFunc("GET /v1/err", handleError)

	mux.HandleFunc("POST /v1/users", cfg.handleMakeUser)
	mux.Handle("GET /v1/users", cfg.middlewareAuth(cfg.handleGetUserByApiKey))

	mux.Handle("POST /v1/feeds", cfg.middlewareAuth(cfg.handleCreateFeed))
	mux.HandleFunc("GET /v1/feeds", cfg.handleGetAllFeeds)

	mux.Handle("POST /v1/feed_follows", cfg.middlewareAuth(cfg.handleCreateFeedFollows))
	mux.Handle("DELETE /v1/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.handleDeleteFeedFollows))
	mux.HandleFunc("GET /v1/feed_follows", cfg.handleGetAllFeedFollows)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server listening on http://localhost:%v", port)
	srv.ListenAndServe()

}
