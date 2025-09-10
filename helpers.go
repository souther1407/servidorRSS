package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func responseWithError(w http.ResponseWriter, code int, message string) {

	if code >= 500 {
		log.Println("Error en el servidor codigo ", code, ": ", message)
	}

	responseWithJSON(w, code, ErrorResponse{
		Error: message,
		Code:  code,
	})

}

func responseWithJSON(w http.ResponseWriter, code int, payload any) {
	jsonData, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		log.Println("error al parsaer el el json: ", err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

}
