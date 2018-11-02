package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRequests(r *mux.Router) {

	// Define all api endpoints here
	r.HandleFunc("/", statusHandler)
	r.HandleFunc("/v1/ranking", rankingHandler)
	r.HandleFunc("/v1/sessions", sessionsHandler)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET status")
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
