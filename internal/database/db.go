package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Database *sqlx.DB
}

type DBSettings struct {
	Address  string
	Port     string
	User     string
	Password string
}

var ddb *DB = nil

func New(s *DBSettings) (*DB, error) {

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/siiliboard?sslmode=disable",
		s.User, s.Password, s.Address, s.Port,
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

func GetDatabase() (*DB, error) {

	if ddb == nil {
		log.Fatalln("Database not initialized")
	}

	// sanity check, remove later on
	err := ddb.Database.Ping()

	if err != nil {
		log.Println("Unable to reach the database")
		return nil, errors.New("Database unavailable")
	}

	return ddb, nil
}

func (*DB) Close() error {

	if ddb == nil {
		return nil
	}

	err := ddb.Database.Close()
	if err != nil {
		return err
	}

	return nil
}
