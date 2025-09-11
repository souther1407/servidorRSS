package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/souther1407/servidorRSS/auth"
	"github.com/souther1407/servidorRSS/internal/database"
)

type NewUserParams struct {
	Name string `name`
}

func (dbConfig *ApiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := NewUserParams{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error al procesar el body %v", err))
		return
	}

	newUser, err := dbConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error al craer el usuario: %v", err))
		return
	}

	responseWithJSON(w, 200, User(newUser))

}

func (dbConfig *ApiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)

	if err != nil {
		responseWithError(w, 401, err.Error())
		return
	}

	user, err := dbConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error al obtener un usuario %v", err))
		return
	}
	responseWithJSON(w, 200, User(user))
}
