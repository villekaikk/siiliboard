package database

import (
	"database/sql"
	"log"
	"siiliboard/internal/domain"
)

func CreateTables() {
	log.Println("Creating database tables if needed")
	createUserTable()
	createBoardTable()
	createBoardMemberTable()
	createTicketTable()
	createCommentTable()

	createDefaultUsers()
}

func createUserTable() {

	db, err := GetDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	query := `CREATE TABLE IF NOT EXISTS app_user (
		user_id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		display_name VARCHAR(100) NOT NULL,
		created timestamp DEFAULT NOW()
	)`

	_, err = db.Database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createBoardTable() {

	db, err := GetDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	query := `CREATE TABLE IF NOT EXISTS board (
		board_id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		created timestamp DEFAULT NOW()
	)`

	_, err = db.Database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createBoardMemberTable() {

	db, err := GetDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	query := `CREATE TABLE IF NOT EXISTS board_member (
		board_id INT REFERENCES board(board_id) ON DELETE CASCADE ON UPDATE CASCADE,
		user_id INT REFERENCES app_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE
	)`

	_, err = db.Database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createCommentTable() {

	db, err := GetDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	query := `CREATE TABLE IF NOT EXISTS comment (
		comment_id SERIAL PRIMARY KEY,
		content VARCHAR(1024) NOT NULL,
		author_id INT REFERENCES app_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
		ticket_id INT REFERENCES ticket(ticket_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL, 
		created timestamp DEFAULT NOW()
	)`

	_, err = db.Database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createTicketTable() {

	db, err := GetDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	query := `CREATE TABLE IF NOT EXISTS ticket (
		ticket_id SERIAL PRIMARY KEY,
		name VARCHAR(256) NOT NULL,
		state VARCHAR(16) NOT NULL DEFAULT 'todo',
		description VARCHAR(2048) NOT NULL,
		author_id INT REFERENCES app_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
		assignee_id INT REFERENCES app_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
		board_id INT REFERENCES board(board_id) ON DELETE CASCADE ON UPDATE CASCADE,
		created timestamp DEFAULT NOW(),
		updated timestamp DEFAULT NOW(),
		closed timestamp DEFAULT '1900-01-01 00:00:00'
	)`

	_, err = db.Database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createDefaultUsers() {
	db, err := GetDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	// Create 'unassigned' user if one does not already exist
	query := "SELECT * FROM app_user WHERE user_id = 0"

	u := &domain.User{}
	err = db.Database.Get(u, query)

	// User found, bail out
	if err == nil {
		return
	}

	if err != sql.ErrNoRows {
		log.Fatalf("Unable to insert default user 'unassigned' - %s", err.Error())
	}

	insert := `INSERT INTO app_user (user_id, name, display_name) VALUES ($1, $2, $3)`

	_, err = db.Database.Exec(insert, 0, "unassigned", "unassigned")

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Inserted user 'unassigned' to the database")
}
