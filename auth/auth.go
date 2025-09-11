package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Obtiene la api key del header de la peticion
// ejemplo: Authorization: Apikey {api}
func GetAPIKey(headers http.Header) (apikey string, err error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("no hay api key")
	}

	values := strings.Split(value, " ")

	if len(values) != 2 || values[0] != "ApiKey" {
		return "", errors.New("api key mal formada")
	}
	return values[1], nil
}
