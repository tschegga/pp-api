package data

// AddSession is closely related to to Session struct is however needed for adding a session to the database.
type AddSession struct {
	UserID  int    `db:"user" json:"userid"`
	Start   string `db:"start" json:"start"`
	Length  int    `db:"length" json:"length"`
	Quality int    `db:"quality" json:"quality"`
}
