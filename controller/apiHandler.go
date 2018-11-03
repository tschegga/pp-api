package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pp-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRequests(r *mux.Router) {

	// Define all api endpoints here
	r.HandleFunc("/", statusHandler)
	r.HandleFunc("/v1/ranking", rankingHandler)
	r.HandleFunc("/v1/sessions", sessionsHandler)
	r.HandleFunc("/v1/users", usersHandler)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET status")

		// Return status
		jsonResponse, jsonError := json.Marshal(data.Status{"pp-api", "0.1.0"})
		if jsonError != nil {
			log.Println(jsonError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		log.Println("GET ranking")

		// Get ranking from database
		result, err := getRanking()
		if err != nil {
			log.Println(err)
			return
		}

		// Return result
		log.Printf("Sending result:%+v", result)
		jsonResponse, jsonError := json.Marshal(result)
		if jsonError != nil {
			log.Println(jsonError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	default:
		log.Println("Method not allowed, only GET is allowed")
	}
}

func sessionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		log.Println("GET sessions")

		// Parse the URL parameter
		q := r.URL.Query()
		userIdString := q.Get("userid")
		if userIdString == "" {
			log.Println("No user was given as parameter")
			return
		}
		userId, intErr := strconv.Atoi(userIdString)
		if intErr != nil {
			log.Printf("Error parsing string to int:%s", intErr)
			return
		}

		// Request the database
		result, err := getSessions(userId)
		if err != nil {
			log.Println(err)
			return
		}

		// Return result
		log.Printf("Sending result:%+v", result)
		jsonResponse, jsonError := json.Marshal(result)
		if jsonError != nil {
			log.Println(jsonError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	case "PUT":
		log.Println("PUT session")
	case "DELETE":
		log.Println("DELETE session")
	default:
		log.Println("Method not allowed, only methods GET, PUT, DELETE allowed")
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("GET users")

		// Check if http headers for username and password are set
		if r.Header.Get("username") == "" || r.Header.Get("password") == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Headers 'username' or 'password' not set")
			return
		}

		// Check if user and password is correct
		var username = r.Header.Get("username")
		validUser, validUserErr := isUserValid(username, r.Header.Get("password"))
		if validUserErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(validUserErr)
			return
		}

		if validUser {
			// Get user object
			user, getUserErr := getUser(username)
			if getUserErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(getUserErr)
				return
			}

			// parse user object
			jsonResponse, jsonErr := json.Marshal(user)
			if jsonErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(jsonErr)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		} else {
			// user and/or password was not correct
			w.WriteHeader(http.StatusForbidden)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
