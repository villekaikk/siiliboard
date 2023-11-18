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

	endpoint := fmt.Sprintf("POST %s", r.URL)
	log.Println(endpoint)
	u := &marshal.UserRequest{}
	err := readBodyToModel(r, u)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, err.Error())
		return
	}

	user, err := database.CreateUser(u)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to create a new user")
		return
	}

	status := http.StatusCreated
	log.Printf("%d - %s - %d\n", status, endpoint, user.Id)
	w.WriteHeader(status)
}

// GET /users/:uid
func GetUserHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("uid")
	endpoint := fmt.Sprintf("GET %s", r.URL)
	log.Println(endpoint)

	idd, err := strconv.Atoi(id)

	if err != nil {
		errCode := http.StatusBadRequest
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Invalid user id %s", id)
		return
	}

	u, err := database.GetUser(idd)

	if err != nil {
		errCode := http.StatusInternalServerError
		log.Printf("%d - %s - %s\n", errCode, endpoint, err.Error())
		w.WriteHeader(errCode)
		fmt.Fprintf(w, "Failed to fetch user")
		return
	}

	status := http.StatusOK
	log.Printf("%d - %s", status, endpoint)
	//w.WriteHeader(status)
	fmt.Fprintf(w, "User: %s, %s", u.Name, u.DisplayName)
}
