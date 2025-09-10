package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func configRouter(r *chi.Mux) {

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Get("/status", okHandler)
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	router := chi.NewRouter()
	configRouter(router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Escuchando en el puerto %s ", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
