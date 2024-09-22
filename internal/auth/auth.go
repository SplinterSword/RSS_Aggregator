package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetKey(r *http.Header) (string, error) {
	authoriztionString := r.Get("Authorization")
	splitString := strings.Split(authoriztionString, " ")
	if len(splitString) != 2 || splitString[0] != "ApiKey" {
		return "", errors.New("invalid authorization header")
	}
	return splitString[1], nil
}
