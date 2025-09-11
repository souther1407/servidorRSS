package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/souther1407/servidorRSS/internal/database"
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
	r.Post("/user", apiConfig.handlerCreateUser)
	r.Get("/user", apiConfig.handlerGetUserByAPIKey)
}

type ApiConfig struct {
	DB *database.Queries
}

var apiConfig ApiConfig

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("La variable PORT no existe o no esta seteada")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalln("La variable DB_URL no existe o no esta seteada")
	}
	router := chi.NewRouter()
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln("Error al conectarse con la base de datos ", err)
	}

	apiConfig = ApiConfig{
		DB: database.New(dbConn),
	}

	configRouter(router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Escuchando en el puerto %s ", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
