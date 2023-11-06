package domain

import (
	"time"
)

type Board struct {
	Name string			`db:"name"`
	Members []User 		`db:"members"`
	Tickets []Ticket 	`db:"tickets"`
	Created time.Time 	`db:"created"`
}