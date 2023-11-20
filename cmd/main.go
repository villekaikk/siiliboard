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

	"siiliboard/internal/api"
	"siiliboard/internal/database"
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

	e, err := utils.FileExists(".env")
	if err != nil {
		log.Println("No .env file to load")
	}
	if e {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
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

	db, err := database.New(s)

	if err != nil {
		log.Fatalf("Error opening database connection - %v\n", err.Error())
	}

	defer db.Close()
	database.CreateTables()

	router := api.NewRouter(rootPath)
	port = fmt.Sprintf(":%v", port)

	log.Printf("Starting the server at http://127.0.0.1%v\n", port)

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
	if !succ || utils.IsStringEmpty(pq_user) {
		return nil, errors.New("PQ_USER not defined")
	}

	pq_passwd, succ := os.LookupEnv("PQ_PASSWD")
	if !succ || utils.IsStringEmpty(pq_passwd) {
		return nil, errors.New("PQ_PASSWD not defined")
	}

	// If given secret is a file, use its contents. Otherwise just use the secret as is
	secret, err := try_read_secret(pq_passwd)

	if err != nil {
		return nil, fmt.Errorf("Unexpected error reading PQ_PASSWD secret - %s", err.Error())
	}

	if !utils.IsStringEmpty(secret) {
		pq_passwd = secret
	}

	pq_addr, succ := os.LookupEnv("PQ_ADDR")
	if !succ || utils.IsStringEmpty(pq_addr) {
		return nil, errors.New("PQ_ADDR not defined")
	}

	pq_port, succ := os.LookupEnv("PQ_PORT")
	if !succ || utils.IsStringEmpty(pq_port) {
		return nil, errors.New("PQ_PORT not defined")
	}

	return &database.DBSettings{
		Address:  pq_addr,
		Port:     pq_port,
		User:     pq_user,
		Password: pq_passwd,
	}, nil
}

func resolveEnv() {

	env, _ := os.LookupEnv("ENV")
	DEBUG = env == "DEBUG"
	api.DEBUG = DEBUG

	if DEBUG {
		log.Println("Application running in DEBUG env")
	}
}

func try_read_secret(secret_path string) (string, error) {

	pq_abs_path, err := filepath.Abs(secret_path)
	if err != nil {
		return "", err
	}
	r, err := utils.FileExists(pq_abs_path)
	if err != nil {
		return "", err
	}

	if r {
		txt, err := os.ReadFile(pq_abs_path)
		if err != nil {
			log.Fatal(err)
		}
		return string(txt), nil
	}

	return "", nil
}
