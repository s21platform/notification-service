package model

type Notification struct {
	Id     int64  `db:"id"`
	Text   string `db:"notification"`
	IsRead bool   `db:"is_read"`
}
