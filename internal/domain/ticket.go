package domain

import (
	"time"
)

type Ticket struct {
	Id          int       `db:"ticket_id"`
	Name        string    `db:"name"`
	State       string    `db:"state"`
	Description string    `db:"description"`
	Author      User      `db:"author"`
	Assignee    User      `db:"assignee"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
	Closed      time.Time `db:"closed"`
	Comments    []Comment `db:"comments"`
}
