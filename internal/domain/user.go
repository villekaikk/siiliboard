package domain

import (
	"time"
)

type User struct {
	Name string			`db:"name"`
	Created time.Time	`db:"created"`
}