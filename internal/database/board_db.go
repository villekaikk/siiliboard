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
		log.Printf("Unable to create new board: %s\n", err.Error())
		return nil, err
	}

	boards := []domain.Board{}
	err = db.database.Select(&boards, "SELECT * FROM board")

	if err != nil {
		log.Printf("Error querying boards from the database: %s\n", err.Error())
		return nil, errors.New("")
	}

	return boards, nil
}

func CreateBoard(br *marshal.BoardRequest) (*domain.Board, error) {

	db, err := GetDatabase()

	if err != nil {
		log.Printf("Unable to create new board: %s", err.Error())
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
