package data

// Rank is the representation of one athlete in the leaderboard.
// It consists of the athlete's name and the rank.
type Rank struct {
	Name  string `db:"name" json:"name"`
	Count int    `db:"COUNT(sessions.athlete)" json:"count"`
}
