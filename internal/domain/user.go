package domain

import (
	"time"
)

type User struct {
	Id          int       `db:"user_id"`
	Name        string    `db:"name"`
	DisplayName string    `db:"display_name"`
	Created     time.Time `db:"created"`
}
