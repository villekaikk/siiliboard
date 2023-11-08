package database

import (
	"log"

	"siiliboard/internal/domain"
	"siiliboard/internal/marshal"
)

func GetBoards() ([]domain.Board, error) {

	_, err := GetDatabase()

	if err != nil {
		log.Printf("Unable to create new board: %s", err.Error())
		return nil, err
	}

	return make([]domain.Board, 0), nil
}

func CreateBoard(br *marshal.BoardRequest) (int, error) {

	db, err := GetDatabase()

	if err != nil {
		log.Printf("Unable to create new board: %s", err.Error())
		return -1, err
	}

	query := `INSERT INTO board (name) VALUES ($1) RETURNING board_id`
	var key int
	err = db.database.QueryRow(query, br.Name).Scan(&key)

	if err != nil {
		log.Printf("Unable to create new board: %s", err.Error())
		return -1, err
	}

	return key, nil
}
