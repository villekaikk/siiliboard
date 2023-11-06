package domain

import (
	"time"
)

type Comment struct {
	Content string		`db:"content"`
	Author User			`db:"author"`
	Created time.Time	`db:"created"`
}