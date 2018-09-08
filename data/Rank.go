package data

type Rank struct {
	Name string `db:"name" json:"name"`
	Count int `db:"COUNT(sessions.user)" json:"count"`
}
