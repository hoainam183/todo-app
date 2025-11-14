package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port != "" {
		fmt.Print(port)
	}

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	v1Router.Get("/heath", handlerReadiness)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("server: %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
