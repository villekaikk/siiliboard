package domain

type BoardMember struct {
	BoardId int `db:"board_id"`
	UserId  int `db:"user_id"`
}
