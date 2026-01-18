package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hoainam183/todo-app/internal/config"
	"github.com/hoainam183/todo-app/internal/database"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error().Err(err).Msg("load env error")
	}
	port := cfg.Port

	dbConfig := cfg.DB
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	log.Info().Msg("connected to database successfully")

	// Store database instance in context if needed
	_ = db

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

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
