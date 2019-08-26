package controller

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type handler func(w http.ResponseWriter, r *http.Request)

// Decodes HTTP basic auth headers for the given request and calls validateUser for authorization check.
func validateBasicAuth(pass handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validateUser(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		pass(w, r)
	}
}

// TODO: Obviously needs to be implemented with real users
// Checks if the given username and password belong to an existing user.
func validateUser(username, password string) bool {
	if username == "test" && password == "test" {
		return true
	}
	return false
}
