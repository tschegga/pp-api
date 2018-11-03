package controller

import (
	"pp-api/data"
)

func getRanking() ([]data.Rank, error) {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "SELECT users.name, COUNT(sessions.user) " +
		"FROM sessions,users WHERE sessions.user = users.idusers " +
		"GROUP BY users.name ORDER BY COUNT(sessions.user) DESC"

	ranking := []data.Rank{}

	rankingError := db.Select(&ranking, query)
	if rankingError != nil {
		return nil, rankingError
	}

	rTx.Commit()

	return ranking, nil
}

func getSessions(user int) ([]data.Session, error) {
	db := GetConnection()

	sTx := db.MustBegin()

	query := "SELECT start, length, quality FROM sessions WHERE user = ?"

	sessions := []data.Session{}

	sessionsError := db.Select(&sessions, query, user)
	if sessionsError != nil {
		return nil, sessionsError
	}

	sTx.Commit()

	return sessions, nil
}

func isUserValid(username string, password string) (bool, error) {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "SELECT idusers,name FROM users WHERE name = ? AND password = ?"

	user := data.User{}

	userError := db.Get(&user, query, username, password)
	if userError != nil {
		// TODO: check on sql error empty result set
		return false, userError
	}

	rTx.Commit()

	// if user or password is not correct function will terminate above
	return true, nil
}

func getUser(username string) (data.User, error) {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "SELECT idusers,name FROM users WHERE name = ?"

	user := data.User{}

	userError := db.Get(&user, query, username)
	if userError != nil {
		return user, userError
	}

	rTx.Commit()

	return user, nil
}

func addUser(username string, password string) error {
	db := GetConnection()

	rTx := db.MustBegin()

	query := "INSERT INTO `users`(`name`, `password`) VALUES (?, ?)"

	_, userError := db.Exec(query, username, password)
	if userError != nil {
		return userError
	}

	rTx.Commit()

	return nil
}

func deleteUserAndSessions(username string) error {
	db := GetConnection()

	rTx := db.MustBegin()

	// Delete all sessions connected to the user
	query := "DELETE FROM `sessions` WHERE user = (SELECT idusers FROM users WHERE name = ?)"
	_, userError := db.Exec(query, username)
	if userError != nil {
		return userError
	}

	// Delete user from users table
	query = "DELETE FROM `users` WHERE name = ?"
	_, userError = db.Exec(query, username)
	if userError != nil {
		return userError
	}

	rTx.Commit()

	return nil
}
