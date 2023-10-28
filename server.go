package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"server/views"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
        log.Fatal("Error loading .env file")
    }

	router := httprouter.New()
    registerViews(router)
	port := 8080
	portStr := fmt.Sprintf(":%v", port)

	fmt.Println()
	fmt.Printf("Starting the server at http:localhost%v\n", portStr)

    log.Fatal(http.ListenAndServe(portStr, router))
}

func registerViews(router *httprouter.Router) {

	router.GET("/", views.Index)
	
}