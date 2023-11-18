package domain

import (
	"time"
)

type Comment struct {
	Id      int       `db:"comment_id"`
	Content string    `db:"content"`
	Author  *User     `db:"author"`
	Created time.Time `db:"created"`
}

func NewComment() *Comment {
	return &Comment{
		Author: NewUser(),
	}
}
