package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the headers of an HTTP request.
// Authorization: ApiKey {insert key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of Authorization header")
	}
	return vals[1], nil
}
