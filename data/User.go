package data

// User represents a user and all his competition related attributes stored in the backend.
type User struct {
	Username string    `db:"name" json:"username"`
	UserID   int       `db:"users_id" json:"userid"`
	Rank     int       `json:"rank"`
	Sessions []Session `json:"sessions"`
}
