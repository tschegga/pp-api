package data

type User struct {
	Username string    `db:"name" json:"username"`
	UserID   int       `db:"idusers" json:"userid"`
	Rank     int       `json:"rank"`
	Sessions []Session `json:"sessions"`
}
