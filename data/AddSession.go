package data

type AddSession struct {
	UserID  int    `db:"user" json:"userid"`
	Start   string `db:"start" json:"start"`
	Length  int    `db:"length" json:"length"`
	Quality int    `db:"quality" json:"quality"`
}
