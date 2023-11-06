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

var ddb *DB = nil

func Connect() (*DB, error) {

	dbAddr := "localhost"
	dbPort := 5432
	dbName := "siiliboard"
	connStr := fmt.Sprintf("postgres://postgres:admin@%s:%d/%s?sslmode=disable", dbAddr, dbPort, dbName)

	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}
	
	log.Printf("PostgreSQL connection established to %s:%d\n", dbAddr, dbPort)
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