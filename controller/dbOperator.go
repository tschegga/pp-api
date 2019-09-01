package controller

import (
	"database/sql"
	"pp-api/data"
)

// getRanking return the current leaderboard.
func getRanking() ([]data.Rank, error) {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "SELECT users.name, COUNT(sessions.athlete) " +
		"FROM sessions,users WHERE sessions.athlete = users.users_id " +
		"GROUP BY users.name ORDER BY COUNT(sessions.athlete) DESC"

	ranking := []data.Rank{}

	rankingError := db.Select(&ranking, query)
	if rankingError != nil {
		return nil, rankingError
	}

	rTx.Commit()

	return ranking, nil
}

// getSessions returns all sessions for the given user.
func getSessions(user int) ([]data.Session, error) {
	db := GetConnection()

	sTx := db.MustBegin()

	query := "SELECT sessions_id, start, length, quality FROM sessions WHERE athlete = ?"

	sessions := []data.Session{}

	sessionsError := db.Select(&sessions, query, user)
	if sessionsError != nil {
		return nil, sessionsError
	}

	sTx.Commit()

	return sessions, nil
}

// addSession adds a session for given user with attributes start, length and quality.
func addSession(userID int, start string, length int, quality int) error {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "INSERT INTO `sessions`(`user`, `start`, `length`, `quality`) VALUES (?, ?, ?, ?)"

	_, sessionErr := db.Exec(query, userID, start, length, quality)
	if sessionErr != nil {
		return sessionErr
	}

	rTx.Commit()

	return nil
}

// deleteSessions deletes a distinct session.
func deleteSession(sessionID int) error {
	db := GetConnection()

	rTx := db.MustBegin()

	// Delete all sessions connected to the user
	query := "DELETE FROM `sessions` WHERE idsessions = ?"
	_, userError := db.Exec(query, sessionID)
	if userError != nil {
		return userError
	}

	rTx.Commit()

	return nil
}

// doesUserExist checks if the given user is present in the database
func doesUserExist(username string) (bool, error) {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "SELECT name FROM users WHERE name = ?"

	var user struct {
		Username string `db:"name"`
	}

	userError := db.Get(&user, query, username)
	if userError != nil {
		var userExists bool

		// user does not exist
		if userError == sql.ErrNoRows {
			userExists = false
			return userExists, nil
		}

		// database or other error
		return userExists, userError
	}

	rTx.Commit()

	// user does exist
	return true, nil
}

// getUser queries the database with the given username.
func getUser(username string) (data.User, error) {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "SELECT users_id,name FROM users WHERE name = ?"

	user := data.User{}

	userError := db.Get(&user, query, username)
	if userError != nil {
		return user, userError
	}

	rTx.Commit()

	return user, nil
}

// addUser adds a user with the given username and password to the database.
func addUser(username string, password string) error {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "INSERT INTO `users`(`name`, `password`, `role`) VALUES (?, ?, ?)"

	// 2 refers to the standard role "athlete" for every new user
	_, userError := db.Exec(query, username, password, 2)
	if userError != nil {
		return userError
	}

	rTx.Commit()

	return nil
}

// deleteUserAndSessions deletes all sessions for a given user and after that the user itself.
func deleteUserAndSessions(username string) error {
	db := GetConnection()

	rTx := db.MustBegin()

	// Delete all sessions connected to the user
	query := "DELETE FROM sessions WHERE athlete = (SELECT users_id FROM users WHERE name = ?)"
	_, userError := db.Exec(query, username)
	if userError != nil {
		return userError
	}

	// Delete user from users table
	query = "DELETE FROM users WHERE name = ?"
	_, userError = db.Exec(query, username)
	if userError != nil {
		return userError
	}

	rTx.Commit()

	return nil
}
