package domain

import (
	"time"
)

type Board struct {
	Id      int       `db:"board_id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created"`
}
