package database

import (
	"errors"
	"log"
	"siiliboard/internal/domain"
	"siiliboard/internal/marshal"
)

func CreateUser(ur *marshal.UserRequest) (*domain.User, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	u := &domain.User{}
	// TODO: Insert succeeds but RETURNING statement fails, causing an error 500
	query := `INSERT INTO app_user (name, display_name) VALUES ($1, $2) RETURNING *`
	err = db.Database.QueryRow(query, ur.Name, ur.DisplayName).Scan(u)

	if err != nil {
		log.Printf("Unable to create new user: %s\n", err.Error())
		return nil, err
	}

	return u, nil
}

func GetUser(user_id int) (*domain.User, error) {

	db, err := GetDatabase()

	if err != nil {
		return nil, err
	}

	u := &domain.User{}
	q := "SELECT * FROM app_user WHERE user_id=($1)"
	err = db.Database.Get(u, q, user_id)

	if err != nil {
		log.Printf("Error querying user from the database: %s\n", err.Error())
		return nil, errors.New("Unable to query users")
	}

	return u, nil
}
