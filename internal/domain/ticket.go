package domain

import (
	"time"
)

type Ticket struct {
	Name string			`db:"name"`
	Description string	`db:"description"`
	Author User			`db:"author"`
	Assignee User		`db:"assignee"`
	Created time.Time	`db:"created"`
	Updated time.Time	`db:"updated"`
	Closed time.Time	`db:"closed"`
	Comments []Comment	`db:"comments"`
}