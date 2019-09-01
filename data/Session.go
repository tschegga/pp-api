package data

// Session represensts one session for a particular athlete.
type Session struct {
	SessionID int    `db:"sessions_id" json:"sessionid"`
	Start     string `db:"start" json:"start"`
	Length    int    `db:"length" json:"length"`
	Quality   int    `db:"quality" json:"quality"`
}
