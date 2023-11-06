package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	
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

	router := routes.NewRouter()
	port = fmt.Sprintf(":%v", port)

	log.Printf("Starting the server at http:127.0.0.1%v\n", port)

    log.Fatal(http.ListenAndServe(port, router))
}

func setupLogger() {
	log.SetFlags(log.Lmsgprefix)
	datetime := "2006-01-02 15:04:05: "
	log.SetPrefix(time.Now().Format(datetime))
}