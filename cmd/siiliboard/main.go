package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"

	"siiliboard/internal/database"
	"siiliboard/internal/routes"
	"siiliboard/internal/utils"
)

var default_port string = "8080"
var DEBUG bool = false

func main() {

	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	logFile := setupLogger(rootPath)

	defer logFile.Close()

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	resolveEnv()

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

	router := routes.NewRouter(rootPath)
	port = fmt.Sprintf(":%v", port)

	log.Printf("Starting the server at http:127.0.0.1%v\n", port)

	log.Fatal(http.ListenAndServe(port, router))
}

func setupLogger(rootPath string) *os.File {

	log.SetFlags(log.Lshortfile)
	datetime := "2006-01-02 15:04:05: "
	log.SetPrefix(time.Now().Format(datetime))
	logFileName := time.Now().Format("2006-01-02.log")
	logDirPath := filepath.Join(rootPath, "log")
	logFileFullPath := filepath.Join(logDirPath, logFileName)

	_, err := os.Stat(logDirPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(logDirPath, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}

	f, err := os.OpenFile(logFileFullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	return f
}

func resolveDBSettings() (*database.DBSettings, error) {

	pq_user, succ := os.LookupEnv("PQ_USER")
	if !succ || utils.IsEmpty(pq_user) {
		return nil, errors.New("PQ_USER not defined")
	}

	pq_passwd, succ := os.LookupEnv("PQ_PASSWD")
	if !succ || utils.IsEmpty(pq_passwd) {
		return nil, errors.New("PQ_PASSWD not defined")
	}

	pq_db, succ := os.LookupEnv("PQ_DB")
	if !succ || utils.IsEmpty(pq_db) {
		return nil, errors.New("PQ_DB not defined")
	}

	pq_addr, succ := os.LookupEnv("PQ_ADDR")
	if !succ || utils.IsEmpty(pq_addr) {
		return nil, errors.New("PQ_ADDR not defined")
	}

	pq_port, succ := os.LookupEnv("PQ_PORT")
	if !succ || utils.IsEmpty(pq_port) {
		return nil, errors.New("PQ_PORT not defined")
	}

	return &database.DBSettings{
		Address:  pq_addr,
		Port:     pq_port,
		User:     pq_user,
		Password: pq_passwd,
		Database: pq_db,
	}, nil
}

func resolveEnv() {

	env, _ := os.LookupEnv("ENV")
	DEBUG = env == "DEBUG"
	routes.DEBUG = DEBUG

	if DEBUG {
		log.Println("Application running in DEBUG env")
	}
}
