package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/hoainam183/todo-app/internal/database"
	"github.com/hoainam183/todo-app/internal/handler"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const (
	defaultPort = "8080"
)

func main() {
	godotenv.Load()

	// Initialize database connection
	dbConfig := database.NewConfig()
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	log.Info().Msg("connected to database successfully")

	// Store database instance in context if needed
	_ = db

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	v1Router.Get("/heath", handler.HandlerReadiness)
	v1Router.Get("/err", handler.HandlerError)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Info().Msgf("http.server: listening at %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Error().Err(err).Msg("server error")
	}
}
