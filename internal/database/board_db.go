package database

import(
	"log"
)

/*
func GetBoards() []domain.Board {


}
*/
func CreateTables() {
	log.Println("Creating database tables if needed")
	createUserTable()
	createCommentTable()
	createTicketTable()
	createBoardTable()
}

func createUserTable() {
	
	db := GetDatabase()

	query := `CREATE TABLE IF NOT EXISTS app_user (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createCommentTable() {

	db := GetDatabase()

	query := `CREATE TABLE IF NOT EXISTS comment (
		id SERIAL PRIMARY KEY,
		content VARCHAR(1024) NOT NULL,
		author app_user NOT NULL,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createTicketTable() {

	db := GetDatabase()

	query := `CREATE TABLE IF NOT EXISTS ticket (
		id SERIAL PRIMARY KEY,
		name VARCHAR(256) NOT NULL,
		description VARCHAR(2048) NOT NULL,
		author app_user NOT NULL,
		assignee app_user,
		comments comment[],
		created timestamp DEFAULT NOW(),
		updated timestamp DEFAULT NOW(),
		closed timestamp DEFAULT '1900-01-01 00:00:00'
	)`

	_, err := db.database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createBoardTable() {

	db := GetDatabase()

	query := `CREATE TABLE IF NOT EXISTS board (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		members app_user[],
		tickets TICKET[],
		created timestamp DEFAULT NOW()
	)`

	_, err := db.database.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}