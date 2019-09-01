package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"pp-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// HandleRequests provides all api endpoints on one place
func HandleRequests(r *mux.Router) {

	// Define all api endpoints here
	r.HandleFunc("/", validateBasicAuth(statusHandler))
	r.HandleFunc("/v1/ranking", validateBasicAuth(rankingHandler))
	r.HandleFunc("/v1/sessions", validateBasicAuth(sessionsHandler))
	r.HandleFunc("/v1/users", validateBasicAuth(usersHandler))
	r.HandleFunc("/v1/users/{username}", validateBasicAuth(usersParameterHandler))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET status")

		// Return status
		jsonResponse, jsonError := json.Marshal(data.Status{Name: "pp-api", Version: "0.1.0"})
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
		userIDString := q.Get("userid")
		if userIDString == "" {
			log.Println("No user was given as parameter")
			return
		}
		userID, intErr := strconv.Atoi(userIDString)
		if intErr != nil {
			log.Printf("Error parsing string to int:%s", intErr)
			return
		}

		// Request the database
		result, err := getSessions(userID)
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

		// Read body
		body, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if readErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(readErr)
			return
		}

		// Unmarshal
		var session data.AddSession
		marshallErr := json.Unmarshal(body, &session)
		if marshallErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(marshallErr)
			return
		}

		addSessionErr := addSession(session.UserID, session.Start, session.Length, session.Quality)
		if addSessionErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(addSessionErr)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "DELETE":
		log.Println("DELETE session")

		// Parse the URL parameter
		q := r.URL.Query()
		sessionIDString := q.Get("sessionid")
		if sessionIDString == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("No session was given as parameter")
			return
		}

		// Convert URL parameter to int
		sessionID, intErr := strconv.Atoi(sessionIDString)
		if intErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error parsing string to int:%s", intErr)
			return
		}

		// Delete session
		deleteSessionErr := deleteSession(sessionID)
		if deleteSessionErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(deleteSessionErr)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log.Println("POST users")

		//
		// Check if user already exists
		//

		// Read body
		body, readErr := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if readErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(readErr)
			return
		}

		// Unmarshal
		var user struct {
			Name     string `json:"username"`
			Password string `json:"password"`
		}
		marshallErr := json.Unmarshal(body, &user)
		if marshallErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(marshallErr)
			return
		}

		// Check in database
		userExists, userExistsErr := doesUserExist(user.Name)
		if userExistsErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(userExistsErr)
			return
		}
		if userExists {
			io.WriteString(w, "user already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		// Create user in database
		usersErr := addUser(user.Name, user.Password)
		if usersErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(usersErr)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func usersParameterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("GET users")

		// Check if user exists
		var username = mux.Vars(r)["username"]
		userExists, userExistsErr := doesUserExist(username)
		if userExistsErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(userExistsErr)
			return
		}

		if userExists {
			// Get user object
			user, getUserErr := getUser(username)
			if getUserErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(getUserErr)
				return
			}

			// Calculate current rank of user
			ranking, rankingErr := getRanking()
			if rankingErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(rankingErr)
				return
			}
			for index, element := range ranking {
				if element.Name == user.Username {
					user.Rank = index + 1
				}
			}

			// Get sessions for user
			var userSessions []data.Session
			var sessionsErr error
			userSessions, sessionsErr = getSessions(user.UserID)
			if sessionsErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(sessionsErr)
				return
			}
			user.Sessions = userSessions

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
			// user does not exist - return 404
			w.WriteHeader(http.StatusNotFound)
		}
	case "DELETE":
		log.Println("DELETE users")

		// Check if user exists
		var username = mux.Vars(r)["username"]
		userExists, userExistsErr := doesUserExist(username)
		if userExistsErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(userExistsErr)
			return
		}

		if userExists {
			deleteUserErr := deleteUserAndSessions(username)
			if deleteUserErr != nil {
				log.Println(deleteUserErr)
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.WriteHeader(http.StatusOK)
		} else {
			// user does not exist - return 404
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
