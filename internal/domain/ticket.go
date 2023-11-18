package domain

import (
	"time"
)

type Ticket struct {
	Id          int       `db:"ticket_id"`
	Board       int       `db:"board_id"`
	Name        string    `db:"name"`
	State       string    `db:"state"`
	Description string    `db:"description"`
	AuthorId    int       `db:"author_id"`
	AssigneeId  int       `db:"assignee_id"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
	Closed      time.Time `db:"closed"`
	Comments    []Comment `db:"comments"`
	Author      *User
	Assignee    *User
}

func NewTicket() *Ticket {
	return &Ticket{
		Author:   NewUser(),
		Assignee: NewUser(),
	}
}
