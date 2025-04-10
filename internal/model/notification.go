package model

import "time"

type Notification struct {
	Id        int64      `db:"id"`
	Text      string     `db:"notification"`
	IsRead    bool       `db:"is_read"`
	ReadAt    *time.Time `db:"read_at"`
	CreatedAt time.Time  `db:"created_at"`
}
