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
