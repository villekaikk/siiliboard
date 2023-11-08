package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"siiliboard/internal/database"
	"siiliboard/internal/routes"
)

var  default_port string = "8080"

func main() {

	setupLogger()
	err := godotenv.Load(".env")
	if err != nil {
        log.Fatal("Error loading .env file")
    }

	port, succ := os.LookupEnv("PORT_SERVE")

	if !succ {
		log.Println("PORT_SERVE not specified, using default port instead")
		port = default_port
	}
	s, err := resolveDBSettings()

	if err != nil {
		log.Fatalf("Unable to resolve database settings - %s\n", err.Error())
	}

    db, err := database.Connect(s)

	if err != nil {
		log.Fatalf("Error opening database connection - %v\n", err.Error())
	}

	defer db.Close()
	database.CreateTables()

	router := routes.NewRouter()
	port = fmt.Sprintf(":%v", port)

	log.Printf("Starting the server at http:127.0.0.1%v\n", port)

    log.Fatal(http.ListenAndServe(port, router))
}

func setupLogger() {
	
	log.SetFlags(log.Lshortfile)
	datetime := "2006-01-02 15:04:05: "
	log.SetPrefix(time.Now().Format(datetime))
}

func resolveDBSettings() (*database.DBSettings, error) {
	
	pq_user, succ := os.LookupEnv("PQ_USER")
	if !succ || len(strings.TrimSpace(pq_user)) == 0 {
		return nil, errors.New("PQ_USER not defined")
	}

	pq_passwd, succ := os.LookupEnv("PQ_PASSWD")
	if !succ || len(strings.TrimSpace(pq_passwd)) == 0 {
		return nil, errors.New("PQ_PASSWD not defined")
	}

	pq_db, succ := os.LookupEnv("PQ_DB")
	if !succ || len(strings.TrimSpace(pq_db)) == 0 {
		return nil, errors.New("PQ_DB not defined")
	}

	pq_addr, succ := os.LookupEnv("PQ_ADDR")
	if !succ || len(strings.TrimSpace(pq_addr)) == 0 {
		return nil, errors.New("PQ_ADDR not defined")
	}

	pq_port, succ := os.LookupEnv("PQ_PORT")
	if !succ || len(strings.TrimSpace(pq_port)) == 0 {
		return nil, errors.New("PQ_PORT not defined")
	}

	return &database.DBSettings{
		Address: pq_addr,
		Port: pq_port,
		User: pq_user,
		Password: pq_passwd,
		Database: pq_db,
		}, nil

}