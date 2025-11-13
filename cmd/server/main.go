package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"

	"github.com/hoainam183/todo-app/internal/config"
	"github.com/hoainam183/todo-app/internal/handlers/rest"
)

func main() {
	cfg := config.LoadPort()
	port := cfg.Port

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	v1Router.Get("/health", rest.HandlerReadiness)
	v1Router.Get("/err", rest.HandlerError)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("http.server:  listening at %s", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}
