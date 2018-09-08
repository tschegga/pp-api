package data

type Session struct {
	Start string `db:"start" json:"start"`
	Length int `db:"length" json:"length"`
	Quality int `db:"quality" json:"quality"`
}
