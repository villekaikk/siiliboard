package api

import (
	"fmt"
	"log"
	"net/http"
	"siiliboard/internal/database"
	"siiliboard/internal/marshal"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// POST /users
func CreateUserHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	u := &marshal.UserRequest{}
	err := readBodyToModel(r, u)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, err.Error())
		return
	}

	_, err = database.CreateUser(u)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to create a new user")
		return
	}

	status := http.StatusCreated
	w.WriteHeader(status)
}

// GET /users/:uid
func GetUserHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("uid")

	idd, err := strconv.Atoi(id)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid user id %s", id)
		return
	}

	u, err := database.GetUser(idd)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("ERROR - %d - %s\n", errCode, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch user")
		return
	}

	status := http.StatusOK
	w.WriteHeader(status)
	fmt.Fprintf(w, "User: %s, %s", u.Name, u.DisplayName)
}
