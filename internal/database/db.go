package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type DB struct {
	database *sqlx.DB
}

type DBSettings struct {
	Address string
	Port string
	User string
	Password string
	Database string
}

var ddb *DB = nil

func Connect(s *DBSettings) (*DB, error) {

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.User, s.Password, s.Address, s.Port, s.Database,
	)

	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}
	
	log.Printf("PostgreSQL connection established to %s:%s\n", s.Address, s.Port)
	ddb = &DB{db}
	return ddb, nil
}

func GetDatabase() *DB {
	
	if ddb == nil {
		log.Fatalln("Database not initialized")
	}

	// sanity check, remove later on
	err := ddb.database.Ping()

	if err != nil {
		log.Fatal("Unable to reach the database")
	}

	return ddb
}

func (*DB) Close() error {

	if ddb == nil {
		return nil
	}

	err := ddb.database.Close()
	if err != nil {
		return err
	}

	return nil
}