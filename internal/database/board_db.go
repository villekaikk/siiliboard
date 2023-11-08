package database

import (
	"errors"
	"log"

	"siiliboard/internal/domain"
	"siiliboard/internal/marshal"
)

func GetBoards() ([]domain.Board, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	boards := []domain.Board{}
	err = db.database.Select(&boards, "SELECT * FROM board")

	if err != nil {
		log.Printf("Error querying boards from the database: %s\n", err.Error())
		return nil, errors.New("Unable to query boards")
	}

	return boards, nil
}

func GetBoard(board_id int) (*domain.Board, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	b := &domain.Board{}
	q := "SELECT * FROM board WHERE board_id=($1)"
	err = db.database.Get(b, q, board_id)

	if err != nil {
		log.Printf("Error querying board from the database: %s\n", err.Error())
		return nil, errors.New("Unable to query boards")
	}

	return b, nil
}

func CreateBoard(br *marshal.BoardRequest) (*domain.Board, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	b := &domain.Board{}
	query := `INSERT INTO board (name) VALUES ($1) RETURNING *`
	err = db.database.QueryRow(query, br.Name).Scan(b)

	if err != nil {
		log.Printf("Unable to create new board: %s\n", err.Error())
		return nil, err
	}

	return b, nil
}
